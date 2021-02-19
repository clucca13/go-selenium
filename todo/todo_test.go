// Copyright 2014 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"testing"

	"sourcegraph.com/sourcegraph/go-selenium"
)

var caps selenium.Capabilities = selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
var executorURL = "http://selenium:4444/wd/hub"

func TestTodo(t *testing.T) {
	wd, err := selenium.NewRemote(caps, executorURL)
	if err != nil {
		t.Error(err)
		return
	}
	wdt := wd.T(t)
	wdt.Get("http://test:8080")

	title := wdt.Title()
	if title != "TODO" {
		t.Errorf("Wanted title TODO, got %s", title)
	}
}
