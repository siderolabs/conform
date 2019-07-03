/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package summarizer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/autonomy/conform/internal/git"
	"github.com/google/go-github/github"
)

// Summarizer describes a hook for send summarized results to a remote API.
type Summarizer interface {
	SetStatus(string, string, string, string) error
}

// GitHub is a summarizer that can be used with GitHub.
type GitHub struct {
	token string
	owner string
	repo  string
	sha   string
}

// Noop is a summarizer that does nothing.
type Noop struct {
}

// SetStatus is a noop func.
func (n *Noop) SetStatus(state, policy, check, message string) error {
	return nil
}

// NewGitHubSummarizer returns a summarizer that posts policy checks as status
// checks on a pull request.
func NewGitHubSummarizer(token string) (*GitHub, error) {
	eventPath, ok := os.LookupEnv("GITHUB_EVENT_PATH")
	if !ok {
		return nil, errors.New("GITHUB_EVENT_PATH is not set")
	}

	data, err := ioutil.ReadFile(eventPath)
	if err != nil {
		return nil, err
	}

	pullRequestEvent := &github.PullRequestEvent{}
	if err = json.Unmarshal(data, pullRequestEvent); err != nil {
		return nil, err
	}

	g, err := git.NewGit()
	if err != nil {
		return nil, err
	}

	if err = g.FetchPullRequest("origin", pullRequestEvent.GetNumber()); err != nil {
		return nil, err
	}

	if err = g.CheckoutPullRequest(pullRequestEvent.GetNumber()); err != nil {
		return nil, err
	}

	sha, err := g.SHA()
	if err != nil {
		log.Fatal(err)
	}

	gh := &GitHub{
		token: token,
		owner: pullRequestEvent.GetRepo().GetOwner().GetLogin(),
		repo:  pullRequestEvent.GetRepo().GetName(),
		sha:   sha,
	}

	return gh, nil
}

// SetStatus sets the status of a GitHub check.
// Valid statuses are "error", "failure", "pending", "success"
func (gh *GitHub) SetStatus(state, policy, check, message string) error {
	if gh.token == "" {
		return errors.New("no token")
	}
	statusCheckContext := strings.ReplaceAll(strings.ToLower(path.Join("conform", policy, check)), " ", "-")
	description := message
	repoStatus := &github.RepoStatus{}
	repoStatus.Context = &statusCheckContext
	repoStatus.Description = &description
	repoStatus.State = &state

	http.DefaultClient.Transport = roundTripper{gh.token}
	githubClient := github.NewClient(http.DefaultClient)

	_, _, err := githubClient.Repositories.CreateStatus(context.Background(), gh.owner, gh.repo, gh.sha, repoStatus)
	if err != nil {
		return err
	}

	return nil
}

type roundTripper struct {
	accessToken string
}

// RoundTrip implements the net/http.RoundTripper interface.
func (rt roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rt.accessToken))
	return http.DefaultTransport.RoundTrip(r)
}
