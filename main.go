package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/mod/semver"
)

var version = "v0.4.0"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("specify exactly one argument/version")
	}

	input := os.Args[1]

	c := semver.Compare(input, version)
	if c > 0 {
		fmt.Printf("%s is a higher version than %s\n", input, version)
	} else if c == 0 {
		fmt.Printf("%s is the same version as %s\n", input, version)
	} else {
		fmt.Printf("%s is a lower version than %s\n", input, version)
	}

}
