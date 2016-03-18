package lisp

import (
	"bytes"
	"reflect"
	"testing"
)

func tokenize(in string) []*token {
	tokenizer := newTokenizer(bytes.NewBufferString(in))
	var got []*token
	for {
		t := tokenizer.next()
		got = append(got, t)
		if t.Type() == _EOF || t.Type() == _ERROR {
			break
		}
	}
	return got
}

func stripLocations(in []*token) {
	for _, t := range in {
		t.loc = 0
	}
}

type testCase struct {
	name string
	in   string
	want []*token
}

func TestTokenizer(t *testing.T) {
	cases := []testCase{
		{
			name: "float-sym",
			in: `.01`,
			want: []*token{
				{t: _NUM_FLOAT, val: []byte(".01")},
				{t: _EOF},
			},
		},
		{
			name: "comment",
			in: `(define ;; bla
x y)`, want: []*token{
				{t: '('},
				{t: _SYMBOL, val: []byte("define")},
				{t: _SYMBOL, val: []byte("x")},
				{t: _SYMBOL, val: []byte("y")},
				{t: ')'},
				{t: _EOF},
			}}, {
			name: "basic",
			in:   `(define -1..0 (+ 2 "lit"))`,
			want: []*token{
				{t: '('},
				{t: _SYMBOL, val: []byte("define")},
				{t: _SYMBOL, val: []byte("-1..0")},
				{t: '('},
				{t: _SYMBOL, val: []byte("+")},
				{t: _NUM_INT, val: []byte("2")},
				{t: _STRING, val: []byte("lit")},
				{t: ')'},
				{t: ')'},
				{t: _EOF},
			}},
		{
			name: "escape-sym",
			in:   `(= ?\) 2)`,
			want: []*token{
				{t: '('},
				{t: _SYMBOL, val: []byte("=")},
				{t: _SYMBOL, val: []byte("?)")},
				{t: _NUM_INT, val: []byte("2")},
				{t: ')'},
				{t: _EOF},
			},
		},
	}

	for _, c := range cases {
		got := tokenize(c.in)
		stripLocations(got)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("case %s: got %v, want %v", c.name, got, c.want)
		}
	}
}
