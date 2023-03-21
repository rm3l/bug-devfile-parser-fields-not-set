package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/rm3l/devfile-lib-issue-parser-default-values-flattened/pkg/devfile"
	"k8s.io/utils/pointer"
)

func main() {
	inputPathFlag := flag.String("input", "./devfile.yaml", "path to the Devfile YAML file")
	flattenFlag := flag.String("flatten", "", "whether to flatten the Devfile parsed or not. Default is true")
	flag.Parse()

	p := *inputPathFlag
	fStr := *flattenFlag
	log.Printf("Parsing Devfile at %q with flatten=%v\n", p, fStr)

	var f *bool
	if fStr != "" {
		b, err := strconv.ParseBool(fStr)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		f = pointer.Bool(b)
	}

	d, err := devfile.Parse(p, f)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	b, err := yaml.Marshal(d.Data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- %s:\n\n%s\n", p, string(b))
}
