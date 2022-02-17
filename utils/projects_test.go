package utils

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/VincentRavera/dango/data"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func TestScanPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    data.Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScanPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScanPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProjectRemote(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProjectRemote(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProjectRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProjectRemote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProjectRemoteUrl(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProjectRemoteUrl(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProjectRemoteUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProjectRemoteUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProjectRevision(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProjectRevision(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProjectRevision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProjectRevision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneProject(t *testing.T) {
	type args struct {
		id      int
		project data.Project
		rc      data.RootConfig
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CloneProject(tt.args.id, tt.args.project, tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("CloneProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CloneProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMatchingRemoteBranch(t *testing.T) {
	// From go-git example_test
	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("", "clone-example")
	if err != nil {
		t.Errorf("findMatchingRemoteBranch() Could create tempDir %v", err)
	}

	defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
	})
	gp := *r

	if err != nil {
		t.Errorf("findMatchingRemoteBranch() Could not clone project %v", err)
	}
	rev, err := gp.ResolveRevision(plumbing.Revision("refs/remotes/origin/branch"))
	if err != nil {
		t.Error("findMatchingRemoteBranch() Could resolve refs/remotes/origin/branch")
	}
	type args struct {
		referenceraw string
		gp           git.Repository
	}
	tests := []struct {
		name    string
		args    args
		want    *git.CheckoutOptions
		wantErr bool
	}{
		{"test",
			args{
				referenceraw: "refs/heads/branch",
				gp: gp,
			},
			&git.CheckoutOptions{
				Hash:   *rev,
				Branch: "refs/heads/branch",
				Create: true,
			}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findMatchingRemoteBranch(tt.args.referenceraw, tt.args.gp)
			if (err != nil) != tt.wantErr {
				t.Errorf("findMatchingRemoteBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatchingRemoteBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHash(t *testing.T) {
	type args struct {
		revision string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{ "Test is hash", args{revision: "XoXoXo"}, true},
		{ "Test is ref", args{revision: "Xo/XoXo"}, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHash(tt.args.revision); got != tt.want {
				t.Errorf("isHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
