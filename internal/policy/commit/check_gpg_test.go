// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//nolint:testpackage
package commit

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	pgp "github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/keybase/go-crypto/openpgp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/siderolabs/conform/internal/git"
	"github.com/siderolabs/conform/internal/policy"
)

type gpgTestRepo struct {
	repo     *gogit.Repository
	worktree *gogit.Worktree
	n        int
}

func newGPGTestRepo(t *testing.T) *gpgTestRepo {
	t.Helper()

	dir := t.TempDir()
	t.Chdir(dir)

	repo, err := gogit.PlainInit(dir, false)
	require.NoError(t, err)

	worktree, err := repo.Worktree()
	require.NoError(t, err)

	return &gpgTestRepo{repo: repo, worktree: worktree}
}

// commit creates a commit signed with the key (unsigned if key is nil) and returns its hash.
func (r *gpgTestRepo) commit(t *testing.T, key *pgp.Entity) string {
	t.Helper()

	r.n++

	require.NoError(t, os.WriteFile("test", fmt.Append(nil, r.n), 0o644))

	_, err := r.worktree.Add("test")
	require.NoError(t, err)

	hash, err := r.worktree.Commit(fmt.Sprintf("feat: change %d", r.n), &gogit.CommitOptions{
		Author: &object.Signature{
			Name:  "test",
			Email: "test@example.com",
			When:  time.Now(),
		},
		SignKey: key,
	})
	require.NoError(t, err)

	return hash.String()
}

// merge creates a merge commit of the parents signed with the key (unsigned if key is nil).
func (r *gpgTestRepo) merge(t *testing.T, key *pgp.Entity, parents ...string) {
	t.Helper()

	parentHashes := make([]plumbing.Hash, 0, len(parents))
	for _, parent := range parents {
		parentHashes = append(parentHashes, plumbing.NewHash(parent))
	}

	_, err := r.worktree.Commit("chore: merge", &gogit.CommitOptions{
		Author: &object.Signature{
			Name:  "test",
			Email: "test@example.com",
			When:  time.Now(),
		},
		Parents:           parentHashes,
		AllowEmptyCommits: true,
		SignKey:           key,
	})
	require.NoError(t, err)
}

// checkout detaches HEAD at the commit.
func (r *gpgTestRepo) checkout(t *testing.T, hash string) {
	t.Helper()

	require.NoError(t, r.worktree.Checkout(&gogit.CheckoutOptions{Hash: plumbing.NewHash(hash), Force: true}))
}

func newSigningKey(t *testing.T, name string) *pgp.Entity {
	t.Helper()

	key, err := pgp.NewEntity(name, "", name+"@example.com", nil)
	require.NoError(t, err)

	return key
}

// useKeyring makes the GPG identity check use the public keys instead of fetching them from GitHub.
func useKeyring(t *testing.T, keys ...*pgp.Entity) {
	t.Helper()

	keyring := make(openpgp.EntityList, 0, len(keys))

	for _, key := range keys {
		var buf bytes.Buffer

		w, err := armor.Encode(&buf, pgp.PublicKeyType, nil)
		require.NoError(t, err)

		require.NoError(t, key.Serialize(w))
		require.NoError(t, w.Close())

		el, err := git.ReadArmoredKeyrings([]string{buf.String()})
		require.NoError(t, err)

		keyring = append(keyring, el...)
	}

	original := fetchOrgKeyring
	fetchCalls := 0

	fetchOrgKeyring = func(context.Context, string) (openpgp.EntityList, error) {
		fetchCalls++

		return keyring, nil
	}

	t.Cleanup(func() {
		fetchOrgKeyring = original

		assert.LessOrEqual(t, fetchCalls, 1, "keyring should be fetched at most once")
	})
}

func gpgPolicy() *Commit {
	return &Commit{
		GPG: &GPG{
			Required: true,
			Identity: &struct {
				GitHubOrganization string `mapstructure:"gitHubOrganization"`
			}{
				GitHubOrganization: "example",
			},
		},
	}
}

// failedChecks returns the error messages of the failed checks with the given name.
func failedChecks(report *policy.Report, name string) []string {
	var failed []string

	for _, check := range report.Checks() {
		if check.Name() == name && len(check.Errors()) > 0 {
			failed = append(failed, check.Message())
		}
	}

	return failed
}

func countChecks(report *policy.Report, name string) int {
	var n int

	for _, check := range report.Checks() {
		if check.Name() == name {
			n++
		}
	}

	return n
}

func TestGPGRevisionRangeChecksEveryCommit(t *testing.T) {
	r := newGPGTestRepo(t)

	member := newSigningKey(t, "member")
	outsider := newSigningKey(t, "outsider")

	useKeyring(t, member)

	base := r.commit(t, member)
	r.commit(t, member)
	unsigned := r.commit(t, nil)
	foreign := r.commit(t, outsider)
	r.commit(t, member) // HEAD is signed by a member

	report, err := gpgPolicy().Compliance(&policy.Options{
		RevisionRange: base + "..HEAD",
	})
	require.NoError(t, err)

	assert.False(t, report.Valid())

	assert.Equal(t, 4, countChecks(report, "GPG"))
	assert.Equal(t, 4, countChecks(report, "GPG Identity"))

	assert.Equal(
		t,
		[]string{fmt.Sprintf("Commit %s does not have a GPG signature", unsigned[:shortHashLength])},
		failedChecks(report, "GPG"),
	)

	failedIdentity := failedChecks(report, "GPG Identity")
	require.Len(t, failedIdentity, 2)
	assert.Contains(t, failedIdentity[0], "commit "+foreign[:shortHashLength]+":")
	assert.Contains(t, failedIdentity[1], "commit "+unsigned[:shortHashLength]+":")
}

