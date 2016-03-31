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

// A parser for Java. It is incomplete, and can fundamentally not work
// unless we introduce heuristics for type names vs. object names.

%{

package java

%}

%union {
	Text string
	Expr Expr
}

%type <expr> expression

%token ABSTRACT
%token BREAK
%token CASE
%token CATCH
%token CLASS
%token DEFAULT
%token ELSE
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

%left LESS_LESS MORE_MORE
%left '=' SLASH_EQUAL PLUS_EQUAL LESS_LESS_EQUAL STAR_EQUAL EXCLAMATION_EQUAL BAR_EQUAL AMPERSAND_EQUAL
%left EQUAL_EQUAL LESS_EQUAL MORE_EQUAL
%left AMPERSAND_AMPERSAND BAR_BAR

%left '?' ':'

%start file_decl

%%

file_decl:
	package_stmt import_stmts annotated_class_decl;
	;

/*
  This is a disaster: imagine

    a.b

  is this the type identifier b from package a, or is this an object
  named a with member b?
*/

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
	annotations class_decl {}
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
/* type casts are a mess up: we'll just have to error out on them.

   (a)(b)

   is this function a called with b as arg, or is this expression b
   cast to type a? The only way around this is to have heuristics
   which guess if a (qualified) identifier is a typename or a instance
   name.
 */
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
	| expression EQUAL_EQUAL expression { }
	| expression EXCLAMATION_EQUAL expression { }
	| expression '<' expression {
	  /* a < b > c

             is this typeA<typeB> c

             or (a < b) > c ?

             again, we must guess if an identifier is a type or an object.
           */
        }
	| expression '>' expression { }
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

	}
	;

identifier: IDENTIFIER
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
	/* */ { }
	| class_body class_element_decl {

	}
	;

opt_method_overrides: | '{' class_body '}'
	;

class_element_decl:
	annotations decl_modifiers type_specifier identifier class_var_init
	| annotations decl_modifiers identifier class_var_init {
                  /* ctor. Sigh. */
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
	TRY '{' statement_list '}' CATCH '(' exception_catch_var ')' '{' statement_list '}' opt_finally
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
	decl_modifiers type_specifier identifier { }
	;

type_specifier:
	identifier
        | '?'
        | identifier gen_type_arg
	| type_specifier '[' ']'
	;

gen_type_arg:
	'<' gen_type_list '>'
        | '<' '>'
        ;

gen_type:
	qualified_id
        | '?'
        ;

gen_type_list:
	gen_type
        | gen_type ',' gen_type
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

type Expr struct {
	Head string
	Children []Expr
}
