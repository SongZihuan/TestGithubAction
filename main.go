// Copyright 2025 TestGithubAction Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	fmt.Printf("1 + 1 = %d\n", Add(1, 1))
}

func Add(a int, b ...int) int {
	for _, r := range b {
		a += r
	}

	return a
}
