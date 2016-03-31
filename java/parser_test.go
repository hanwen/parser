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
	"bytes"
	"reflect"
	"testing"
)

type testCase struct {
	name  string
	in    string
	debug bool
	full bool
}

func TestParse(t *testing.T) {
	cases := []testCase{
		{
			name: "nest",
			in: `
  public TimeCache() {
    f();
  }
`,
		},
		{
			name: "return inside",
			in:   `GitilesServlet servlet = new A(new B(), new C());`,
		},
		{
			name: "generic method",
			in:   `List<URL> x = ImmutableList.<URL>of();`,
		},
		{
			name: "switch",
			in: `void f() {
    switch (format.get()) {
      case HTML:
	break;
      default:
        break;
    }
}`,
		},
		{
			name: "else",
			in: `void f() { if (true <= MAX_SYMLINK_TARGET_LENGTH) {
      return target;
    } else {
      return 1;
} }`,
		},
		{
			name: "final var",
			in:   `void f() { final String x = "1"; }`,
		},
		{
			name: "multithrow",
			in:   `abstract void f() throws A, B;`,
		},
		{
			name: "negative int",
			in:   `int x = -1;`,
		},
		{
			name: "return inside",
			in: `GitilesServlet servlet =
        new GitilesServlet(new Config(), new DefaultRenderer(GitilesServlet.STATIC_PREFIX,
              ImmutableList.<URL> of(), repoName + " test site"),
            TestGitilesUrls.URLS, new TestGitilesAccess(repo.getRepository()),
            new RepositoryResolver<HttpServletRequest>() {
              @Override
              public Repository open(HttpServletRequest req, String name)
                  throws RepositoryNotFoundException {
                if (!repoName.equals(name)) {
                  throw new RepositoryNotFoundException(name);
                }
                return repo.getRepository();
              }
            }, null, null, null, gitwebRedirect);
`,
		},
		{
			name: "override",
			in:   `Object in = new Object() { void foo() { } };`,
		},
		{
			name: "if",
			in:   `void f() { if (1) { } }`,
		},
		{
			name: "var init",
			in: `
  int variable = 2;
`,
		},
		{
			name: "decl/assignment",
			in: `
  Map<A> m = new Map<B>();
  void func() {
    long j = 2;
    k = 3;
  }
`,
		},
		{
			name: "assign to expr",
			in:   `void f() { this.y = 42; }`,
		},
		{
			name: "template",
			in: `
	 public static CacheBuilder<Object, Object> defaultBuilder() {
    return CacheBuilder.newBuilder().maximumSize(10 * 10);
  }
  Long getTime(final RevWalk walk, final ObjectId id) throws IOException {
  }
`,
		},
		{
			name: "try catch",
			in: `
  void x() {
    try {
      x = 1;
    } catch (Exception | A e) {
    } finally {
    }
  }
`,
		},
		{
			name: "funcall",
			in: `void f() {
  g();
}`,
		},
		{
			name: "method",
			in:   `int func(byte arg1) throws Exception { }`,
		},
		{
			name: "member",
			in: `
  int x;
  public byte []bla;
`,
		},
		{
			name: "enum",
			full: true,
		in: `package foo;
enum P {
 A(1), B;
 void f() {}
}`,
		},
		{
			name: "nested class",
			in: `static class Inner { } int outer = 1;`,
		},{
			name: "...",
			in: `static ImmutableSet<Field> setOf(Iterable<Field> base, Field... fields) {
      return Sets.immutableEnumSet(Iterables.concat(base, Arrays.asList(fields)));
    }`,
		},
	}

	for _, c := range cases {
		full := c.in
		if !c.full {
			full = "package pkg; class Klass {" + c.in + "}"
		}

		_, err := Parse(bytes.NewBufferString(full), c.debug)
		if err != nil {
			t.Errorf("%s: yyParse failed", c.name)
		}
	}
}

func TestDecls(t *testing.T) {
	in := `package foo; class Klass {
  int x = 1;
  Object j() { return null; }
}`
	got, err := Parse(bytes.NewBufferString(in), false)
	if err != nil {
		t.Errorf("parse: %v", err)
	}

	want := &ClassDecl{
		Decl: Decl{
			Name: "Klass",
			Type: "class",
		},
		Members: []Node{
			&Decl{Name: "x", Type: "int"},
			&Decl{Name: "j", Type: "Object"},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}

	wantSyms := []string{"Klass:class", "x:int", "j:Object"}
	gotSyms := Symbols(got)
	if !reflect.DeepEqual(gotSyms, wantSyms) {
		t.Errorf("got syms %v, want %v", gotSyms, wantSyms)
	}
}
