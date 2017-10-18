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
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fileMain = "github.com/dlespiau/covertool/examples/calc/main.go"
)

func TestMatches(t *testing.T) {
	tests := []struct {
		pattern, filename string
		expected          bool
	}{
		{"main.go", fileMain, false},
		{"*.go", fileMain, true},
		{"ma*.go", fileMain, false},
		{"*/main.*", fileMain, true},
	}

	for _, test := range tests {
		t.Log(test)
		got, err := matches(test.pattern, test.filename)
		assert.Equal(t, test.expected, got)
		if test.expected {
			assert.Nil(t, err)
			continue
		}
		assert.NotNil(t, err)
	}
}
