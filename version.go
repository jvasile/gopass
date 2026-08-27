package main

import (
	"strings"

	"github.com/blang/semver/v4"
	"jamesvasile.com/go/gopass/v2/pkg/debug"
)

func getVersion() semver.Version {
	sv, err := semver.Parse(strings.TrimPrefix(version, "v"))
	if err == nil {
		return sv
	}

	if sv := debug.ModuleVersion("jamesvasile.com/go/gopass/v2"); sv.String() != "0.0.0" {
		return sv
	}

	return semver.Version{
		Major: 1,
		Minor: 16,
		Patch: 1,
		Pre: []semver.PRVersion{
			{VersionStr: "git"},
		},
		Build: []string{"d601d3ef"},
	}
}
