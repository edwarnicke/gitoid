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

package gitoid

import (
	"fmt"
	"hash"
)

// GitObjectType type of git object - current values are "blob", "commit", "tag", "tree".
type GitObjectType string

const (
	BLOB   GitObjectType = "blob"
	COMMIT GitObjectType = "commit"
	TAG    GitObjectType = "tag"
	TREE   GitObjectType = "tree"
)

// New - new GitOID hash.Hash.  Writes the git object header to the provided hash and then returns that Hash.
//       From there the contents themselves can be written to the hash.Hash and the Hash value itself can be
//       computed using hash.Hash.Sum()
func New(gitObjectType GitObjectType, contentLength int, h hash.Hash) hash.Hash {
	h.Write(Header(gitObjectType, contentLength))

	return h
}

// Header - returns the git object header from the gitObjectType and contentLength.
func Header(gitObjectType GitObjectType, contentLength int) []byte {
	return []byte(fmt.Sprintf("%s %d\000", gitObjectType, contentLength))
}
