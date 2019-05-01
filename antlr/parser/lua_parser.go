// Code generated from Lua.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lua

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 455,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 93, 10, 3, 12, 3, 14, 3, 96, 11, 3, 3, 3,
	5, 3, 99, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 132,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 158, 10, 4, 5, 4, 160, 10, 4, 3, 5, 3, 5, 5, 5, 164, 10, 5,
	3, 5, 5, 5, 167, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7,
	176, 10, 7, 12, 7, 14, 7, 179, 11, 7, 3, 7, 3, 7, 5, 7, 183, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 190, 10, 8, 12, 8, 14, 8, 193, 11, 8,
	3, 8, 5, 8, 196, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 216, 10, 11, 3, 12, 3, 12, 3, 12, 7, 12, 221, 10, 12, 12, 12, 14, 12,
	224, 11, 12, 3, 13, 3, 13, 3, 13, 7, 13, 229, 10, 13, 12, 13, 14, 13, 232,
	11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 238, 10, 13, 12, 13, 14, 13,
	241, 11, 13, 5, 13, 243, 10, 13, 3, 14, 3, 14, 3, 14, 7, 14, 248, 10, 14,
	12, 14, 14, 14, 251, 11, 14, 3, 15, 3, 15, 3, 15, 7, 15, 256, 10, 15, 12,
	15, 14, 15, 259, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 273, 10, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 307,
	10, 16, 12, 16, 14, 16, 310, 11, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 7, 19, 318, 10, 19, 12, 19, 14, 19, 321, 11, 19, 3, 20, 3, 20, 6,
	20, 325, 10, 20, 13, 20, 14, 20, 326, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	5, 21, 334, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 342,
	10, 22, 3, 22, 7, 22, 345, 10, 22, 12, 22, 14, 22, 348, 11, 22, 3, 23,
	3, 23, 3, 23, 3, 24, 7, 24, 354, 10, 24, 12, 24, 14, 24, 357, 11, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 365, 10, 24, 3, 25, 3, 25,
	5, 25, 369, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 5, 26, 375, 10, 26, 3,
	26, 3, 26, 3, 26, 5, 26, 380, 10, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28,
	5, 28, 387, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5,
	29, 396, 10, 29, 3, 29, 5, 29, 399, 10, 29, 3, 30, 3, 30, 5, 30, 403, 10,
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 411, 10, 31, 12, 31,
	14, 31, 414, 11, 31, 3, 31, 5, 31, 417, 10, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 429, 10, 32, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43,
	3, 44, 3, 44, 3, 44, 2, 3, 30, 45, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 2, 12, 3, 2,
	27, 28, 3, 2, 29, 31, 4, 2, 3, 3, 13, 13, 3, 2, 40, 45, 3, 2, 47, 48, 3,
	2, 49, 52, 3, 2, 53, 57, 5, 2, 48, 48, 55, 55, 58, 59, 3, 2, 65, 68, 3,
	2, 62, 64, 2, 477, 2, 88, 3, 2, 2, 2, 4, 94, 3, 2, 2, 2, 6, 159, 3, 2,
	2, 2, 8, 161, 3, 2, 2, 2, 10, 168, 3, 2, 2, 2, 12, 172, 3, 2, 2, 2, 14,
	184, 3, 2, 2, 2, 16, 199, 3, 2, 2, 2, 18, 204, 3, 2, 2, 2, 20, 215, 3,
	2, 2, 2, 22, 217, 3, 2, 2, 2, 24, 242, 3, 2, 2, 2, 26, 244, 3, 2, 2, 2,
	28, 252, 3, 2, 2, 2, 30, 272, 3, 2, 2, 2, 32, 311, 3, 2, 2, 2, 34, 313,
	3, 2, 2, 2, 36, 315, 3, 2, 2, 2, 38, 322, 3, 2, 2, 2, 40, 333, 3, 2, 2,
	2, 42, 341, 3, 2, 2, 2, 44, 349, 3, 2, 2, 2, 46, 355, 3, 2, 2, 2, 48, 368,
	3, 2, 2, 2, 50, 379, 3, 2, 2, 2, 52, 381, 3, 2, 2, 2, 54, 384, 3, 2, 2,
	2, 56, 398, 3, 2, 2, 2, 58, 400, 3, 2, 2, 2, 60, 406, 3, 2, 2, 2, 62, 428,
	3, 2, 2, 2, 64, 430, 3, 2, 2, 2, 66, 432, 3, 2, 2, 2, 68, 434, 3, 2, 2,
	2, 70, 436, 3, 2, 2, 2, 72, 438, 3, 2, 2, 2, 74, 440, 3, 2, 2, 2, 76, 442,
	3, 2, 2, 2, 78, 444, 3, 2, 2, 2, 80, 446, 3, 2, 2, 2, 82, 448, 3, 2, 2,
	2, 84, 450, 3, 2, 2, 2, 86, 452, 3, 2, 2, 2, 88, 89, 5, 4, 3, 2, 89, 90,
	7, 2, 2, 3, 90, 3, 3, 2, 2, 2, 91, 93, 5, 6, 4, 2, 92, 91, 3, 2, 2, 2,
	93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 98, 3,
	2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 99, 5, 8, 5, 2, 98, 97, 3, 2, 2, 2, 98,
	99, 3, 2, 2, 2, 99, 5, 3, 2, 2, 2, 100, 160, 7, 3, 2, 2, 101, 160, 5, 20,
	11, 2, 102, 160, 5, 38, 20, 2, 103, 160, 5, 10, 6, 2, 104, 160, 7, 4, 2,
	2, 105, 106, 7, 5, 2, 2, 106, 160, 7, 61, 2, 2, 107, 108, 7, 6, 2, 2, 108,
	109, 5, 4, 3, 2, 109, 110, 7, 7, 2, 2, 110, 160, 3, 2, 2, 2, 111, 112,
	7, 8, 2, 2, 112, 113, 5, 30, 16, 2, 113, 114, 7, 6, 2, 2, 114, 115, 5,
	4, 3, 2, 115, 116, 7, 7, 2, 2, 116, 160, 3, 2, 2, 2, 117, 118, 7, 9, 2,
	2, 118, 119, 5, 4, 3, 2, 119, 120, 7, 10, 2, 2, 120, 121, 5, 30, 16, 2,
	121, 160, 3, 2, 2, 2, 122, 160, 5, 14, 8, 2, 123, 124, 7, 11, 2, 2, 124,
	125, 7, 61, 2, 2, 125, 126, 7, 12, 2, 2, 126, 127, 5, 30, 16, 2, 127, 128,
	7, 13, 2, 2, 128, 131, 5, 30, 16, 2, 129, 130, 7, 13, 2, 2, 130, 132, 5,
	30, 16, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 134, 7, 6, 2, 2, 134, 135, 5, 4, 3, 2, 135, 136, 7, 7, 2, 2,
	136, 160, 3, 2, 2, 2, 137, 138, 7, 11, 2, 2, 138, 139, 5, 26, 14, 2, 139,
	140, 7, 14, 2, 2, 140, 141, 5, 28, 15, 2, 141, 142, 7, 6, 2, 2, 142, 143,
	5, 4, 3, 2, 143, 144, 7, 7, 2, 2, 144, 160, 3, 2, 2, 2, 145, 146, 7, 15,
	2, 2, 146, 147, 5, 12, 7, 2, 147, 148, 5, 54, 28, 2, 148, 160, 3, 2, 2,
	2, 149, 150, 7, 16, 2, 2, 150, 151, 7, 15, 2, 2, 151, 152, 7, 61, 2, 2,
	152, 160, 5, 54, 28, 2, 153, 154, 7, 16, 2, 2, 154, 157, 5, 26, 14, 2,
	155, 156, 7, 12, 2, 2, 156, 158, 5, 28, 15, 2, 157, 155, 3, 2, 2, 2, 157,
	158, 3, 2, 2, 2, 158, 160, 3, 2, 2, 2, 159, 100, 3, 2, 2, 2, 159, 101,
	3, 2, 2, 2, 159, 102, 3, 2, 2, 2, 159, 103, 3, 2, 2, 2, 159, 104, 3, 2,
	2, 2, 159, 105, 3, 2, 2, 2, 159, 107, 3, 2, 2, 2, 159, 111, 3, 2, 2, 2,
	159, 117, 3, 2, 2, 2, 159, 122, 3, 2, 2, 2, 159, 123, 3, 2, 2, 2, 159,
	137, 3, 2, 2, 2, 159, 145, 3, 2, 2, 2, 159, 149, 3, 2, 2, 2, 159, 153,
	3, 2, 2, 2, 160, 7, 3, 2, 2, 2, 161, 163, 7, 17, 2, 2, 162, 164, 5, 28,
	15, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2,
	165, 167, 7, 3, 2, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167,
	9, 3, 2, 2, 2, 168, 169, 7, 18, 2, 2, 169, 170, 7, 61, 2, 2, 170, 171,
	7, 18, 2, 2, 171, 11, 3, 2, 2, 2, 172, 177, 7, 61, 2, 2, 173, 174, 7, 19,
	2, 2, 174, 176, 7, 61, 2, 2, 175, 173, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2,
	177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 182, 3, 2, 2, 2, 179,
	177, 3, 2, 2, 2, 180, 181, 7, 20, 2, 2, 181, 183, 7, 61, 2, 2, 182, 180,
	3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 13, 3, 2, 2, 2, 184, 185, 7, 21,
	2, 2, 185, 186, 5, 30, 16, 2, 186, 187, 7, 22, 2, 2, 187, 191, 5, 4, 3,
	2, 188, 190, 5, 16, 9, 2, 189, 188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191,
	189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191,
	3, 2, 2, 2, 194, 196, 5, 18, 10, 2, 195, 194, 3, 2, 2, 2, 195, 196, 3,
	2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198, 7, 7, 2, 2, 198, 15, 3, 2, 2,
	2, 199, 200, 7, 23, 2, 2, 200, 201, 5, 30, 16, 2, 201, 202, 7, 22, 2, 2,
	202, 203, 5, 4, 3, 2, 203, 17, 3, 2, 2, 2, 204, 205, 7, 24, 2, 2, 205,
	206, 5, 4, 3, 2, 206, 19, 3, 2, 2, 2, 207, 208, 5, 22, 12, 2, 208, 209,
	7, 12, 2, 2, 209, 210, 5, 28, 15, 2, 210, 216, 3, 2, 2, 2, 211, 212, 5,
	24, 13, 2, 212, 213, 7, 12, 2, 2, 213, 214, 5, 28, 15, 2, 214, 216, 3,
	2, 2, 2, 215, 207, 3, 2, 2, 2, 215, 211, 3, 2, 2, 2, 216, 21, 3, 2, 2,
	2, 217, 222, 5, 42, 22, 2, 218, 219, 7, 13, 2, 2, 219, 221, 5, 42, 22,
	2, 220, 218, 3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222,
	223, 3, 2, 2, 2, 223, 23, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 230, 5,
	44, 23, 2, 226, 227, 7, 13, 2, 2, 227, 229, 5, 44, 23, 2, 228, 226, 3,
	2, 2, 2, 229, 232, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2,
	2, 231, 243, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 233, 234, 5, 34, 18, 2,
	234, 239, 5, 42, 22, 2, 235, 236, 7, 13, 2, 2, 236, 238, 5, 42, 22, 2,
	237, 235, 3, 2, 2, 2, 238, 241, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239,
	240, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 242, 225,
	3, 2, 2, 2, 242, 233, 3, 2, 2, 2, 243, 25, 3, 2, 2, 2, 244, 249, 7, 61,
	2, 2, 245, 246, 7, 13, 2, 2, 246, 248, 7, 61, 2, 2, 247, 245, 3, 2, 2,
	2, 248, 251, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250,
	27, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 252, 257, 5, 30, 16, 2, 253, 254,
	7, 13, 2, 2, 254, 256, 5, 30, 16, 2, 255, 253, 3, 2, 2, 2, 256, 259, 3,
	2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 29, 3, 2, 2,
	2, 259, 257, 3, 2, 2, 2, 260, 261, 8, 16, 1, 2, 261, 273, 7, 25, 2, 2,
	262, 273, 5, 32, 17, 2, 263, 273, 5, 84, 43, 2, 264, 273, 5, 86, 44, 2,
	265, 273, 7, 26, 2, 2, 266, 273, 5, 52, 27, 2, 267, 273, 5, 36, 19, 2,
	268, 273, 5, 58, 30, 2, 269, 270, 5, 80, 41, 2, 270, 271, 5, 30, 16, 10,
	271, 273, 3, 2, 2, 2, 272, 260, 3, 2, 2, 2, 272, 262, 3, 2, 2, 2, 272,
	263, 3, 2, 2, 2, 272, 264, 3, 2, 2, 2, 272, 265, 3, 2, 2, 2, 272, 266,
	3, 2, 2, 2, 272, 267, 3, 2, 2, 2, 272, 268, 3, 2, 2, 2, 272, 269, 3, 2,
	2, 2, 273, 308, 3, 2, 2, 2, 274, 275, 12, 11, 2, 2, 275, 276, 5, 82, 42,
	2, 276, 277, 5, 30, 16, 11, 277, 307, 3, 2, 2, 2, 278, 279, 12, 9, 2, 2,
	279, 280, 5, 76, 39, 2, 280, 281, 5, 30, 16, 10, 281, 307, 3, 2, 2, 2,
	282, 283, 12, 8, 2, 2, 283, 284, 5, 74, 38, 2, 284, 285, 5, 30, 16, 9,
	285, 307, 3, 2, 2, 2, 286, 287, 12, 7, 2, 2, 287, 288, 5, 72, 37, 2, 288,
	289, 5, 30, 16, 7, 289, 307, 3, 2, 2, 2, 290, 291, 12, 6, 2, 2, 291, 292,
	5, 70, 36, 2, 292, 293, 5, 30, 16, 7, 293, 307, 3, 2, 2, 2, 294, 295, 12,
	5, 2, 2, 295, 296, 5, 68, 35, 2, 296, 297, 5, 30, 16, 6, 297, 307, 3, 2,
	2, 2, 298, 299, 12, 4, 2, 2, 299, 300, 5, 66, 34, 2, 300, 301, 5, 30, 16,
	5, 301, 307, 3, 2, 2, 2, 302, 303, 12, 3, 2, 2, 303, 304, 5, 78, 40, 2,
	304, 305, 5, 30, 16, 4, 305, 307, 3, 2, 2, 2, 306, 274, 3, 2, 2, 2, 306,
	278, 3, 2, 2, 2, 306, 282, 3, 2, 2, 2, 306, 286, 3, 2, 2, 2, 306, 290,
	3, 2, 2, 2, 306, 294, 3, 2, 2, 2, 306, 298, 3, 2, 2, 2, 306, 302, 3, 2,
	2, 2, 307, 310, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2,
	309, 31, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 311, 312, 9, 2, 2, 2, 312, 33,
	3, 2, 2, 2, 313, 314, 9, 3, 2, 2, 314, 35, 3, 2, 2, 2, 315, 319, 5, 40,
	21, 2, 316, 318, 5, 48, 25, 2, 317, 316, 3, 2, 2, 2, 318, 321, 3, 2, 2,
	2, 319, 317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 37, 3, 2, 2, 2, 321,
	319, 3, 2, 2, 2, 322, 324, 5, 40, 21, 2, 323, 325, 5, 48, 25, 2, 324, 323,
	3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 326, 327, 3, 2,
	2, 2, 327, 39, 3, 2, 2, 2, 328, 334, 5, 42, 22, 2, 329, 330, 7, 32, 2,
	2, 330, 331, 5, 30, 16, 2, 331, 332, 7, 33, 2, 2, 332, 334, 3, 2, 2, 2,
	333, 328, 3, 2, 2, 2, 333, 329, 3, 2, 2, 2, 334, 41, 3, 2, 2, 2, 335, 342,
	7, 61, 2, 2, 336, 337, 7, 32, 2, 2, 337, 338, 5, 30, 16, 2, 338, 339, 7,
	33, 2, 2, 339, 340, 5, 46, 24, 2, 340, 342, 3, 2, 2, 2, 341, 335, 3, 2,
	2, 2, 341, 336, 3, 2, 2, 2, 342, 346, 3, 2, 2, 2, 343, 345, 5, 46, 24,
	2, 344, 343, 3, 2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 43, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 350, 5,
	34, 18, 2, 350, 351, 5, 42, 22, 2, 351, 45, 3, 2, 2, 2, 352, 354, 5, 48,
	25, 2, 353, 352, 3, 2, 2, 2, 354, 357, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2,
	355, 356, 3, 2, 2, 2, 356, 364, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 358,
	359, 7, 34, 2, 2, 359, 360, 5, 30, 16, 2, 360, 361, 7, 35, 2, 2, 361, 365,
	3, 2, 2, 2, 362, 363, 7, 19, 2, 2, 363, 365, 7, 61, 2, 2, 364, 358, 3,
	2, 2, 2, 364, 362, 3, 2, 2, 2, 365, 47, 3, 2, 2, 2, 366, 367, 7, 20, 2,
	2, 367, 369, 7, 61, 2, 2, 368, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369,
	370, 3, 2, 2, 2, 370, 371, 5, 50, 26, 2, 371, 49, 3, 2, 2, 2, 372, 374,
	7, 32, 2, 2, 373, 375, 5, 28, 15, 2, 374, 373, 3, 2, 2, 2, 374, 375, 3,
	2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 380, 7, 33, 2, 2, 377, 380, 5, 58,
	30, 2, 378, 380, 5, 86, 44, 2, 379, 372, 3, 2, 2, 2, 379, 377, 3, 2, 2,
	2, 379, 378, 3, 2, 2, 2, 380, 51, 3, 2, 2, 2, 381, 382, 7, 15, 2, 2, 382,
	383, 5, 54, 28, 2, 383, 53, 3, 2, 2, 2, 384, 386, 7, 32, 2, 2, 385, 387,
	5, 56, 29, 2, 386, 385, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 3,
	2, 2, 2, 388, 389, 7, 33, 2, 2, 389, 390, 5, 4, 3, 2, 390, 391, 7, 7, 2,
	2, 391, 55, 3, 2, 2, 2, 392, 395, 5, 26, 14, 2, 393, 394, 7, 13, 2, 2,
	394, 396, 7, 26, 2, 2, 395, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396,
	399, 3, 2, 2, 2, 397, 399, 7, 26, 2, 2, 398, 392, 3, 2, 2, 2, 398, 397,
	3, 2, 2, 2, 399, 57, 3, 2, 2, 2, 400, 402, 7, 36, 2, 2, 401, 403, 5, 60,
	31, 2, 402, 401, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2,
	404, 405, 7, 37, 2, 2, 405, 59, 3, 2, 2, 2, 406, 412, 5, 62, 32, 2, 407,
	408, 5, 64, 33, 2, 408, 409, 5, 62, 32, 2, 409, 411, 3, 2, 2, 2, 410, 407,
	3, 2, 2, 2, 411, 414, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 412, 413, 3, 2,
	2, 2, 413, 416, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 415, 417, 5, 64, 33,
	2, 416, 415, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 61, 3, 2, 2, 2, 418,
	419, 7, 34, 2, 2, 419, 420, 5, 30, 16, 2, 420, 421, 7, 35, 2, 2, 421, 422,
	7, 12, 2, 2, 422, 423, 5, 30, 16, 2, 423, 429, 3, 2, 2, 2, 424, 425, 7,
	61, 2, 2, 425, 426, 7, 12, 2, 2, 426, 429, 5, 30, 16, 2, 427, 429, 5, 30,
	16, 2, 428, 418, 3, 2, 2, 2, 428, 424, 3, 2, 2, 2, 428, 427, 3, 2, 2, 2,
	429, 63, 3, 2, 2, 2, 430, 431, 9, 4, 2, 2, 431, 65, 3, 2, 2, 2, 432, 433,
	7, 38, 2, 2, 433, 67, 3, 2, 2, 2, 434, 435, 7, 39, 2, 2, 435, 69, 3, 2,
	2, 2, 436, 437, 9, 5, 2, 2, 437, 71, 3, 2, 2, 2, 438, 439, 7, 46, 2, 2,
	439, 73, 3, 2, 2, 2, 440, 441, 9, 6, 2, 2, 441, 75, 3, 2, 2, 2, 442, 443,
	9, 7, 2, 2, 443, 77, 3, 2, 2, 2, 444, 445, 9, 8, 2, 2, 445, 79, 3, 2, 2,
	2, 446, 447, 9, 9, 2, 2, 447, 81, 3, 2, 2, 2, 448, 449, 7, 60, 2, 2, 449,
	83, 3, 2, 2, 2, 450, 451, 9, 10, 2, 2, 451, 85, 3, 2, 2, 2, 452, 453, 9,
	11, 2, 2, 453, 87, 3, 2, 2, 2, 40, 94, 98, 131, 157, 159, 163, 166, 177,
	182, 191, 195, 215, 222, 230, 239, 242, 249, 257, 272, 306, 308, 319, 326,
	333, 341, 346, 355, 364, 368, 374, 379, 386, 395, 398, 402, 412, 416, 428,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'break'", "'goto'", "'do'", "'end'", "'while'", "'repeat'",
	"'until'", "'for'", "'='", "','", "'in'", "'function'", "'local'", "'return'",
	"'::'", "'.'", "':'", "'if'", "'then'", "'elseif'", "'else'", "'nil'",
	"'...'", "'true'", "'false'", "'int'", "'float'", "'string'", "'('", "')'",
	"'['", "']'", "'{'", "'}'", "'or'", "'and'", "'<'", "'>'", "'<='", "'>='",
	"'~='", "'=='", "'..'", "'+'", "'-'", "'*'", "'/'", "'%'", "'//'", "'&'",
	"'|'", "'~'", "'<<'", "'>>'", "'not'", "'#'", "'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "NAME", "NORMALSTRING", "CHARSTRING", "LONGSTRING",
	"INT", "HEX", "FLOAT", "HEX_FLOAT", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"chunk", "block", "stat", "retstat", "label", "funcname", "ifstat", "elseifstat",
	"elsestat", "assign", "varlist", "typedvarlist", "namelist", "explist",
	"exp", "boolLiteral", "typeLiteral", "prefixexp", "functioncall", "varOrExp",
	"varId", "typedvar", "varSuffix", "nameAndArgs", "args", "functiondef",
	"funcbody", "parlist", "tableconstructor", "fieldlist", "field", "fieldsep",
	"operatorOr", "operatorAnd", "operatorComparison", "operatorStrcat", "operatorAddSub",
	"operatorMulDivMod", "operatorBitwise", "operatorUnary", "operatorPower",
	"numberLiteral", "stringLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LuaParser struct {
	*antlr.BaseParser
}

