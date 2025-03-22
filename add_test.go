// Copyright 2025 TestGithubAction Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestAdd(t *testing.T) {
	test1 := t.Run("1+1", func(t *testing.T) {
		if Add(1, 1) != 2 {
			t.Fatalf("Running test %s error!", t.Name())
		}
	})

	test2 := t.Run("1+2", func(t *testing.T) {
		if Add(1, 2) != 3 {
			t.Fatalf("Running test %s error!", t.Name())
		}
	})

	if !test1 || !test2 {
		t.Fail()
	}

	if t.Failed() {
		t.Logf("Run test %s failed.", t.Name())
	} else {
		t.Logf("Run test %s success.", t.Name())
	}

}
