// Copyright (c) 2017 Intel Corporation
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

import "os"

var atexitFuncs []func()

// atexit registers a function f that will be run when exit is called.
func atexit(f func()) {
	atexitFuncs = append(atexitFuncs, f)
}

// exit calls all atexit handlers before exiting the process with status.
func exit(status int) {
	for _, f := range atexitFuncs {
		f()
	}
	os.Exit(status)
}