func NewLuaParser(input antlr.TokenStream) *LuaParser {
	this := new(LuaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Lua.g4"

	return this
}

// LuaParser tokens.
const (
	LuaParserEOF          = antlr.TokenEOF
	LuaParserT__0         = 1
	LuaParserT__1         = 2
	LuaParserT__2         = 3
	LuaParserT__3         = 4
	LuaParserT__4         = 5
	LuaParserT__5         = 6
	LuaParserT__6         = 7
	LuaParserT__7         = 8
	LuaParserT__8         = 9
	LuaParserT__9         = 10
	LuaParserT__10        = 11
	LuaParserT__11        = 12
	LuaParserT__12        = 13
	LuaParserT__13        = 14
	LuaParserT__14        = 15
	LuaParserT__15        = 16
	LuaParserT__16        = 17
	LuaParserT__17        = 18
	LuaParserT__18        = 19
	LuaParserT__19        = 20
	LuaParserT__20        = 21
	LuaParserT__21        = 22
	LuaParserT__22        = 23
	LuaParserT__23        = 24
	LuaParserT__24        = 25
	LuaParserT__25        = 26
	LuaParserT__26        = 27
	LuaParserT__27        = 28
	LuaParserT__28        = 29
	LuaParserT__29        = 30
	LuaParserT__30        = 31
	LuaParserT__31        = 32
	LuaParserT__32        = 33
	LuaParserT__33        = 34
	LuaParserT__34        = 35
	LuaParserT__35        = 36
	LuaParserT__36        = 37
	LuaParserT__37        = 38
	LuaParserT__38        = 39
	LuaParserT__39        = 40
	LuaParserT__40        = 41
	LuaParserT__41        = 42
	LuaParserT__42        = 43
	LuaParserT__43        = 44
	LuaParserT__44        = 45
	LuaParserT__45        = 46
	LuaParserT__46        = 47
	LuaParserT__47        = 48
	LuaParserT__48        = 49
	LuaParserT__49        = 50
	LuaParserT__50        = 51
	LuaParserT__51        = 52
	LuaParserT__52        = 53
	LuaParserT__53        = 54
	LuaParserT__54        = 55
	LuaParserT__55        = 56
	LuaParserT__56        = 57
	LuaParserT__57        = 58
	LuaParserNAME         = 59
	LuaParserNORMALSTRING = 60
	LuaParserCHARSTRING   = 61
	LuaParserLONGSTRING   = 62
	LuaParserINT          = 63
	LuaParserHEX          = 64
	LuaParserFLOAT        = 65
	LuaParserHEX_FLOAT    = 66
	LuaParserCOMMENT      = 67
	LuaParserLINE_COMMENT = 68
	LuaParserWS           = 69
	LuaParserSHEBANG      = 70
)

// LuaParser rules.
const (
	LuaParserRULE_chunk              = 0
	LuaParserRULE_block              = 1
	LuaParserRULE_stat               = 2
	LuaParserRULE_retstat            = 3
	LuaParserRULE_label              = 4
	LuaParserRULE_funcname           = 5
	LuaParserRULE_ifstat             = 6
	LuaParserRULE_elseifstat         = 7
	LuaParserRULE_elsestat           = 8
	LuaParserRULE_assign             = 9
	LuaParserRULE_varlist            = 10
	LuaParserRULE_typedvarlist       = 11
	LuaParserRULE_namelist           = 12
	LuaParserRULE_explist            = 13
	LuaParserRULE_exp                = 14
	LuaParserRULE_boolLiteral        = 15
	LuaParserRULE_typeLiteral        = 16
	LuaParserRULE_prefixexp          = 17
	LuaParserRULE_functioncall       = 18
	LuaParserRULE_varOrExp           = 19
	LuaParserRULE_varId              = 20
	LuaParserRULE_typedvar           = 21
	LuaParserRULE_varSuffix          = 22
	LuaParserRULE_nameAndArgs        = 23
	LuaParserRULE_args               = 24
	LuaParserRULE_functiondef        = 25
	LuaParserRULE_funcbody           = 26
	LuaParserRULE_parlist            = 27
	LuaParserRULE_tableconstructor   = 28
	LuaParserRULE_fieldlist          = 29
	LuaParserRULE_field              = 30
	LuaParserRULE_fieldsep           = 31
	LuaParserRULE_operatorOr         = 32
	LuaParserRULE_operatorAnd        = 33
	LuaParserRULE_operatorComparison = 34
	LuaParserRULE_operatorStrcat     = 35
	LuaParserRULE_operatorAddSub     = 36
	LuaParserRULE_operatorMulDivMod  = 37
	LuaParserRULE_operatorBitwise    = 38
	LuaParserRULE_operatorUnary      = 39
	LuaParserRULE_operatorPower      = 40
	LuaParserRULE_numberLiteral      = 41
	LuaParserRULE_stringLiteral      = 42
)

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_chunk
	return p
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ChunkContext) EOF() antlr.TerminalNode {
	return s.GetToken(LuaParserEOF, 0)
}

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitChunk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Chunk() (localctx IChunkContext) {
	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LuaParserRULE_chunk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Block()
	}
	{
		p.SetState(87)
		p.Match(LuaParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) Retstat() IRetstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetstatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LuaParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__0)|(1<<LuaParserT__1)|(1<<LuaParserT__2)|(1<<LuaParserT__3)|(1<<LuaParserT__5)|(1<<LuaParserT__6)|(1<<LuaParserT__8)|(1<<LuaParserT__12)|(1<<LuaParserT__13)|(1<<LuaParserT__15)|(1<<LuaParserT__18)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28)|(1<<LuaParserT__29))) != 0) || _la == LuaParserNAME {
		{
			p.SetState(89)
			p.Stat()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__14 {
		{
			p.SetState(95)
			p.Retstat()
		}

	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatContext) Functioncall() IFunctioncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctioncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *StatContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *StatContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *StatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *StatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatContext) Ifstat() IIfstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *StatContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *StatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *StatContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *StatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LuaParserRULE_stat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(LuaParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Functioncall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(LuaParserT__1)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(103)
			p.Match(LuaParserT__2)
		}
		{
			p.SetState(104)
			p.Match(LuaParserNAME)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(105)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(106)
			p.Block()
		}
		{
			p.SetState(107)
			p.Match(LuaParserT__4)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(109)
			p.Match(LuaParserT__5)
		}
		{
			p.SetState(110)
			p.exp(0)
		}
		{
			p.SetState(111)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(112)
			p.Block()
		}
		{
			p.SetState(113)
			p.Match(LuaParserT__4)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(115)
			p.Match(LuaParserT__6)
		}
		{
			p.SetState(116)
			p.Block()
		}
		{
			p.SetState(117)
			p.Match(LuaParserT__7)
		}
		{
			p.SetState(118)
			p.exp(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(120)
			p.Ifstat()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(121)
			p.Match(LuaParserT__8)
		}
		{
			p.SetState(122)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(123)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(124)
			p.exp(0)
		}
		{
			p.SetState(125)
			p.Match(LuaParserT__10)
		}
		{
			p.SetState(126)
			p.exp(0)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__10 {
			{
				p.SetState(127)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(128)
				p.exp(0)
			}

		}
		{
			p.SetState(131)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(132)
			p.Block()
		}
		{
			p.SetState(133)
			p.Match(LuaParserT__4)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(135)
			p.Match(LuaParserT__8)
		}
		{
			p.SetState(136)
			p.Namelist()
		}
		{
			p.SetState(137)
			p.Match(LuaParserT__11)
		}
		{
			p.SetState(138)
			p.Explist()
		}
		{
			p.SetState(139)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(140)
			p.Block()
		}
		{
			p.SetState(141)
			p.Match(LuaParserT__4)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(143)
			p.Match(LuaParserT__12)
		}
		{
			p.SetState(144)
			p.Funcname()
		}
		{
			p.SetState(145)
			p.Funcbody()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(147)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(148)
			p.Match(LuaParserT__12)
		}
		{
			p.SetState(149)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(150)
			p.Funcbody()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(151)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(152)
			p.Namelist()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__9 {
			{
				p.SetState(153)
				p.Match(LuaParserT__9)
			}
			{
				p.SetState(154)
				p.Explist()
			}

		}

	}

	return localctx
}

// IRetstatContext is an interface to support dynamic dispatch.
type IRetstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetstatContext differentiates from other interfaces.
	IsRetstatContext()
}

type RetstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetstatContext() *RetstatContext {
	var p = new(RetstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_retstat
	return p
}

func (*RetstatContext) IsRetstatContext() {}

func NewRetstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetstatContext {
	var p = new(RetstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_retstat

	return p
}

func (s *RetstatContext) GetParser() antlr.Parser { return s.parser }

func (s *RetstatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *RetstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitRetstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Retstat() (localctx IRetstatContext) {
	localctx = NewRetstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LuaParserRULE_retstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(LuaParserT__14)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(LuaParserT__12-13))|(1<<(LuaParserT__22-13))|(1<<(LuaParserT__23-13))|(1<<(LuaParserT__24-13))|(1<<(LuaParserT__25-13))|(1<<(LuaParserT__29-13))|(1<<(LuaParserT__33-13)))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46))|(1<<(LuaParserNAME-46))|(1<<(LuaParserNORMALSTRING-46))|(1<<(LuaParserCHARSTRING-46))|(1<<(LuaParserLONGSTRING-46))|(1<<(LuaParserINT-46))|(1<<(LuaParserHEX-46))|(1<<(LuaParserFLOAT-46))|(1<<(LuaParserHEX_FLOAT-46)))) != 0) {
		{
			p.SetState(160)
			p.Explist()
		}

	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 {
		{
			p.SetState(163)
			p.Match(LuaParserT__0)
		}

	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LuaParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(LuaParserT__15)
	}
	{
		p.SetState(167)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(168)
		p.Match(LuaParserT__15)
	}

	return localctx
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *FuncnameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFuncname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LuaParserRULE_funcname)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(LuaParserNAME)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__16 {
		{
			p.SetState(171)
			p.Match(LuaParserT__16)
		}
		{
			p.SetState(172)
			p.Match(LuaParserNAME)
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__17 {
		{
			p.SetState(178)
			p.Match(LuaParserT__17)
		}
		{
			p.SetState(179)
			p.Match(LuaParserNAME)
		}

	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IfstatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstatContext) AllElseifstat() []IElseifstatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseifstatContext)(nil)).Elem())
	var tst = make([]IElseifstatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseifstatContext)
		}
	}

	return tst
}

