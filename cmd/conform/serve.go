// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
)

const (
	path = "/github"
)

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		if err := os.MkdirAll("/tmp", 0o700); err != nil {
			log.Fatal(err)
		}

		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			payload, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("failed to read payload: %+v\n", err)

				return
			}

			go func() {
				dir, err := os.MkdirTemp("", "conform")
				if err != nil {
					log.Printf("failed to create temporary directory: %+v\n", err)

					return
				}

				defer os.RemoveAll(dir) //nolint:errcheck

				if err = os.MkdirAll(filepath.Join(dir, "github"), 0o700); err != nil {
					log.Printf("failed to create github directory: %+v\n", err)

					return
				}
				if err = os.MkdirAll(filepath.Join(dir, "repo"), 0o700); err != nil {
					log.Printf("failed to create repo directory: %+v\n", err)

					return
				}

				event := filepath.Join(dir, "github", "event.json")
				pullRequestEvent := &github.PullRequestEvent{}
				if err = json.Unmarshal(payload, pullRequestEvent); err != nil {
					log.Printf("failed to parse pull_request event: %+v\n", err)

					return
				}

				cloneRepo := filepath.Join(dir, "repo")
				cloneURL := pullRequestEvent.GetPullRequest().GetBase().GetRepo().GetCloneURL()

				log.Printf("Cloning %s", cloneURL)

				repo, err := git.PlainClone(cloneRepo, false, &git.CloneOptions{
					SingleBranch: false,
					NoCheckout:   true,
					URL:          cloneURL,
					Progress:     os.Stdout,
				})
				if err != nil {
					log.Printf("failed to clone repo: %+v\n", err)

					return
				}

				id := pullRequestEvent.GetPullRequest().GetNumber()

				ref := pullRequestEvent.GetPullRequest().GetHead().GetRef()

				refSpec := fmt.Sprintf("refs/pull/%d/head:%s", id, ref)

				err = repo.Fetch(&git.FetchOptions{
					RefSpecs: []config.RefSpec{
						config.RefSpec("refs/heads/*:refs/heads/*"),
						config.RefSpec(refSpec),
					},
					Progress: os.Stdout,
				})
				if err != nil {
					log.Printf("failed to fetch %q: %v", refSpec, err)

					return
				}

				worktree, err := repo.Worktree()
				if err != nil {
					log.Printf("failed to get working tree: %v", err)

					return
				}

				err = worktree.Checkout(&git.CheckoutOptions{
					Branch: plumbing.NewBranchReferenceName(ref),
				})
				if err != nil {
					log.Printf("failed to checkout %q: %v", ref, err)

					return
				}

				log.Printf("writing %s to disk", event)

				if err = os.WriteFile(event, payload, 0o600); err != nil {
					log.Printf("failed to write event to disk: %+v\n", err)

					return
				}
				cmd := exec.Command("/proc/self/exe", "enforce", "--reporter=github", "--commit-ref=refs/heads/"+pullRequestEvent.GetPullRequest().GetBase().GetRef())
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stdout
				cmd.Dir = cloneRepo
				cmd.Env = []string{fmt.Sprintf("INPUT_TOKEN=%s", os.Getenv("INPUT_TOKEN")), fmt.Sprintf("GITHUB_EVENT_PATH=%s", event)}
				err = cmd.Start()
				if err != nil {
					log.Printf("failed to start command: %+v\n", err)

					return
				}
				err = cmd.Wait()
				if err != nil {
					log.Printf("command failed: %+v\n", err)

					return
				}
			}()

			w.WriteHeader(http.StatusOK)
		})

		log.Fatal(http.ListenAndServe(":3000", nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}
