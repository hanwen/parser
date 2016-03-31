/* -*- Go -*-
Yaccable grammar based on
https://docs.oracle.com/javase/specs/jls/se8/html/jls-19.html

Included for amusement purposes. Yacc detects 18 unused reductions, ~1600 S/R conflicts and 526 R/R conflicts.
 */
%{

package java

%}

%token integerliteral floatingpointliteral booleanliteral characterliteral stringliteral nullliteral
%token Identifier

%token abstract
%token ampersandampersand
%token ampersandequal
%token arrow
%token assert
%token barbar
%token barequal
%token boolean
%token brackets
%token BREAK
%token byte
%token CASE
%token catch
%token char
%token class
%token CONTINUE
%token DEFAULT
%token diamond
%token do
%token dotdotdot
%token double
%token doublecolon
%token ELSE
%token enum
%token equalequal
%token exclamationequal
%token extends
%token final
%token finally
%token float
%token FOR
%token hatequal
%token identifier
%token IF
%token implements
%token IMPORT
%token instanceof
%token int
%token INTERFACE
%token lessequal
%token lessless
%token lesslessequal
%token literal
%token long
%token minusequal
%token minusminus
%token moreequal
%token moremore
%token moremoreequal
%token moremoremore
%token moremoremoreequal
%token native
%token new
%token PACKAGE
%token percentequal
%token plusequal
%token plusplus
%token private
%token protected
%token public
%token RETURN
%token short
%token slashequal
%token starequal
%token static
%token strictfp
%token super
%token SWITCH
%token synchronized
%token this
%token throw
%token throws
%token transient
%token try
%token void
%token volatile
%token while

%left '*' '/' '%'
%left '+' '-'
%left '>' '<'
%left '&'
%left '^'
%left '|'
%left barbar ampersandampersand


%start CompilationUnit

%%

Type:
	PrimitiveType
	| ReferenceType
	;

PrimitiveType:
	rep_Annotation NumericType
	| rep_Annotation boolean
	;

NumericType:
	IntegralType
	| FloatingPointType
	;

IntegralType:
	byte | short | int | long | char
	;

FloatingPointType:
	float | double
	;

ReferenceType:
	ClassOrInterfaceType
	| TypeVariable
	| ArrayType
	;

ClassOrInterfaceType:
	ClassType
	| InterfaceType
	;

ClassType:
	rep_Annotation Identifier opt_TypeArguments
	| ClassOrInterfaceType '.' rep_Annotation Identifier opt_TypeArguments
	;

InterfaceType:
	ClassType
	;

TypeVariable:
	rep_Annotation Identifier
	;

ArrayType:
	PrimitiveType Dims
	| ClassOrInterfaceType Dims
	| TypeVariable Dims
	;

Dims:
	DimsElt | Dims DimsElt
	;

DimsElt:
	rep_Annotation '[' ']'
	;

TypeParameter:
	rep_TypeParameterModifier Identifier opt_TypeBound
	;

TypeParameterModifier:
	Annotation
	;

TypeBound:
	extends TypeVariable
	| extends ClassOrInterfaceType rep_AdditionalBound
	;

AdditionalBound:
	'&' InterfaceType
	;

TypeArguments:
	'<' TypeArgumentList '>'
	;

TypeArgumentList:
	TypeArgument rep_comma_TypeArgument
	;

TypeArgument:
	ReferenceType
	| Wildcard
	;

Wildcard:
	rep_Annotation '?' opt_WildcardBounds
	;

WildcardBounds:
	extends ReferenceType
	| super ReferenceType
	;

TypeName:
	Identifier
	| PackageOrTypeName '.' Identifier
	;

PackageOrTypeName:
	Identifier
	| PackageOrTypeName '.' Identifier
	;

ExpressionName:
	Identifier
	| AmbiguousName '.' Identifier
	;

MethodName:
	Identifier
	;

PackageName:
	Identifier
	| PackageName '.' Identifier
	;

AmbiguousName:
	Identifier
	| AmbiguousName '.' Identifier
	;

