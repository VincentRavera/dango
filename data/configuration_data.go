package data

// Build Context
type Build struct {
    init string
	test string
	build string
	release string
}

// Project Context
type Project struct {
    Name string
	Remote string
	Url string
	Revision string
	Type string
	BuildContext Build
	Dependencies *string
}

// Current Context
type CurrentConfig struct {
    Name string
	Projects *Project
}

// Root configuration
// Loads external context such as env var
type RootConfig struct {
	Configuration CurrentConfig
	RootPath string
	Batch bool
}
