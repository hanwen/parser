package main

import (
	"flag"
	"log"
	"io/ioutil"
	"bytes"
	"os"

	"github.com/hanwen/parser/java"
)

func main() {
	debug := flag.Bool("debug", false, "dump debug info")
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

		_, err = java.Parse(bytes.NewBuffer(c), *debug)
		log.Printf("Parse(%s): %v\n", a, err)
	}
	os.Exit(exitCode)
}
