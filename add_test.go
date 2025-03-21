// Copyright 2025 TestGithubAction Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Errorf("Running test %s error!", t.Name())
	}

	t.Logf("Running test %s success!", t.Name())
}
