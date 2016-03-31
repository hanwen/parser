//line parser.y:19
package java

import __yyfmt__ "fmt"

//line parser.y:20
//line parser.y:24
type yySymType struct {
	yys       int
	text      string
	node      Node
	classDecl *ClassDecl
}

const ABSTRACT = 57346
const BREAK = 57347
const CASE = 57348
const CATCH = 57349
const CLASS = 57350
const DEFAULT = 57351
const ELSE = 57352
const ENUM = 57353
const EXTENDS = 57354
const FINAL = 57355
const FINALLY = 57356
const FOR = 57357
const IF = 57358
const IMPLEMENTS = 57359
const IMPORT = 57360
const INSTANCEOF = 57361
const NEW = 57362
const PACKAGE = 57363
const PRIVATE = 57364
const PROTECTED = 57365
const PUBLIC = 57366
const RETURN = 57367
const STATIC = 57368
const SUPER = 57369
const SYNCHRONIZED = 57370
const SWITCH = 57371
const THROW = 57372
const THROWS = 57373
const TRY = 57374
const WHILE = 57375
const DOT_DOT_DOT = 57376
const DOT_DOT = 57377
const STRING = 57378
const NUM = 57379
const IDENTIFIER = 57380
const SHL = 57381
const SHR = 57382
const SHRNUM = 57383
const LT = 57384
const GT = 57385
const LESS_LESS = 57386
const SLASH_EQUAL = 57387
const PLUS_EQUAL = 57388
const LESS_LESS_EQUAL = 57389
const STAR_EQUAL = 57390
const EXCLAMATION_EQUAL = 57391
const BAR_EQUAL = 57392
const AMPERSAND_EQUAL = 57393
const EQUAL_EQUAL = 57394
const LESS_EQUAL = 57395
const MORE_EQUAL = 57396
const AMPERSAND_AMPERSAND = 57397
const BAR_BAR = 57398

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
	"ENUM",
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
	"DOT_DOT_DOT",
	"DOT_DOT",
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

//line parser.y:478

var compilationUnitNode Node

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 149,
	73, 122,
	-2, 131,
	-1, 159,
	75, 124,
	-2, 11,
	-1, 273,
	6, 105,
	9, 105,
	78, 105,
	-2, 0,
	-1, 277,
	6, 104,
	9, 104,
	78, 104,
	-2, 0,
}

const yyNprod = 143
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1058

