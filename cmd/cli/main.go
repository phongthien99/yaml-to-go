package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/PrasadG193/yaml2go"
)

const helpMsg = `yaml2go converts YAML specs to Go type definitions

Usage:
    yaml2go < /path/to/yamlspec.yaml

Examples:
    yaml2go < test/example1.yaml
    yaml2go < test/example1.yaml > example1.go
`

func printHelp(f string) {
	helpArgs := []string{"-h", "--help", "help"}
	for _, m := range helpArgs {
		if f == m {
			fmt.Println(helpMsg)
			os.Exit(0)
		}
	}
}

func main() {
	inputFilePath := flag.String("i", "", "File path to read")

	outFilePath := flag.String("o", "", "File path to read")
	flag.Parse()

	// Create yaml2go object and invoke Convert()
	y2g := yaml2go.New()
	if *inputFilePath == "" {
		fmt.Println("Please provide a file path using -i  flag")
		return
	}

	// Read the contents of the file
	content, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	result, err := y2g.Convert("EnvironmentVariable", []byte(content))

	result = fmt.Sprintf(`package types %s`, result)
	if err != nil {
		log.Fatal("Invalid YAML")
	}

	fmt.Printf(result)
	err = ioutil.WriteFile(*outFilePath, []byte(result), 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
	return
}
