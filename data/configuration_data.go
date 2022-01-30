package data

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
	// Remote url to checkout the project
	Url string
	// Revision to set, can be a branch name or a tag
	Revision string
	// Build receipe of a project
	BuildContext Build
	// Dependencies, will hint at build order
	Dependencies []string
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
