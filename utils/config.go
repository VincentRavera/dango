package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/VincentRavera/dango/data"
)

var rootConfig *data.RootConfig
var originalRootConfig *data.RootConfig
var l = log.Default()

func ProcessSystemErrors(e error) {
	if e != nil {
		l.Fatal(e)
		os.Exit(1)
	}
}

func GetConfig() data.RootConfig {
	if rootConfig != nil {
		return *rootConfig
	}
	rootpath, workspace, batch := ReadEnviron()

	currentJsonPath := filepath.Join(rootpath, "current.json")

	rootConfig = &data.RootConfig{
		RootPath: rootpath,
		WorkPath: workspace,
		Batch: batch,
		Configuration: ParseConfig(currentJsonPath),
	}
	if originalRootConfig == nil {
		originalRootConfig = &data.RootConfig{
		RootPath: rootpath,
		WorkPath: workspace,
		Batch: batch,
		Configuration: ParseConfig(currentJsonPath),
		}
	}
	return *rootConfig
}

func ReadEnviron() (string, string, bool) {
	// Root path
	rootpath := os.Getenv("DANGO_ROOT")
	if len(rootpath) == 0 {
		var e error
		rootpath, e = os.Getwd()
		ProcessSystemErrors(e)
	}
	l.Printf("DANGO_ROOT=%s\n", rootpath)

	// Workspace path
	workspace := os.Getenv("DANGO_WORKSPACE")
	if len(rootpath) == 0 {
		workspace = fmt.Sprintf("%s/workspace", rootpath)
	}
	l.Printf("DANGO_WORKSPACE=%s\n", workspace)

	batch := false
	batch_mode := os.Getenv("DANGO_BATCH")
	if len(batch_mode) > 0 {
		batch = true
	}
	l.Printf("DANGO_BATCH=%t\n", batch)
	return rootpath, workspace, batch
}

// https://tutorialedge.net/golang/parsing-json-with-golang/
func ParseConfig(path string) data.CurrentConfig {
	jsonFile, err := os.Open(path)
	ProcessSystemErrors(err)
	byteValue, err := ioutil.ReadAll(jsonFile)
	ProcessSystemErrors(err)
	var currentConfig data.CurrentConfig
	json.Unmarshal(byteValue, &currentConfig)
	jsonFile.Close()
	return currentConfig
}

func UpdateConfig(newRootConfig data.RootConfig) error {
	rootConfig.RootPath = newRootConfig.RootPath
	rootConfig.Configuration.Type = newRootConfig.Configuration.Type
	rootConfig.Configuration.Name = newRootConfig.Configuration.Name
	return mergeProjectList(newRootConfig.Configuration)
}

func mergeProjectList(currentConfig data.CurrentConfig) error {
	// TODO Think about merging
	// newProjectList := currentConfig.Projects
	// currentProjectList := GetConfig().Configuration.Projects
	// oriProjectList := originalRootConfig.Configuration.Projects
	// panic("un")
	return nil

}

func SaveConfig() error {
	cfbytes, err := json.Marshal(&rootConfig.Configuration)
	if err != nil {
		return fmt.Errorf("SavingConfig:MarshalError: %s", err)
	}
	currentJsonPath := filepath.Join(rootConfig.RootPath, "current.json")
	currentConfFile, err := os.Create(currentJsonPath)
	if err != nil {
		return fmt.Errorf("SavingConfig:OpeningConfigError: %s", err)
	}
	_, err = currentConfFile.Write(cfbytes)
	if err != nil {
		return fmt.Errorf("SavingConfig:SavingError: %s", err)
	}
	currentConfFile.Close()
	return nil

}
