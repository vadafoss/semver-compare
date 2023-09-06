package main

import (
	"fmt"

	"golang.org/x/mod/semver"
)

func main() {
	v1 := "v0.3.0"
	v2 := "v0.4.0"
	c := semver.Compare(v2, v1)
	if c > 1 || c == 1 {
		fmt.Println("v2 is greater than v1")
	} else {
		fmt.Println("v1 is greater than v2")
	}

}