CompilationUnit:
	opt_PackageDeclaration rep_ImportDeclaration rep_TypeDeclaration
	;

PackageDeclaration:
	rep_PackageModifier PACKAGE Identifier rep_dot_Identifier ';'
	;

PackageModifier:
	Annotation
	;

ImportDeclaration:
	SingleTypeImportDeclaration
	| TypeImportOnDemandDeclaration
	| SingleStaticImportDeclaration
	| StaticImportOnDemandDeclaration
	;

SingleTypeImportDeclaration:
	IMPORT TypeName ';'
	;

TypeImportOnDemandDeclaration:
	IMPORT PackageOrTypeName '.' '*' ';'
	;

SingleStaticImportDeclaration:
	IMPORT static TypeName '.' Identifier ';'
	;

StaticImportOnDemandDeclaration:
	IMPORT static TypeName '.' '*' ';'
	;

TypeDeclaration:
	ClassDeclaration
	| InterfaceDeclaration
	;

ClassDeclaration:
	NormalClassDeclaration
	| EnumDeclaration
	;

NormalClassDeclaration:
	rep_ClassModifier class Identifier opt_TypeParameters opt_Superclass opt_Superinterfaces ClassBody
	;

ClassModifier:
	Annotation | public | protected | private
	| abstract | static | final | strictfp
	;

TypeParameters:
	'<' TypeParameterList '>'
	;

TypeParameterList:
	TypeParameter rep_comma_TypeParameter
	;

Superclass:
	extends ClassType
	;

Superinterfaces:
	implements InterfaceTypeList
	;

InterfaceTypeList:
	InterfaceType rep_comma_InterfaceType
	;

ClassBody:
	'{' rep_ClassBodyDeclaration '}'
	;

ClassBodyDeclaration:
	ClassMemberDeclaration
	| InstanceInitializer
	| StaticInitializer
	| ConstructorDeclaration
	;

ClassMemberDeclaration:
	FieldDeclaration
	| MethodDeclaration
	| ClassDeclaration
	| InterfaceDeclaration
;

FieldDeclaration:
	rep_FieldModifier UnannType VariableDeclaratorList
	;

FieldModifier:
	Annotation | public | protected | private
	| static | final| transient| volatile
	;

VariableDeclaratorList:
	VariableDeclarator rep_comma_VariableDeclarator
	;

VariableDeclarator:
	VariableDeclaratorId opt_eq_VariableInitializer
	;

VariableDeclaratorId:
	Identifier opt_Dims
	;

VariableInitializer:
	Expression
	| ArrayInitializer
	;

UnannType:
	UnannPrimitiveType
	| UnannReferenceType
	;

UnannPrimitiveType:
	NumericType
	| boolean
	;

UnannReferenceType:
	UnannClassOrInterfaceType
	| UnannTypeVariable
	| UnannArrayType
	;

UnannClassOrInterfaceType:
	UnannClassType
	| UnannInterfaceType
	;

UnannClassType:
	Identifier opt_TypeArguments
	| UnannClassOrInterfaceType '.' rep_Annotation Identifier opt_TypeArguments
	;

UnannInterfaceType:
	UnannClassType
	;

UnannTypeVariable:
	Identifier
	;

UnannArrayType:
	UnannPrimitiveType Dims
	| UnannClassOrInterfaceType Dims
	| UnannTypeVariable Dims
	;

MethodDeclaration:
	rep_MethodModifier MethodHeader MethodBody
	;

MethodModifier:
	Annotation | public | protected | private
	| abstract | static | final | synchronized | native | strictfp
	;

MethodHeader:
	Result MethodDeclarator opt_Throws
	| TypeParameters rep_Annotation Result MethodDeclarator opt_Throws
	;

Result:
	UnannType
	| void
	;

MethodDeclarator:
	Identifier '(' opt_FormalParameterList ')' opt_Dims
	;

FormalParameterList:
	ReceiverParameter
	| FormalParameters ',' LastFormalParameter
	| LastFormalParameter
	;

