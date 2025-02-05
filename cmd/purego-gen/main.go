package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	puregogen "github.com/Zyko0/purego-gen"
)

var (
	inputfile    string
	extrafiles   string
	embedLoaders bool
	platforms    string

	dry        bool
	nowarnings bool
)

/* Directives
puregogen:library path=opengl alias=gl path:windows=opengl32.dll
puregogen:function symbol=clCreateContext
*/

func main() {
	flag.StringVar(&inputfile, "input", "", "The input .go file to parse")
	flag.StringVar(&extrafiles, "extra", "", "The extra input .go files to parse for definitions, as a comma-separated list")
	flag.BoolVar(&embedLoaders, "embed-loaders", false, "Generate a single file by linking unexported loading methods from ebitengine/purego")
	flag.StringVar(&platforms, "platforms", "windows,darwin,linux,freebsd", "A list of comma separated platforms supported for library loading")
	flag.BoolVar(&dry, "dry", false, "Outputs the generated code to stdout instead of a file")
	flag.BoolVar(&nowarnings, "no-warnings", false, "Prevent printing warnings to sderr")
	flag.Parse()

	if inputfile == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "\nerror: missing input file")
		os.Exit(1)
	}

	var p *puregogen.Parser
	if extrafiles != "" {
		p = puregogen.NewParser(inputfile, strings.Split(extrafiles, ",")...)
	} else {
		p = puregogen.NewParser(inputfile)
	}
	g, err := p.Parse()
	// Warnings printed by default
	if !nowarnings {
		for _, w := range p.Warnings() {
			fmt.Fprintln(os.Stderr, w)
		}
	}
	// Display errors and abort
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		for _, err := range p.Errors() {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
	// Generate output file
	files, err := g.Generate(embedLoaders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		for _, err := range g.Errors() {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
	// If dry mode enabled, print the generated code to stdout
	if dry {
		for _, f := range files {
			fmt.Println("file:", f.Filename)
			fmt.Println(f.Content)
		}
		return
	}
	// By default write the generated code to a file named after the go:generate filename + "_impl.go" suffix
	for _, f := range files {
		err = os.WriteFile(f.Filename, []byte(f.Content), 0644)
		if err != nil {
			log.Fatalf("error: couldn't right to file '%s': %v", f.Filename, err)
		}
	}
}