var yyAct = [...]int{

	48, 16, 191, 190, 6, 158, 94, 84, 240, 168,
	130, 10, 6, 20, 186, 157, 59, 106, 6, 260,
	6, 7, 261, 35, 36, 153, 144, 128, 128, 127,
	181, 175, 145, 20, 180, 187, 92, 275, 262, 244,
	45, 185, 6, 60, 225, 163, 96, 89, 42, 97,
	173, 85, 83, 141, 87, 88, 6, 61, 62, 17,
	131, 269, 128, 172, 6, 45, 103, 101, 128, 108,
	109, 110, 111, 112, 113, 114, 115, 116, 117, 118,
	119, 120, 122, 123, 124, 125, 128, 155, 267, 150,
	140, 258, 5, 12, 237, 32, 132, 12, 60, 230,
	215, 108, 50, 213, 18, 136, 85, 165, 53, 7,
	31, 64, 34, 139, 133, 129, 134, 30, 46, 47,
	7, 51, 49, 44, 142, 39, 28, 33, 52, 13,
	108, 7, 6, 149, 56, 24, 26, 25, 151, 27,
	148, 29, 86, 138, 154, 143, 239, 159, 91, 156,
	160, 7, 7, 105, 85, 212, 100, 80, 69, 70,
	152, 161, 64, 170, 86, 160, 169, 12, 146, 7,
	271, 174, 164, 38, 171, 85, 37, 12, 6, 254,
	169, 104, 179, 178, 86, 188, 99, 192, 6, 220,
	219, 218, 217, 209, 85, 98, 7, 211, 177, 78,
	67, 210, 68, 214, 102, 81, 72, 71, 80, 69,
	70, 222, 11, 64, 3, 221, 55, 19, 85, 159,
	253, 227, 228, 229, 147, 226, 93, 233, 231, 7,
	105, 41, 249, 192, 7, 192, 192, 160, 241, 241,
	232, 246, 247, 95, 243, 182, 137, 105, 8, 184,
	192, 192, 167, 241, 264, 6, 263, 250, 256, 252,
	245, 216, 248, 266, 259, 270, 251, 201, 197, 192,
	183, 196, 6, 195, 192, 176, 166, 268, 192, 192,
	207, 199, 273, 58, 202, 90, 57, 277, 54, 40,
	278, 63, 193, 21, 162, 204, 107, 43, 15, 50,
	9, 4, 2, 1, 194, 14, 0, 0, 206, 200,
	0, 203, 205, 0, 0, 46, 47, 7, 51, 49,
	0, 0, 79, 82, 0, 52, 0, 0, 0, 78,
	67, 199, 68, 0, 202, 81, 72, 71, 80, 69,
	70, 0, 193, 64, 0, 204, 0, 265, 0, 50,
	0, 0, 198, 0, 194, 0, 0, 279, 206, 200,
	199, 203, 205, 202, 276, 46, 47, 7, 51, 49,
	0, 193, 0, 0, 204, 52, 0, 0, 50, 0,
	0, 0, 0, 194, 0, 0, 0, 206, 200, 0,
	203, 205, 0, 0, 46, 47, 7, 51, 49, 0,
	0, 63, 198, 67, 52, 68, 0, 274, 81, 72,
	71, 80, 69, 70, 0, 0, 64, 0, 0, 0,
	0, 0, 66, 65, 75, 76, 77, 73, 74, 0,
	0, 198, 79, 82, 0, 0, 257, 0, 0, 78,
	67, 199, 68, 0, 202, 81, 72, 71, 80, 69,
	70, 0, 193, 64, 0, 204, 0, 255, 0, 50,
	0, 0, 0, 0, 194, 0, 0, 0, 206, 200,
	50, 203, 205, 0, 0, 46, 47, 7, 51, 49,
	0, 0, 63, 0, 0, 52, 46, 47, 7, 51,
	49, 0, 0, 0, 0, 0, 52, 0, 0, 0,
	121, 0, 0, 66, 65, 75, 76, 77, 73, 74,
	0, 0, 198, 79, 82, 0, 0, 238, 63, 0,
	78, 67, 0, 68, 0, 0, 81, 72, 71, 80,
	69, 70, 0, 0, 64, 0, 0, 0, 236, 66,
	65, 75, 76, 77, 73, 74, 0, 0, 0, 79,
	82, 0, 0, 0, 63, 0, 78, 67, 0, 68,
	0, 0, 81, 72, 71, 80, 69, 70, 0, 0,
	64, 0, 0, 0, 235, 66, 65, 75, 76, 77,
	73, 74, 0, 0, 0, 79, 82, 0, 0, 0,
	0, 0, 78, 67, 199, 68, 0, 202, 81, 72,
	71, 80, 69, 70, 0, 193, 64, 0, 204, 0,
	234, 0, 50, 0, 0, 0, 0, 194, 0, 0,
	0, 206, 200, 50, 203, 205, 0, 0, 46, 47,
	7, 51, 49, 0, 0, 63, 0, 0, 52, 46,
	47, 7, 51, 49, 0, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 66, 65, 75, 76,
	77, 73, 74, 0, 0, 198, 79, 82, 0, 0,
	189, 0, 0, 78, 67, 199, 68, 0, 202, 81,
	72, 71, 80, 69, 70, 0, 193, 64, 0, 204,
	0, 126, 0, 50, 0, 0, 0, 0, 194, 0,
	0, 0, 206, 200, 0, 203, 205, 0, 0, 46,
	47, 7, 51, 49, 0, 0, 63, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 65, 75,
	76, 77, 73, 74, 0, 0, 198, 79, 82, 0,
	242, 63, 0, 0, 78, 67, 0, 68, 0, 0,
	81, 72, 71, 80, 69, 70, 0, 0, 64, 0,
	224, 0, 66, 65, 75, 76, 77, 73, 74, 0,
	0, 0, 79, 82, 0, 63, 0, 0, 0, 78,
	67, 0, 68, 0, 0, 81, 72, 71, 80, 69,
	70, 0, 0, 64, 0, 223, 66, 65, 75, 76,
	77, 73, 74, 0, 0, 0, 79, 82, 0, 63,
	0, 0, 0, 78, 67, 0, 68, 0, 0, 81,
	72, 71, 80, 69, 70, 0, 0, 64, 0, 208,
	66, 65, 75, 76, 77, 73, 74, 0, 0, 0,
	79, 82, 63, 0, 0, 0, 0, 78, 67, 0,
	68, 0, 0, 81, 72, 71, 80, 69, 70, 0,
	0, 64, 272, 66, 65, 75, 76, 77, 73, 74,
	0, 0, 0, 79, 82, 63, 0, 0, 0, 0,
	78, 67, 0, 68, 0, 0, 81, 72, 71, 80,
	69, 70, 0, 0, 64, 135, 66, 65, 75, 76,
	77, 73, 74, 0, 0, 0, 79, 82, 0, 0,
	0, 0, 0, 78, 67, 199, 68, 0, 202, 81,
	72, 71, 80, 69, 70, 0, 193, 64, 0, 204,
	63, 0, 0, 50, 0, 0, 0, 0, 194, 0,
	0, 0, 206, 200, 0, 203, 205, 0, 0, 46,
	47, 7, 51, 49, 0, 0, 73, 74, 0, 52,
	0, 79, 82, 0, 0, 0, 0, 0, 78, 67,
	0, 68, 0, 0, 81, 72, 71, 80, 69, 70,
	30, 0, 64, 0, 22, 0, 198, 23, 30, 28,
	0, 0, 22, 0, 0, 23, 0, 28, 24, 26,
	25, 0, 27, 0, 29, 0, 24, 26, 25, 0,
	27, 0, 29, 0, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 86,
}
var yyPact = [...]int{

	193, -1000, -1000, 158, 194, 56, -1000, -1000, -1000, -1000,
	-15, 191, 158, -1000, -1000, -1000, 994, 158, 54, 158,
	-1000, -1000, 158, 158, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 136, 131, -1000, 52, 219, -29, 603, 35, -1000,
	199, 158, 158, -18, -1000, 866, -1000, -1000, -1000, 603,
	71, 603, 603, -1000, -30, 158, 126, -42, -27, -1000,
	155, -1000, 603, 158, 603, 196, 603, 603, 603, 603,
	603, 603, 603, 603, 603, 603, 603, 603, 603, 450,
	603, 603, 603, 616, -11, 102, -1000, 866, 272, -1000,
	-16, 126, -1000, -1000, -15, -1000, -1000, 158, 603, -1000,
	-1000, 833, -1000, -1000, 158, 93, 15, -23, 866, 91,
	91, 40, 40, 91, 91, 272, 272, 921, 921, 921,
	345, 603, 142, 40, 91, 142, -1000, 603, -54, -1000,
	-46, 158, 986, -1000, 14, 603, -1000, 110, -1000, -51,
	-1000, 603, 142, 12, -1000, -1000, -1000, 126, -17, 179,
	-1000, -1000, -1000, 71, 866, -32, 107, -1000, 34, -1000,
	603, -52, -1000, -1000, -1000, -1000, -12, -26, -1000, -15,
	866, -47, 167, -1000, 113, -1000, -43, 158, -1000, 7,
	-1000, -1000, -41, 126, 158, -1000, 592, 158, -1000, -1000,
	-1000, 766, 158, 71, 82, -1000, -1000, -1000, -1000, 30,
	603, -1000, 27, 152, 151, 150, 149, 126, -1000, 92,
	-17, 732, -1000, -1000, 697, -1000, -33, 71, 603, 603,
	603, 26, 92, -1000, -1000, -1000, -17, 535, 499, 463,
	-1000, 21, 439, 88, 673, 673, -38, -1000, 234, 603,
	222, -1000, -1000, -1000, -1000, 206, 139, 382, -1000, 673,
	358, 13, -1000, -39, 158, -1000, -1000, -1000, -1000, -1000,
	603, 16, -1000, -14, 114, 126, 800, -1000, 329, -40,
	-1000, 158, -1000, 923, -1000, -1000, 126, 923, 279, -1000,
}
var yyPgo = [...]int{

	0, 305, 10, 243, 0, 168, 7, 303, 302, 301,
	92, 300, 6, 298, 297, 123, 2, 17, 296, 115,
	294, 293, 1, 289, 288, 286, 285, 283, 16, 15,
	5, 276, 275, 14, 3, 273, 271, 268, 267, 266,
	264, 8, 262, 261, 260, 259, 256, 254, 252, 9,
	249, 246, 245,
}
var yyR1 = [...]int{

	0, 7, 10, 10, 8, 9, 9, 11, 11, 11,
	3, 12, 12, 13, 13, 14, 14, 15, 17, 17,
	18, 18, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 21, 21, 21, 21, 21, 21, 21, 22,
	22, 1, 1, 4, 23, 23, 24, 24, 26, 26,
	2, 2, 20, 20, 25, 25, 27, 27, 28, 28,
	5, 5, 5, 29, 29, 29, 33, 33, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 38, 39, 39, 40, 40, 37, 41, 41, 36,
	42, 42, 35, 44, 44, 43, 43, 45, 45, 46,
	47, 47, 30, 30, 31, 31, 48, 48, 49, 50,
	50, 6, 6, 6, 6, 19, 19, 51, 51, 32,
	32, 52, 52,
}
var yyR2 = [...]int{

	0, 3, 1, 3, 3, 0, 2, 3, 4, 5,
	2, 0, 2, 2, 5, 1, 3, 1, 0, 1,
	1, 3, 1, 1, 1, 3, 3, 5, 3, 4,
	3, 4, 3, 3, 3, 3, 3, 3, 6, 2,
	3, 3, 2, 3, 3, 3, 3, 4, 3, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 0,
	2, 8, 6, 1, 0, 2, 0, 2, 1, 3,
	0, 2, 0, 3, 2, 2, 1, 3, 1, 4,
	5, 4, 1, 2, 7, 5, 0, 2, 2, 4,
	5, 3, 2, 1, 1, 1, 1, 2, 3, 1,
	2, 7, 0, 2, 4, 3, 5, 1, 3, 6,
	0, 2, 7, 0, 7, 0, 6, 0, 4, 2,
	1, 3, 0, 2, 0, 1, 1, 3, 5, 0,
	1, 1, 1, 2, 3, 3, 2, 1, 3, 0,
	2, 1, 3,
}
var yyChk = [...]int{

	-1000, -7, -8, 21, -9, -10, -4, 38, -3, -11,
	-12, 18, 41, 73, -1, -13, -22, 74, -10, 26,
	-4, -21, 8, 11, 22, 24, 23, 26, 13, 28,
	4, -10, 41, 73, -10, -4, -4, 40, 42, 73,
	-23, 12, 77, -14, -15, -16, 36, 37, -4, 40,
	20, 39, 46, 73, -24, 17, -10, -25, -27, -28,
	-4, 75, 76, 19, 71, 41, 40, 58, 60, 67,
	68, 65, 64, 45, 46, 42, 43, 44, 57, 50,
	66, 63, 51, -16, -6, -4, 71, -16, -16, 77,
	-26, -10, 78, -5, -12, -3, 73, 76, 40, -15,
	-10, -16, 8, -4, -19, 51, -17, -18, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, 50, -16, -16, -16, -16, 75, 40, 79, -19,
	-2, 76, -22, -28, -17, 72, -4, -51, 50, -6,
	75, 76, -16, -17, 80, 78, -5, -10, -6, -4,
	75, -16, 50, 76, -16, 75, -4, -29, -30, 40,
	58, -6, -20, 77, -29, 73, -31, -48, -49, -12,
	-16, -2, 75, 76, -22, 78, -32, 31, -49, -6,
	77, 73, -52, -10, -50, 34, -33, 76, -4, 78,
	-34, -16, -4, 13, 25, -35, -36, -37, 73, 2,
	30, -38, 5, 32, 16, 33, 29, -10, 73, -4,
	-6, -16, 73, 73, -16, 73, -43, 40, 40, 40,
	40, -30, -4, 73, 73, 77, -6, -16, -16, -16,
	73, -30, -33, -4, 75, 75, 75, 73, 78, 58,
	-41, -34, 77, -41, 77, -44, 7, -16, -42, 10,
	-33, -39, -45, 14, 40, 75, -41, 78, 78, -40,
	6, 9, 77, -46, -47, -10, -16, 72, -33, 75,
	-4, 56, 72, -33, 78, 77, -10, -33, -33, 78,
}
var yyDef = [...]int{

	0, -2, 5, 0, 11, 0, 2, 63, 1, 6,
	59, 0, 0, 4, 10, 12, 0, 0, 0, 0,
	3, 60, 0, 0, 52, 53, 54, 55, 56, 57,
	58, 13, 0, 7, 0, 64, 0, 0, 0, 8,
	66, 0, 0, 0, 15, 17, 22, 23, 24, 0,
	0, 0, 0, 9, 0, 0, 65, 11, 0, 76,
	78, 14, 0, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 131, 132, 39, 42, 70,
	67, 68, 62, 75, 59, 82, 74, 0, 18, 16,
	26, 0, 28, 30, 0, 0, 0, 19, 20, 32,
	33, 34, 35, 36, 37, 40, 41, 43, 44, 45,
	46, 0, 51, 48, 49, 50, 25, 18, 0, 133,
	11, 0, 0, 77, 0, 0, 31, 0, 136, 137,
	29, 0, 47, 0, 134, 61, 71, 69, 0, -2,
	79, 27, 135, 0, 21, 72, 122, 81, 0, -2,
	0, 138, 38, 70, 80, 83, 0, 125, 126, 59,
	123, 11, 139, 11, 0, 73, 0, 0, 127, 129,
	86, 85, 140, 141, 0, 130, 0, 0, 128, 84,
	87, 0, 24, 0, 0, 93, 94, 95, 96, 0,
	0, 99, 0, 115, 0, 0, 0, 142, 88, 122,
	0, 0, 92, 97, 0, 100, 0, 0, 0, 0,
	0, 0, 122, 91, 98, 86, 0, 0, 0, 0,
	89, 0, 0, 0, 0, 0, 0, 90, 113, 0,
	110, 107, 86, 106, 102, 117, 0, 0, 109, 0,
	0, 0, 112, 0, 0, 116, 111, 108, 101, 103,
	0, 0, 86, 0, 0, 120, 0, 86, 0, 0,
	119, 0, 86, -2, 118, 86, 121, -2, 0, 114,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 39, 3, 3, 3, 44, 54, 3,
	40, 75, 42, 45, 76, 46, 41, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 72, 73,
	51, 58, 50, 71, 74, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 79, 3, 80, 55, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 77, 56, 78,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 47, 48, 49,
	52, 53, 57, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 69, 70,
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

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:93
		{
			compilationUnitNode = yyDollar[3].classDecl
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:108
		{
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:117
		{
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:118
		{
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:122
		{
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:123
		{
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:124
		{
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:128
		{
			yyVAL.classDecl = yyDollar[2].classDecl
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:134
		{
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:135
		{
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:139
		{
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:140
		{
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:145
		{
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:146
		{

		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:153
		{
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:158
		{
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:167
		{

		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:170
		{

		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:173
		{

		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:176
		{
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:186
		{

		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:190
		{

		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:193
		{
			/* dubious: expression is usually not callable? */
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:195
		{

		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:198
		{

		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:201
		{
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:202
		{
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:203
		{
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:204
		{
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:205
		{
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:206
		{
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:207
		{
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:209
		{
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:210
		{

		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:213
		{

		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:216
		{
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:217
		{

		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:220
		{

		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:223
		{

		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:226
		{

		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:229
		{
			/* WTF:  >> is not a token.
			private Map<AnyObjectId, Set<Ref>> refsById;
			*/
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:233
		{
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:234
		{
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:235
		{
			/* a < b > c

			   is this typeA<typeB> c

			   or (a < b) > c ?

			   again, we must guess if an identifier is a type or an object.
			*/
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:244
		{
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:250
		{
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:251
		{
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:252
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:253
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:254
		{
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:255
		{
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:256
		{
		}
	case 61:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:264
		{
			yyVAL.classDecl = yyDollar[7].classDecl
			yyVAL.classDecl.Decl = Decl{
				Type: "class",
				Name: yyDollar[3].text,
			}
		}
	case 62:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:271
		{
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:275
		{
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:279
		{
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:283
		{
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:289
		{
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:294
		{
			yyVAL.classDecl = &ClassDecl{}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:295
		{
			yyVAL.classDecl.Members = append(yyVAL.classDecl.Members, yyDollar[2].node)
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:319
		{
			yyVAL.node = &Decl{
				Name: yyDollar[4].text,
				Type: yyDollar[3].text,
			}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:325
		{
			/* ctor. Sigh. */
			yyVAL.node = &Decl{
				Name: yyDollar[3].text,
				Type: "void",
			}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:331
		{
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:336
		{
		}
	case 84:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:337
		{
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:339
		{
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:344
		{
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:345
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:349
		{
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:350
		{
			/* can't use type_specifier, since '<' is ambiguous with expression. */
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:352
		{
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:353
		{
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:354
		{
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:355
		{
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:356
		{
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:357
		{
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
		{
			yylex.Error("statement error")
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:362
		{

		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:365
		{
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:412
		{
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:426
		{
		}
	case 124:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:431
		{
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:432
		{
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:436
		{
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:437
		{

		}
	case 128:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:443
		{
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:452
		{
			yyVAL.text = "?"
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:453
		{
			yyVAL.text = yyDollar[1].text
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:454
		{
			yyVAL.text = yyDollar[1].text + "[]"
		}
	case 139:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:468
		{
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:469
		{
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:473
		{
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:474
		{
		}
	}
	goto yystack /* stack new state and value */
}