FormalParameters:
	FormalParameter rep_comma_FormalParameter
	| ReceiverParameter rep_comma_FormalParameter
	;

FormalParameter:
	rep_VariableModifier UnannType VariableDeclaratorId
	;

VariableModifier:
	Annotation | final
	;

LastFormalParameter:
	rep_VariableModifier UnannType rep_Annotation dotdotdot VariableDeclaratorId
	| FormalParameter
	;

ReceiverParameter:
	rep_Annotation UnannType opt_Identifier_dot this
	;

Throws:
	throws ExceptionTypeList
	;

ExceptionTypeList:
	ExceptionType rep_comma_ExceptionType
	;

ExceptionType:
	ClassType
	| TypeVariable
	;

MethodBody:
	Block
	| ';'
	;

InstanceInitializer:
	Block
	;

StaticInitializer:
	static Block
	;

ConstructorDeclaration:
	rep_ConstructorModifier ConstructorDeclarator opt_Throws ConstructorBody
	;

ConstructorModifier:
	Annotation | public | protected | private
	;

ConstructorDeclarator:
	opt_TypeParameters SimpleTypeName '(' opt_FormalParameterList ')'
	;

SimpleTypeName:
	Identifier
	;

ConstructorBody:
	'{' opt_ExplicitConstructorInvocation opt_BlockStatements '}'
	;

ExplicitConstructorInvocation:
	opt_TypeArguments this ParenArgumentList ';'
	| opt_TypeArguments super ParenArgumentList ';'
	| ExpressionName '.' opt_TypeArguments super ParenArgumentList ';'
	| Primary '.' opt_TypeArguments super ParenArgumentList ';'
	;

EnumDeclaration:
	rep_ClassModifier enum Identifier opt_Superinterfaces EnumBody
	;

EnumBody:
	'{' opt_EnumConstantList opt_comma opt_EnumBodyDeclarations '}'
	;

EnumConstantList:
	EnumConstant rep_comma_EnumConstant
	;

EnumConstant:
	rep_EnumConstantModifier Identifier opt_ParenArgumentList opt_ClassBody
	;

ParenArgumentList: '(' ArgumentList ')'
	;

EnumConstantModifier:
	Annotation
	;

EnumBodyDeclarations:
	';' rep_ClassBodyDeclaration
	;

InterfaceDeclaration:
	NormalInterfaceDeclaration
	| AnnotationTypeDeclaration
	;

NormalInterfaceDeclaration:
	rep_InterfaceModifier INTERFACE Identifier opt_TypeParameters opt_ExtendsInterfaces InterfaceBody
	;

InterfaceModifier:
	Annotation |  public | protected | private |
abstract | static | strictfp
	;

ExtendsInterfaces:
	extends InterfaceTypeList
	;

InterfaceBody:
	'{' rep_InterfaceMemberDeclaration '}'
	;

InterfaceMemberDeclaration:
	ConstantDeclaration
	| InterfaceMethodDeclaration
	| ClassDeclaration
	| InterfaceDeclaration
	';'
	;

ConstantDeclaration:
	rep_ConstantModifier UnannType VariableDeclaratorList ;
	;

ConstantModifier:
	Annotation | public
	| static | final
	;

InterfaceMethodDeclaration:
	rep_InterfaceMethodModifier MethodHeader MethodBody
	;

InterfaceMethodModifier:
	Annotation | public |
	abstract | DEFAULT | static | strictfp
	;

AnnotationTypeDeclaration:
	rep_InterfaceModifier '@' INTERFACE Identifier AnnotationTypeBody
	;

AnnotationTypeBody:
	'{' rep_AnnotationTypeMemberDeclaration '}'
	;

AnnotationTypeMemberDeclaration:
	AnnotationTypeElementDeclaration
	| ConstantDeclaration
	| ClassDeclaration
	| InterfaceDeclaration
	| ';'
	;