func (s *IfstatContext) Elseifstat(i int) IElseifstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseifstatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseifstatContext)
}

func (s *IfstatContext) Elsestat() IElsestatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElsestatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElsestatContext)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitIfstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LuaParserRULE_ifstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(LuaParserT__18)
	}
	{
		p.SetState(183)
		p.exp(0)
	}
	{
		p.SetState(184)
		p.Match(LuaParserT__19)
	}
	{
		p.SetState(185)
		p.Block()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__20 {
		{
			p.SetState(186)
			p.Elseifstat()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__21 {
		{
			p.SetState(192)
			p.Elsestat()
		}

	}
	{
		p.SetState(195)
		p.Match(LuaParserT__4)
	}

	return localctx
}

// IElseifstatContext is an interface to support dynamic dispatch.
type IElseifstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseifstatContext differentiates from other interfaces.
	IsElseifstatContext()
}

type ElseifstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifstatContext() *ElseifstatContext {
	var p = new(ElseifstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_elseifstat
	return p
}

func (*ElseifstatContext) IsElseifstatContext() {}

func NewElseifstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifstatContext {
	var p = new(ElseifstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_elseifstat

	return p
}

func (s *ElseifstatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifstatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ElseifstatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseifstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitElseifstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Elseifstat() (localctx IElseifstatContext) {
	localctx = NewElseifstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LuaParserRULE_elseifstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(LuaParserT__20)
	}
	{
		p.SetState(198)
		p.exp(0)
	}
	{
		p.SetState(199)
		p.Match(LuaParserT__19)
	}
	{
		p.SetState(200)
		p.Block()
	}

	return localctx
}

// IElsestatContext is an interface to support dynamic dispatch.
type IElsestatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElsestatContext differentiates from other interfaces.
	IsElsestatContext()
}

type ElsestatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsestatContext() *ElsestatContext {
	var p = new(ElsestatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_elsestat
	return p
}

func (*ElsestatContext) IsElsestatContext() {}

func NewElsestatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsestatContext {
	var p = new(ElsestatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_elsestat

	return p
}

func (s *ElsestatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsestatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElsestatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsestatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsestatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitElsestat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Elsestat() (localctx IElsestatContext) {
	localctx = NewElsestatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LuaParserRULE_elsestat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(LuaParserT__21)
	}
	{
		p.SetState(203)
		p.Block()
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Varlist() IVarlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarlistContext)
}

func (s *AssignContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *AssignContext) Typedvarlist() ITypedvarlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedvarlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedvarlistContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LuaParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29, LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Varlist()
		}
		{
			p.SetState(206)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(207)
			p.Explist()
		}

	case LuaParserT__26, LuaParserT__27, LuaParserT__28:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Typedvarlist()
		}
		{
			p.SetState(210)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(211)
			p.Explist()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVarlistContext is an interface to support dynamic dispatch.
type IVarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarlistContext differentiates from other interfaces.
	IsVarlistContext()
}

type VarlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarlistContext() *VarlistContext {
	var p = new(VarlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_varlist
	return p
}

func (*VarlistContext) IsVarlistContext() {}

func NewVarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarlistContext {
	var p = new(VarlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varlist

	return p
}

func (s *VarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *VarlistContext) AllVarId() []IVarIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarIdContext)(nil)).Elem())
	var tst = make([]IVarIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarIdContext)
		}
	}

	return tst
}

