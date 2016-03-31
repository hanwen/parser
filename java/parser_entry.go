// Copyright 2016 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package java

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func Parse(r io.Reader, debug bool) (Node, error) {
	tok := newTokenizer(r)
	tok.debug = debug
	if debug {
		yyDebug = 5
	} else {
		yyDebug = 0
	}
	if exitCode := yyParse(tok); exitCode != 0 {
		return nil, fmt.Errorf("yyParse failed")
	}
	if len(tok.errors) > 0 {
		return compilationUnitNode, errors.New(strings.Join(tok.errors, "\n"))
	}
	return compilationUnitNode, nil
}


func Symbols(node Node) []string {
	var syms []string
	switch n := node.(type) {
	case *ClassDecl:
		syms = append(syms, n.Decl.String())
		for _, mem := range n.Members {
			syms = append(syms, Symbols(mem)...)
		}
	case *Decl:
		syms = append(syms, n.String())
	}
	return syms
}
