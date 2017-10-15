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

// This files contains profile-related functions. They are not in profile.go as
// that file taken from the go source code and have it stay pristine makes
// synching with upstream easier.

package main

import (
	"fmt"
	"io"
)

// WriteProfiles writes profiles out to w.
func WriteProfiles(w io.Writer, profiles []*Profile) error {
	if len(profiles) == 0 {
		return nil
	}

	// We've checked that all input profiles have compatible modes, so we just
	// write out the first one.
	fmt.Fprintf(w, "mode: %s\n", profiles[0].Mode)
	for _, profile := range profiles {
		blocks := profile.Blocks
		for i := range profile.Blocks {
			fmt.Fprintf(w, "%s:%d.%d,%d.%d %d %d\n", profile.FileName,
				blocks[i].StartLine, blocks[i].StartCol,
				blocks[i].EndLine, blocks[i].EndCol,
				blocks[i].NumStmt, blocks[i].Count)
		}
	}

	return nil
}
