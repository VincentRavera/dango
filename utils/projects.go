package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/VincentRavera/dango/data"
	"github.com/go-git/go-git/v5"
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
	return ref.Name().String(), nil
}

// TODO: delegate writing
func AddProject(project data.Project, rootC data.RootConfig) error {

	rootC.Configuration.Projects = append(rootC.Configuration.Projects, project)
	cfbytes, e := json.Marshal(rootC.Configuration)
	if e != nil {
		return fmt.Errorf("AddProject: %s", e)
	}
	currentJsonPath := fmt.Sprintf("%s/current.json", rootC.RootPath)
	currentConfFile, e := os.Create(currentJsonPath)
	defer currentConfFile.Close()
	if e != nil {
		return fmt.Errorf("AddProject: %s", e)
	}
	currentConfFile.Write(cfbytes)
	return nil
}