func (s *VarlistContext) VarId(i int) IVarIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarIdContext)
}

func (s *VarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Varlist() (localctx IVarlistContext) {
	localctx = NewVarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LuaParserRULE_varlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.VarId()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__10 {
		{
			p.SetState(216)
			p.Match(LuaParserT__10)
		}
		{
			p.SetState(217)
			p.VarId()
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypedvarlistContext is an interface to support dynamic dispatch.
type ITypedvarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedvarlistContext differentiates from other interfaces.
	IsTypedvarlistContext()
}

type TypedvarlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedvarlistContext() *TypedvarlistContext {
	var p = new(TypedvarlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_typedvarlist
	return p
}

func (*TypedvarlistContext) IsTypedvarlistContext() {}

func NewTypedvarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedvarlistContext {
	var p = new(TypedvarlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_typedvarlist

	return p
}

func (s *TypedvarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedvarlistContext) AllTypedvar() []ITypedvarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypedvarContext)(nil)).Elem())
	var tst = make([]ITypedvarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypedvarContext)
		}
	}

	return tst
}

func (s *TypedvarlistContext) Typedvar(i int) ITypedvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedvarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypedvarContext)
}

func (s *TypedvarlistContext) TypeLiteral() ITypeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLiteralContext)
}

func (s *TypedvarlistContext) AllVarId() []IVarIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarIdContext)(nil)).Elem())
	var tst = make([]IVarIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarIdContext)
		}
	}

	return tst
}

