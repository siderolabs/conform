// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// THIS FILE WAS AUTOMATICALLY GENERATED, PLEASE DO NOT EDIT.
//
// Generated on 2023-09-12T10:42:20Z by kres f9e39ef-dirty.

// Package version contains variables such as project name, tag and sha. It's a proper alternative to using
// -ldflags '-X ...'.
package version

import (
	_ "embed"
	"runtime/debug"
	"strings"
)

var (
	// Tag declares project git tag.
	//go:embed data/tag
	Tag string
	// SHA declares project git SHA.
	//go:embed data/sha
	SHA string
	// Name declares project name.
	Name = func() string {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			panic("cannot read build info, something is very wrong")
		}

		// Check if siderolabs project
		if strings.HasPrefix(info.Path, "github.com/siderolabs/") {
			return info.Path[strings.LastIndex(info.Path, "/")+1:]
		}

		// We could return a proper full path here, but it could be seen as a privacy violation.
		return "community-project"
	}()
)
