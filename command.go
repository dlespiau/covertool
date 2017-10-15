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
	"io"
	"os"
)

// WriteProfilesToFile is a higher level WriteProfiles with a few extra features:
//  - write to a file (!)
//  - support '-' for stdout
func WriteProfilesToFile(filename string, profiles []*Profile) error {
	out := io.Writer(os.Stdout)
	if filename != "-" {
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}

	if err := WriteProfiles(out, profiles); err != nil {
		return err
	}

	return nil
}
