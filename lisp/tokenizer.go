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
	"bufio"
	"fmt"
	"io"
)

type byteReader struct {
	r      io.Reader
	unread bool
	last   byte
	idx    int
}

func newByteReader(r io.Reader) *byteReader {
	return &byteReader{r: r}
}

func (b *byteReader) Loc() int {
	return b.idx
}

func (b *byteReader) UnreadByte() {
	b.unread = true
	b.idx--
}

func (b *byteReader) ReadByte() (byte, error) {
	if b.unread {
		b.unread = false
		b.idx++
		return b.last, nil
	}

	var c [1]byte
	n, err := b.r.Read(c[:])
	b.last = c[0]
	b.idx += n
	return b.last, err
}

type tokenizer struct {
	r *byteReader
}

func newTokenizer(r io.Reader) *tokenizer {
	return &tokenizer{newByteReader(bufio.NewReaderSize(r, 8192))}
}

func (t *tokenizer) skipComment() {
	for {
		c, err := t.r.ReadByte()
		if err != nil || c == '\n' {
			return
		}
	}
}

type token struct {
	t   byte
	val []byte
	loc int
}

func (t *token) Loc() int {
	return t.loc
}

const (
	_SYMBOL    = iota
	_NUM_INT   = iota
	_NUM_FLOAT = iota
	_STRING    = iota
	_EOF       = iota
	_ERROR     = iota
)

func isRange(c, a, b byte) bool {
	return c >= a && c <= b
}

func isDigit(c byte) bool {
	return isRange(c, '0', '9')
}

func isLetter(c byte) bool {
	return isRange(c, 'a', 'z') || isRange(c, 'A', 'Z') || c == '_'
}

func isSpace(c byte) bool {
	switch c {
	case ' ', '\n', '\f', '\t', '\r':
		return true
	}

	return false

}

func (t *tokenizer) next() *token {
	tok := t.inext()
	return tok
}

func (t *tokenizer) inext() *token {
	for {
		loc := t.r.Loc()
		c, err := t.r.ReadByte()
		if err != nil {
			return &token{t: _EOF, loc: loc}
		}
		if c == ';' {
			t.skipComment()
			continue
		}
		if isSpace(c) {
			continue
		}

		var tok *token
		if c == '"' {
			tok = t.lexString(loc)
			return tok
		}

		switch c {
		case '(', ')', '\'', '`', ',':
			return &token{t: c, loc: loc}
		}

		return t.lexSymNum(c, loc)
	}
}

// numType returns numtype if this is a numeric constant. Doesn't handle
// complex numbers.
func (t *token) numType() (r byte) {
	if len(t.val) == 1 && t.val[0] == '.' {
		return '.'
	}

	r = _SYMBOL

	var numType byte
	numType = _NUM_INT
	seenExp := false
	seenDot := false
	seenMantissa := false
	for i, c := range t.val {
		switch c {
		case '.':
			if seenDot || seenExp {
				return
			}
			seenDot = true
			numType = _NUM_FLOAT
		case '-':
			if i > 0 {
				return
			}
		case 'e':
			if !seenMantissa {
				return
			}
			if seenExp {
				return
			}
			seenExp = true
			numType = _NUM_FLOAT
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			seenMantissa = true
		default:
			return
		}
	}
	if !seenMantissa {
		return
	}

	return numType
}

func (t *token) Type() int {
	return int(t.t)
}

func (t *token) String() string {
	k := ""
	switch t.t {
	case _SYMBOL:
		k = "sym"
	case _NUM_INT:
		k = "int"
	case _NUM_FLOAT:
		k = "flt"
	case _STRING:
		k = "str"
	case _EOF:
		k = "eof"
	case _ERROR:
		k = "error"
	default:
		k = fmt.Sprintf("'%c'", t.t)
	}
	if t.val != nil {
		k += fmt.Sprintf(":%q", t.val)
	}
	return k
}

func (t *tokenizer) lexString(loc int) *token {
	tok := &token{
		t:   _STRING,
		loc: loc,
	}

	for {
		c, err := t.r.ReadByte()
		if err != nil {
			return &token{t: _ERROR}
		}

		switch c {
		case '\\':
			c, err = t.r.ReadByte()
			if err != nil {
				return &token{t: _ERROR, loc: loc}
			}

			c = unescape(c)
		case '"':
			return tok
		}

		tok.val = append(tok.val, c)
	}
}

func unescape(c byte) byte {
	switch c {
	case 'a':
		return '\a'
	case 'b':
		return '\b'
	case 'f':
		return '\f'
	case 'n':
		return '\n'
	case 'r':
		return '\r'
	case 't':
		return '\t'
	case 'v':
		return '\v'
	}
	return c
}

func (t *tokenizer) lexSymNum(c byte, loc int) *token {
	tok := t.lexSymbol(c, loc)
	tok.t = tok.numType()
	return tok
}

func (t *tokenizer) lexSymbol(c byte, loc int) *token {
	tok := &token{
		val: []byte{c},
		t:   _SYMBOL,
		loc: loc,
	}

	for {
		c, err := t.r.ReadByte()
		if err != nil {
			return tok
		}

		if c == '\\' {
			// Emacs lisp supports this, but GUILE doesn't.
			c, err = t.r.ReadByte()
			if err != nil {
 				return &token{t: _ERROR}
			}
		} else if c == ';' || c == '(' || c == ')' {
			t.r.UnreadByte()
			return tok
		} else if isSpace(c) {
			return tok
		}

		tok.val = append(tok.val, c)
	}
}