func TestGPGRevisionRangeAllSigned(t *testing.T) {
	r := newGPGTestRepo(t)

	member1 := newSigningKey(t, "member1")
	member2 := newSigningKey(t, "member2")

	useKeyring(t, member1, member2)

	base := r.commit(t, nil) // base is not part of the range
	r.commit(t, member1)
	r.commit(t, member2)

	report, err := gpgPolicy().Compliance(&policy.Options{
		RevisionRange: base + "..HEAD",
	})
	require.NoError(t, err)

	assert.True(t, report.Valid())
	assert.Equal(t, 2, countChecks(report, "GPG Identity"))
}

func TestGPGHead(t *testing.T) {
	r := newGPGTestRepo(t)

	member := newSigningKey(t, "member")

	useKeyring(t, member)

	r.commit(t, member)

	c := gpgPolicy()

	report, err := c.Compliance(&policy.Options{})
	require.NoError(t, err)

	assert.True(t, report.Valid())

	r.commit(t, nil)

	report, err = c.Compliance(&policy.Options{})
	require.NoError(t, err)

	assert.False(t, report.Valid())
	assert.Len(t, failedChecks(report, "GPG"), 1)
}

func TestDivergedRevisionRange(t *testing.T) {
	r := newGPGTestRepo(t)

	member := newSigningKey(t, "member")

	useKeyring(t, member)

	base := r.commit(t, member)

	// the PR branch
	require.NoError(t, r.worktree.Checkout(&gogit.CheckoutOptions{Branch: plumbing.NewBranchReferenceName("pr"), Create: true}))
	r.commit(t, member)
	r.commit(t, member)

	// the main branch moves ahead with an unsigned commit, which is not part of the PR
	require.NoError(t, r.worktree.Checkout(&gogit.CheckoutOptions{Hash: plumbing.NewHash(base), Force: true}))
	mainHead := r.commit(t, nil)

	require.NoError(t, r.worktree.Checkout(&gogit.CheckoutOptions{Branch: plumbing.NewBranchReferenceName("pr")}))

	report, err := gpgPolicy().Compliance(&policy.Options{
		RevisionRange: mainHead + "..HEAD",
	})
	require.NoError(t, err)

	assert.True(t, report.Valid())
	assert.Equal(t, 2, countChecks(report, "GPG"))
}

func TestEmptyRevisionRange(t *testing.T) {
	r := newGPGTestRepo(t)

	useKeyring(t)

	r.commit(t, nil)

	report, err := gpgPolicy().Compliance(&policy.Options{
		RevisionRange: "HEAD..HEAD",
	})
	require.NoError(t, err)

	assert.True(t, report.Valid())
	assert.Empty(t, report.Checks())
}

func TestCommitsRange(t *testing.T) {
	r := newGPGTestRepo(t)

	base := r.commit(t, nil)
	second := r.commit(t, nil)
	third := r.commit(t, nil)

	g, err := git.NewGit()
	require.NoError(t, err)

	commits, err := g.Commits(base, "HEAD")
	require.NoError(t, err)

	hashes := make([]string, 0, len(commits))
	for _, c := range commits {
		hashes = append(hashes, c.Hash.String())
	}

	assert.Equal(t, []string{third, second}, hashes)

	head, err := r.repo.ResolveRevision(plumbing.Revision("HEAD"))
	require.NoError(t, err)
	assert.Equal(t, third, head.String())
}

func TestGPGRevisionRangeBaseMergedIntoBranch(t *testing.T) {
	r := newGPGTestRepo(t)

	member := newSigningKey(t, "member")

	useKeyring(t, member)

	base := r.commit(t, member)
	pr := r.commit(t, member)

	// the main branch moves ahead with an unsigned commit, which is not part of the PR
	r.checkout(t, base)
	mainHead := r.commit(t, nil)

	// main is merged into the PR branch, so it is reachable only via the second parent
	r.checkout(t, pr)
	r.merge(t, member, pr, mainHead)

	report, err := gpgPolicy().Compliance(&policy.Options{
		RevisionRange: mainHead + "..HEAD",
	})
	require.NoError(t, err)

	assert.True(t, report.Valid())
	assert.Equal(t, 2, countChecks(report, "GPG")) // the merge commit and pr
}

func TestGPGRevisionRangeUnsignedCommitBehindMerge(t *testing.T) {
	r := newGPGTestRepo(t)

	member := newSigningKey(t, "member")

	useKeyring(t, member)

	base := r.commit(t, member)
	signed := r.commit(t, member)

	// the unsigned commit is merged into the branch via the second parent of a signed merge commit
	r.checkout(t, base)
	unsigned := r.commit(t, nil)

	r.checkout(t, signed)
	r.merge(t, member, signed, unsigned)

	report, err := gpgPolicy().Compliance(&policy.Options{
		RevisionRange: base + "..HEAD",
	})
	require.NoError(t, err)

	assert.False(t, report.Valid())
	assert.Equal(t, 3, countChecks(report, "GPG"))
	assert.Equal(
		t,
		[]string{fmt.Sprintf("Commit %s does not have a GPG signature", unsigned[:shortHashLength])},
		failedChecks(report, "GPG"),
	)
}

func TestCommitsReversedRange(t *testing.T) {
	r := newGPGTestRepo(t)

	base := r.commit(t, nil)
	head := r.commit(t, nil)

	g, err := git.NewGit()
	require.NoError(t, err)

	_, err = g.Commits(head, base)
	require.ErrorContains(t, err, "invalid revision range")
}
