package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/n1c00o/go.nc0.fr/pkg/tmpl"

	"github.com/n1c00o/go.nc0.fr/pkg/config"

	"gopkg.in/yaml.v3"
)

var (
	inputFile = flag.String("input", "", "The input configuration file (YAML).")
	outputDir = flag.String("output", "", "The output directory.")
	verbose   = flag.Bool("verbose", false, "On true, prints debug messages.")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("gen: ")

	if *inputFile == "" {
		log.Fatalln("Missing input configuration file.")
	}
	if *outputDir == "" {
		log.Fatalln("Missing output directory.")
	}

	// Read the input configuration file and parse it.
	b, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v\n", *inputFile, err)
	}

	cfg := new(config.Config)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		log.Fatalf("Failed to parse configuration file: %v\n", err)
	}

	// Create output directory and files
	// The index file ($OUTPUT/index.html) should contain an index of the
	// modules listed in cfg.modules.
	// Modules from cfg.modules should have a specific file
	// ($OUTPUT/$MODULE/index.html) with the required meta tags.
	if err := os.Mkdir(*outputDir, 0777); err != nil {
		log.Fatalf("Unable to create output directory %s: %v\n", *outputDir, err)
	}
	if *verbose {
		log.Printf("Created output directory %s.\n", *outputDir)
	}

	// Create a list of all modules' prefixes with the hostname.
	var pfs []string
	for _, module := range cfg.Modules {
		pfs = append(pfs, fmt.Sprintf("%s/%s", cfg.Hostname, module.Prefix))
	}

	createFile(
		fmt.Sprintf("%s/index.html", *outputDir),
		tmpl.IndexTmpl,
		struct {
			Hostname string
			Prefixes []string
		}{
			Hostname: cfg.Hostname,
			Prefixes: pfs,
		},
	)
	if *verbose {
		log.Printf("Created index file at %s/index.html.\n", *outputDir)
	}

	for _, module := range cfg.Modules {
		if err := os.Mkdir(fmt.Sprintf("%s/%s", *outputDir, module.Prefix), 0777); err != nil {
			log.Fatalf("Unable to create directory %s/%s: %v\n", *outputDir, module.Prefix, err)
		}

		createFile(
			fmt.Sprintf("%s/%s/index.html", *outputDir, module.Prefix),
			tmpl.ModuleTmpl,
			struct {
				Hostname string
				Prefix   string
				Repo     string
				VCS      string
				Dir      string
				File     string
			}{
				Hostname: cfg.Hostname,
				Prefix:   fmt.Sprintf("%s/%s", cfg.Hostname, module.Prefix),
				Repo:     module.Repo,
				VCS:      module.VCS,
				Dir:      module.Dir,
				File:     module.File,
			},
		)

		if *verbose {
			log.Printf("Created module file at %s/%s/index.html.\n", *outputDir, module.Prefix)
		}
	}
}

// createFile creates a file and write the template inside it.
func createFile(path string, tpl *template.Template, vars interface{}) {
	fl, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to create file %s: %v\n", path, err)
	}

	/*fl, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file %s: %v\n", path, err)
	}*/
	defer func(fl *os.File) {
		err := fl.Close()
		if err != nil {
			log.Fatalf("Cannot close file %s: %v\n", path, err)
		}
	}(fl)

	if err := tpl.Execute(fl, vars); err != nil {
		log.Fatalf("Cannot execute template %s on file %s: %v\n", tpl.Name(), path, err)
	}
}
