package main

import (
	"flag"
	"log"
	"io/ioutil"
	"bytes"
	"os"
	"path/filepath"

	"github.com/hanwen/parser/java"
)

func main() {
	debug := flag.Bool("debug", false, "dump debug info");
	syms := flag.Bool("syms", false, "dump symbols found");

	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("must supply argument")
	}

	exitCode := 0
	for _, a := range flag.Args() {
		c, err := ioutil.ReadFile(a)
		if err != nil {
			log.Printf("%s: %v\n", a, err)
			exitCode = 1
			continue
		}

		node, err := java.Parse(bytes.NewBuffer(c), *debug)
		log.Println("parse", a, err)
		if *syms && node != nil {
			log.Printf("%s: %s", filepath.Base(a), java.Symbols(node))
		}
	}
	os.Exit(exitCode)
}
