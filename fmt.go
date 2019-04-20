// yamlfmt formats YAML. It reads from standard input or any files given on the command line.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	strict = flag.Bool("strict", false, "be strict when parsing the YAML")
)

func main() {
	flag.Parse()
	decoders := []*yaml.Decoder{}

	if flag.NArg() == 0 {
		decoders = append(decoders, yaml.NewDecoder(os.Stdin))
	}
	for _, f := range flag.Args() {
		r, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		decoders = append(decoders, yaml.NewDecoder(r))
	}

	for _, d := range decoders {
		in := yaml.Node{}
		err := d.Decode(&in)
		for err == nil {
			e := yaml.NewEncoder(os.Stdout)
			if err := e.Encode(&in); err != nil {
				log.Fatal(err)
			}
			e.Close()

			if err = d.Decode(&in); err == nil {
				fmt.Println("---")
			}
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
}
