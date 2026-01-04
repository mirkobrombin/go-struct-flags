package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mirkobrombin/go-struct-flags/v2/pkg/structflags"
)

type Options struct {
	Verbose bool   `flag:"short:v, long:verbose, name:Enable Verbose Output"`
	Output  string `flag:"short:o, long:output, default:json"`
}

func main() {
	opts := &Options{}

	// Bind flags to the struct
	structflags.Bind(opts)

	// Parse standard flags
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: example [flags]")
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Verbose: %v\n", opts.Verbose)
	fmt.Printf("Output: %v\n", opts.Output)
}