AnnotationTypeElementDeclaration:
	rep_AnnotationTypeElementModifier UnannType Identifier '(' ')' opt_Dims opt_DefaultValue ';'

AnnotationTypeElementModifier:
	Annotation | public | abstract
	;

DefaultValue:
	DEFAULT ElementValue
	;

Annotation:
	NormalAnnotation
	| MarkerAnnotation
	| SingleElementAnnotation
	;

NormalAnnotation:
	'@' TypeName '(' opt_ElementValuePairList ')'
	;

ElementValuePairList:
	ElementValuePair rep_comma_ElementValuePair
	;

ElementValuePair:
	Identifier '=' ElementValue
	;

ElementValue:
	ConditionalExpression
	| ElementValueArrayInitializer
	| Annotation
	;

ElementValueArrayInitializer:
	'{' rep_ElementValueArrayInitializer '}'
	;

ElementValueArrayInitializer:
	opt_ElementValueList opt_comma
	;

ElementValueList:
	ElementValue rep_comma_ElementValue
	;

MarkerAnnotation:
	'@' TypeName
	;

SingleElementAnnotation:
	'@' TypeName '(' ElementValue ')'
	;

ArrayInitializer:
	'{' opt_VariableInitializerList opt_comma '}'
	;

VariableInitializerList:
	VariableInitializer rep_comma_VariableInitializer
	;

Block:
	'{' opt_BlockStatements '}'
	;

BlockStatements:
	BlockStatement rep_BlockStatement
	;

BlockStatement:
	LocalVariableDeclarationStatement
	| ClassDeclaration
	| Statement
	;

LocalVariableDeclarationStatement:
	LocalVariableDeclaration ;
	;

LocalVariableDeclaration:
	rep_VariableModifier UnannType VariableDeclaratorList
	;

Statement:
	StatementWithoutTrailingSubstatement
	| LabeledStatement
	| IfThenStatement
	| IfThenElseStatement
	| WhileStatement
	| ForStatement
	;

StatementNoShortIf:
	StatementWithoutTrailingSubstatement
	| LabeledStatementNoShortIf
	| IfThenElseStatementNoShortIf
	| WhileStatementNoShortIf
	| ForStatementNoShortIf
	;

StatementWithoutTrailingSubstatement:
	Block
	| EmptyStatement
	| ExpressionStatement
	| AssertStatement
	| SwitchStatement
	| DoStatement
	| BreakStatement
	| ContinueStatement
	| ReturnStatement
	| SynchronizedStatement
	| ThrowStatement
	| TryStatement
	;

EmptyStatement:
	;

LabeledStatement:
	Identifier ':' Statement
	;

LabeledStatementNoShortIf:
	Identifier ':' StatementNoShortIf
	;

ExpressionStatement:
	StatementExpression ;
	;

StatementExpression:
	Assignment
	| PreIncrementExpression
	| PreDecrementExpression
	| PostIncrementExpression
	| PostDecrementExpression
	| MethodInvocation
	| ClassInstanceCreationExpression
	;

IfThenStatement:
	IF '(' Expression ')' Statement
	;

IfThenElseStatement:
	IF '(' Expression ')' StatementNoShortIf ELSE Statement
	;

IfThenElseStatementNoShortIf:
	IF '(' Expression ')' StatementNoShortIf ELSE StatementNoShortIf
	;

AssertStatement:
	assert Expression ';'
assert Expression ':' Expression ';'
	;

SwitchStatement:
	SWITCH '(' Expression ')' SwitchBlock
	;

SwitchBlock:
	'{' rep_SwitchBlockStatementGroup rep_SwitchLabel '}'
	;

SwitchBlockStatementGroup:
	SwitchLabels BlockStatements
	;

SwitchLabels:
	SwitchLabel rep_SwitchLabel
	;

SwitchLabel:
	CASE ConstantExpression ':'
	| CASE EnumConstantName ':'
	| DEFAULT ':'
	;

EnumConstantName:
	Identifier
	;

WhileStatement:
	while '(' Expression ')' Statement
	;

