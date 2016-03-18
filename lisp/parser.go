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

package lisp

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type BaseToken struct {
	Loc int
}

type IntToken struct {
	BaseToken
	Value int64
}

func (t *IntToken) Print(w io.Writer) {
	fmt.Fprintf(w, "%d", t.Value)
}

type FloatToken struct {
	BaseToken
	Value float64
}

func (t *FloatToken) Print(w io.Writer) {
	fmt.Fprintf(w, "%f", t.Value)
}

type SymToken struct {
	BaseToken
	Value string
}

func (t *SymToken) Print(w io.Writer) {
	w.Write([]byte(t.Value))
}

type StringToken struct {
	BaseToken
	Value string
}

func (t *StringToken) Print(w io.Writer) {
	fmt.Fprintf(w, "%q", t.Value)
}

type ConsExpr struct {
	Car Expr
	Cdr Expr
}

type Expr interface {
	Print(w io.Writer)
}

func ExprString(e Expr) string {
	var b bytes.Buffer
	e.Print(&b)
	return b.String()
}

func (e *ConsExpr) Print(w io.Writer) {
	w.Write([]byte{'('})
	for e != nil {
		if e.Car == nil {
			w.Write([]byte("()"))
		} else {
			e.Car.Print(w)
		}
		if e.Cdr == nil {
			break
		}

		w.Write([]byte{' '})

		next, ok := e.Cdr.(*ConsExpr)
		if !ok {
			w.Write([]byte(". "))
			e.Cdr.Print(w)
			break
		}

		e = next
	}
	w.Write([]byte(")"))
}

// ParseExpr parses a single lisp expression.
func ParseExpr(r io.Reader) (Expr, error) {
	t := newTokenizer(r)
	expr, err := parseExpr(t)
	if err != nil {
		return nil, err
	}

	eof := t.next()
	if eof.Type() != _EOF {
		return nil, fmt.Errorf("trailing junk at position %d: %v", eof.Loc(), eof)
	}

	return expr, err
}

// Parse parses a sequence of LISP expressions.
func Parse(r io.Reader) ([]Expr, error) {
	t := newTokenizer(r)

	var result []Expr
	for {
		expr, err := parseExpr(t)
		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}

		result = append(result, expr)
	}
	return result, nil
}

var errParenClose = fmt.Errorf("end of list")
var errDot = fmt.Errorf("dot")

func parseExpr(t *tokenizer) (Expr, error) {
	n := t.next()
	switch n.Type() {
	case _ERROR:
		return nil, fmt.Errorf("tokenizer error at %d", n.Loc())
	case _EOF:
		return nil, io.EOF
	case _NUM_INT:
		i, err := strconv.ParseInt(string(n.val), 10, 64)
		if err != nil {
			return nil, err
		}
		return &IntToken{
			BaseToken: BaseToken{Loc: n.loc},
			Value:     i,
		}, nil

	case _NUM_FLOAT:
		i, err := strconv.ParseFloat(string(n.val), 64)
		if err != nil {
			return nil, err
		}
		return &FloatToken{
			BaseToken: BaseToken{Loc: n.loc},
			Value:     i,
		}, nil

	case _SYMBOL:
		return &SymToken{
			BaseToken: BaseToken{Loc: n.loc},
			Value:     string(n.val),
		}, nil
	case _STRING:
		return &StringToken{
			BaseToken: BaseToken{Loc: n.loc},
			Value:     string(n.val),
		}, nil
	case '.':
		return nil, errDot

	case '\'', '`', ',':
		ch, err := parseExpr(t)
		if err != nil {
			return nil, err
		}

		sym := map[int]string{
			'\'': "quote",
			'`':  "quasiquote",
			',':  "unquote",
		}[n.Type()]

		return makeList(
			&SymToken{Value: sym},
			ch), nil
	case '(':
		cons, err := parseCons(t)
		return cons, err
	case ')':
		return nil, errParenClose

	default:
		return nil, fmt.Errorf("unexpected token %c", n.Type())
	}

}

func makeList(exprs ...Expr) Expr {
	var head Expr
	var tail *Expr
	tail = &head
	for _, e := range exprs {
		newTail := &ConsExpr{Car: e}
		*tail = newTail
		tail = &newTail.Cdr
	}

	return head
}

func parseCons(t *tokenizer) (Expr, error) {
	var head Expr
	var tail *Expr
	tail = &head

	for {
		car, err := parseExpr(t)
		if err == errParenClose {
			return head, nil
		}

		if err == errDot {
			cdr, err := parseExpr(t)
			if err != nil {
				return nil, err
			}
			*tail = cdr

			parenClose := t.next()
			if parenClose.Type() != ')' {
				return nil, fmt.Errorf("expect ')' after '.' in list, got %v", parenClose)
			}

			return head, nil
		}

		if err != nil {
			return nil, err
		}

		newTail := &ConsExpr{
			Car: car,
		}
		*tail = newTail
		tail = &newTail.Cdr
	}
}
