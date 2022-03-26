package main

//go:generate go run go-exp/build-info
//
// ref: https://tip.golang.org/doc/go1.18#debug/buildinfo
//
// New debug/buildinfo package
//
// The new debug/buildinfo package provides access to module versions, version control
// information, and build flags embedded in executable files built by the go command.
//
// The same information is also available via runtime/debug.ReadBuildInfo for the currently
// running binary and via go version -m on the command line.

import (
	"github.com/davecgh/go-spew/spew"
	"runtime/debug"
)

func main() {
	if bi, ok := debug.ReadBuildInfo(); ok {
		spew.Config.Indent = "    "
		spew.Config.DisableMethods = true
		spew.Dump(bi)
	}
}

/*
(*debug.BuildInfo)(0xc000120000)({
    GoVersion: (string) (len=6) "go1.18",
    Path: (string) (len=16) "go-exp/buildinfo",
    Main: (debug.Module) {
        Path: (string) (len=6) "go-exp",
        Version: (string) (len=7) "(devel)",
        Sum: (string) "",
        Replace: (*debug.Module)(<nil>)
    },
    Deps: ([]*debug.Module) (len=1 cap=1) {
        (*debug.Module)(0xc000022080)({
            Path: (string) (len=26) "github.com/davecgh/go-spew",
            Version: (string) (len=6) "v1.1.1",
            Sum: (string) (len=47) "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
            Replace: (*debug.Module)(<nil>)
        })
    },
    Settings: ([]debug.BuildSetting) (len=13 cap=16) {
        (debug.BuildSetting) {
            Key: (string) (len=9) "-compiler",
            Value: (string) (len=2) "gc"
        },
        (debug.BuildSetting) {
            Key: (string) (len=11) "CGO_ENABLED",
            Value: (string) (len=1) "1"
        },
        (debug.BuildSetting) {
            Key: (string) (len=10) "CGO_CFLAGS",
            Value: (string) ""
        },
        (debug.BuildSetting) {
            Key: (string) (len=12) "CGO_CPPFLAGS",
            Value: (string) ""
        },
        (debug.BuildSetting) {
            Key: (string) (len=12) "CGO_CXXFLAGS",
            Value: (string) ""
        },
        (debug.BuildSetting) {
            Key: (string) (len=11) "CGO_LDFLAGS",
            Value: (string) ""
        },
        (debug.BuildSetting) {
            Key: (string) (len=6) "GOARCH",
            Value: (string) (len=5) "amd64"
        },
        (debug.BuildSetting) {
            Key: (string) (len=4) "GOOS",
            Value: (string) (len=6) "darwin"
        },
        (debug.BuildSetting) {
            Key: (string) (len=7) "GOAMD64",
            Value: (string) (len=2) "v1"
        },
        (debug.BuildSetting) {
            Key: (string) (len=3) "vcs",
            Value: (string) (len=3) "git"
        },
        (debug.BuildSetting) {
            Key: (string) (len=12) "vcs.revision",
            Value: (string) (len=40) "9fe078e9246c733ed9e5b985194f582f382ebbf3"
        },
        (debug.BuildSetting) {
            Key: (string) (len=8) "vcs.time",
            Value: (string) (len=20) "2022-03-21T12:50:44Z"
        },
        (debug.BuildSetting) {
            Key: (string) (len=12) "vcs.modified",
            Value: (string) (len=4) "true"
        }
    }
})
*/