WhileStatementNoShortIf:
	while '(' Expression ')' StatementNoShortIf
	;

DoStatement:
	do Statement while '(' Expression ')' ;
	;

ForStatement:
	BasicForStatement
	| EnhancedForStatement
	;

ForStatementNoShortIf:
	BasicForStatementNoShortIf
	| EnhancedForStatementNoShortIf
	;

BasicForStatement:
	FOR '(' opt_ForInit ';' opt_Expression ';' opt_ForUpdate ')' Statement
	;

BasicForStatementNoShortIf:
	FOR '(' opt_ForInit ';' opt_Expression ';' opt_ForUpdate ')' StatementNoShortIf
	;

ForInit:
	StatementExpressionList
	| LocalVariableDeclaration
	;

ForUpdate:
	StatementExpressionList
	;

StatementExpressionList:
	StatementExpression rep_comma_StatementExpression
	;

EnhancedForStatement:
	FOR '(' rep_VariableModifier UnannType VariableDeclaratorId ':' Expression ')' Statement
	;

EnhancedForStatementNoShortIf:
	FOR '(' rep_VariableModifier UnannType VariableDeclaratorId ':' Expression ')' StatementNoShortIf
	;

BreakStatement:
	BREAK opt_Identifier ;
	;

ContinueStatement:
	CONTINUE opt_Identifier ;
	;

ReturnStatement:
	RETURN opt_Expression ;
	;

ThrowStatement:
	throw Expression ;
	;

SynchronizedStatement:
	synchronized '(' Expression ')' Block
	;

TryStatement:
	try Block Catches
	| try Block opt_Catches Finally
	| TryWithResourcesStatement
	;

Catches:
	CatchClause rep_CatchClause
	;

CatchClause:
	catch '(' CatchFormalParameter ')' Block
	;

CatchFormalParameter:
	rep_VariableModifier CatchType VariableDeclaratorId
	;

CatchType:
	UnannClassType rep_bar_ClassType
	;

Finally:
	finally Block
	;

TryWithResourcesStatement:
	try ResourceSpecification Block opt_Catches opt_Finally
	;

ResourceSpecification:
	'(' ResourceList opt_semicolon ')'
	;

ResourceList:
	Resource rep_semicolon_Resource
	;

Resource:
	rep_VariableModifier UnannType VariableDeclaratorId '=' Expression
	;

Primary:
	PrimaryNoNewArray
	| ArrayCreationExpression
	;

Literal:
	integerliteral
	| floatingpointliteral
	| booleanliteral
	| characterliteral
	| stringliteral
	| nullliteral
	;

PrimaryNoNewArray:
	Literal
	| ClassLiteral
	| this
	| TypeName '.' this
	| '(' Expression ')'
	| ClassInstanceCreationExpression
	| FieldAccess
	| ArrayAccess
	| MethodInvocation
	| MethodReference
	;

ClassLiteral:
	TypeName rep_brackets '.' class
	| NumericType rep_brackets '.' class
	| boolean rep_brackets '.' class
	| void '.' class
	;

ClassInstanceCreationExpression:
	UnqualifiedClassInstanceCreationExpression
	| ExpressionName '.' UnqualifiedClassInstanceCreationExpression
	| Primary '.' UnqualifiedClassInstanceCreationExpression
	;

UnqualifiedClassInstanceCreationExpression:
	new opt_TypeArguments ClassOrInterfaceTypeToInstantiate ParenArgumentList opt_ClassBody
	;

ClassOrInterfaceTypeToInstantiate:
	rep_Annotation Identifier rep_dot_rep_Annotation_Identifier opt_TypeArgumentsOrDiamond
	;
dot_rep_Annotation_Identifier:
	'.' rep_Annotation Identifier
	;

TypeArgumentsOrDiamond:
	TypeArguments
	| diamond
	;

FieldAccess:
	Primary '.' Identifier
	| super '.' Identifier
	| TypeName '.' super '.' Identifier
	;

