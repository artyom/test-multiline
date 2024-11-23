package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	return nil
}

const (
	groupStart = "::group::"
	groupEnd   = "::endgroup::"
)
