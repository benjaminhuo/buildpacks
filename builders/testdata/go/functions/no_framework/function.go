// Copyright 2020 Google LLC
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

// Package function contains a test function.
package function

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

// Func is a test function.
func Func(w http.ResponseWriter, r *http.Request) {
	if quote.Hello() == "Hello, world." {
		fmt.Fprintf(w, "PASS")
	} else {
		fmt.Fprintln(w, "FAIL")
	}
}