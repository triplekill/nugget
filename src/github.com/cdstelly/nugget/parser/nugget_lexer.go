// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// import "../NTypes"

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 47, 386,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 330, 10, 39, 3, 40,
	3, 40, 3, 40, 3, 41, 6, 41, 336, 10, 41, 13, 41, 14, 41, 337, 3, 42, 6,
	42, 341, 10, 42, 13, 42, 14, 42, 342, 3, 42, 3, 42, 6, 42, 347, 10, 42,
	13, 42, 14, 42, 348, 5, 42, 351, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 7,
	43, 357, 10, 43, 12, 43, 14, 43, 360, 11, 43, 3, 43, 3, 43, 3, 44, 6, 44,
	365, 10, 44, 13, 44, 14, 44, 366, 3, 44, 3, 44, 3, 45, 5, 45, 372, 10,
	45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 7, 46, 380, 10, 46, 12, 46,
	14, 46, 383, 11, 46, 3, 46, 3, 46, 2, 2, 47, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43,
	85, 44, 87, 45, 89, 46, 91, 47, 3, 2, 8, 4, 2, 62, 62, 64, 64, 3, 2, 50,
	59, 4, 2, 67, 92, 99, 124, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34,
	4, 2, 12, 12, 15, 15, 2, 397, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53,
	3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2,
	61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2,
	2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2,
	2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2,
	2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3,
	2, 2, 2, 3, 93, 3, 2, 2, 2, 5, 100, 3, 2, 2, 2, 7, 102, 3, 2, 2, 2, 9,
	104, 3, 2, 2, 2, 11, 106, 3, 2, 2, 2, 13, 108, 3, 2, 2, 2, 15, 113, 3,
	2, 2, 2, 17, 119, 3, 2, 2, 2, 19, 124, 3, 2, 2, 2, 21, 130, 3, 2, 2, 2,
	23, 137, 3, 2, 2, 2, 25, 141, 3, 2, 2, 2, 27, 148, 3, 2, 2, 2, 29, 153,
	3, 2, 2, 2, 31, 157, 3, 2, 2, 2, 33, 162, 3, 2, 2, 2, 35, 167, 3, 2, 2,
	2, 37, 174, 3, 2, 2, 2, 39, 179, 3, 2, 2, 2, 41, 188, 3, 2, 2, 2, 43, 197,
	3, 2, 2, 2, 45, 204, 3, 2, 2, 2, 47, 209, 3, 2, 2, 2, 49, 220, 3, 2, 2,
	2, 51, 232, 3, 2, 2, 2, 53, 246, 3, 2, 2, 2, 55, 254, 3, 2, 2, 2, 57, 259,
	3, 2, 2, 2, 59, 266, 3, 2, 2, 2, 61, 281, 3, 2, 2, 2, 63, 290, 3, 2, 2,
	2, 65, 296, 3, 2, 2, 2, 67, 303, 3, 2, 2, 2, 69, 307, 3, 2, 2, 2, 71, 310,
	3, 2, 2, 2, 73, 313, 3, 2, 2, 2, 75, 315, 3, 2, 2, 2, 77, 329, 3, 2, 2,
	2, 79, 331, 3, 2, 2, 2, 81, 335, 3, 2, 2, 2, 83, 340, 3, 2, 2, 2, 85, 352,
	3, 2, 2, 2, 87, 364, 3, 2, 2, 2, 89, 371, 3, 2, 2, 2, 91, 375, 3, 2, 2,
	2, 93, 94, 7, 118, 2, 2, 94, 95, 7, 119, 2, 2, 95, 96, 7, 114, 2, 2, 96,
	97, 7, 110, 2, 2, 97, 98, 7, 103, 2, 2, 98, 99, 7, 93, 2, 2, 99, 4, 3,
	2, 2, 2, 100, 101, 7, 46, 2, 2, 101, 6, 3, 2, 2, 2, 102, 103, 7, 95, 2,
	2, 103, 8, 3, 2, 2, 2, 104, 105, 7, 63, 2, 2, 105, 10, 3, 2, 2, 2, 106,
	107, 7, 126, 2, 2, 107, 12, 3, 2, 2, 2, 108, 109, 7, 118, 2, 2, 109, 110,
	7, 123, 2, 2, 110, 111, 7, 114, 2, 2, 111, 112, 7, 103, 2, 2, 112, 14,
	3, 2, 2, 2, 113, 114, 7, 114, 2, 2, 114, 115, 7, 116, 2, 2, 115, 116, 7,
	107, 2, 2, 116, 117, 7, 112, 2, 2, 117, 118, 7, 118, 2, 2, 118, 16, 3,
	2, 2, 2, 119, 120, 7, 117, 2, 2, 120, 121, 7, 107, 2, 2, 121, 122, 7, 124,
	2, 2, 122, 123, 7, 103, 2, 2, 123, 18, 3, 2, 2, 2, 124, 125, 7, 118, 2,
	2, 125, 126, 7, 123, 2, 2, 126, 127, 7, 114, 2, 2, 127, 128, 7, 103, 2,
	2, 128, 129, 7, 122, 2, 2, 129, 20, 3, 2, 2, 2, 130, 131, 7, 114, 2, 2,
	131, 132, 7, 116, 2, 2, 132, 133, 7, 107, 2, 2, 133, 134, 7, 112, 2, 2,
	134, 135, 7, 118, 2, 2, 135, 136, 7, 122, 2, 2, 136, 22, 3, 2, 2, 2, 137,
	138, 7, 116, 2, 2, 138, 139, 7, 99, 2, 2, 139, 140, 7, 121, 2, 2, 140,
	24, 3, 2, 2, 2, 141, 142, 7, 117, 2, 2, 142, 143, 7, 118, 2, 2, 143, 144,
	7, 116, 2, 2, 144, 145, 7, 107, 2, 2, 145, 146, 7, 112, 2, 2, 146, 147,
	7, 105, 2, 2, 147, 26, 3, 2, 2, 2, 148, 149, 7, 117, 2, 2, 149, 150, 7,
	106, 2, 2, 150, 151, 7, 99, 2, 2, 151, 152, 7, 51, 2, 2, 152, 28, 3, 2,
	2, 2, 153, 154, 7, 111, 2, 2, 154, 155, 7, 102, 2, 2, 155, 156, 7, 55,
	2, 2, 156, 30, 3, 2, 2, 2, 157, 158, 7, 112, 2, 2, 158, 159, 7, 118, 2,
	2, 159, 160, 7, 104, 2, 2, 160, 161, 7, 117, 2, 2, 161, 32, 3, 2, 2, 2,
	162, 163, 7, 104, 2, 2, 163, 164, 7, 107, 2, 2, 164, 165, 7, 110, 2, 2,
	165, 166, 7, 103, 2, 2, 166, 34, 3, 2, 2, 2, 167, 168, 7, 114, 2, 2, 168,
	169, 7, 99, 2, 2, 169, 170, 7, 101, 2, 2, 170, 171, 7, 109, 2, 2, 171,
	172, 7, 103, 2, 2, 172, 173, 7, 118, 2, 2, 173, 36, 3, 2, 2, 2, 174, 175,
	7, 114, 2, 2, 175, 176, 7, 101, 2, 2, 176, 177, 7, 99, 2, 2, 177, 178,
	7, 114, 2, 2, 178, 38, 3, 2, 2, 2, 179, 180, 7, 103, 2, 2, 180, 181, 7,
	122, 2, 2, 181, 182, 7, 107, 2, 2, 182, 183, 7, 104, 2, 2, 183, 184, 7,
	107, 2, 2, 184, 185, 7, 112, 2, 2, 185, 186, 7, 104, 2, 2, 186, 187, 7,
	113, 2, 2, 187, 40, 3, 2, 2, 2, 188, 189, 7, 102, 2, 2, 189, 190, 7, 99,
	2, 2, 190, 191, 7, 118, 2, 2, 191, 192, 7, 103, 2, 2, 192, 193, 7, 118,
	2, 2, 193, 194, 7, 107, 2, 2, 194, 195, 7, 111, 2, 2, 195, 196, 7, 103,
	2, 2, 196, 42, 3, 2, 2, 2, 197, 198, 7, 111, 2, 2, 198, 199, 7, 103, 2,
	2, 199, 200, 7, 111, 2, 2, 200, 201, 7, 113, 2, 2, 201, 202, 7, 116, 2,
	2, 202, 203, 7, 123, 2, 2, 203, 44, 3, 2, 2, 2, 204, 205, 7, 106, 2, 2,
	205, 206, 7, 118, 2, 2, 206, 207, 7, 118, 2, 2, 207, 208, 7, 114, 2, 2,
	208, 46, 3, 2, 2, 2, 209, 210, 7, 110, 2, 2, 210, 211, 7, 107, 2, 2, 211,
	212, 7, 117, 2, 2, 212, 213, 7, 118, 2, 2, 213, 214, 7, 113, 2, 2, 214,
	215, 7, 104, 2, 2, 215, 216, 7, 47, 2, 2, 216, 217, 7, 111, 2, 2, 217,
	218, 7, 102, 2, 2, 218, 219, 7, 55, 2, 2, 219, 48, 3, 2, 2, 2, 220, 221,
	7, 110, 2, 2, 221, 222, 7, 107, 2, 2, 222, 223, 7, 117, 2, 2, 223, 224,
	7, 118, 2, 2, 224, 225, 7, 113, 2, 2, 225, 226, 7, 104, 2, 2, 226, 227,
	7, 47, 2, 2, 227, 228, 7, 117, 2, 2, 228, 229, 7, 106, 2, 2, 229, 230,
	7, 99, 2, 2, 230, 231, 7, 51, 2, 2, 231, 50, 3, 2, 2, 2, 232, 233, 7, 110,
	2, 2, 233, 234, 7, 107, 2, 2, 234, 235, 7, 117, 2, 2, 235, 236, 7, 118,
	2, 2, 236, 237, 7, 113, 2, 2, 237, 238, 7, 104, 2, 2, 238, 239, 7, 47,
	2, 2, 239, 240, 7, 117, 2, 2, 240, 241, 7, 106, 2, 2, 241, 242, 7, 99,
	2, 2, 242, 243, 7, 52, 2, 2, 243, 244, 7, 55, 2, 2, 244, 245, 7, 56, 2,
	2, 245, 52, 3, 2, 2, 2, 246, 247, 7, 103, 2, 2, 247, 248, 7, 122, 2, 2,
	248, 249, 7, 118, 2, 2, 249, 250, 7, 116, 2, 2, 250, 251, 7, 99, 2, 2,
	251, 252, 7, 101, 2, 2, 252, 253, 7, 118, 2, 2, 253, 54, 3, 2, 2, 2, 254,
	255, 7, 117, 2, 2, 255, 256, 7, 113, 2, 2, 256, 257, 7, 116, 2, 2, 257,
	258, 7, 118, 2, 2, 258, 56, 3, 2, 2, 2, 259, 260, 7, 117, 2, 2, 260, 261,
	7, 106, 2, 2, 261, 262, 7, 99, 2, 2, 262, 263, 7, 52, 2, 2, 263, 264, 7,
	55, 2, 2, 264, 265, 7, 56, 2, 2, 265, 58, 3, 2, 2, 2, 266, 267, 7, 105,
	2, 2, 267, 268, 7, 103, 2, 2, 268, 269, 7, 118, 2, 2, 269, 270, 7, 73,
	2, 2, 270, 271, 7, 103, 2, 2, 271, 272, 7, 118, 2, 2, 272, 273, 7, 84,
	2, 2, 273, 274, 7, 103, 2, 2, 274, 275, 7, 115, 2, 2, 275, 276, 7, 119,
	2, 2, 276, 277, 7, 103, 2, 2, 277, 278, 7, 117, 2, 2, 278, 279, 7, 118,
	2, 2, 279, 280, 7, 117, 2, 2, 280, 60, 3, 2, 2, 2, 281, 282, 7, 102, 2,
	2, 282, 283, 7, 107, 2, 2, 283, 284, 7, 117, 2, 2, 284, 285, 7, 109, 2,
	2, 285, 286, 7, 107, 2, 2, 286, 287, 7, 112, 2, 2, 287, 288, 7, 104, 2,
	2, 288, 289, 7, 113, 2, 2, 289, 62, 3, 2, 2, 2, 290, 291, 7, 119, 2, 2,
	291, 292, 7, 112, 2, 2, 292, 293, 7, 107, 2, 2, 293, 294, 7, 113, 2, 2,
	294, 295, 7, 112, 2, 2, 295, 64, 3, 2, 2, 2, 296, 297, 7, 114, 2, 2, 297,
	298, 7, 117, 2, 2, 298, 299, 7, 110, 2, 2, 299, 300, 7, 107, 2, 2, 300,
	301, 7, 117, 2, 2, 301, 302, 7, 118, 2, 2, 302, 66, 3, 2, 2, 2, 303, 304,
	7, 39, 2, 2, 304, 305, 7, 39, 2, 2, 305, 306, 7, 39, 2, 2, 306, 68, 3,
	2, 2, 2, 307, 308, 7, 99, 2, 2, 308, 309, 7, 117, 2, 2, 309, 70, 3, 2,
	2, 2, 310, 311, 7, 100, 2, 2, 311, 312, 7, 123, 2, 2, 312, 72, 3, 2, 2,
	2, 313, 314, 7, 93, 2, 2, 314, 74, 3, 2, 2, 2, 315, 316, 7, 104, 2, 2,
	316, 317, 7, 107, 2, 2, 317, 318, 7, 110, 2, 2, 318, 319, 7, 118, 2, 2,
	319, 320, 7, 103, 2, 2, 320, 321, 7, 116, 2, 2, 321, 76, 3, 2, 2, 2, 322,
	330, 9, 2, 2, 2, 323, 324, 7, 64, 2, 2, 324, 330, 7, 63, 2, 2, 325, 326,
	7, 62, 2, 2, 326, 330, 7, 63, 2, 2, 327, 328, 7, 63, 2, 2, 328, 330, 7,
	63, 2, 2, 329, 322, 3, 2, 2, 2, 329, 323, 3, 2, 2, 2, 329, 325, 3, 2, 2,
	2, 329, 327, 3, 2, 2, 2, 330, 78, 3, 2, 2, 2, 331, 332, 7, 93, 2, 2, 332,
	333, 7, 95, 2, 2, 333, 80, 3, 2, 2, 2, 334, 336, 9, 3, 2, 2, 335, 334,
	3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 337, 338, 3, 2,
	2, 2, 338, 82, 3, 2, 2, 2, 339, 341, 9, 4, 2, 2, 340, 339, 3, 2, 2, 2,
	341, 342, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343,
	350, 3, 2, 2, 2, 344, 346, 7, 48, 2, 2, 345, 347, 9, 4, 2, 2, 346, 345,
	3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 349, 3, 2,
	2, 2, 349, 351, 3, 2, 2, 2, 350, 344, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2,
	351, 84, 3, 2, 2, 2, 352, 358, 7, 36, 2, 2, 353, 354, 7, 36, 2, 2, 354,
	357, 7, 36, 2, 2, 355, 357, 10, 5, 2, 2, 356, 353, 3, 2, 2, 2, 356, 355,
	3, 2, 2, 2, 357, 360, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2,
	2, 2, 359, 361, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 362, 7, 36, 2, 2,
	362, 86, 3, 2, 2, 2, 363, 365, 9, 6, 2, 2, 364, 363, 3, 2, 2, 2, 365, 366,
	3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 368, 3, 2,
	2, 2, 368, 369, 8, 44, 2, 2, 369, 88, 3, 2, 2, 2, 370, 372, 7, 15, 2, 2,
	371, 370, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373,
	374, 7, 12, 2, 2, 374, 90, 3, 2, 2, 2, 375, 376, 7, 49, 2, 2, 376, 377,
	7, 49, 2, 2, 377, 381, 3, 2, 2, 2, 378, 380, 10, 7, 2, 2, 379, 378, 3,
	2, 2, 2, 380, 383, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2,
	2, 382, 384, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 384, 385, 8, 46, 2, 2, 385,
	92, 3, 2, 2, 2, 13, 2, 329, 337, 342, 348, 350, 356, 358, 366, 371, 381,
	3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'tuple['", "','", "']'", "'='", "'|'", "'type'", "'print'", "'size'",
	"'typex'", "'printx'", "'raw'", "'string'", "'sha1'", "'md5'", "'ntfs'",
	"'file'", "'packet'", "'pcap'", "'exifinfo'", "'datetime'", "'memory'",
	"'http'", "'listof-md5'", "'listof-sha1'", "'listof-sha256'", "'extract'",
	"'sort'", "'sha256'", "'getGetRequests'", "'diskinfo'", "'union'", "'pslist'",
	"'%%%'", "'as'", "'by'", "'['", "'filter'", "", "'[]'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS", "NL", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "COMPOP", "LISTOP", "INT", "ID", "STRING",
	"WS", "NL", "LINE_COMMENT",
}

type NuggetLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNuggetLexer(input antlr.CharStream) *NuggetLexer {

	l := new(NuggetLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Nugget.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NuggetLexer tokens.
const (
	NuggetLexerT__0         = 1
	NuggetLexerT__1         = 2
	NuggetLexerT__2         = 3
	NuggetLexerT__3         = 4
	NuggetLexerT__4         = 5
	NuggetLexerT__5         = 6
	NuggetLexerT__6         = 7
	NuggetLexerT__7         = 8
	NuggetLexerT__8         = 9
	NuggetLexerT__9         = 10
	NuggetLexerT__10        = 11
	NuggetLexerT__11        = 12
	NuggetLexerT__12        = 13
	NuggetLexerT__13        = 14
	NuggetLexerT__14        = 15
	NuggetLexerT__15        = 16
	NuggetLexerT__16        = 17
	NuggetLexerT__17        = 18
	NuggetLexerT__18        = 19
	NuggetLexerT__19        = 20
	NuggetLexerT__20        = 21
	NuggetLexerT__21        = 22
	NuggetLexerT__22        = 23
	NuggetLexerT__23        = 24
	NuggetLexerT__24        = 25
	NuggetLexerT__25        = 26
	NuggetLexerT__26        = 27
	NuggetLexerT__27        = 28
	NuggetLexerT__28        = 29
	NuggetLexerT__29        = 30
	NuggetLexerT__30        = 31
	NuggetLexerT__31        = 32
	NuggetLexerT__32        = 33
	NuggetLexerT__33        = 34
	NuggetLexerT__34        = 35
	NuggetLexerT__35        = 36
	NuggetLexerT__36        = 37
	NuggetLexerCOMPOP       = 38
	NuggetLexerLISTOP       = 39
	NuggetLexerINT          = 40
	NuggetLexerID           = 41
	NuggetLexerSTRING       = 42
	NuggetLexerWS           = 43
	NuggetLexerNL           = 44
	NuggetLexerLINE_COMMENT = 45
)
