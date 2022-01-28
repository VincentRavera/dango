package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/VincentRavera/dango/data"
)

func ProcessSystemErrors(e error, l log.Logger) {
	if e != nil {
		l.Fatal(e)
		os.Exit(1)
	}
}

func GetConfig(l log.Logger) data.RootConfig {
	rootpath, batch := ReadEnviron(l)

	currentJsonPath := fmt.Sprintf("%s/current.json", rootpath)

	config := data.RootConfig{
		RootPath: rootpath,
		Batch: batch,
		Configuration: ParseConfig(currentJsonPath, l),
	}
	l.Print(config)

	return config
}

func ReadEnviron(l log.Logger) (string, bool) {
	rootpath := os.Getenv("DANGO_ROOT")
	if len(rootpath) == 0 {
		var e error
		rootpath, e = os.Getwd()
		ProcessSystemErrors(e, l)
	}
	l.Printf("DANGO_ROOT=%s\n", rootpath)

	batch := false
	batch_mode := os.Getenv("DANGO_BATCH")
	if len(batch_mode) > 0 {
		batch = true
	}
	l.Printf("DANGO_BATCH=%t\n", batch)
	return rootpath, batch
}

// https://tutorialedge.net/golang/parsing-json-with-golang/
func ParseConfig(path string, l log.Logger) data.CurrentConfig {
	jsonFile, err := os.Open(path)
	ProcessSystemErrors(err, l)
	byteValue, err := ioutil.ReadAll(jsonFile)
	ProcessSystemErrors(err, l)
	var currentConfig data.CurrentConfig
	json.Unmarshal(byteValue, &currentConfig)
	return currentConfig
}
