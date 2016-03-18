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
