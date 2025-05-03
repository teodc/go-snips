package main

// The "must" pattern is a common pattern in Go to handle errors and reduce boilerplate code.
// It's also used in the standard library.
// The idea is to wrap the error handling in a function that panics if the error is not nil.
// Best practices:
// - Use when panics are acceptable (e.g., in main(), during initialization)
// - You don't necessarily have to panic, you can also log.Fatal() or use your own "shutdown" function

import (
	"fmt"
	"os"
	"regexp"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Not a big fan of the name "CheckErr" but seems to be the standard in the Go community
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Example 1
	invalidRegex := "["
	r := Must(regexp.Compile(invalidRegex))
	fmt.Println(r)

	// Example 2
	src := "./prout.txt"
	dst := "./out/prout.txt"

	fr := Must(os.Open(src))
	defer func(fr *os.File) {
		CheckErr(fr.Close())
	}(fr)

	fw := Must(os.Create(dst))
	CheckErr(fw.Close())
}
