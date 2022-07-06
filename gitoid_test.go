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
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/edwarnicke/gitoid"
)

const (
	filename = "LICENSE"
)

func Example_gitoid_sha1() {
	file, _ := os.Open(filename)
	defer file.Close()

	gitoidHash, _ := gitoid.New(file)
	fmt.Println(gitoidHash)
	// Output: 261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
}

func Example_gitoid_uri_sha1() {
	file, _ := os.Open(filename)
	defer file.Close()

	gitoidHash, _ := gitoid.New(file)
	fmt.Println(gitoidHash.URI())
	// Output: gitoid:blob:sha1:261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
}

func Example_gitoid_sha256() {
	file, _ := os.Open(filename)
	defer file.Close()

	gitoidHash, _ := gitoid.New(file, gitoid.WithSha256())
	fmt.Println(gitoidHash)
	// Output: ed43975fbdc3084195eb94723b5f6df44eeeed1cdda7db0c7121edf5d84569ab
}

func Example_gitoid_uri_sha256() {
	file, _ := os.Open(filename)
	defer file.Close()

	gitoidHash, _ := gitoid.New(file, gitoid.WithSha256())
	fmt.Println(gitoidHash.URI())
	// Output: gitoid:blob:sha256:ed43975fbdc3084195eb94723b5f6df44eeeed1cdda7db0c7121edf5d84569ab
}

func Example_gitoid_bytes_sha1() {
	input := []byte("example")
	gitoidHash, _ := gitoid.New(bytes.NewBuffer(input))
	fmt.Println(gitoidHash)
	// Output: 96236f8158b12701d5e75c14fb876c4a0f31b963
}

func Example_gitoid_sha1_content_length() {
	file, _ := os.Open(filename)
	defer file.Close()
	fi, _ := file.Stat()

	gitoidHash, _ := gitoid.New(file, gitoid.WithContentLength(fi.Size()))
	fmt.Println(gitoidHash)
	// Output: 261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
}

func Test_gitoid_sha1_content_length(t *testing.T) {
	t.Parallel()

	file, _ := os.Open(filename)
	defer file.Close()
	fi, _ := file.Stat()

	_, err := gitoid.New(file, gitoid.WithContentLength(fi.Size()+1))
	if err == nil {
		t.Fatalf("expected error specifying contentLength in excess of available bytes.  no error detected.")
	}
}
