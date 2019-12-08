/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"

	git "gopkg.in/src-d/go-git.v4"
)

const (
	path = "/github"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll("/tmp", 0700); err != nil {
			log.Fatal(err)
		}

		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			payload, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("failed to read payload: %+v\n", err)
				return
			}

			go func() {
				dir, err := ioutil.TempDir("/tmp", "conform")
				if err != nil {
					log.Printf("failed to create temporary directory: %+v\n", err)
					return
				}
				// nolint: errcheck
				defer os.RemoveAll(dir)

				if err := os.MkdirAll(filepath.Join(dir, "github"), 0700); err != nil {
					log.Printf("failed to create github directory: %+v\n", err)
					return
				}
				if err := os.MkdirAll(filepath.Join(dir, "repo"), 0700); err != nil {
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
				_, err = git.PlainClone(cloneRepo, false, &git.CloneOptions{
					URL:      pullRequestEvent.GetRepo().GetCloneURL(),
					Progress: os.Stdout,
				})

				if err = ioutil.WriteFile(event, payload, 0600); err != nil {
					log.Printf("failed to clone repo: %+v\n", err)
					return
				}

				log.Printf("writing %s to disk", event)
				if err = ioutil.WriteFile(event, payload, 0600); err != nil {
					log.Printf("failed to write event to disk: %+v\n", err)
					return
				}
				cmd := exec.Command("/proc/self/exe", "enforce", "--summary=github")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stdout
				cmd.Dir = cloneRepo
				cmd.Env = []string{fmt.Sprintf("GITHUB_TOKEN=%s", os.Getenv("GITHUB_TOKEN")), fmt.Sprintf("GITHUB_EVENT_PATH=%s", event)}
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

		http.ListenAndServe(":3000", nil)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
