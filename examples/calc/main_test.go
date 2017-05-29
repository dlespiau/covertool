//
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
//

package main

import (
	"flag"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/dlespiau/covertool/pkg/exit"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{0, 0, 0},
		{1, 1, 2},
		{-2, 1, -1},
		{-1, -10, -11},
	}

	for i := range tests {
		test := &tests[i]
		got := add(test.a, test.b)
		if test.expected != got {
			t.Fatalf("expected %d but got %d", test.expected, got)
		}
	}
}

type dummyTestDeps struct{}

func (d dummyTestDeps) MatchString(pat, str string) (bool, error)   { return false, nil }
func (d dummyTestDeps) StartCPUProfile(io.Writer) error             { return nil }
func (d dummyTestDeps) StopCPUProfile()                             {}
func (d dummyTestDeps) WriteHeapProfile(io.Writer) error            { return nil }
func (d dummyTestDeps) WriteProfileTo(string, io.Writer, int) error { return nil }

// FlushProfiles flushes test profiles to disk. It works by build and executing
// a dummy list of 1 test. This is to ensure we execute the M.after() function
// (a function internal to the testing package) where all reporting (cpu, mem,
// cover, ... profiles) is flushed to disk.
func FlushProfiles() {
	// We define a dummy test function so we don't get the "there's no test
	// defined!" warning from the testing package.
	tests := []testing.InternalTest{
		testing.InternalTest{"TestDummy", func(*testing.T) {}},
	}
	benchmarks := []testing.InternalBenchmark{}
	examples := []testing.InternalExample{}
	dummyM := testing.MainStart(dummyTestDeps{}, tests, benchmarks, examples)
	dummyM.Run()
}

func TestMain(m *testing.M) {
	// Parse the command line using the stdlib flag package so the flags defined
	// in the testing package get populated.
	flag.Parse()

	// Strip command line arguments that were for the testing package and are now
	// handled.
	var runtimeArgs []string
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") ||
			strings.HasPrefix(arg, "-httptest.") {
			continue
		}
		runtimeArgs = append(runtimeArgs, arg)
	}
	os.Args = runtimeArgs

	// Make sure we have the opportunity to flush the coverage report to disk when
	// terminating the process.
	exit.Atexit(FlushProfiles)

	// If the test binary name is cc-runtime.coverage, we've are being asked to
	// run the coverage-instrumented cc-runtime.
	if path.Base(os.Args[0]) == "calc" {
		main()
		exit.Exit(0)
	}

	os.Exit(m.Run())
}
