package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/VincentRavera/dango/data"
)

func TestGetConfig(t *testing.T) {
	rootPath := "../test_resources/utils/FAKE_ROOT"
	os.Setenv("DANGO_ROOT", rootPath)
	os.Setenv("DANGO_BATCH", "bb")
	os.Setenv("DANGO_WORKSPACE", "cc")
	tests := []struct {
		name string
		want data.RootConfig
	}{
		{"Test0", data.RootConfig{
			RootPath: rootPath,
			Batch:    true,
			WorkPath: "cc",
			Configuration: data.CurrentConfig{
				Type:     "development",
				Name:     "main",
				Projects: []data.Project{},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadEnviron(t *testing.T) {
	os.Setenv("DANGO_ROOT", "aa")
	os.Setenv("DANGO_BATCH", "bb")
	os.Setenv("DANGO_WORKSPACE", "cc")
	tests := []struct {
		name  string
		want  string
		want1 string
		want2 bool
	}{
		{"Test0", "aa", "cc", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ReadEnviron()
			if got != tt.want {
				t.Errorf("ReadEnviron() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadEnviron() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ReadEnviron() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestParseConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want data.CurrentConfig
	}{
		{"Test0", args{
			path: "../test_resources/utils/FAKE_ROOT/current.json",
		}, data.CurrentConfig{
			Type:     "development",
			Name:     "main",
			Projects: []data.Project{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateConfig(t *testing.T) {
	rootPath := "../test_resources/utils/FAKE_ROOT"
	os.Setenv("DANGO_ROOT", rootPath)
	os.Setenv("DANGO_BATCH", "bb")
	os.Setenv("DANGO_WORKSPACE", "cc")
	rc := GetConfig()
	rc.Configuration.Type = "Test"
	type args struct {
		newRootConfig data.RootConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test0", args{newRootConfig: rc}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateConfig(tt.args.newRootConfig); (err != nil) != tt.wantErr {
				t.Errorf("UpdateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if rc.Configuration.Type != *&rootConfig.Configuration.Type {
		t.Errorf("rootConfig: expected %s, got: %s",
			rc.Configuration.Type, rootConfig.Configuration.Type)
	}
}

func TestSaveConfig(t *testing.T) {
	rootPath := "../test_resources/utils/FAKE_ROOT"
	os.Setenv("DANGO_ROOT", rootPath)
	os.Setenv("DANGO_BATCH", "bb")
	os.Setenv("DANGO_WORKSPACE", "cc")
	rc := GetConfig()
	rc.Configuration.Type = "Test"
	dir, err := ioutil.TempDir("./", "prefix")
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("Dire::%s\n", dir)
	// f, _ := ioutil.ReadDir("./")
	// for _, value := range f {
	// 	fmt.Println(value.Name(), value.IsDir())
	// }
	defer os.RemoveAll(dir)
	rc.RootPath = dir
	UpdateConfig(rc)
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Test0", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveConfig(); (err != nil) != tt.wantErr {
				t.Errorf("SaveConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	got := ParseConfig(filepath.Join(dir, "current.json"))
	expected := ParseConfig(filepath.Join(rootPath, "current.json"))

	if got.Type == expected.Type {
		t.Errorf("Expected %s to be diffrent of %s", expected.Type, got.Type)
	}

}

func TestProcessSystemErrors(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProcessSystemErrors(tt.args.e)
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Test relative", args{path: "../test_resources/"}, true, false },
		{"Test fail relative", args{path: "../XXXX"}, false, false },
		{"Test Absolute", args{path: "/home"}, true, false },
		{"Test Absolute", args{path: "/XXXX"}, false, false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateProject(t *testing.T) {
	rootPath := "../test_resources/utils/FAKE_ROOT"
	os.Setenv("DANGO_ROOT", rootPath)
	os.Setenv("DANGO_BATCH", "bb")
	os.Setenv("DANGO_WORKSPACE", "cc")
	GetConfig()
	project, err := ScanPath("../")
	if err != nil {
		t.Error("Could not scan a project")
	}
	AddProject(project)
	alteredProject := project
	alteredProject.Name = "Altered"
	type args struct {
		id      int
		updated data.Project
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test0", args{0, alteredProject}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateProject(tt.args.id, tt.args.updated)
		})
	}
	newRC := GetConfig()
	if newRC.Configuration.Projects[0].Name != alteredProject.Name {
		t.Errorf("UpdateProject() = %v, want %v", newRC.Configuration.Projects[0], alteredProject)
	}
}

func TestAddProject(t *testing.T) {
	rootPath := "../test_resources/utils/FAKE_ROOT"
	os.Setenv("DANGO_ROOT", rootPath)
	os.Setenv("DANGO_BATCH", "bb")
	os.Setenv("DANGO_WORKSPACE", "cc")
	GetConfig()
	project, err := ScanPath("../")
	if err != nil {
		t.Error("Could not scan a project")
	}
	type args struct {
		newProject data.Project
	}
	tests := []struct {
		name string
		args args
	}{
		{ "Test 0", args{newProject: project} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddProject(tt.args.newProject)
		})
	}
	newRC := GetConfig()
	if newRC.Configuration.Projects[0].URL != project.URL {
		t.Errorf("Expected %s but got %s", project.URL, newRC.Configuration.Projects[0].URL)
	}
}
