// Generated from Nugget.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 31, 306,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 100,
	10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 112, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 126, 10, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 136, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 144, 10, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 154, 10, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	168, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 180, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 194, 10, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 214, 10, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 230, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 5, 23, 240, 10, 23, 3, 24, 6, 24, 243, 10, 24, 13, 24,
	14, 24, 244, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 253, 10,
	25, 12, 25, 14, 25, 256, 11, 25, 3, 25, 3, 25, 3, 26, 6, 26, 261, 10, 26,
	13, 26, 14, 26, 262, 3, 27, 3, 27, 6, 27, 267, 10, 27, 13, 27, 14, 27,
	268, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 277, 10, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33,
	300, 10, 33, 12, 33, 14, 33, 303, 11, 33, 3, 33, 3, 33, 3, 254, 2, 34,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	2, 61, 2, 63, 2, 65, 31, 3, 2, 6, 5, 2, 11, 12, 15, 15, 34, 34, 7, 2, 44,
	44, 48, 48, 60, 60, 67, 92, 99, 124, 4, 2, 62, 62, 64, 64, 4, 2, 12, 12,
	15, 15, 2, 323, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 67, 3, 2, 2, 2,
	5, 69, 3, 2, 2, 2, 7, 71, 3, 2, 2, 2, 9, 73, 3, 2, 2, 2, 11, 78, 3, 2,
	2, 2, 13, 81, 3, 2, 2, 2, 15, 83, 3, 2, 2, 2, 17, 85, 3, 2, 2, 2, 19, 87,
	3, 2, 2, 2, 21, 99, 3, 2, 2, 2, 23, 111, 3, 2, 2, 2, 25, 113, 3, 2, 2,
	2, 27, 125, 3, 2, 2, 2, 29, 135, 3, 2, 2, 2, 31, 143, 3, 2, 2, 2, 33, 153,
	3, 2, 2, 2, 35, 167, 3, 2, 2, 2, 37, 179, 3, 2, 2, 2, 39, 193, 3, 2, 2,
	2, 41, 213, 3, 2, 2, 2, 43, 229, 3, 2, 2, 2, 45, 239, 3, 2, 2, 2, 47, 242,
	3, 2, 2, 2, 49, 248, 3, 2, 2, 2, 51, 260, 3, 2, 2, 2, 53, 266, 3, 2, 2,
	2, 55, 276, 3, 2, 2, 2, 57, 278, 3, 2, 2, 2, 59, 289, 3, 2, 2, 2, 61, 291,
	3, 2, 2, 2, 63, 293, 3, 2, 2, 2, 65, 295, 3, 2, 2, 2, 67, 68, 7, 63, 2,
	2, 68, 4, 3, 2, 2, 2, 69, 70, 7, 126, 2, 2, 70, 6, 3, 2, 2, 2, 71, 72,
	7, 46, 2, 2, 72, 8, 3, 2, 2, 2, 73, 74, 7, 117, 2, 2, 74, 75, 7, 99, 2,
	2, 75, 76, 7, 120, 2, 2, 76, 77, 7, 103, 2, 2, 77, 10, 3, 2, 2, 2, 78,
	79, 7, 118, 2, 2, 79, 80, 7, 113, 2, 2, 80, 12, 3, 2, 2, 2, 81, 82, 7,
	36, 2, 2, 82, 14, 3, 2, 2, 2, 83, 84, 7, 41, 2, 2, 84, 16, 3, 2, 2, 2,
	85, 86, 7, 42, 2, 2, 86, 18, 3, 2, 2, 2, 87, 88, 7, 43, 2, 2, 88, 20, 3,
	2, 2, 2, 89, 90, 7, 101, 2, 2, 90, 91, 7, 118, 2, 2, 91, 92, 7, 107, 2,
	2, 92, 93, 7, 111, 2, 2, 93, 100, 7, 103, 2, 2, 94, 95, 7, 69, 2, 2, 95,
	96, 7, 86, 2, 2, 96, 97, 7, 75, 2, 2, 97, 98, 7, 79, 2, 2, 98, 100, 7,
	71, 2, 2, 99, 89, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2, 100, 22, 3, 2, 2, 2,
	101, 102, 7, 111, 2, 2, 102, 103, 7, 118, 2, 2, 103, 104, 7, 107, 2, 2,
	104, 105, 7, 111, 2, 2, 105, 112, 7, 103, 2, 2, 106, 107, 7, 79, 2, 2,
	107, 108, 7, 86, 2, 2, 108, 109, 7, 75, 2, 2, 109, 110, 7, 79, 2, 2, 110,
	112, 7, 71, 2, 2, 111, 101, 3, 2, 2, 2, 111, 106, 3, 2, 2, 2, 112, 24,
	3, 2, 2, 2, 113, 114, 7, 117, 2, 2, 114, 115, 7, 107, 2, 2, 115, 116, 7,
	112, 2, 2, 116, 26, 3, 2, 2, 2, 117, 118, 7, 110, 2, 2, 118, 119, 7, 113,
	2, 2, 119, 120, 7, 99, 2, 2, 120, 126, 7, 102, 2, 2, 121, 122, 7, 78, 2,
	2, 122, 123, 7, 81, 2, 2, 123, 124, 7, 67, 2, 2, 124, 126, 7, 70, 2, 2,
	125, 117, 3, 2, 2, 2, 125, 121, 3, 2, 2, 2, 126, 28, 3, 2, 2, 2, 127, 128,
	7, 104, 2, 2, 128, 129, 7, 116, 2, 2, 129, 130, 7, 113, 2, 2, 130, 136,
	7, 111, 2, 2, 131, 132, 7, 72, 2, 2, 132, 133, 7, 84, 2, 2, 133, 134, 7,
	81, 2, 2, 134, 136, 7, 79, 2, 2, 135, 127, 3, 2, 2, 2, 135, 131, 3, 2,
	2, 2, 136, 30, 3, 2, 2, 2, 137, 138, 7, 99, 2, 2, 138, 139, 7, 110, 2,
	2, 139, 144, 7, 110, 2, 2, 140, 141, 7, 67, 2, 2, 141, 142, 7, 78, 2, 2,
	142, 144, 7, 78, 2, 2, 143, 137, 3, 2, 2, 2, 143, 140, 3, 2, 2, 2, 144,
	32, 3, 2, 2, 2, 145, 146, 7, 106, 2, 2, 146, 147, 7, 99, 2, 2, 147, 148,
	7, 117, 2, 2, 148, 154, 7, 106, 2, 2, 149, 150, 7, 74, 2, 2, 150, 151,
	7, 67, 2, 2, 151, 152, 7, 85, 2, 2, 152, 154, 7, 74, 2, 2, 153, 145, 3,
	2, 2, 2, 153, 149, 3, 2, 2, 2, 154, 34, 3, 2, 2, 2, 155, 156, 7, 117, 2,
	2, 156, 157, 7, 103, 2, 2, 157, 158, 7, 110, 2, 2, 158, 159, 7, 103, 2,
	2, 159, 160, 7, 101, 2, 2, 160, 168, 7, 118, 2, 2, 161, 162, 7, 85, 2,
	2, 162, 163, 7, 71, 2, 2, 163, 164, 7, 78, 2, 2, 164, 165, 7, 71, 2, 2,
	165, 166, 7, 69, 2, 2, 166, 168, 7, 86, 2, 2, 167, 155, 3, 2, 2, 2, 167,
	161, 3, 2, 2, 2, 168, 36, 3, 2, 2, 2, 169, 170, 7, 104, 2, 2, 170, 171,
	7, 107, 2, 2, 171, 172, 7, 110, 2, 2, 172, 173, 7, 103, 2, 2, 173, 180,
	7, 117, 2, 2, 174, 175, 7, 72, 2, 2, 175, 176, 7, 75, 2, 2, 176, 177, 7,
	78, 2, 2, 177, 178, 7, 71, 2, 2, 178, 180, 7, 85, 2, 2, 179, 169, 3, 2,
	2, 2, 179, 174, 3, 2, 2, 2, 180, 38, 3, 2, 2, 2, 181, 182, 7, 107, 2, 2,
	182, 183, 7, 111, 2, 2, 183, 184, 7, 99, 2, 2, 184, 185, 7, 105, 2, 2,
	185, 186, 7, 103, 2, 2, 186, 194, 7, 117, 2, 2, 187, 188, 7, 75, 2, 2,
	188, 189, 7, 79, 2, 2, 189, 190, 7, 67, 2, 2, 190, 191, 7, 73, 2, 2, 191,
	192, 7, 71, 2, 2, 192, 194, 7, 85, 2, 2, 193, 181, 3, 2, 2, 2, 193, 187,
	3, 2, 2, 2, 194, 40, 3, 2, 2, 2, 195, 196, 7, 102, 2, 2, 196, 197, 7, 113,
	2, 2, 197, 198, 7, 101, 2, 2, 198, 199, 7, 119, 2, 2, 199, 200, 7, 111,
	2, 2, 200, 201, 7, 103, 2, 2, 201, 202, 7, 112, 2, 2, 202, 203, 7, 118,
	2, 2, 203, 214, 7, 117, 2, 2, 204, 205, 7, 70, 2, 2, 205, 206, 7, 81, 2,
	2, 206, 207, 7, 69, 2, 2, 207, 208, 7, 87, 2, 2, 208, 209, 7, 79, 2, 2,
	209, 210, 7, 71, 2, 2, 210, 211, 7, 80, 2, 2, 211, 212, 7, 86, 2, 2, 212,
	214, 7, 85, 2, 2, 213, 195, 3, 2, 2, 2, 213, 204, 3, 2, 2, 2, 214, 42,
	3, 2, 2, 2, 215, 216, 7, 103, 2, 2, 216, 217, 7, 122, 2, 2, 217, 218, 7,
	118, 2, 2, 218, 219, 7, 116, 2, 2, 219, 220, 7, 99, 2, 2, 220, 221, 7,
	101, 2, 2, 221, 230, 7, 118, 2, 2, 222, 223, 7, 71, 2, 2, 223, 224, 7,
	90, 2, 2, 224, 225, 7, 86, 2, 2, 225, 226, 7, 84, 2, 2, 226, 227, 7, 67,
	2, 2, 227, 228, 7, 69, 2, 2, 228, 230, 7, 86, 2, 2, 229, 215, 3, 2, 2,
	2, 229, 222, 3, 2, 2, 2, 230, 44, 3, 2, 2, 2, 231, 232, 7, 108, 2, 2, 232,
	233, 7, 113, 2, 2, 233, 234, 7, 107, 2, 2, 234, 240, 7, 112, 2, 2, 235,
	236, 7, 76, 2, 2, 236, 237, 7, 81, 2, 2, 237, 238, 7, 75, 2, 2, 238, 240,
	7, 80, 2, 2, 239, 231, 3, 2, 2, 2, 239, 235, 3, 2, 2, 2, 240, 46, 3, 2,
	2, 2, 241, 243, 9, 2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2,
	244, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246,
	247, 8, 24, 2, 2, 247, 48, 3, 2, 2, 2, 248, 249, 7, 49, 2, 2, 249, 250,
	7, 49, 2, 2, 250, 254, 3, 2, 2, 2, 251, 253, 11, 2, 2, 2, 252, 251, 3,
	2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 254, 252, 3, 2, 2,
	2, 255, 257, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 257, 258, 8, 25, 2, 2, 258,
	50, 3, 2, 2, 2, 259, 261, 5, 59, 30, 2, 260, 259, 3, 2, 2, 2, 261, 262,
	3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 52, 3, 2,
	2, 2, 264, 267, 9, 3, 2, 2, 265, 267, 5, 59, 30, 2, 266, 264, 3, 2, 2,
	2, 266, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268,
	269, 3, 2, 2, 2, 269, 54, 3, 2, 2, 2, 270, 277, 9, 4, 2, 2, 271, 272, 7,
	62, 2, 2, 272, 277, 7, 63, 2, 2, 273, 274, 7, 64, 2, 2, 274, 277, 7, 63,
	2, 2, 275, 277, 7, 63, 2, 2, 276, 270, 3, 2, 2, 2, 276, 271, 3, 2, 2, 2,
	276, 273, 3, 2, 2, 2, 276, 275, 3, 2, 2, 2, 277, 56, 3, 2, 2, 2, 278, 279,
	5, 59, 30, 2, 279, 280, 5, 59, 30, 2, 280, 281, 7, 49, 2, 2, 281, 282,
	5, 59, 30, 2, 282, 283, 5, 59, 30, 2, 283, 284, 7, 49, 2, 2, 284, 285,
	5, 59, 30, 2, 285, 286, 5, 59, 30, 2, 286, 287, 5, 59, 30, 2, 287, 288,
	5, 59, 30, 2, 288, 58, 3, 2, 2, 2, 289, 290, 4, 50, 59, 2, 290, 60, 3,
	2, 2, 2, 291, 292, 7, 46, 2, 2, 292, 62, 3, 2, 2, 2, 293, 294, 7, 44, 2,
	2, 294, 64, 3, 2, 2, 2, 295, 296, 7, 49, 2, 2, 296, 297, 7, 49, 2, 2, 297,
	301, 3, 2, 2, 2, 298, 300, 10, 5, 2, 2, 299, 298, 3, 2, 2, 2, 300, 303,
	3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 304, 3, 2,
	2, 2, 303, 301, 3, 2, 2, 2, 304, 305, 8, 33, 2, 2, 305, 66, 3, 2, 2, 2,
	22, 2, 99, 111, 125, 135, 143, 153, 167, 179, 193, 213, 229, 239, 244,
	254, 262, 266, 268, 276, 301, 3, 8, 2, 2,
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
	"", "'='", "'|'", "','", "'save'", "'to'", "'\"'", "'''", "'('", "')'",
	"", "", "'sin'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "CTIME", "MTIME", "SIN", "LOAD",
	"FROM", "ALL", "HASH", "SELECT", "FILES", "IMAGES", "DOCUMENTS", "EXTRACT",
	"JOIN", "WS", "COMMENT", "NUMBER", "StrLit", "OPERATION", "DATE", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"CTIME", "MTIME", "SIN", "LOAD", "FROM", "ALL", "HASH", "SELECT", "FILES",
	"IMAGES", "DOCUMENTS", "EXTRACT", "JOIN", "WS", "COMMENT", "NUMBER", "StrLit",
	"OPERATION", "DATE", "DIGIT", "DLMT", "WILDCARD", "LINE_COMMENT",
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
	NuggetLexerCTIME        = 10
	NuggetLexerMTIME        = 11
	NuggetLexerSIN          = 12
	NuggetLexerLOAD         = 13
	NuggetLexerFROM         = 14
	NuggetLexerALL          = 15
	NuggetLexerHASH         = 16
	NuggetLexerSELECT       = 17
	NuggetLexerFILES        = 18
	NuggetLexerIMAGES       = 19
	NuggetLexerDOCUMENTS    = 20
	NuggetLexerEXTRACT      = 21
	NuggetLexerJOIN         = 22
	NuggetLexerWS           = 23
	NuggetLexerCOMMENT      = 24
	NuggetLexerNUMBER       = 25
	NuggetLexerStrLit       = 26
	NuggetLexerOPERATION    = 27
	NuggetLexerDATE         = 28
	NuggetLexerLINE_COMMENT = 29
)
