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

/*

A parser for Java. It is incomplete, and can fundamentally not work in
its current form. The biggest problem is that we can't tell from
identifiers whether they refer to types or objects: does "a.b" mean
"varA.memberB" or "packageA.typeB"?

This has the following repercussions:

 * a < b > c

  can be "typeA<typeB> c" or (varA < varB) > varC.

 * (a)(b)

  could be (funcA)(varB) or (typeA)(varB)

 * Since a.b can be either a type or variable (expression), a
   simplistic grammar such as

   expression:
      identifier
      ;
   type_specifier:
      identifier
      ;
   statement:
      expression ';'
      | type_specifier identifier ';'
      ;

   is ambiguous: with a single token lookahead, you can't decide which
   production to pick if you encounterf an identifier trying to parse
   a statement. I tried to work around this by having "type_specifier
   identifier" be an expression, but this makes type casts problematic
   (see previous point.)

The parser in org.eclipse.jdt.core seems to handle this with lots of
extra rules to disambiguate certain constructions, and from inside the
lexer a separate parser is started to disambiguate some tokens. In
particular, it tries a parse to emit TokenNameBeginTypeArguments if is
at a '<' token. (see
compiler/org/eclipse/jdt/internal/compiler/parser/Scanner.java), to
avoid the < (less than) vs < (generic type argument) ambiguity.

*/

%{

package java

%}

%union {
	text string
	node Node
        classDecl *ClassDecl
}

%type <classDecl> class_decl class_body annotated_class_decl
%type <text> identifier
%type <node> class_element_decl
%type <text> type_specifier

%token ABSTRACT
%token BREAK
%token CASE
%token CATCH
%token CLASS
%token DEFAULT
%token ELSE
%token ENUM
%token EXTENDS
%token FINAL
%token FINALLY
%token FOR
%token IF
%token IMPLEMENTS
%token IMPORT
%token INSTANCEOF
%token NEW
%token PACKAGE
%token PRIVATE
%token PROTECTED
%token PUBLIC
%token RETURN
%token STATIC
%token SUPER
%token SYNCHRONIZED
%token SWITCH
%token THROW
%token THROWS
%token TRY
%token WHILE

%token DOT_DOT_DOT DOT_DOT

%token <Text> STRING NUM IDENTIFIER

%left '!'
%left '('
%left '.'
%left '*' '/' '%'
%left '+' '-'
%left SHL SHR SHRNUM
%left '>' '<' LT GT INSTANCEOF
%left '&'
%left '^'
%left '|'

%left LESS_LESS
%left '=' SLASH_EQUAL PLUS_EQUAL LESS_LESS_EQUAL STAR_EQUAL EXCLAMATION_EQUAL BAR_EQUAL AMPERSAND_EQUAL
%left EQUAL_EQUAL LESS_EQUAL MORE_EQUAL
%left AMPERSAND_AMPERSAND BAR_BAR

%left '?' ':'

%start file_decl

%%

file_decl:
	package_stmt import_stmts annotated_class_decl {
		compilationUnitNode = $3;
	}
	;

qualified_id:
	identifier {}
	| qualified_id '.' identifier
	;

package_stmt:
	PACKAGE qualified_id ';'
	;

import_stmts:
	/* empty */ {}
	| import_stmts import_stmt {}
	;

import_stmt:
	IMPORT qualified_id ';' { }
	| IMPORT STATIC qualified_id ';' { }
	| IMPORT qualified_id '.' '*' ';' { }
	;

annotated_class_decl:
	annotations class_decl {
		$$ = $2;
	}
	;

annotations:
	/* empty */ { }
	| annotations annotation { }
	;

annotation:
	'@' qualified_id { }
	| '@' qualified_id '(' named_arg_list ')' {
	}
	;

named_arg_list:
	named_arg { }
	| named_arg_list ',' named_arg {

	}
	;

named_arg:
	/* expression already allows assignment */
	expression { }
	;

arg_list:
	/* */
	| non_empty_arg_list { }
	;

non_empty_arg_list:
	expression
	| non_empty_arg_list ',' expression
	;

expression:
	STRING {

	}
	| NUM  {

	}
	| identifier {

	}
	| '(' expression ')' { }
	| expression INSTANCEOF qualified_id {

        }
	| expression '?' expression ':' expression
	| expression '.' CLASS {

	}
	| expression '(' arg_list ')' {
          /* dubious: expression is usually not callable? */
	}
	| expression '.' identifier {

        }
	| expression '.' gen_type_arg identifier {

        }
	| expression '=' expression {	}
	| expression PLUS_EQUAL expression {	}
	| expression LESS_EQUAL expression {	}
	| expression MORE_EQUAL expression {	}
	| expression AMPERSAND_EQUAL expression {	}
	| expression BAR_EQUAL expression {	}
	| NEW type_specifier '(' arg_list ')' opt_method_overrides {
        }
	| '!' expression {}
	| expression '+' expression {

	}
	| expression '-' expression {

	}
	| '-' expression { }
	| expression '*' expression {

	}
	| expression '/' expression {

	}
	| expression '%' expression {

	}
	| expression LESS_LESS expression {

        }
	| expression '>' '>' expression {
		/* WTF:  >> is not a token.
		private Map<AnyObjectId, Set<Ref>> refsById;
		*/
        }
	| expression EQUAL_EQUAL expression { }
	| expression EXCLAMATION_EQUAL expression { }
	| expression '<' expression {
        }
	| expression '>' expression {
	}

	;

