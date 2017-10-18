// Copyright (c) 2017 Damien Lespiau
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli"
)

func matches(pattern, filename string) (bool, error) {
	match, err := filepath.Match(pattern, filename)
	if err != nil {
		return false, err
	}

	return match, nil
}

func rm(ctx *cli.Context) error {
	args := ctx.Args()
	if len(args) != 2 {
		return fmt.Errorf("expecting two arguments, got %d", len(args))
	}

	pattern := args[0]
	profiles, err := ParseProfiles(args[1])
	if err != nil {
		return err
	}

	var newProfiles []*Profile
	for _, profile := range profiles {
		match, err := matches(pattern, profile.FileName)
		if err != nil {
			return err
		}
		if match {
			// Don't include the files matching pattern.
			continue
		}
		newProfiles = append(newProfiles, profile)
	}

	output := ctx.String("output")
	if err := WriteProfilesToFile(output, newProfiles); err != nil {
		return err
	}

	return nil
}

var rmCommand = cli.Command{

	Name:      "rm",
	Usage:     "remove file(s) from a coverage profile",
	ArgsUsage: "pattern profile",
	Action:    rm,
}
