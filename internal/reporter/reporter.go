/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// Package reporter provides check result reporting.
package reporter

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

	"github.com/google/go-github/github"

	"github.com/talos-systems/conform/internal/git"
)

// Reporter describes a hook for sending summarized results to a remote API.
type Reporter interface {
	SetStatus(string, string, string, string) error
}

// GitHub is a reporter that summarizes policy statuses as GitHub statuses.
type GitHub struct {
	token string
	owner string
	repo  string
	sha   string
}

// Noop is a reporter that does nothing.
type Noop struct{}

// SetStatus is a noop func.
func (n *Noop) SetStatus(state, policy, check, message string) error {
	return nil
}

// NewGitHubReporter returns a reporter that posts policy checks as
// status checks on a pull request.
func NewGitHubReporter() (*GitHub, error) {
	token, ok := os.LookupEnv("INPUT_TOKEN")
	if !ok {
		return nil, errors.New("missing INPUT_TOKEN")
	}

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
//
// Valid statuses are "error", "failure", "pending", "success".
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
