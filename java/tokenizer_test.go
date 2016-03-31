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
	"text/scanner"
)

var testSource = `
package com.google.gitiles;

import com.google.common.base.Function;

import java.util.Map;

/** Renderer that precompiles Soy and uses static precompiled CSS. */
public class DefaultRenderer extends Renderer {
  private final SoyTofu tofu;
  DefaultRenderer() {
    this("", ImmutableList.<URL> of(), "");
  }

  public DefaultRenderer(String staticPrefix, Iterable<URL> customTemplates, String siteTitle) {
    this(ImmutableMap.<String, String> of(), staticPrefix, customTemplates, siteTitle);
  }

  public DefaultRenderer(Map<String, String> globals, String staticPrefix,
      Iterable<URL> customTemplates, String siteTitle) {
    super(
        new Function<String, URL>() {
          @Override
          public URL apply(String name) {
            return Resources.getResource(Renderer.class, "templates/" + name);
          }
        },
        globals, staticPrefix, customTemplates, siteTitle);
    SoyFileSet.Builder builder = SoyFileSet.builder()
        .setCompileTimeGlobals(this.globals);
    for (URL template : templates.values()) {
      builder.add(template);
    }
    tofu = builder.build().compileToTofu();
    int x;
    x += (1 << 2);
  }

  @Override
  protected SoyTofu getTofu() {
    return tofu;
  }
}
`

type tokData struct {
	code int
	text string
}

func (d tokData) String() string {
	return tokenString(d.code, d.text)
}

func tokenCodes(in string) []tokData {
	var res []tokData
	tok := newTokenizer(bytes.NewBufferString(in))

	for {
		var sym yySymType
		r := tok.Lex(&sym)
		if r == scanner.EOF {
			break
		}

		res = append(res, tokData{r, sym.text})
	}
	return res
}

type tokCase struct {
	in   string
	want []tokData
}

func TestTokenizerCompound(t *testing.T) {
	cases := []tokCase{
		{
			in: "1 << 2 < 3",
			want: []tokData{
				{NUM, "1"},
				{LESS_LESS, "<<"},
				{NUM, "2"},
				{'<', "<"},
				{NUM, "3"},
			},
		},
		{
			in: "1L",
			want: []tokData{
				{NUM, "1L"},
			},
		},
	}

	for _, c := range cases {
		got := tokenCodes(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: got %s, want %s", c.in, got, c.want)
		}
	}
}
