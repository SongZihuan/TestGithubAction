// Copyright 2025 TestGithubAction Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	fmt.Printf("1 + 1 = %d\n", Add(1, 1))
	fmt.Printf("3 - 2 = %d\n", Add(3, 2))
}

func Add(a ...int) (res int) {
	for _, r := range a {
		res += r
	}
	return res
}

func Sub(a int, b ...int) int {
	for _, r := range b {
		a -= r
	}

	return a
}
