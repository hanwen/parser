package lisp

import (
	"bytes"
	"testing"
)

func TestConsPrint(t *testing.T) {
	c := &ConsExpr{
		Car: &SymToken{Value: "x"},
		Cdr: &IntToken{Value: 20},
	}

	got := ExprString(c)
	want := "(x . 20)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDefParse(t *testing.T) {
	ex, err := ParseExpr(bytes.NewBufferString("(define x 'y)"))
	if err != nil {
		t.Fatalf("ParseExpr: %v", err)
	}

	got := ExprString(ex)
	want := "(define x (quote y))"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConsParse(t *testing.T) {
	ex, err := ParseExpr(bytes.NewBufferString("(1 . 2)"))
	if err != nil {
		t.Fatalf("ParseExpr: %v", err)
	}

	got := ExprString(ex)
	want := "(1 . 2)"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestListParse(t *testing.T) {
	ex, err := ParseExpr(bytes.NewBufferString("(1 2 3 . (4 5))"))
	if err != nil {
		t.Fatalf("ParseExpr: %v", err)
	}

	got := ExprString(ex)
	want := "(1 2 3 4 5)"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
