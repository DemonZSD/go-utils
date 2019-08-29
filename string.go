// Copyright 2019 Weshzhu
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// author : weshzhu
// the string ops util
package go_utils

func StringIsEmpty(args string) bool {
	return IsNil(args) || len(args) == 0
}

// check every arg is nil or ""
func StringsNotEmpty(args ...string) bool {
	for _, val := range args {
		if StringIsEmpty(val) {
			return false
		}
	}
	return true
}