func (s *TypedvarlistContext) VarId(i int) IVarIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarIdContext)
}

func (s *TypedvarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedvarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedvarlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTypedvarlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Typedvarlist() (localctx ITypedvarlistContext) {
	localctx = NewTypedvarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LuaParserRULE_typedvarlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Typedvar()
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__10 {
			{
				p.SetState(224)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(225)
				p.Typedvar()
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.TypeLiteral()
		}
		{
			p.SetState(232)
			p.VarId()
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__10 {
			{
				p.SetState(233)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(234)
				p.VarId()
			}

			p.SetState(239)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// INamelistContext is an interface to support dynamic dispatch.
type INamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamelistContext differentiates from other interfaces.
	IsNamelistContext()
}

type NamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamelistContext() *NamelistContext {
	var p = new(NamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_namelist
	return p
}

func (*NamelistContext) IsNamelistContext() {}

func NewNamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamelistContext {
	var p = new(NamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_namelist

	return p
}

func (s *NamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *NamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *NamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamelistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNamelist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Namelist() (localctx INamelistContext) {
	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LuaParserRULE_namelist)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(LuaParserNAME)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(243)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(244)
				p.Match(LuaParserNAME)
			}

		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IExplistContext is an interface to support dynamic dispatch.
type IExplistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExplistContext differentiates from other interfaces.
	IsExplistContext()
}

type ExplistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExplistContext() *ExplistContext {
	var p = new(ExplistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_explist
	return p
}

func (*ExplistContext) IsExplistContext() {}

func NewExplistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExplistContext {
	var p = new(ExplistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_explist

	return p
}

func (s *ExplistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExplistContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExplistContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExplistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExplistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExplist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Explist() (localctx IExplistContext) {
	localctx = NewExplistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LuaParserRULE_explist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.exp(0)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__10 {
		{
			p.SetState(251)
			p.Match(LuaParserT__10)
		}
		{
			p.SetState(252)
			p.exp(0)
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) BoolLiteral() IBoolLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolLiteralContext)
}

func (s *ExpContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *ExpContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ExpContext) Functiondef() IFunctiondefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiondefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiondefContext)
}

func (s *ExpContext) Prefixexp() IPrefixexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixexpContext)
}

func (s *ExpContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ExpContext) OperatorUnary() IOperatorUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorUnaryContext)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) OperatorPower() IOperatorPowerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorPowerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorPowerContext)
}