ArrayAccess:
	ExpressionName '[' Expression ']'
	| PrimaryNoNewArray '[' Expression ']'
	;

MethodInvocation:
	MethodName ParenArgumentList
	| TypeName '.' opt_TypeArguments Identifier ParenArgumentList
	| ExpressionName '.' opt_TypeArguments Identifier ParenArgumentList
	| Primary '.' opt_TypeArguments Identifier ParenArgumentList
	| super '.' opt_TypeArguments Identifier ParenArgumentList
	| TypeName '.' super '.' opt_TypeArguments Identifier ParenArgumentList
	;

ArgumentList:
	Expression rep_comma_Expression
	;

MethodReference:
	ExpressionName doublecolon opt_TypeArguments Identifier
	| ReferenceType doublecolon opt_TypeArguments Identifier
	| Primary doublecolon opt_TypeArguments Identifier
	| super doublecolon opt_TypeArguments Identifier
	| TypeName '.' super doublecolon opt_TypeArguments Identifier
	| ClassType doublecolon opt_TypeArguments new
	| ArrayType doublecolon new
	;

ArrayCreationExpression:
	new PrimitiveType DimExprs opt_Dims
	| new ClassOrInterfaceType DimExprs opt_Dims
	| new PrimitiveType Dims ArrayInitializer
	| new ClassOrInterfaceType Dims ArrayInitializer
	;

DimExprs:
	DimExpr rep_DimExpr
	;

DimExpr:
	rep_Annotation '[' Expression ']'
	;

Expression:
	LambdaExpression
	| AssignmentExpression
	;

LambdaExpression:
	LambdaParameters arrow LambdaBody
	;

LambdaParameters:
	Identifier
'(' opt_FormalParameterList ')'
'(' InferredFormalParameterList ')'
	;

InferredFormalParameterList:
	Identifier rep_comma_Identifier
	;

LambdaBody:
	Expression
	| Block
	;

AssignmentExpression:
	ConditionalExpression
	| Assignment
	;

Assignment:
	LeftHandSide AssignmentOperator Expression
	;

LeftHandSide:
	ExpressionName
	| FieldAccess
	| ArrayAccess
	;

AssignmentOperator:
	'=' | starequal | slashequal | percentequal | plusequal | minusequal | lesslessequal | moremoreequal | moremoremoreequal | ampersandequal | hatequal | barequal
	;

ConditionalExpression:
	ConditionalOrExpression
	| ConditionalOrExpression '?' Expression ':' ConditionalExpression
	| ConditionalOrExpression '?' Expression ':' LambdaExpression
	;

ConditionalOrExpression:
	ConditionalAndExpression
	| ConditionalOrExpression barbar ConditionalAndExpression
	;

ConditionalAndExpression:
	InclusiveOrExpression
	| ConditionalAndExpression ampersandampersand InclusiveOrExpression
	;

InclusiveOrExpression:
	ExclusiveOrExpression
	| InclusiveOrExpression '|' ExclusiveOrExpression
	;

ExclusiveOrExpression:
	AndExpression
	| ExclusiveOrExpression '^' AndExpression
	;

AndExpression:
	EqualityExpression
	| AndExpression '&' EqualityExpression
	;

EqualityExpression:
	RelationalExpression
	| EqualityExpression equalequal RelationalExpression
	| EqualityExpression exclamationequal RelationalExpression
	;

RelationalExpression:
	ShiftExpression
	| RelationalExpression '<' ShiftExpression
	| RelationalExpression '>' ShiftExpression
	| RelationalExpression lessequal ShiftExpression
	| RelationalExpression moreequal ShiftExpression
	| RelationalExpression instanceof ReferenceType
	;

ShiftExpression:
	AdditiveExpression
	| ShiftExpression lessless AdditiveExpression
	| ShiftExpression moremore AdditiveExpression
	| ShiftExpression moremoremore AdditiveExpression
	;

AdditiveExpression:
	MultiplicativeExpression
	| AdditiveExpression '+' MultiplicativeExpression
	| AdditiveExpression '-' MultiplicativeExpression
	;

