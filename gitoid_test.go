// Copyright (c) 2022 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gitoid_test

import (
	"crypto/sha1" // #nosec
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/edwarnicke/gitoid"
)

func Test_gitoid_sha1(t *testing.T) {
	t.Parallel()

	filename := "LICENSE"
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Fatalf("error opening %s: %s", filename, err)
	}

	gitoidHash := gitoid.New(gitoid.BLOB, len(input), sha1.New()) // #nosec

	gitoidHash.Write(input)
	result := fmt.Sprintf("%x", gitoidHash.Sum(nil))
	expected := "261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64"

	if result != expected {
		t.Fatalf("unexpected result.  Actual: %s Expected: %s", result, expected)
	}
}

func Example_gitoid_sha1() {
	filename := "LICENSE"
	input, _ := ioutil.ReadFile(filename)
	gitoidHash := gitoid.New(gitoid.BLOB, len(input), sha1.New()) // #nosec
	gitoidHash.Write(input)
	gitObjectID := gitoidHash.Sum(nil)
	fmt.Printf("%x", gitObjectID)
	// Output: 261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
}
