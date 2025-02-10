package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	inputfile    string
	extrafiles   string
	openLibrary  bool
	platforms    string
	functionName string

	dry        bool
	nowarnings bool
)

/* Directives
puregogen:library path=opengl alias=gl path:windows=opengl32.dll
puregogen:function symbol=clCreateContext
*/

func main() {
	flag.StringVar(&inputfile, "input", "", "The input .go file to parse")
	flag.StringVar(&extrafiles, "extra", "", "An optional list of comma separated .go files to parse for definitions")
	flag.BoolVar(&openLibrary, "open-library", false, "If true, the generated init() function will handle loading the library from disk")
	flag.StringVar(&platforms, "platforms", "windows,darwin,linux,freebsd", "A list of comma separated platforms supported for library loading")
	flag.StringVar(&functionName, "function-name", "init", "The name given to the function loading all the symbols")
	flag.BoolVar(&dry, "dry", false, "Outputs the generated code to stdout instead of a file")
	flag.BoolVar(&nowarnings, "no-warnings", false, "Prevent printing warnings to sderr")
	flag.Parse()

	if inputfile == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "\nerror: missing input file")
		os.Exit(1)
	}

	var p *Parser
	if extrafiles != "" {
		p = NewParser(inputfile, strings.Split(extrafiles, ",")...)
	} else {
		p = NewParser(inputfile)
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
	files, err := g.Generate(&GenerateOptions{
		OpenLibrary:  openLibrary,
		FunctionName: functionName,
	})
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