MultiplicativeExpression:
	UnaryExpression
	| MultiplicativeExpression '*' UnaryExpression
	| MultiplicativeExpression '/' UnaryExpression
	| MultiplicativeExpression '%' UnaryExpression
	;

UnaryExpression:
	PreIncrementExpression
	| PreDecrementExpression
	| '+' UnaryExpression
	| '-' UnaryExpression
	| UnaryExpressionNotPlusMinus
	;

PreIncrementExpression:
	plusplus UnaryExpression
	;

PreDecrementExpression:
	minusminus UnaryExpression
	;

UnaryExpressionNotPlusMinus:
	PostfixExpression
	| '~' UnaryExpression
	| '!' UnaryExpression
	| CastExpression
	;

PostfixExpression:
	Primary
	| ExpressionName
	| PostIncrementExpression
	| PostDecrementExpression
	;

PostIncrementExpression:
	PostfixExpression plusplus
	;

PostDecrementExpression:
	PostfixExpression minusminus
	;

CastExpression:
	'(' PrimitiveType ')' UnaryExpression
	| '(' ReferenceType rep_AdditionalBound ')' UnaryExpressionNotPlusMinus
	| '(' ReferenceType rep_AdditionalBound ')' LambdaExpression
	;

ConstantExpression:
	Expression

/* repeated exprs */
rep_Annotation: | rep_Annotation Annotation ;
rep_TypeParameterModifier: | rep_TypeParameterModifier TypeParameterModifier ;
rep_AdditionalBound: | rep_AdditionalBound AdditionalBound ;
rep_comma_TypeArgument: | rep_comma_TypeArgument comma_TypeArgument ;
rep_ImportDeclaration: | rep_ImportDeclaration ImportDeclaration ;
rep_TypeDeclaration: | rep_TypeDeclaration TypeDeclaration ;
rep_PackageModifier: | rep_PackageModifier PackageModifier ;
rep_dot_Identifier: | rep_dot_Identifier dot_Identifier ;
rep_ClassModifier: | rep_ClassModifier ClassModifier ;
rep_comma_TypeParameter: | rep_comma_TypeParameter comma_TypeParameter ;
rep_comma_InterfaceType: | rep_comma_InterfaceType comma_InterfaceType ;
rep_ClassBodyDeclaration: | rep_ClassBodyDeclaration ClassBodyDeclaration ;
rep_FieldModifier: | rep_FieldModifier FieldModifier ;
rep_comma_VariableDeclarator: | rep_comma_VariableDeclarator comma_VariableDeclarator ;
rep_MethodModifier: | rep_MethodModifier MethodModifier ;
rep_comma_FormalParameter: | rep_comma_FormalParameter comma_FormalParameter ;
rep_VariableModifier: | rep_VariableModifier VariableModifier ;
rep_comma_ExceptionType: | rep_comma_ExceptionType comma_ExceptionType ;
rep_ConstructorModifier: | rep_ConstructorModifier ConstructorModifier ;
rep_comma_EnumConstant: | rep_comma_EnumConstant comma_EnumConstant ;
rep_EnumConstantModifier: | rep_EnumConstantModifier EnumConstantModifier ;
rep_InterfaceModifier: | rep_InterfaceModifier InterfaceModifier ;
rep_ConstantModifier: | rep_ConstantModifier ConstantModifier ;
rep_InterfaceMethodModifier: | rep_InterfaceMethodModifier InterfaceMethodModifier ;
rep_AnnotationTypeElementModifier: | rep_AnnotationTypeElementModifier AnnotationTypeElementModifier ;
rep_comma_ElementValuePair: | rep_comma_ElementValuePair comma_ElementValuePair ;
rep_comma_ElementValue: | rep_comma_ElementValue comma_ElementValue ;
rep_comma_VariableInitializer: | rep_comma_VariableInitializer comma_VariableInitializer ;
rep_BlockStatement: | rep_BlockStatement BlockStatement ;
rep_SwitchBlockStatementGroup: | rep_SwitchBlockStatementGroup SwitchBlockStatementGroup ;
rep_SwitchLabel: | rep_SwitchLabel SwitchLabel ;
rep_comma_StatementExpression: | rep_comma_StatementExpression comma_StatementExpression ;
rep_CatchClause: | rep_CatchClause CatchClause ;
rep_bar_ClassType: | rep_bar_ClassType bar_ClassType ;
rep_semicolon_Resource: | rep_semicolon_Resource semicolon_Resource ;
rep_brackets: | rep_brackets brackets ;
rep_dot_rep_Annotation_Identifier: | rep_dot_rep_Annotation_Identifier dot_rep_Annotation_Identifier ;
rep_comma_Expression: | rep_comma_Expression comma_Expression ;
rep_DimExpr: | rep_DimExpr DimExpr ;
rep_comma_Identifier: | rep_comma_Identifier comma_Identifier ;
rep_InterfaceMemberDeclaration: | rep_InterfaceMemberDeclaration InterfaceMemberDeclaration ;
rep_AnnotationTypeMemberDeclaration: | rep_AnnotationTypeMemberDeclaration AnnotationTypeMemberDeclaration ;
rep_ElementValueArrayInitializer: | rep_ElementValueArrayInitializer ElementValueArrayInitializer ;


