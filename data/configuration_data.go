package data

import "fmt"

// Build Context
type Build struct {
	// Init phase, should always succeed
    init []string
	// Run unit test, can fail, i won't judge
	test []string
	// Run build phase, should always succeed
	build []string
	// Run publishing phase, should always succeed
	publish []string
}

// Project Context
type Project struct {
	// short name of the project
    Name string
	// Path to the project usualy in DANGO_WORKSPACE
    Location string
	// Git remote name $(git remote)
	Remote string
	// Git remote Url $(git remote origin get-url)
	URL string
	// Revision to set, can be a branch name or a tag
	Revision string
	// Build receipe of a project
	BuildContext Build
	// Dependencies, will hint at build order
	Dependencies []string
}

func (pr Project) String() string {
	return fmt.Sprintf("Name: %s\nLocation: %s\n" +
		"URL: %s\nRemote: %s\nRevision: %s\n",
		pr.Name, pr.Location, pr.URL, pr.Remote, pr.Revision)
}

// Current Context
type CurrentConfig struct {
	// Type of the current work e.g. development or release
	Type string
	// Name of the work e.g.
	// 1.2.3 for a type release
	// fix_mem_leak for a devel
    Name string
	// List of projects descriptions
	Projects []Project
}

func (cc CurrentConfig) String() string {
	sep := "====%d====\n"
	projects := ""
	for i, pr := range cc.Projects {
		projects += fmt.Sprintf(sep, i) + pr.String()
	}
	return fmt.Sprintf("\tType: %s\n\tName: %s\nProjects:\n%s",
		cc.Type, cc.Name,projects)
}

// Root configuration
// Loads external context such as env var
type RootConfig struct {
	// Configuration parsed from the current.json
	Configuration CurrentConfig
	// Change root path
	RootPath string
	// Where to clone applications
	WorkPath string
	// Enables batch mode
	Batch bool
}

func (rc RootConfig)String() string {
	out := fmt.Sprintf("DANGO_ROOT=%s\nDANGO_WORKSPACE=%s\nDANGO_BATCH=%t\nCurrent:\n%s",
		rc.RootPath, rc.WorkPath, rc.Batch, rc.Configuration.String())
	return out
}
