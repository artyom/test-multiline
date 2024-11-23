package main

import (
	"flag"
	"fmt"
	"iter"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	if err := run(flag.Args()); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	fmt.Println(groupStart, "Command line arguments")
	for i, s := range args {
		fmt.Printf("%d\t%q\n", i, s)
	}
	fmt.Println(groupEnd)
	fmt.Println(groupStart, "Environment variables")
	for i, s := range os.Environ() {
		fmt.Printf("%d\t%q\n", i, s)
	}
	fmt.Println(groupEnd)
	fmt.Println(groupStart, "Explicit handling of input three")
	for line := range Lines(os.Getenv("INPUT_THREE")) {
		fmt.Printf("raw value: %q\n", line)
		if k, v, ok := strings.Cut(line, "="); ok {
			k = strings.TrimSpace(k)
			v = strings.TrimSpace(v)
			fmt.Printf("as key-value pair: %s: %s\n", k, v)
		}
	}
	fmt.Println(groupEnd)
	return nil
}

const (
	groupStart = "::group::"
	groupEnd   = "::endgroup::"
)

// TODO: replace with strings.Lines once Go 1.24 is out
func Lines(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for len(s) > 0 {
			var line string
			if i := strings.IndexByte(s, '\n'); i >= 0 {
				line, s = s[:i+1], s[i+1:]
			} else {
				line, s = s, ""
			}
			if !yield(line) {
				return
			}
		}
	}
}