func (s *ExpContext) OperatorMulDivMod() IOperatorMulDivModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorMulDivModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorMulDivModContext)
}

func (s *ExpContext) OperatorAddSub() IOperatorAddSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAddSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAddSubContext)
}

func (s *ExpContext) OperatorStrcat() IOperatorStrcatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorStrcatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorStrcatContext)
}

func (s *ExpContext) OperatorComparison() IOperatorComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorComparisonContext)
}

func (s *ExpContext) OperatorAnd() IOperatorAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAndContext)
}

func (s *ExpContext) OperatorOr() IOperatorOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorOrContext)
}

func (s *ExpContext) OperatorBitwise() IOperatorBitwiseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorBitwiseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorBitwiseContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *LuaParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, LuaParserRULE_exp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__22:
		{
			p.SetState(259)
			p.Match(LuaParserT__22)
		}

	case LuaParserT__24, LuaParserT__25:
		{
			p.SetState(260)
			p.BoolLiteral()
		}

	case LuaParserINT, LuaParserHEX, LuaParserFLOAT, LuaParserHEX_FLOAT:
		{
			p.SetState(261)
			p.NumberLiteral()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		{
			p.SetState(262)
			p.StringLiteral()
		}

	case LuaParserT__23:
		{
			p.SetState(263)
			p.Match(LuaParserT__23)
		}

	case LuaParserT__12:
		{
			p.SetState(264)
			p.Functiondef()
		}

	case LuaParserT__29, LuaParserNAME:
		{
			p.SetState(265)
			p.Prefixexp()
		}

	case LuaParserT__33:
		{
			p.SetState(266)
			p.Tableconstructor()
		}

	case LuaParserT__45, LuaParserT__52, LuaParserT__55, LuaParserT__56:
		{
			p.SetState(267)
			p.OperatorUnary()
		}
		{
			p.SetState(268)
			p.exp(8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(273)
					p.OperatorPower()
				}
				{
					p.SetState(274)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(277)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(278)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(281)
					p.OperatorAddSub()
				}
				{
					p.SetState(282)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(285)
					p.OperatorStrcat()
				}
				{
					p.SetState(286)
					p.exp(5)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(289)
					p.OperatorComparison()
				}
				{
					p.SetState(290)
					p.exp(5)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(293)
					p.OperatorAnd()
				}
				{
					p.SetState(294)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(297)
					p.OperatorOr()
				}
				{
					p.SetState(298)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(301)
					p.OperatorBitwise()
				}
				{
					p.SetState(302)
					p.exp(2)
				}

			}

		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolLiteralContext is an interface to support dynamic dispatch.
type IBoolLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolLiteralContext differentiates from other interfaces.
	IsBoolLiteralContext()
}

type BoolLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLiteralContext() *BoolLiteralContext {
	var p = new(BoolLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_boolLiteral
	return p
}

func (*BoolLiteralContext) IsBoolLiteralContext() {}

func NewBoolLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_boolLiteral

	return p
}

func (s *BoolLiteralContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) BoolLiteral() (localctx IBoolLiteralContext) {
	localctx = NewBoolLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LuaParserRULE_boolLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__24 || _la == LuaParserT__25) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeLiteralContext is an interface to support dynamic dispatch.
type ITypeLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLiteralContext differentiates from other interfaces.
	IsTypeLiteralContext()
}

type TypeLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLiteralContext() *TypeLiteralContext {
	var p = new(TypeLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_typeLiteral
	return p
}

func (*TypeLiteralContext) IsTypeLiteralContext() {}

func NewTypeLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLiteralContext {
	var p = new(TypeLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_typeLiteral

	return p
}

func (s *TypeLiteralContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTypeLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) TypeLiteral() (localctx ITypeLiteralContext) {
	localctx = NewTypeLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LuaParserRULE_typeLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrefixexpContext is an interface to support dynamic dispatch.
type IPrefixexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixexpContext differentiates from other interfaces.
	IsPrefixexpContext()
}

type PrefixexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixexpContext() *PrefixexpContext {
	var p = new(PrefixexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_prefixexp
	return p
}

func (*PrefixexpContext) IsPrefixexpContext() {}

func NewPrefixexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixexpContext {
	var p = new(PrefixexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_prefixexp

	return p
}

func (s *PrefixexpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixexpContext) VarOrExp() IVarOrExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *PrefixexpContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *PrefixexpContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *PrefixexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitPrefixexp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Prefixexp() (localctx IPrefixexpContext) {
	localctx = NewPrefixexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LuaParserRULE_prefixexp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.VarOrExp()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(314)
				p.NameAndArgs()
			}

		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_functioncall
	return p
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) VarOrExp() IVarOrExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *FunctioncallContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *FunctioncallContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFunctioncall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Functioncall() (localctx IFunctioncallContext) {
	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LuaParserRULE_functioncall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.VarOrExp()
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(321)
				p.NameAndArgs()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IVarOrExpContext is an interface to support dynamic dispatch.
type IVarOrExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarOrExpContext differentiates from other interfaces.
	IsVarOrExpContext()
}

type VarOrExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarOrExpContext() *VarOrExpContext {
	var p = new(VarOrExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_varOrExp
	return p
}

func (*VarOrExpContext) IsVarOrExpContext() {}

func NewVarOrExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrExpContext {
	var p = new(VarOrExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varOrExp

	return p
}

func (s *VarOrExpContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrExpContext) VarId() IVarIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarIdContext)
}

func (s *VarOrExpContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarOrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarOrExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) VarOrExp() (localctx IVarOrExpContext) {
	localctx = NewVarOrExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LuaParserRULE_varOrExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.VarId()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(328)
			p.exp(0)
		}
		{
			p.SetState(329)
			p.Match(LuaParserT__30)
		}

	}

	return localctx
}

// IVarIdContext is an interface to support dynamic dispatch.
type IVarIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarIdContext differentiates from other interfaces.
	IsVarIdContext()
}

type VarIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarIdContext() *VarIdContext {
	var p = new(VarIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_varId
	return p
}

func (*VarIdContext) IsVarIdContext() {}

func NewVarIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarIdContext {
	var p = new(VarIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varId

	return p
}

func (s *VarIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VarIdContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *VarIdContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarIdContext) AllVarSuffix() []IVarSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarSuffixContext)(nil)).Elem())
	var tst = make([]IVarSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarSuffixContext)
		}
	}

	return tst
}

func (s *VarIdContext) VarSuffix(i int) IVarSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarSuffixContext)
}

func (s *VarIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) VarId() (localctx IVarIdContext) {
	localctx = NewVarIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LuaParserRULE_varId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		{
			p.SetState(333)
			p.Match(LuaParserNAME)
		}

	case LuaParserT__29:
		{
			p.SetState(334)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(335)
			p.exp(0)
		}
		{
			p.SetState(336)
			p.Match(LuaParserT__30)
		}
		{
			p.SetState(337)
			p.VarSuffix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(341)
				p.VarSuffix()
			}

		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// ITypedvarContext is an interface to support dynamic dispatch.
type ITypedvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedvarContext differentiates from other interfaces.
	IsTypedvarContext()
}

type TypedvarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedvarContext() *TypedvarContext {
	var p = new(TypedvarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_typedvar
	return p
}

func (*TypedvarContext) IsTypedvarContext() {}

func NewTypedvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedvarContext {
	var p = new(TypedvarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_typedvar

	return p
}

func (s *TypedvarContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedvarContext) TypeLiteral() ITypeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLiteralContext)
}

func (s *TypedvarContext) VarId() IVarIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarIdContext)
}

func (s *TypedvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedvarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTypedvar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Typedvar() (localctx ITypedvarContext) {
	localctx = NewTypedvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LuaParserRULE_typedvar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.TypeLiteral()
	}
	{
		p.SetState(348)
		p.VarId()
	}

	return localctx
}

// IVarSuffixContext is an interface to support dynamic dispatch.
type IVarSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarSuffixContext differentiates from other interfaces.
	IsVarSuffixContext()
}

type VarSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSuffixContext() *VarSuffixContext {
	var p = new(VarSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_varSuffix
	return p
}

func (*VarSuffixContext) IsVarSuffixContext() {}

func NewVarSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSuffixContext {
	var p = new(VarSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_varSuffix

	return p
}

func (s *VarSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSuffixContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarSuffixContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *VarSuffixContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *VarSuffixContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *VarSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVarSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) VarSuffix() (localctx IVarSuffixContext) {
	localctx = NewVarSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LuaParserRULE_varSuffix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__17 || _la == LuaParserT__29 || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(LuaParserT__33-34))|(1<<(LuaParserNORMALSTRING-34))|(1<<(LuaParserCHARSTRING-34))|(1<<(LuaParserLONGSTRING-34)))) != 0) {
		{
			p.SetState(350)
			p.NameAndArgs()
		}

		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__31:
		{
			p.SetState(356)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(357)
			p.exp(0)
		}
		{
			p.SetState(358)
			p.Match(LuaParserT__32)
		}

	case LuaParserT__16:
		{
			p.SetState(360)
			p.Match(LuaParserT__16)
		}
		{
			p.SetState(361)
			p.Match(LuaParserNAME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameAndArgsContext is an interface to support dynamic dispatch.
type INameAndArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameAndArgsContext differentiates from other interfaces.
	IsNameAndArgsContext()
}

type NameAndArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameAndArgsContext() *NameAndArgsContext {
	var p = new(NameAndArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_nameAndArgs
	return p
}

func (*NameAndArgsContext) IsNameAndArgsContext() {}

func NewNameAndArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAndArgsContext {
	var p = new(NameAndArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_nameAndArgs

	return p
}

func (s *NameAndArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *NameAndArgsContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *NameAndArgsContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *NameAndArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameAndArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameAndArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNameAndArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) NameAndArgs() (localctx INameAndArgsContext) {
	localctx = NewNameAndArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LuaParserRULE_nameAndArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__17 {
		{
			p.SetState(364)
			p.Match(LuaParserT__17)
		}
		{
			p.SetState(365)
			p.Match(LuaParserNAME)
		}

	}
	{
		p.SetState(368)
		p.Args()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ArgsContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ArgsContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LuaParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(370)
			p.Match(LuaParserT__29)
		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(LuaParserT__12-13))|(1<<(LuaParserT__22-13))|(1<<(LuaParserT__23-13))|(1<<(LuaParserT__24-13))|(1<<(LuaParserT__25-13))|(1<<(LuaParserT__29-13))|(1<<(LuaParserT__33-13)))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46))|(1<<(LuaParserNAME-46))|(1<<(LuaParserNORMALSTRING-46))|(1<<(LuaParserCHARSTRING-46))|(1<<(LuaParserLONGSTRING-46))|(1<<(LuaParserINT-46))|(1<<(LuaParserHEX-46))|(1<<(LuaParserFLOAT-46))|(1<<(LuaParserHEX_FLOAT-46)))) != 0) {
			{
				p.SetState(371)
				p.Explist()
			}

		}
		{
			p.SetState(374)
			p.Match(LuaParserT__30)
		}

	case LuaParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Tableconstructor()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(376)
			p.StringLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctiondefContext is an interface to support dynamic dispatch.
type IFunctiondefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctiondefContext differentiates from other interfaces.
	IsFunctiondefContext()
}

type FunctiondefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiondefContext() *FunctiondefContext {
	var p = new(FunctiondefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_functiondef
	return p
}

func (*FunctiondefContext) IsFunctiondefContext() {}

func NewFunctiondefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiondefContext {
	var p = new(FunctiondefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_functiondef

	return p
}

func (s *FunctiondefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiondefContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiondefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFunctiondef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Functiondef() (localctx IFunctiondefContext) {
	localctx = NewFunctiondefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LuaParserRULE_functiondef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Match(LuaParserT__12)
	}
	{
		p.SetState(380)
		p.Funcbody()
	}

	return localctx
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_funcbody
	return p
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncbodyContext) Parlist() IParlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParlistContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFuncbody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Funcbody() (localctx IFuncbodyContext) {
	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LuaParserRULE_funcbody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(LuaParserT__29)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__23 || _la == LuaParserNAME {
		{
			p.SetState(383)
			p.Parlist()
		}

	}
	{
		p.SetState(386)
		p.Match(LuaParserT__30)
	}
	{
		p.SetState(387)
		p.Block()
	}
	{
		p.SetState(388)
		p.Match(LuaParserT__4)
	}

	return localctx
}

// IParlistContext is an interface to support dynamic dispatch.
type IParlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParlistContext differentiates from other interfaces.
	IsParlistContext()
}

type ParlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParlistContext() *ParlistContext {
	var p = new(ParlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_parlist
	return p
}

func (*ParlistContext) IsParlistContext() {}

func NewParlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParlistContext {
	var p = new(ParlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_parlist

	return p
}

func (s *ParlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParlistContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *ParlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitParlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Parlist() (localctx IParlistContext) {
	localctx = NewParlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LuaParserRULE_parlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(396)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.Namelist()
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__10 {
			{
				p.SetState(391)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(392)
				p.Match(LuaParserT__23)
			}

		}

	case LuaParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(LuaParserT__23)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableconstructorContext is an interface to support dynamic dispatch.
type ITableconstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableconstructorContext differentiates from other interfaces.
	IsTableconstructorContext()
}

type TableconstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableconstructorContext() *TableconstructorContext {
	var p = new(TableconstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_tableconstructor
	return p
}

func (*TableconstructorContext) IsTableconstructorContext() {}

func NewTableconstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableconstructorContext {
	var p = new(TableconstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_tableconstructor

	return p
}

func (s *TableconstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableconstructorContext) Fieldlist() IFieldlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldlistContext)
}

func (s *TableconstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableconstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableconstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTableconstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Tableconstructor() (localctx ITableconstructorContext) {
	localctx = NewTableconstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LuaParserRULE_tableconstructor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(LuaParserT__33)
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(LuaParserT__12-13))|(1<<(LuaParserT__22-13))|(1<<(LuaParserT__23-13))|(1<<(LuaParserT__24-13))|(1<<(LuaParserT__25-13))|(1<<(LuaParserT__29-13))|(1<<(LuaParserT__31-13))|(1<<(LuaParserT__33-13)))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46))|(1<<(LuaParserNAME-46))|(1<<(LuaParserNORMALSTRING-46))|(1<<(LuaParserCHARSTRING-46))|(1<<(LuaParserLONGSTRING-46))|(1<<(LuaParserINT-46))|(1<<(LuaParserHEX-46))|(1<<(LuaParserFLOAT-46))|(1<<(LuaParserHEX_FLOAT-46)))) != 0) {
		{
			p.SetState(399)
			p.Fieldlist()
		}

	}
	{
		p.SetState(402)
		p.Match(LuaParserT__34)
	}

	return localctx
}

// IFieldlistContext is an interface to support dynamic dispatch.
type IFieldlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldlistContext differentiates from other interfaces.
	IsFieldlistContext()
}

type FieldlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldlistContext() *FieldlistContext {
	var p = new(FieldlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_fieldlist
	return p
}

func (*FieldlistContext) IsFieldlistContext() {}

func NewFieldlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldlistContext {
	var p = new(FieldlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_fieldlist

	return p
}

func (s *FieldlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldlistContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldlistContext) AllFieldsep() []IFieldsepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldsepContext)(nil)).Elem())
	var tst = make([]IFieldsepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldsepContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Fieldsep(i int) IFieldsepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldsepContext)
}

func (s *FieldlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFieldlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Fieldlist() (localctx IFieldlistContext) {
	localctx = NewFieldlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LuaParserRULE_fieldlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Field()
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(405)
				p.Fieldsep()
			}
			{
				p.SetState(406)
				p.Field()
			}

		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 || _la == LuaParserT__10 {
		{
			p.SetState(413)
			p.Fieldsep()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *FieldContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LuaParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(417)
			p.exp(0)
		}
		{
			p.SetState(418)
			p.Match(LuaParserT__32)
		}
		{
			p.SetState(419)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(420)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(423)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(424)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(425)
			p.exp(0)
		}

	}

	return localctx
}

