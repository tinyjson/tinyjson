package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvePackagePath(t *testing.T) {
	gopath = "/home/gopath"
	goroot = "/var/goroot"

	exist = func(dirPath string) bool {
		return dirPath == "/home/gopath/src/foo/vendor/a/a" ||
			dirPath == "/home/gopath/src/foo/vendor/b/b" ||
			dirPath == "/home/gopath/src/bar/vendor/b/b" ||
			dirPath == "/home/gopath/src/vendor/c/c" ||
			dirPath == "/var/goroot/src/d/d" ||
			dirPath == "/home/gopath/src/d/d" ||
			dirPath == "/home/gopath/src/e/e"
	}

	for _, tt := range []struct {
		name string
		file string
		pkg  string
		want string
	}{
		{
			name: "vendor",
			file: "/home/gopath/src/foo/main.go",
			pkg:  "a/a",
			want: "/home/gopath/src/foo/vendor/a/a",
		},
		{
			name: "vendor",
			file: "/home/gopath/src/foo/main.go",
			pkg:  "b/b",
			want: "/home/gopath/src/foo/vendor/b/b",
		},
		{
			name: "up vendor",
			file: "/home/gopath/src/foo/main.go",
			pkg:  "c/c",
			want: "/home/gopath/src/vendor/c/c",
		},
		{
			name: "goroot",
			file: "/home/gopath/src/foo/main.go",
			pkg:  "d/d",
			want: "/var/goroot/src/d/d",
		},
		{
			name: "gopath",
			file: "/home/gopath/src/foo/main.go",
			pkg:  "e/e",
			want: "/home/gopath/src/e/e",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, resolvePackagePath(tt.file, tt.pkg))
		})
	}
}
