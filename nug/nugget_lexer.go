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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 48, 405,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 5, 40, 357, 10, 40, 3, 41, 3, 41, 3, 41, 3, 42, 6, 42,
	363, 10, 42, 13, 42, 14, 42, 364, 3, 43, 6, 43, 368, 10, 43, 13, 43, 14,
	43, 369, 3, 44, 3, 44, 3, 44, 3, 44, 7, 44, 376, 10, 44, 12, 44, 14, 44,
	379, 11, 44, 3, 44, 3, 44, 3, 45, 6, 45, 384, 10, 45, 13, 45, 14, 45, 385,
	3, 45, 3, 45, 3, 46, 5, 46, 391, 10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 47, 7, 47, 399, 10, 47, 12, 47, 14, 47, 402, 11, 47, 3, 47, 3, 47,
	2, 2, 48, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38,
	75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47,
	93, 48, 3, 2, 8, 4, 2, 62, 62, 64, 64, 3, 2, 50, 59, 4, 2, 67, 92, 99,
	124, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15,
	2, 414, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3,
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63,
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2,
	71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2,
	2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2,
	2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2,
	2, 2, 3, 95, 3, 2, 2, 2, 5, 102, 3, 2, 2, 2, 7, 104, 3, 2, 2, 2, 9, 106,
	3, 2, 2, 2, 11, 108, 3, 2, 2, 2, 13, 110, 3, 2, 2, 2, 15, 112, 3, 2, 2,
	2, 17, 114, 3, 2, 2, 2, 19, 119, 3, 2, 2, 2, 21, 125, 3, 2, 2, 2, 23, 130,
	3, 2, 2, 2, 25, 136, 3, 2, 2, 2, 27, 143, 3, 2, 2, 2, 29, 147, 3, 2, 2,
	2, 31, 154, 3, 2, 2, 2, 33, 159, 3, 2, 2, 2, 35, 163, 3, 2, 2, 2, 37, 168,
	3, 2, 2, 2, 39, 173, 3, 2, 2, 2, 41, 180, 3, 2, 2, 2, 43, 185, 3, 2, 2,
	2, 45, 194, 3, 2, 2, 2, 47, 203, 3, 2, 2, 2, 49, 210, 3, 2, 2, 2, 51, 221,
	3, 2, 2, 2, 53, 233, 3, 2, 2, 2, 55, 247, 3, 2, 2, 2, 57, 255, 3, 2, 2,
	2, 59, 260, 3, 2, 2, 2, 61, 267, 3, 2, 2, 2, 63, 282, 3, 2, 2, 2, 65, 291,
	3, 2, 2, 2, 67, 297, 3, 2, 2, 2, 69, 304, 3, 2, 2, 2, 71, 309, 3, 2, 2,
	2, 73, 336, 3, 2, 2, 2, 75, 339, 3, 2, 2, 2, 77, 342, 3, 2, 2, 2, 79, 356,
	3, 2, 2, 2, 81, 358, 3, 2, 2, 2, 83, 362, 3, 2, 2, 2, 85, 367, 3, 2, 2,
	2, 87, 371, 3, 2, 2, 2, 89, 383, 3, 2, 2, 2, 91, 390, 3, 2, 2, 2, 93, 394,
	3, 2, 2, 2, 95, 96, 7, 118, 2, 2, 96, 97, 7, 119, 2, 2, 97, 98, 7, 114,
	2, 2, 98, 99, 7, 110, 2, 2, 99, 100, 7, 103, 2, 2, 100, 101, 7, 93, 2,
	2, 101, 4, 3, 2, 2, 2, 102, 103, 7, 46, 2, 2, 103, 6, 3, 2, 2, 2, 104,
	105, 7, 95, 2, 2, 105, 8, 3, 2, 2, 2, 106, 107, 7, 63, 2, 2, 107, 10, 3,
	2, 2, 2, 108, 109, 7, 126, 2, 2, 109, 12, 3, 2, 2, 2, 110, 111, 7, 42,
	2, 2, 111, 14, 3, 2, 2, 2, 112, 113, 7, 43, 2, 2, 113, 16, 3, 2, 2, 2,
	114, 115, 7, 118, 2, 2, 115, 116, 7, 123, 2, 2, 116, 117, 7, 114, 2, 2,
	117, 118, 7, 103, 2, 2, 118, 18, 3, 2, 2, 2, 119, 120, 7, 114, 2, 2, 120,
	121, 7, 116, 2, 2, 121, 122, 7, 107, 2, 2, 122, 123, 7, 112, 2, 2, 123,
	124, 7, 118, 2, 2, 124, 20, 3, 2, 2, 2, 125, 126, 7, 117, 2, 2, 126, 127,
	7, 107, 2, 2, 127, 128, 7, 124, 2, 2, 128, 129, 7, 103, 2, 2, 129, 22,
	3, 2, 2, 2, 130, 131, 7, 118, 2, 2, 131, 132, 7, 123, 2, 2, 132, 133, 7,
	114, 2, 2, 133, 134, 7, 103, 2, 2, 134, 135, 7, 122, 2, 2, 135, 24, 3,
	2, 2, 2, 136, 137, 7, 114, 2, 2, 137, 138, 7, 116, 2, 2, 138, 139, 7, 107,
	2, 2, 139, 140, 7, 112, 2, 2, 140, 141, 7, 118, 2, 2, 141, 142, 7, 122,
	2, 2, 142, 26, 3, 2, 2, 2, 143, 144, 7, 116, 2, 2, 144, 145, 7, 99, 2,
	2, 145, 146, 7, 121, 2, 2, 146, 28, 3, 2, 2, 2, 147, 148, 7, 117, 2, 2,
	148, 149, 7, 118, 2, 2, 149, 150, 7, 116, 2, 2, 150, 151, 7, 107, 2, 2,
	151, 152, 7, 112, 2, 2, 152, 153, 7, 105, 2, 2, 153, 30, 3, 2, 2, 2, 154,
	155, 7, 117, 2, 2, 155, 156, 7, 106, 2, 2, 156, 157, 7, 99, 2, 2, 157,
	158, 7, 51, 2, 2, 158, 32, 3, 2, 2, 2, 159, 160, 7, 111, 2, 2, 160, 161,
	7, 102, 2, 2, 161, 162, 7, 55, 2, 2, 162, 34, 3, 2, 2, 2, 163, 164, 7,
	112, 2, 2, 164, 165, 7, 118, 2, 2, 165, 166, 7, 104, 2, 2, 166, 167, 7,
	117, 2, 2, 167, 36, 3, 2, 2, 2, 168, 169, 7, 104, 2, 2, 169, 170, 7, 107,
	2, 2, 170, 171, 7, 110, 2, 2, 171, 172, 7, 103, 2, 2, 172, 38, 3, 2, 2,
	2, 173, 174, 7, 114, 2, 2, 174, 175, 7, 99, 2, 2, 175, 176, 7, 101, 2,
	2, 176, 177, 7, 109, 2, 2, 177, 178, 7, 103, 2, 2, 178, 179, 7, 118, 2,
	2, 179, 40, 3, 2, 2, 2, 180, 181, 7, 114, 2, 2, 181, 182, 7, 101, 2, 2,
	182, 183, 7, 99, 2, 2, 183, 184, 7, 114, 2, 2, 184, 42, 3, 2, 2, 2, 185,
	186, 7, 103, 2, 2, 186, 187, 7, 122, 2, 2, 187, 188, 7, 107, 2, 2, 188,
	189, 7, 104, 2, 2, 189, 190, 7, 107, 2, 2, 190, 191, 7, 112, 2, 2, 191,
	192, 7, 104, 2, 2, 192, 193, 7, 113, 2, 2, 193, 44, 3, 2, 2, 2, 194, 195,
	7, 102, 2, 2, 195, 196, 7, 99, 2, 2, 196, 197, 7, 118, 2, 2, 197, 198,
	7, 103, 2, 2, 198, 199, 7, 118, 2, 2, 199, 200, 7, 107, 2, 2, 200, 201,
	7, 111, 2, 2, 201, 202, 7, 103, 2, 2, 202, 46, 3, 2, 2, 2, 203, 204, 7,
	111, 2, 2, 204, 205, 7, 103, 2, 2, 205, 206, 7, 111, 2, 2, 206, 207, 7,
	113, 2, 2, 207, 208, 7, 116, 2, 2, 208, 209, 7, 123, 2, 2, 209, 48, 3,
	2, 2, 2, 210, 211, 7, 110, 2, 2, 211, 212, 7, 107, 2, 2, 212, 213, 7, 117,
	2, 2, 213, 214, 7, 118, 2, 2, 214, 215, 7, 113, 2, 2, 215, 216, 7, 104,
	2, 2, 216, 217, 7, 47, 2, 2, 217, 218, 7, 111, 2, 2, 218, 219, 7, 102,
	2, 2, 219, 220, 7, 55, 2, 2, 220, 50, 3, 2, 2, 2, 221, 222, 7, 110, 2,
	2, 222, 223, 7, 107, 2, 2, 223, 224, 7, 117, 2, 2, 224, 225, 7, 118, 2,
	2, 225, 226, 7, 113, 2, 2, 226, 227, 7, 104, 2, 2, 227, 228, 7, 47, 2,
	2, 228, 229, 7, 117, 2, 2, 229, 230, 7, 106, 2, 2, 230, 231, 7, 99, 2,
	2, 231, 232, 7, 51, 2, 2, 232, 52, 3, 2, 2, 2, 233, 234, 7, 110, 2, 2,
	234, 235, 7, 107, 2, 2, 235, 236, 7, 117, 2, 2, 236, 237, 7, 118, 2, 2,
	237, 238, 7, 113, 2, 2, 238, 239, 7, 104, 2, 2, 239, 240, 7, 47, 2, 2,
	240, 241, 7, 117, 2, 2, 241, 242, 7, 106, 2, 2, 242, 243, 7, 99, 2, 2,
	243, 244, 7, 52, 2, 2, 244, 245, 7, 55, 2, 2, 245, 246, 7, 56, 2, 2, 246,
	54, 3, 2, 2, 2, 247, 248, 7, 103, 2, 2, 248, 249, 7, 122, 2, 2, 249, 250,
	7, 118, 2, 2, 250, 251, 7, 116, 2, 2, 251, 252, 7, 99, 2, 2, 252, 253,
	7, 101, 2, 2, 253, 254, 7, 118, 2, 2, 254, 56, 3, 2, 2, 2, 255, 256, 7,
	117, 2, 2, 256, 257, 7, 113, 2, 2, 257, 258, 7, 116, 2, 2, 258, 259, 7,
	118, 2, 2, 259, 58, 3, 2, 2, 2, 260, 261, 7, 117, 2, 2, 261, 262, 7, 106,
	2, 2, 262, 263, 7, 99, 2, 2, 263, 264, 7, 52, 2, 2, 264, 265, 7, 55, 2,
	2, 265, 266, 7, 56, 2, 2, 266, 60, 3, 2, 2, 2, 267, 268, 7, 105, 2, 2,
	268, 269, 7, 103, 2, 2, 269, 270, 7, 118, 2, 2, 270, 271, 7, 73, 2, 2,
	271, 272, 7, 103, 2, 2, 272, 273, 7, 118, 2, 2, 273, 274, 7, 84, 2, 2,
	274, 275, 7, 103, 2, 2, 275, 276, 7, 115, 2, 2, 276, 277, 7, 119, 2, 2,
	277, 278, 7, 103, 2, 2, 278, 279, 7, 117, 2, 2, 279, 280, 7, 118, 2, 2,
	280, 281, 7, 117, 2, 2, 281, 62, 3, 2, 2, 2, 282, 283, 7, 102, 2, 2, 283,
	284, 7, 107, 2, 2, 284, 285, 7, 117, 2, 2, 285, 286, 7, 109, 2, 2, 286,
	287, 7, 107, 2, 2, 287, 288, 7, 112, 2, 2, 288, 289, 7, 104, 2, 2, 289,
	290, 7, 113, 2, 2, 290, 64, 3, 2, 2, 2, 291, 292, 7, 119, 2, 2, 292, 293,
	7, 112, 2, 2, 293, 294, 7, 107, 2, 2, 294, 295, 7, 113, 2, 2, 295, 296,
	7, 112, 2, 2, 296, 66, 3, 2, 2, 2, 297, 298, 7, 114, 2, 2, 298, 299, 7,
	117, 2, 2, 299, 300, 7, 110, 2, 2, 300, 301, 7, 107, 2, 2, 301, 302, 7,
	117, 2, 2, 302, 303, 7, 118, 2, 2, 303, 68, 3, 2, 2, 2, 304, 305, 7, 105,
	2, 2, 305, 306, 7, 116, 2, 2, 306, 307, 7, 103, 2, 2, 307, 308, 7, 114,
	2, 2, 308, 70, 3, 2, 2, 2, 309, 310, 7, 80, 2, 2, 310, 311, 7, 87, 2, 2,
	311, 312, 7, 73, 2, 2, 312, 313, 7, 73, 2, 2, 313, 314, 7, 71, 2, 2, 314,
	315, 7, 86, 2, 2, 315, 316, 7, 73, 2, 2, 316, 317, 7, 71, 2, 2, 317, 318,
	7, 80, 2, 2, 318, 319, 7, 71, 2, 2, 319, 320, 7, 84, 2, 2, 320, 321, 7,
	67, 2, 2, 321, 322, 7, 86, 2, 2, 322, 323, 7, 81, 2, 2, 323, 324, 7, 84,
	2, 2, 324, 325, 7, 82, 2, 2, 325, 326, 7, 78, 2, 2, 326, 327, 7, 67, 2,
	2, 327, 328, 7, 69, 2, 2, 328, 329, 7, 71, 2, 2, 329, 330, 7, 74, 2, 2,
	330, 331, 7, 81, 2, 2, 331, 332, 7, 78, 2, 2, 332, 333, 7, 70, 2, 2, 333,
	334, 7, 71, 2, 2, 334, 335, 7, 84, 2, 2, 335, 72, 3, 2, 2, 2, 336, 337,
	7, 99, 2, 2, 337, 338, 7, 117, 2, 2, 338, 74, 3, 2, 2, 2, 339, 340, 7,
	100, 2, 2, 340, 341, 7, 123, 2, 2, 341, 76, 3, 2, 2, 2, 342, 343, 7, 104,
	2, 2, 343, 344, 7, 107, 2, 2, 344, 345, 7, 110, 2, 2, 345, 346, 7, 118,
	2, 2, 346, 347, 7, 103, 2, 2, 347, 348, 7, 116, 2, 2, 348, 78, 3, 2, 2,
	2, 349, 357, 9, 2, 2, 2, 350, 351, 7, 64, 2, 2, 351, 357, 7, 63, 2, 2,
	352, 353, 7, 62, 2, 2, 353, 357, 7, 63, 2, 2, 354, 355, 7, 63, 2, 2, 355,
	357, 7, 63, 2, 2, 356, 349, 3, 2, 2, 2, 356, 350, 3, 2, 2, 2, 356, 352,
	3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 357, 80, 3, 2, 2, 2, 358, 359, 7, 93,
	2, 2, 359, 360, 7, 95, 2, 2, 360, 82, 3, 2, 2, 2, 361, 363, 9, 3, 2, 2,
	362, 361, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364,
	365, 3, 2, 2, 2, 365, 84, 3, 2, 2, 2, 366, 368, 9, 4, 2, 2, 367, 366, 3,
	2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2,
	2, 370, 86, 3, 2, 2, 2, 371, 377, 7, 36, 2, 2, 372, 373, 7, 36, 2, 2, 373,
	376, 7, 36, 2, 2, 374, 376, 10, 5, 2, 2, 375, 372, 3, 2, 2, 2, 375, 374,
	3, 2, 2, 2, 376, 379, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 378, 3, 2,
	2, 2, 378, 380, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 380, 381, 7, 36, 2, 2,
	381, 88, 3, 2, 2, 2, 382, 384, 9, 6, 2, 2, 383, 382, 3, 2, 2, 2, 384, 385,
	3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 387, 3, 2,
	2, 2, 387, 388, 8, 45, 2, 2, 388, 90, 3, 2, 2, 2, 389, 391, 7, 15, 2, 2,
	390, 389, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392,
	393, 7, 12, 2, 2, 393, 92, 3, 2, 2, 2, 394, 395, 7, 49, 2, 2, 395, 396,
	7, 49, 2, 2, 396, 400, 3, 2, 2, 2, 397, 399, 10, 7, 2, 2, 398, 397, 3,
	2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 400, 401, 3, 2, 2,
	2, 401, 403, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 403, 404, 8, 47, 2, 2, 404,
	94, 3, 2, 2, 2, 11, 2, 356, 364, 369, 375, 377, 385, 390, 400, 3, 8, 2,
	2,
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
	"", "'tuple['", "','", "']'", "'='", "'|'", "'('", "')'", "'type'", "'print'",
	"'size'", "'typex'", "'printx'", "'raw'", "'string'", "'sha1'", "'md5'",
	"'ntfs'", "'file'", "'packet'", "'pcap'", "'exifinfo'", "'datetime'", "'memory'",
	"'listof-md5'", "'listof-sha1'", "'listof-sha256'", "'extract'", "'sort'",
	"'sha256'", "'getGetRequests'", "'diskinfo'", "'union'", "'pslist'", "'grep'",
	"'NUGGETGENERATORPLACEHOLDER'", "'as'", "'by'", "'filter'", "", "'[]'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS", "NL", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "COMPOP", "LISTOP", "INT",
	"ID", "STRING", "WS", "NL", "LINE_COMMENT",
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
	NuggetLexerT__37        = 38
	NuggetLexerCOMPOP       = 39
	NuggetLexerLISTOP       = 40
	NuggetLexerINT          = 41
	NuggetLexerID           = 42
	NuggetLexerSTRING       = 43
	NuggetLexerWS           = 44
	NuggetLexerNL           = 45
	NuggetLexerLINE_COMMENT = 46
)
