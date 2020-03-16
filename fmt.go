// yamlfmt formats YAML. It reads from standard input or any files given on the command line.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	overwrite := flag.Bool("w", false, "overwrite the input file")
	indent := flag.Int("indent", 2, "default indent")
	flag.Parse()

	for _, f := range flag.Args() {
		formatFile(f, *indent, *overwrite)
	}
}

func formatFile(f string, indent int, overwrite bool) {
	r, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	if e := formatStream(r, &out, indent); e != nil {
		log.Fatalf("Failed formatting YAML stream: %v", e)
	}

	r.Close()

	if e := dumpStream(&out, f, overwrite); e != nil {
		log.Fatalf("Cannot overwrite: %v", e)
	}
}

func formatStream(r io.Reader, out io.Writer, indent int) error {
	d := yaml.NewDecoder(r)
	in := yaml.Node{}
	err := d.Decode(&in)
	for err == nil {
		e := yaml.NewEncoder(out)
		e.SetIndent(indent)
		if err := e.Encode(&in); err != nil {
			log.Fatal(err)
		}
		e.Close()

		if err = d.Decode(&in); err == nil {
			fmt.Fprintln(out, "---\n")
		}
	}

	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func dumpStream(out *bytes.Buffer, f string, overwrite bool) error {
	if overwrite {
		return ioutil.WriteFile(f, out.Bytes(), 0744)
	}
	_, err := io.Copy(os.Stdout, out)
	return err
}
