//line parser.y:3
package java

import __yyfmt__ "fmt"

//line parser.y:4
//line parser.y:8
type yySymType struct {
	yys  int
	Text string
	Expr Expr
}

const ABSTRACT = 57346
const BREAK = 57347
const CASE = 57348
const CATCH = 57349
const CLASS = 57350
const DEFAULT = 57351
const ELSE = 57352
const EXTENDS = 57353
const FINAL = 57354
const FINALLY = 57355
const FOR = 57356
const IF = 57357
const IMPLEMENTS = 57358
const IMPORT = 57359
const INSTANCEOF = 57360
const NEW = 57361
const PACKAGE = 57362
const PRIVATE = 57363
const PROTECTED = 57364
const PUBLIC = 57365
const RETURN = 57366
const STATIC = 57367
const SUPER = 57368
const SYNCHRONIZED = 57369
const SWITCH = 57370
const THROW = 57371
const THROWS = 57372
const TRY = 57373
const WHILE = 57374
const STRING = 57375
const NUM = 57376
const IDENTIFIER = 57377
const SHL = 57378
const SHR = 57379
const SHRNUM = 57380
const LT = 57381
const GT = 57382
const LESS_LESS = 57383
const MORE_MORE = 57384
const SLASH_EQUAL = 57385
const PLUS_EQUAL = 57386
const LESS_LESS_EQUAL = 57387
const STAR_EQUAL = 57388
const EXCLAMATION_EQUAL = 57389
const BAR_EQUAL = 57390
const AMPERSAND_EQUAL = 57391
const EQUAL_EQUAL = 57392
const LESS_EQUAL = 57393
const MORE_EQUAL = 57394
const AMPERSAND_AMPERSAND = 57395
const BAR_BAR = 57396

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ABSTRACT",
	"BREAK",
	"CASE",
	"CATCH",
	"CLASS",
	"DEFAULT",
	"ELSE",
	"EXTENDS",
	"FINAL",
	"FINALLY",
	"FOR",
	"IF",
	"IMPLEMENTS",
	"IMPORT",
	"INSTANCEOF",
	"NEW",
	"PACKAGE",
	"PRIVATE",
	"PROTECTED",
	"PUBLIC",
	"RETURN",
	"STATIC",
	"SUPER",
	"SYNCHRONIZED",
	"SWITCH",
	"THROW",
	"THROWS",
	"TRY",
	"WHILE",
	"STRING",
	"NUM",
	"IDENTIFIER",
	"'!'",
	"'('",
	"'.'",
	"'*'",
	"'/'",
	"'%'",
	"'+'",
	"'-'",
	"SHL",
	"SHR",
	"SHRNUM",
	"'>'",
	"'<'",
	"LT",
	"GT",
	"'&'",
	"'^'",
	"'|'",
	"LESS_LESS",
	"MORE_MORE",
	"'='",
	"SLASH_EQUAL",
	"PLUS_EQUAL",
	"LESS_LESS_EQUAL",
	"STAR_EQUAL",
	"EXCLAMATION_EQUAL",
	"BAR_EQUAL",
	"AMPERSAND_EQUAL",
	"EQUAL_EQUAL",
	"LESS_EQUAL",
	"MORE_EQUAL",
	"AMPERSAND_AMPERSAND",
	"BAR_BAR",
	"'?'",
	"':'",
	"';'",
	"'@'",
	"')'",
	"','",
	"'{'",
	"'}'",
	"'['",
	"']'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:406
type Expr struct {
	Head     string
	Children []Expr
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 142,
	71, 109,
	-2, 116,
	-1, 147,
	73, 111,
	-2, 58,
	-1, 244,
	6, 96,
	9, 96,
	76, 96,
	-2, 0,
	-1, 246,
	6, 95,
	9, 95,
	76, 95,
	-2, 0,
}

const yyNprod = 130
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 955

