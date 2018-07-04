package release

import (
	"context"

	"github.com/autonomy/conform/pkg/metadata"
	"github.com/google/go-github/github"

	"golang.org/x/oauth2"
)

// Release is an interface that describes a release.
type Release interface {
	Create(*metadata.Metadata) error
	Upload() error
}

// GitHub implements the Release interface. It provides release functionality
// specific to GitHub.
type GitHub struct {
	Owner      string `mapstructure:"owner"`
	Repository string `mapstructure:"repository"`
	Client     *github.Client
	ctx        context.Context
}

// NewGitHub initializes and returns a GitHub struct.
func NewGitHub(owner, repo, token string) *GitHub {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GitHub{
		Owner:      owner,
		Repository: repo,
		Client:     client,
		ctx:        ctx,
	}
}

// Create implements the Release interface.
func (g *GitHub) Create(meta *metadata.Metadata) (err error) {
	release := &github.RepositoryRelease{
		Name:            &meta.Git.Tag,
		TagName:         &meta.Git.Tag,
		TargetCommitish: &meta.Git.SHA,
		Prerelease:      &meta.Version.IsPrerelease,
	}
	_, _, err = g.Client.Repositories.CreateRelease(g.ctx, g.Owner, g.Repository, release)
	if err != nil {
		return err
	}

	return nil
}

// Upload implements the Release interface.
func (g *GitHub) Upload() (err error) {
	return err
}
