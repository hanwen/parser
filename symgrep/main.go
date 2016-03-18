package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/hanwen/parser/lisp"
)

func toSym(e lisp.Expr) string {
	if e == nil {
		return ""
	}
	if c, ok := e.(*lisp.SymToken); ok {
		return c.Value
	}
	return ""
}

func car(e lisp.Expr) lisp.Expr {
	return e.(*lisp.ConsExpr).Car
}
func cdr(e lisp.Expr) lisp.Expr {
	return e.(*lisp.ConsExpr).Cdr
}

func consp(e lisp.Expr) bool {
	_, ok := e.(*lisp.ConsExpr)
	return ok
}

// return byte indices of matching symbols.
func grepSymbols(exprs []lisp.Expr, re *regexp.Regexp) []int {
	var indices []int
	for _, e := range exprs {
		if !consp(e) || !consp(cdr(e)) {
			continue
		}

		switch toSym(car(e)) {
		case "defvar", "defalias", "defun":
			name := toSym(car(cdr(e)))
			if re.FindString(name) != "" {
				i := car(cdr(e)).(*lisp.SymToken).Loc
				indices = append(indices, i)
			}
		}
	}
	return indices
}

func main() {
	flag.Parse()

	re, err := regexp.Compile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range flag.Args()[1:] {
		content, err := ioutil.ReadFile(a)
		if err != nil {
			log.Printf("Open: %v", err)
			continue
		}

		exprs, err := lisp.Parse(bytes.NewBuffer(content))
		if err != nil {
			log.Fatalf("Parse(%s): %v last %v", a, err,
				lisp.ExprString(exprs[len(exprs)-1]))
		}

		for _, l := range grepSymbols(exprs, re) {
			before := content[:l]
			after := content[l:]
			n := bytes.Count(before, []byte{'\n'})
			start := bytes.LastIndexByte(before, '\n')
			end := bytes.IndexByte(after, '\n')

			fmt.Printf("%s:%d: %s%s\n", a, n,
				string(before[start+1:]),
				string(after[:end]))

		}
	}
}