var yyAct = [...]int{

	45, 77, 172, 171, 6, 217, 146, 168, 154, 145,
	121, 115, 6, 20, 92, 127, 155, 149, 6, 234,
	6, 112, 235, 34, 165, 128, 250, 16, 164, 241,
	221, 197, 20, 140, 82, 54, 55, 7, 42, 169,
	6, 158, 134, 125, 116, 236, 157, 136, 78, 76,
	124, 80, 81, 6, 17, 12, 215, 6, 42, 89,
	87, 113, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 113,
	72, 62, 63, 209, 47, 57, 31, 57, 37, 232,
	196, 118, 6, 122, 194, 151, 50, 5, 43, 44,
	7, 48, 46, 240, 12, 18, 7, 7, 49, 114,
	148, 30, 147, 33, 71, 94, 60, 6, 61, 32,
	132, 73, 65, 64, 72, 62, 63, 126, 135, 57,
	7, 148, 7, 53, 91, 6, 193, 13, 142, 141,
	79, 123, 144, 56, 120, 138, 84, 137, 238, 133,
	86, 156, 143, 88, 150, 41, 78, 159, 12, 7,
	163, 7, 6, 36, 35, 12, 123, 162, 90, 173,
	6, 222, 75, 74, 190, 78, 191, 200, 192, 71,
	7, 60, 161, 61, 195, 147, 73, 65, 64, 72,
	62, 63, 202, 91, 57, 199, 91, 201, 198, 3,
	11, 206, 207, 208, 148, 205, 173, 52, 249, 210,
	131, 85, 19, 173, 173, 39, 218, 218, 47, 220,
	224, 216, 7, 6, 166, 173, 173, 225, 218, 237,
	230, 119, 43, 44, 7, 48, 46, 239, 153, 6,
	180, 228, 49, 183, 248, 173, 173, 173, 244, 245,
	174, 246, 173, 185, 56, 167, 130, 47, 251, 227,
	223, 10, 175, 188, 233, 226, 187, 181, 182, 184,
	186, 43, 44, 7, 48, 46, 178, 177, 66, 67,
	176, 49, 160, 75, 74, 152, 129, 83, 51, 38,
	71, 21, 60, 180, 61, 139, 183, 73, 65, 64,
	72, 62, 63, 174, 93, 57, 185, 40, 15, 179,
	47, 14, 9, 8, 252, 175, 229, 4, 2, 187,
	181, 1, 184, 186, 43, 44, 7, 48, 46, 0,
	0, 0, 242, 0, 49, 0, 0, 0, 60, 180,
	61, 0, 183, 73, 65, 64, 72, 62, 63, 174,
	0, 57, 185, 0, 0, 0, 47, 0, 0, 0,
	0, 175, 179, 0, 0, 187, 181, 247, 184, 186,
	43, 44, 7, 48, 46, 0, 56, 0, 0, 0,
	49, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 59, 58, 68, 69, 70,
	66, 67, 0, 0, 0, 75, 74, 0, 179, 0,
	0, 0, 71, 231, 60, 56, 61, 0, 0, 73,
	65, 64, 72, 62, 63, 0, 0, 57, 0, 0,
	0, 214, 0, 0, 59, 58, 68, 69, 70, 66,
	67, 0, 0, 0, 75, 74, 0, 0, 0, 0,
	0, 71, 56, 60, 0, 61, 0, 0, 73, 65,
	64, 72, 62, 63, 0, 0, 57, 0, 0, 0,
	213, 59, 58, 68, 69, 70, 66, 67, 0, 0,
	0, 75, 74, 0, 0, 0, 0, 0, 71, 0,
	60, 180, 61, 0, 183, 73, 65, 64, 72, 62,
	63, 174, 0, 57, 185, 0, 0, 212, 47, 0,
	0, 0, 0, 175, 0, 0, 0, 187, 181, 0,
	184, 186, 43, 44, 7, 48, 46, 0, 0, 0,
	0, 0, 49, 0, 29, 0, 0, 180, 22, 0,
	183, 0, 27, 0, 0, 0, 0, 174, 0, 0,
	185, 23, 25, 24, 47, 26, 0, 28, 0, 175,
	179, 0, 0, 187, 181, 211, 184, 186, 43, 44,
	7, 48, 46, 0, 56, 0, 0, 0, 49, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 59, 58, 68, 69, 70, 66, 67,
	0, 0, 0, 75, 74, 0, 179, 0, 0, 0,
	71, 170, 60, 180, 61, 0, 183, 73, 65, 64,
	72, 62, 63, 174, 0, 57, 185, 0, 0, 111,
	47, 0, 0, 0, 0, 175, 0, 0, 0, 187,
	181, 0, 184, 186, 43, 44, 7, 48, 46, 56,
	0, 0, 0, 0, 49, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 59, 58,
	68, 69, 70, 66, 67, 0, 0, 0, 75, 74,
	0, 0, 179, 0, 56, 71, 219, 60, 0, 61,
	0, 0, 73, 65, 64, 72, 62, 63, 0, 0,
	57, 0, 204, 59, 58, 68, 69, 70, 66, 67,
	0, 0, 0, 75, 74, 0, 0, 0, 0, 56,
	71, 0, 60, 0, 61, 0, 0, 73, 65, 64,
	72, 62, 63, 0, 0, 57, 0, 203, 59, 58,
	68, 69, 70, 66, 67, 0, 0, 0, 75, 74,
	0, 0, 0, 0, 56, 71, 0, 60, 0, 61,
	0, 0, 73, 65, 64, 72, 62, 63, 0, 0,
	57, 0, 189, 59, 58, 68, 69, 70, 66, 67,
	0, 0, 0, 75, 74, 0, 0, 0, 56, 0,
	71, 0, 60, 0, 61, 0, 0, 73, 65, 64,
	72, 62, 63, 0, 0, 57, 243, 59, 58, 68,
	69, 70, 66, 67, 0, 0, 0, 75, 74, 0,
	0, 0, 56, 0, 71, 0, 60, 0, 61, 0,
	0, 73, 65, 64, 72, 62, 63, 0, 0, 57,
	117, 59, 58, 68, 69, 70, 66, 67, 0, 0,
	0, 75, 74, 0, 0, 0, 0, 0, 71, 0,
	60, 180, 61, 0, 183, 73, 65, 64, 72, 62,
	63, 174, 0, 57, 185, 0, 0, 0, 47, 0,
	0, 0, 0, 175, 0, 0, 0, 187, 181, 29,
	184, 186, 43, 44, 7, 48, 46, 27, 0, 0,
	0, 0, 49, 0, 0, 0, 23, 25, 24, 0,
	26, 0, 28, 0, 0, 0, 0, 0, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	179, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 79,
}
var yyPact = [...]int{

	179, -1000, -1000, 126, 183, 66, -1000, -1000, -1000, -1000,
	-18, 187, 126, -1000, -1000, -1000, 530, 126, 48, 126,
	-1000, -1000, 126, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	127, 124, -1000, 17, 204, 199, 25, -1000, 191, 126,
	-38, -1000, 804, -1000, -1000, -1000, 199, 71, 199, 199,
	-1000, -41, 126, 120, -1000, 199, 126, 199, 145, 199,
	199, 199, 199, 199, 199, 199, 199, 199, 199, 199,
	199, 199, 199, 199, 199, 199, 556, -16, 86, -1000,
	804, 125, -1000, -30, 120, -1000, -1000, 770, -1000, -1000,
	126, 97, -23, -31, 804, 16, 16, 18, 18, 16,
	16, 125, 125, 236, 236, 236, 282, 18, 16, 60,
	60, -1000, 199, -63, -1000, -51, 126, 199, -1000, 102,
	-1000, -32, 120, -1000, -1000, 199, -26, -1000, -1000, -1000,
	-18, 120, -1000, -1000, 72, 804, -42, 885, -1000, -1000,
	-1000, 2, 148, -59, 75, -1000, 24, -1000, 199, -1000,
	-1000, -1000, -27, -33, -1000, 885, 804, 152, -1000, 2,
	-47, 126, -1000, -1000, -1000, -1000, -35, 120, 535, 126,
	-1000, -1000, 701, 126, 71, 65, -1000, -1000, -1000, -1000,
	23, 199, -1000, 19, -44, 161, 158, 140, 120, -1000,
	54, 2, 666, -1000, -1000, 631, -1000, -1000, 199, 199,
	199, 12, 54, -1000, -1000, 489, 434, 397, 358, -1000,
	-15, 214, 611, 611, -45, -1000, 134, 210, -1000, -1000,
	-1000, -1000, 126, -1000, 611, 337, 13, -28, 95, 120,
	-1000, -1000, -1000, -1000, 199, 33, -46, -1000, 126, 736,
	-1000, -1000, 120, -1000, 859, 291, 859, 195, -1000, -49,
	-1000, 238, -1000,
}
var yyPgo = [...]int{

	0, 2, 321, 318, 317, 313, 93, 0, 312, 256,
	311, 308, 307, 155, 14, 304, 109, 1, 295, 291,
	16, 289, 288, 11, 287, 286, 9, 6, 285, 282,
	7, 3, 280, 277, 276, 268, 265, 264, 5, 260,
	259, 244, 241, 238, 8, 231, 10, 224,
}
var yyR1 = [...]int{

	0, 2, 6, 6, 3, 4, 4, 8, 8, 8,
	5, 9, 9, 11, 11, 12, 12, 13, 14, 14,
	15, 15, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 19, 19, 19, 19, 19, 19, 19, 20, 20,
	10, 7, 21, 21, 22, 22, 24, 24, 23, 23,
	18, 18, 25, 25, 26, 26, 26, 30, 30, 31,
	31, 31, 31, 31, 31, 31, 31, 31, 31, 31,
	31, 31, 35, 36, 36, 37, 37, 34, 38, 38,
	33, 39, 39, 32, 41, 41, 40, 42, 42, 27,
	27, 28, 28, 43, 43, 44, 17, 17, 17, 17,
	16, 16, 46, 46, 45, 45, 29, 29, 47, 47,
}
var yyR2 = [...]int{

	0, 3, 1, 3, 3, 0, 2, 3, 4, 5,
	2, 0, 2, 2, 5, 1, 3, 1, 0, 1,
	1, 3, 1, 1, 1, 3, 3, 5, 3, 4,
	3, 4, 3, 3, 3, 3, 3, 3, 6, 2,
	3, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 0, 2,
	8, 1, 0, 2, 0, 2, 1, 3, 0, 2,
	0, 3, 5, 4, 2, 7, 5, 0, 2, 2,
	4, 5, 3, 2, 1, 1, 1, 1, 2, 3,
	1, 2, 7, 0, 2, 4, 3, 5, 1, 3,
	6, 0, 2, 12, 0, 4, 2, 1, 3, 0,
	2, 0, 1, 1, 3, 3, 1, 1, 2, 3,
	3, 2, 1, 1, 1, 3, 0, 2, 1, 3,
}
var yyChk = [...]int{

	-1000, -2, -3, 20, -4, -6, -7, 35, -5, -8,
	-9, 17, 38, 71, -10, -11, -20, 72, -6, 25,
	-7, -19, 8, 21, 23, 22, 25, 12, 27, 4,
	-6, 38, 71, -6, -7, 37, 39, 71, -21, 11,
	-12, -13, -1, 33, 34, -7, 37, 19, 36, 43,
	71, -22, 16, -6, 73, 74, 18, 69, 38, 37,
	56, 58, 65, 66, 63, 62, 42, 43, 39, 40,
	41, 54, 64, 61, 48, 47, -1, -17, -7, 69,
	-1, -1, 75, -24, -6, -13, -6, -1, 8, -7,
	-16, 48, -14, -15, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, 73, 37, 77, -16, -23, 74, 70, -7, -45,
	47, -46, -6, 69, 73, 74, -14, 78, 76, -25,
	-9, -6, -1, 47, 74, -1, 73, -20, -46, -18,
	75, -17, -7, -23, -7, -26, -27, 37, 56, 76,
	-26, 71, -28, -43, -44, -20, -1, 73, 74, -17,
	-29, 30, -44, -7, 75, 71, -47, -6, -30, 74,
	76, -31, -1, -7, 12, 24, -32, -33, -34, 71,
	2, 29, -35, 5, 31, 15, 32, 28, -6, 71,
	-7, -17, -1, 71, 71, -1, 71, 75, 37, 37,
	37, -27, -7, 71, 71, -30, -1, -1, -1, 71,
	-27, 76, 73, 73, 73, 71, 7, -38, -31, 75,
	-38, 75, 37, -39, 10, -30, -36, -40, -42, -6,
	-38, 76, 76, -37, 6, 9, 73, -7, 53, -1,
	70, 75, -6, 70, -30, -30, -30, 76, -41, 13,
	75, -30, 76,
}
var yyDef = [...]int{

	0, -2, 5, 0, 11, 0, 2, 61, 1, 6,
	58, 0, 0, 4, 10, 12, 0, 0, 0, 0,
	3, 59, 0, 51, 52, 53, 54, 55, 56, 57,
	13, 0, 7, 0, 62, 0, 0, 8, 64, 0,
	0, 15, 17, 22, 23, 24, 0, 0, 0, 0,
	9, 0, 0, 63, 14, 0, 0, 0, 0, 18,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 116, 117,
	39, 42, 68, 65, 66, 16, 26, 0, 28, 30,
	0, 0, 0, 19, 20, 32, 33, 34, 35, 36,
	37, 40, 41, 43, 44, 45, 46, 47, 48, 49,
	50, 25, 18, 0, 118, 11, 0, 0, 31, 0,
	121, 124, 122, 123, 29, 0, 0, 119, 60, 69,
	58, 67, 27, 120, 0, 21, 70, 0, 125, 38,
	68, 0, -2, 11, 109, 73, 0, -2, 0, 71,
	72, 74, 0, 112, 113, 0, 110, 126, 58, 0,
	0, 0, 114, 115, 77, 76, 127, 128, 0, 0,
	75, 78, 0, 24, 0, 0, 84, 85, 86, 87,
	0, 0, 90, 0, 0, 0, 0, 0, 129, 79,
	109, 0, 0, 83, 88, 0, 91, 77, 0, 0,
	0, 0, 109, 82, 89, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 81, 0, 101, 98, 77,
	97, 93, 0, 100, 0, 0, 0, 0, 0, 107,
	102, 99, 92, 94, 0, 0, 0, 106, 0, 0,
	77, 77, 108, 77, -2, 0, -2, 104, 103, 0,
	77, 0, 105,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 36, 3, 3, 3, 41, 51, 3,
	37, 73, 39, 42, 74, 43, 38, 40, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 70, 71,
	48, 56, 47, 69, 72, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 77, 3, 78, 52, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 75, 53, 76,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 44, 45, 46, 49, 50, 54,
	55, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:83
		{
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:92
		{
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:93
		{
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:97
		{
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:98
		{
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:99
		{
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:103
		{
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:107
		{
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:108
		{
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:112
		{
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:113
		{
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:118
		{
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:119
		{

		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:126
		{
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:131
		{
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:140
		{

		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:143
		{

		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:146
		{

		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:149
		{
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:159
		{

		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:163
		{

		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:166
		{
			/* dubious: expression is usually not callable? */
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:168
		{

		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:171
		{

		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:174
		{
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:175
		{
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:176
		{
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:177
		{
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:178
		{
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:179
		{
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:180
		{
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:182
		{
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:183
		{

		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:186
		{

		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:189
		{
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:190
		{

		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:193
		{

		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:196
		{

		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:199
		{

		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:202
		{
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:203
		{
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:204
		{
			/* a < b > c

			   is this typeA<typeB> c

			   or (a < b) > c ?

			   again, we must guess if an identifier is a type or an object.
			*/
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:213
		{
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:217
		{
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:218
		{
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:219
		{
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:220
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:221
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:222
		{
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:223
		{
		}
	case 60:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:231
		{

		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:240
		{
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:244
		{
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:250
		{
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:255
		{
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:256
		{

		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:266
		{
			/* ctor. Sigh. */
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:271
		{
		}
	case 75:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:272
		{
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:274
		{
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:279
		{
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:280
		{
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:284
		{
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:285
		{
			/* can't use type_specifier, since '<' is ambiguous with expression. */
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:287
		{
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:288
		{
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:289
		{
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:290
		{
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:291
		{
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:292
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:294
		{
			yylex.Error("statement error")
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:297
		{

		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:300
		{
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:339
		{
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:353
		{
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:358
		{
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:359
		{
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:363
		{
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:364
		{

		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:370
		{
		}
	case 126:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:396
		{
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:397
		{
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:401
		{
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:402
		{
		}
	}
	goto yystack /* stack new state and value */
}