/* opt */

opt_TypeArguments: | TypeArguments ;
opt_TypeBound: | TypeBound ;
opt_WildcardBounds: | WildcardBounds ;
opt_PackageDeclaration: | PackageDeclaration ;
opt_TypeParameters: | TypeParameters ;
opt_Superclass: | Superclass ;
opt_Superinterfaces: | Superinterfaces ;
opt_eq_VariableInitializer: | eq_VariableInitializer ;
opt_Dims: | Dims ;
opt_Throws: | Throws ;
opt_FormalParameterList: | FormalParameterList ;
opt_Identifier_dot: | Identifier_dot ;
opt_ExplicitConstructorInvocation: | ExplicitConstructorInvocation ;
opt_BlockStatements: | BlockStatements ;
opt_EnumConstantList: | EnumConstantList ;
opt_comma: | ',' ;
opt_EnumBodyDeclarations: | EnumBodyDeclarations ;
opt_ParenArgumentList: | ParenArgumentList ;
opt_ClassBody: | ClassBody ;
opt_ExtendsInterfaces: | ExtendsInterfaces ;
opt_DefaultValue: | DefaultValue ;
opt_ElementValuePairList: | ElementValuePairList ;
opt_ElementValueList: | ElementValueList ;
opt_VariableInitializerList: | VariableInitializerList ;
opt_ForInit: | ForInit ;
opt_Expression: | Expression ;
opt_ForUpdate: | ForUpdate ;
opt_Identifier: | Identifier ;
opt_Catches: | Catches ;
opt_Finally: | Finally ;
opt_semicolon: | ';' ;
opt_TypeArgumentsOrDiamond: | TypeArgumentsOrDiamond ;

/* comma */
comma_ElementValue: ',' ElementValue ;
comma_ElementValuePair: ',' ElementValuePair ;
comma_EnumConstant: ',' EnumConstant ;
comma_ExceptionType: ',' ExceptionType ;
comma_Expression: ',' Expression ;
comma_FormalParameter: ',' FormalParameter ;
comma_Identifier: ',' Identifier ;
comma_InterfaceType: ',' InterfaceType ;
comma_StatementExpression: ',' StatementExpression ;
comma_TypeArgument: ',' TypeArgument ;
comma_TypeParameter: ',' TypeParameter ;
comma_VariableDeclarator: ',' VariableDeclarator ;
comma_VariableInitializer: ',' VariableInitializer ;

dot_Identifier: '.' Identifier ;
bar_ClassType: '|' ClassType ;
semicolon_Resource: ';' Resource ;
eq_VariableInitializer: '=' VariableInitializer ;

Identifier_dot: Identifier '.' ;

%%