// IFieldsepContext is an interface to support dynamic dispatch.
type IFieldsepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsepContext differentiates from other interfaces.
	IsFieldsepContext()
}

type FieldsepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsepContext() *FieldsepContext {
	var p = new(FieldsepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_fieldsep
	return p
}

func (*FieldsepContext) IsFieldsepContext() {}

func NewFieldsepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsepContext {
	var p = new(FieldsepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_fieldsep

	return p
}

func (s *FieldsepContext) GetParser() antlr.Parser { return s.parser }
func (s *FieldsepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitFieldsep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Fieldsep() (localctx IFieldsepContext) {
	localctx = NewFieldsepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LuaParserRULE_fieldsep)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(428)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__0 || _la == LuaParserT__10) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorOrContext is an interface to support dynamic dispatch.
type IOperatorOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorOrContext differentiates from other interfaces.
	IsOperatorOrContext()
}

type OperatorOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorOrContext() *OperatorOrContext {
	var p = new(OperatorOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorOr
	return p
}

func (*OperatorOrContext) IsOperatorOrContext() {}

func NewOperatorOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorOrContext {
	var p = new(OperatorOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorOr

	return p
}

func (s *OperatorOrContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorOr() (localctx IOperatorOrContext) {
	localctx = NewOperatorOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LuaParserRULE_operatorOr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(LuaParserT__35)
	}

	return localctx
}

// IOperatorAndContext is an interface to support dynamic dispatch.
type IOperatorAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAndContext differentiates from other interfaces.
	IsOperatorAndContext()
}

type OperatorAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAndContext() *OperatorAndContext {
	var p = new(OperatorAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAnd
	return p
}

func (*OperatorAndContext) IsOperatorAndContext() {}

func NewOperatorAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAndContext {
	var p = new(OperatorAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorAnd

	return p
}

func (s *OperatorAndContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorAnd() (localctx IOperatorAndContext) {
	localctx = NewOperatorAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LuaParserRULE_operatorAnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(LuaParserT__36)
	}

	return localctx
}

// IOperatorComparisonContext is an interface to support dynamic dispatch.
type IOperatorComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorComparisonContext differentiates from other interfaces.
	IsOperatorComparisonContext()
}

type OperatorComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorComparisonContext() *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorComparison
	return p
}

func (*OperatorComparisonContext) IsOperatorComparisonContext() {}

func NewOperatorComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorComparison

	return p
}

func (s *OperatorComparisonContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorComparison() (localctx IOperatorComparisonContext) {
	localctx = NewOperatorComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LuaParserRULE_operatorComparison)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(LuaParserT__37-38))|(1<<(LuaParserT__38-38))|(1<<(LuaParserT__39-38))|(1<<(LuaParserT__40-38))|(1<<(LuaParserT__41-38))|(1<<(LuaParserT__42-38)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorStrcatContext is an interface to support dynamic dispatch.
type IOperatorStrcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorStrcatContext differentiates from other interfaces.
	IsOperatorStrcatContext()
}

type OperatorStrcatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorStrcatContext() *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorStrcat
	return p
}

func (*OperatorStrcatContext) IsOperatorStrcatContext() {}

func NewOperatorStrcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorStrcat

	return p
}

func (s *OperatorStrcatContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorStrcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorStrcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorStrcatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorStrcat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorStrcat() (localctx IOperatorStrcatContext) {
	localctx = NewOperatorStrcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LuaParserRULE_operatorStrcat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.Match(LuaParserT__43)
	}

	return localctx
}

// IOperatorAddSubContext is an interface to support dynamic dispatch.
type IOperatorAddSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAddSubContext differentiates from other interfaces.
	IsOperatorAddSubContext()
}

type OperatorAddSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAddSubContext() *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAddSub
	return p
}

func (*OperatorAddSubContext) IsOperatorAddSubContext() {}

func NewOperatorAddSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorAddSub

	return p
}

func (s *OperatorAddSubContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAddSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorAddSub() (localctx IOperatorAddSubContext) {
	localctx = NewOperatorAddSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, LuaParserRULE_operatorAddSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__44 || _la == LuaParserT__45) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorMulDivModContext is an interface to support dynamic dispatch.
type IOperatorMulDivModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorMulDivModContext differentiates from other interfaces.
	IsOperatorMulDivModContext()
}

type OperatorMulDivModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorMulDivModContext() *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorMulDivMod
	return p
}

func (*OperatorMulDivModContext) IsOperatorMulDivModContext() {}

func NewOperatorMulDivModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorMulDivMod

	return p
}

func (s *OperatorMulDivModContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorMulDivModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorMulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorMulDivMod() (localctx IOperatorMulDivModContext) {
	localctx = NewOperatorMulDivModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, LuaParserRULE_operatorMulDivMod)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(LuaParserT__46-47))|(1<<(LuaParserT__47-47))|(1<<(LuaParserT__48-47))|(1<<(LuaParserT__49-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorBitwiseContext is an interface to support dynamic dispatch.
type IOperatorBitwiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorBitwiseContext differentiates from other interfaces.
	IsOperatorBitwiseContext()
}

type OperatorBitwiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorBitwiseContext() *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorBitwise
	return p
}

func (*OperatorBitwiseContext) IsOperatorBitwiseContext() {}

func NewOperatorBitwiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorBitwise

	return p
}

func (s *OperatorBitwiseContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorBitwiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorBitwiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorBitwiseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorBitwise(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorBitwise() (localctx IOperatorBitwiseContext) {
	localctx = NewOperatorBitwiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, LuaParserRULE_operatorBitwise)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(LuaParserT__50-51))|(1<<(LuaParserT__51-51))|(1<<(LuaParserT__52-51))|(1<<(LuaParserT__53-51))|(1<<(LuaParserT__54-51)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorUnaryContext is an interface to support dynamic dispatch.
type IOperatorUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorUnaryContext differentiates from other interfaces.
	IsOperatorUnaryContext()
}

type OperatorUnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorUnaryContext() *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorUnary
	return p
}

func (*OperatorUnaryContext) IsOperatorUnaryContext() {}

func NewOperatorUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorUnary

	return p
}

func (s *OperatorUnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorUnary() (localctx IOperatorUnaryContext) {
	localctx = NewOperatorUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, LuaParserRULE_operatorUnary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorPowerContext is an interface to support dynamic dispatch.
type IOperatorPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorPowerContext differentiates from other interfaces.
	IsOperatorPowerContext()
}

type OperatorPowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorPowerContext() *OperatorPowerContext {
	var p = new(OperatorPowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorPower
	return p
}

func (*OperatorPowerContext) IsOperatorPowerContext() {}

func NewOperatorPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorPowerContext {
	var p = new(OperatorPowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorPower

	return p
}

func (s *OperatorPowerContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorPowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorPowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitOperatorPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) OperatorPower() (localctx IOperatorPowerContext) {
	localctx = NewOperatorPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, LuaParserRULE_operatorPower)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(LuaParserT__57)
	}

	return localctx
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_numberLiteral
	return p
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(LuaParserINT, 0)
}

func (s *NumberLiteralContext) HEX() antlr.TerminalNode {
	return s.GetToken(LuaParserHEX, 0)
}

func (s *NumberLiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(LuaParserFLOAT, 0)
}

func (s *NumberLiteralContext) HEX_FLOAT() antlr.TerminalNode {
	return s.GetToken(LuaParserHEX_FLOAT, 0)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) NumberLiteral() (localctx INumberLiteralContext) {
	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, LuaParserRULE_numberLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(LuaParserINT-63))|(1<<(LuaParserHEX-63))|(1<<(LuaParserFLOAT-63))|(1<<(LuaParserHEX_FLOAT-63)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) NORMALSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserNORMALSTRING, 0)
}

func (s *StringLiteralContext) CHARSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserCHARSTRING, 0)
}

func (s *StringLiteralContext) LONGSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserLONGSTRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, LuaParserRULE_stringLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(LuaParserNORMALSTRING-60))|(1<<(LuaParserCHARSTRING-60))|(1<<(LuaParserLONGSTRING-60)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *LuaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LuaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