decl_modifier:
	PRIVATE { }
	| PUBLIC { }
	| PROTECTED { }
	| STATIC { }
	| FINAL { }
	| SYNCHRONIZED { }
	| ABSTRACT { }
	;

decl_modifiers:
	| decl_modifiers decl_modifier
	;

class_decl:
	decl_modifiers CLASS identifier class_inheritance class_impl '{' class_body '}' {
          $$ = $7
	  $$.Decl = Decl{
			  Type: "class",
			  Name: $3,
	  }
	}
	| decl_modifiers ENUM identifier '{' enum_body '}' {}
	;


identifier: IDENTIFIER { }
	;

class_inheritance:
	/* */ { }
	| EXTENDS qualified_id
	;
class_impl:
	/* */ { }
	| IMPLEMENTS interfaces
	;

interfaces:
	qualified_id
	| interfaces ',' qualified_id {
	}
	;

class_body:
	/* */ { $$ = &ClassDecl{} }
	| class_body class_element_decl {
		$$.Members = append($$.Members, $2)
	}
	;

opt_method_overrides: | '{' class_body '}'
	;

enum_body:
	enum_elts ';'
	| enum_body class_element_decl
	;

enum_elts:
	enum_elt
	| enum_elts ',' enum_elt
	;

enum_elt:
	identifier
	| identifier '(' arg_list ')'
	;

class_element_decl:
	annotations decl_modifiers type_specifier identifier class_var_init {
		$$ = &Decl{
			Name: $4,
			Type: $3,
		}
	}
	| annotations decl_modifiers identifier class_var_init {
                  /* ctor. Sigh. */
		$$ = &Decl{
			Name: $3,
			Type: "void",
		}
        }
	| annotated_class_decl {
	}
	;

class_var_init:
	var_init ';' { }
	|  '(' formal_arg_list ')' throws_decl '{' statement_list '}' {
        }
	|  '(' formal_arg_list ')' throws_decl ';' {
	}
	;

statement_list:
	/**/ { }
	| statement_list statement { }
	;

statement:
	expression ';' {}
	| identifier identifier var_init ';' {
	  /* can't use type_specifier, since '<' is ambiguous with expression. */
        }
	| FINAL type_specifier identifier var_init  ';' {}
	| RETURN expression ';' {}
	| RETURN ';' {}
	| try_catch {}
	| if_statement {}
	| while_statement {}
	| ';'
        | error ';' {
                    yylex.Error("statement error")
        }
	| THROW expression ';' {

        }
	| switch_statement {}
	| BREAK ';'
	;

switch_statement:
	SWITCH '(' expression ')' '{' switch_block '}'
        ;

switch_block:
	| switch_block switch_block_elt
        ;

switch_block_elt:
	CASE expression ':' statement_list
        | DEFAULT ':' statement_list
        ;

while_statement:
	WHILE '(' expression ')'  sub_block
        ;

sub_block:
	statement
        | '{' statement_list '}'
        ;

if_statement:
	IF '(' expression ')' sub_block opt_else
        ;

opt_else:
	| ELSE sub_block
	;

try_catch:
	TRY opt_resource '{' statement_list '}' opt_catch  opt_finally
	;

opt_catch:
	| CATCH '(' exception_catch_var ')' '{' statement_list '}'
	;

opt_resource:
	| '(' type_specifier identifier '=' expression ')'
	;

opt_finally:
	| FINALLY '{' statement_list '}' {}
	;

exception_catch_var:
	exception_type_catch_list identifier
        ;

exception_type_catch_list:
	qualified_id
        | exception_type_catch_list '|' qualified_id
        ;


var_init:
	/**/ {}
	| '=' expression
	;

formal_arg_list:
	/* */ { }
	| non_empty_formal_arg_list { }
	;

non_empty_formal_arg_list:
	formal_arg_elt { }
	| non_empty_formal_arg_list ',' formal_arg_elt {

	}
	;

formal_arg_elt:
	annotations decl_modifiers type_specifier opt_dotdotdot identifier { }
	;

opt_dotdotdot:
	| DOT_DOT_DOT
	;

type_specifier:
	identifier
	| '?' { $$ = "?" }
	| identifier gen_type_arg { $$ = $1 }
	| type_specifier '[' ']' { $$ = $1 + "[]" }
	;

gen_type_arg:
	'<' gen_type_list '>'
        | '<' '>'
        ;

gen_type_list:
	type_specifier
        | type_specifier ',' type_specifier
        ;

throws_decl:
	/**/  { }
	| THROWS nonempty_qualified_identifier_list { }
	;

nonempty_qualified_identifier_list:
	qualified_id { }
	| nonempty_qualified_identifier_list ',' qualified_id {
	}
	;

%%

var compilationUnitNode Node
