package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/VincentRavera/dango/data"
)

func ScanPath(path string) data.Project {
    return data.Project{
    	Name:         filepath.Base(path),
    	Location:     path,
    	Url:          "",
    	Revision:     "",
    	BuildContext: data.Build{},
    	Dependencies: []string{},
    }
}

func AddProject(project data.Project, rootC data.RootConfig) {

	rootC.Configuration.Projects = append(rootC.Configuration.Projects, project)
	cfbytes, e := json.Marshal(rootC.Configuration)
	if e != nil {
		fmt.Errorf("Error %s", e)
	}
	currentJsonPath := fmt.Sprintf("%s/current.json", rootC.RootPath)
	currentConfFile, e := os.Create(currentJsonPath)
	defer currentConfFile.Close()
	if e != nil {
		fmt.Errorf("Error %s", e)
	}
	currentConfFile.Write(cfbytes)
}
