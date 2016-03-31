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
	"fmt"
	"io"
	"strings"
	"text/scanner"
)

type tokenizer struct {
	*scanner.Scanner
	debug  bool
	errors []string
}

var reservedWords = map[string]int{
	"abstract":     ABSTRACT,
	"case":         CASE,
	"catch":        CATCH,
	"class":        CLASS,
	"default":      DEFAULT,
	"else":         ELSE,
	"extends":      EXTENDS,
	"final":        FINAL,
	"finally":      FINALLY,
	"for":          FOR,
	"if":           IF,
	"implements":   IMPLEMENTS,
	"import":       IMPORT,
	"instanceof":   INSTANCEOF,
	"new":          NEW,
	"package":      PACKAGE,
	"private":      PRIVATE,
	"protected":    PROTECTED,
	"public":       PUBLIC,
	"return":       RETURN,
	"static":       STATIC,
	"super":        SUPER,
	"synchronized": SYNCHRONIZED,
	"switch":       SWITCH,
	"throw":        THROW,
	"throws":       THROWS,
	"try":          TRY,
	"while":        WHILE,
}

var reservedWordStrings = map[int]string{}

var compoundTokens = map[string]int{
	"*=":  STAR_EQUAL,
	"+=":  PLUS_EQUAL,
	"/=":  SLASH_EQUAL,
	"<=":  LESS_EQUAL,
	">=":  MORE_EQUAL,
	"&&":  AMPERSAND_AMPERSAND,
	"&=":  AMPERSAND_EQUAL,
	"||":  BAR_BAR,
	"|=":  BAR_EQUAL,
	"<<":  LESS_LESS,
	"<<=": LESS_LESS_EQUAL,
	">>":  MORE_MORE,
	"!=":  EXCLAMATION_EQUAL,
	"==":  EQUAL_EQUAL,
}

func init() {
	for k, v := range reservedWords {
		reservedWordStrings[v] = k
	}

	for k := range compoundTokens {
		compoundTokens[k[:1]] = int(k[0])
	}
}

func newTokenizer(r io.Reader) *tokenizer {
	t := &tokenizer{
		Scanner: &scanner.Scanner{
			Mode: scanner.GoTokens,
		},
	}
	t.Init(r)
	return t
}

func tokenString(r int, t string) string {
	if r < 256 && r >= 32 {
		return fmt.Sprintf("'%c'", r)
	}

	if n, ok := reservedWordStrings[r]; ok {
		return strings.ToUpper(n)
	}
	switch r {
	case IDENTIFIER:
		return fmt.Sprintf("id:%s", t)
	case NUM:
		return fmt.Sprintf("num:%s", t)
	default:
		return fmt.Sprintf("t_%d:%q", r, t)
	}
}

func (t *tokenizer) Error(e string) {
	t.errors = append(t.errors, e)
}

func (t *tokenizer) Lex(l *yySymType) int {
	r := t.iLex(l)
	if t.debug {
		fmt.Printf("pos %v: %v\n", t.Scanner.Position, tokenString(r, l.Text))
	}
	return r
}

func (t *tokenizer) iLex(l *yySymType) int {
	l.Text = ""
	for {
		r := t.Scanner.Scan()
		l.Text += t.Scanner.TokenText()

		switch r {
		case scanner.Ident:
			if v, ok := reservedWords[l.Text]; ok {
				return v
			}

			return IDENTIFIER
		case scanner.Int:
			p := t.Scanner.Peek()
			if p == 'l' || p == 'L' {
				t.Scanner.Next()
				l.Text += string([]rune{p})
				return NUM
			}

			return NUM
		case scanner.String:
			return STRING
		}

		if curCode, ok := compoundTokens[l.Text]; ok {
			p := t.Scanner.Peek()
			if _, next := compoundTokens[l.Text+string([]rune{p})]; next {
				continue
			}

			return curCode
		}

		return int(r)
	}
}
