package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/VincentRavera/dango/data"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func ScanPath(path string) (data.Project, error) {
	abpath, err := filepath.Abs(path)
	if err != nil {
		return data.Project{}, err
	}
	remoteName, err := getProjectRemote(abpath)
	if err != nil {
		return data.Project{}, err
	}
	revision, err := getProjectRevision(abpath)
	if err != nil {
		return data.Project{}, err
	}
	url, err := getProjectRemoteUrl(abpath)
	if err != nil {
		return data.Project{}, err
	}
    return data.Project{
    	Name:         filepath.Base(path),
    	Location:     abpath,
		Remote: remoteName,
		URL: url,
    	Revision:     revision,
    	BuildContext: data.Build{},
    	Dependencies: []string{},
    }, nil
}

func getProjectRemote(path string) (string, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", fmt.Errorf("GitOpenError: %s", err)
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return "", fmt.Errorf("GitOpenRemoteError: %s", err)
	}
	remotename := remotes[0].Config().Name
	return remotename, nil
}

func getProjectRemoteUrl(path string) (string, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", fmt.Errorf("GitOpenError: %s", err)
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return "", fmt.Errorf("GitOpenRemoteError: %s", err)
	}
	remoteurl := remotes[0].Config().URLs[0]
	return remoteurl, nil
}

func getProjectRevision(path string) (string, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", fmt.Errorf("GitOpenError: %s", err)
	}
	ref, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("GitHeadError: %s", err)
	}
	if strings.HasPrefix(ref.Name().String(), "ref") {
		return ref.Name().String(), nil
	} else {
		return ref.Hash().String(), nil
	}
}

// Clone the project, returns a log, or an error
func CloneProject(id int, project data.Project, rc data.RootConfig) (string, error) {
	output := ""
	// Test location if is already present
	isExists, err := Exists(project.Location)
	if err != nil {
		return output, err
	}
	if isExists  {
		return "", nil
	}
	// Clone
	newLocation := filepath.Join(rc.WorkPath, project.Name)
	output += fmt.Sprintf("Cloning %s to %s at %s",
		project.Name, newLocation, project.Revision)
	gp, err := git.PlainClone(newLocation, false, &git.CloneOptions{
		URL:      project.URL,
		RemoteName: project.Remote,
	})
	if err != nil {
		return "", fmt.Errorf("Cannot clone %s: %v", project.Name, err)
	}
	// Worktree is used to checkout
	wt, err := gp.Worktree()
	if err != nil {
		return "", fmt.Errorf("Worktree failure %s: %v", project.Name, err)
	}
	// Checkout
	// Check if is a reference or a hash
	var cop *git.CheckoutOptions
	isHash := isHash(project.Revision)
	if isHash {
		revisionHash, err := gp.ResolveRevision(plumbing.Revision(project.Revision))
		if err != nil {
			return "", fmt.Errorf("Cannot resolve hash value %s for %s: %v",
				project.Revision, project.Name, err)
		}
		cop = &git.CheckoutOptions{
			Hash: *revisionHash,
		}
	} else {
		// Not a hash
		cop, err = findMatchingRemoteBranch(project.Revision, *gp)
		if err != nil {
			return "", fmt.Errorf("Cannot solve reference %s in project %s: %s", project.Revision, project.Name, err)
		}
	}
	// Actual checkout
	wt.Checkout(cop)
	// Update location
	project.Location = newLocation
	UpdateProject(id, project)
	return output, nil
}

// Clumsy function to create a branch from a remote branch
func findMatchingRemoteBranch(referenceraw string, gp git.Repository) (*git.CheckoutOptions, error) {
	refn := plumbing.ReferenceName(referenceraw)
	var output *git.CheckoutOptions
	_, err := gp.ResolveRevision(plumbing.Revision(referenceraw))
	if err != nil {
		// Reference is not created
		// This happen if the branch exist on remote but not local
		// This won't happen on tags
		// Let's guess
		refs, _ := gp.References()
		hasMatchReference := false
		err = refs.ForEach(func(ref *plumbing.Reference) error {
			if (strings.HasSuffix(ref.Name().Short(), refn.Short())) {
				// Match found
				if (ref.Name().IsRemote()) {
					output = &git.CheckoutOptions{
						Hash:   ref.Hash(),
						Branch: refn,
						Create: true,
					}
					hasMatchReference = true
				} else {
					return fmt.Errorf(
						"Did not expect to match %s with %s",
						refn, ref)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		if ! hasMatchReference {
			return nil, errors.New("Could not resolve revision.")
		}
	} else {
		// reference already exist
		// maybe remote or tags
		output = &git.CheckoutOptions{
			Branch: refn,
		}
	}

	return output, nil
}

// Check if revision in a commit hash
func isHash(revision string) bool {
	return ! strings.Contains(revision, "/")
}
