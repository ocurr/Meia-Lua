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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 433,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 83, 10, 3, 12, 3, 14, 3, 86, 11, 3, 3,
	3, 5, 3, 89, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 125, 10, 4, 12, 4, 14, 4, 128, 11, 4, 3, 4, 3, 4, 5, 4,
	132, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 144, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 170, 10, 4, 5, 4, 172, 10, 4, 3, 5, 3, 5, 5, 5,
	176, 10, 5, 3, 5, 5, 5, 179, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 7, 7, 188, 10, 7, 12, 7, 14, 7, 191, 11, 7, 3, 7, 3, 7, 5, 7, 195,
	10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 200, 10, 8, 12, 8, 14, 8, 203, 11, 8, 3,
	9, 3, 9, 3, 9, 7, 9, 208, 10, 9, 12, 9, 14, 9, 211, 11, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 7, 9, 217, 10, 9, 12, 9, 14, 9, 220, 11, 9, 5, 9, 222, 10,
	9, 3, 10, 3, 10, 3, 10, 7, 10, 227, 10, 10, 12, 10, 14, 10, 230, 11, 10,
	3, 11, 3, 11, 3, 11, 7, 11, 235, 10, 11, 12, 11, 14, 11, 238, 11, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 253, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 287, 10, 12, 12, 12, 14,
	12, 290, 11, 12, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 296, 10, 14, 12, 14,
	14, 14, 299, 11, 14, 3, 15, 3, 15, 6, 15, 303, 10, 15, 13, 15, 14, 15,
	304, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 312, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 320, 10, 17, 3, 17, 7, 17, 323, 10,
	17, 12, 17, 14, 17, 326, 11, 17, 3, 18, 3, 18, 3, 18, 3, 19, 7, 19, 332,
	10, 19, 12, 19, 14, 19, 335, 11, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 5, 19, 343, 10, 19, 3, 20, 3, 20, 5, 20, 347, 10, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 5, 21, 353, 10, 21, 3, 21, 3, 21, 3, 21, 5, 21, 358,
	10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 365, 10, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 374, 10, 24, 3, 24, 5, 24,
	377, 10, 24, 3, 25, 3, 25, 5, 25, 381, 10, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 7, 26, 389, 10, 26, 12, 26, 14, 26, 392, 11, 26, 3, 26,
	5, 26, 395, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 5, 27, 407, 10, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 2, 3,
	22, 40, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 2, 11, 3, 2, 29, 31, 4, 2, 3, 3, 17, 17, 3, 2, 40, 45, 3, 2,
	47, 48, 3, 2, 49, 52, 3, 2, 53, 57, 5, 2, 48, 48, 55, 55, 58, 59, 3, 2,
	65, 68, 3, 2, 62, 64, 2, 460, 2, 78, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6,
	171, 3, 2, 2, 2, 8, 173, 3, 2, 2, 2, 10, 180, 3, 2, 2, 2, 12, 184, 3, 2,
	2, 2, 14, 196, 3, 2, 2, 2, 16, 221, 3, 2, 2, 2, 18, 223, 3, 2, 2, 2, 20,
	231, 3, 2, 2, 2, 22, 252, 3, 2, 2, 2, 24, 291, 3, 2, 2, 2, 26, 293, 3,
	2, 2, 2, 28, 300, 3, 2, 2, 2, 30, 311, 3, 2, 2, 2, 32, 319, 3, 2, 2, 2,
	34, 327, 3, 2, 2, 2, 36, 333, 3, 2, 2, 2, 38, 346, 3, 2, 2, 2, 40, 357,
	3, 2, 2, 2, 42, 359, 3, 2, 2, 2, 44, 362, 3, 2, 2, 2, 46, 376, 3, 2, 2,
	2, 48, 378, 3, 2, 2, 2, 50, 384, 3, 2, 2, 2, 52, 406, 3, 2, 2, 2, 54, 408,
	3, 2, 2, 2, 56, 410, 3, 2, 2, 2, 58, 412, 3, 2, 2, 2, 60, 414, 3, 2, 2,
	2, 62, 416, 3, 2, 2, 2, 64, 418, 3, 2, 2, 2, 66, 420, 3, 2, 2, 2, 68, 422,
	3, 2, 2, 2, 70, 424, 3, 2, 2, 2, 72, 426, 3, 2, 2, 2, 74, 428, 3, 2, 2,
	2, 76, 430, 3, 2, 2, 2, 78, 79, 5, 4, 3, 2, 79, 80, 7, 2, 2, 3, 80, 3,
	3, 2, 2, 2, 81, 83, 5, 6, 4, 2, 82, 81, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2,
	84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3,
	2, 2, 2, 87, 89, 5, 8, 5, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89,
	5, 3, 2, 2, 2, 90, 172, 7, 3, 2, 2, 91, 92, 5, 16, 9, 2, 92, 93, 7, 4,
	2, 2, 93, 94, 5, 20, 11, 2, 94, 172, 3, 2, 2, 2, 95, 172, 5, 28, 15, 2,
	96, 172, 5, 10, 6, 2, 97, 172, 7, 5, 2, 2, 98, 99, 7, 6, 2, 2, 99, 172,
	7, 61, 2, 2, 100, 101, 7, 7, 2, 2, 101, 102, 5, 4, 3, 2, 102, 103, 7, 8,
	2, 2, 103, 172, 3, 2, 2, 2, 104, 105, 7, 9, 2, 2, 105, 106, 5, 22, 12,
	2, 106, 107, 7, 7, 2, 2, 107, 108, 5, 4, 3, 2, 108, 109, 7, 8, 2, 2, 109,
	172, 3, 2, 2, 2, 110, 111, 7, 10, 2, 2, 111, 112, 5, 4, 3, 2, 112, 113,
	7, 11, 2, 2, 113, 114, 5, 22, 12, 2, 114, 172, 3, 2, 2, 2, 115, 116, 7,
	12, 2, 2, 116, 117, 5, 22, 12, 2, 117, 118, 7, 13, 2, 2, 118, 126, 5, 4,
	3, 2, 119, 120, 7, 14, 2, 2, 120, 121, 5, 22, 12, 2, 121, 122, 7, 13, 2,
	2, 122, 123, 5, 4, 3, 2, 123, 125, 3, 2, 2, 2, 124, 119, 3, 2, 2, 2, 125,
	128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 131,
	3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130, 7, 15, 2, 2, 130, 132, 5, 4,
	3, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2,
	133, 134, 7, 8, 2, 2, 134, 172, 3, 2, 2, 2, 135, 136, 7, 16, 2, 2, 136,
	137, 7, 61, 2, 2, 137, 138, 7, 4, 2, 2, 138, 139, 5, 22, 12, 2, 139, 140,
	7, 17, 2, 2, 140, 143, 5, 22, 12, 2, 141, 142, 7, 17, 2, 2, 142, 144, 5,
	22, 12, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2,
	2, 2, 145, 146, 7, 7, 2, 2, 146, 147, 5, 4, 3, 2, 147, 148, 7, 8, 2, 2,
	148, 172, 3, 2, 2, 2, 149, 150, 7, 16, 2, 2, 150, 151, 5, 18, 10, 2, 151,
	152, 7, 18, 2, 2, 152, 153, 5, 20, 11, 2, 153, 154, 7, 7, 2, 2, 154, 155,
	5, 4, 3, 2, 155, 156, 7, 8, 2, 2, 156, 172, 3, 2, 2, 2, 157, 158, 7, 19,
	2, 2, 158, 159, 5, 12, 7, 2, 159, 160, 5, 44, 23, 2, 160, 172, 3, 2, 2,
	2, 161, 162, 7, 20, 2, 2, 162, 163, 7, 19, 2, 2, 163, 164, 7, 61, 2, 2,
	164, 172, 5, 44, 23, 2, 165, 166, 7, 20, 2, 2, 166, 169, 5, 18, 10, 2,
	167, 168, 7, 4, 2, 2, 168, 170, 5, 20, 11, 2, 169, 167, 3, 2, 2, 2, 169,
	170, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171, 90, 3, 2, 2, 2, 171, 91, 3,
	2, 2, 2, 171, 95, 3, 2, 2, 2, 171, 96, 3, 2, 2, 2, 171, 97, 3, 2, 2, 2,
	171, 98, 3, 2, 2, 2, 171, 100, 3, 2, 2, 2, 171, 104, 3, 2, 2, 2, 171, 110,
	3, 2, 2, 2, 171, 115, 3, 2, 2, 2, 171, 135, 3, 2, 2, 2, 171, 149, 3, 2,
	2, 2, 171, 157, 3, 2, 2, 2, 171, 161, 3, 2, 2, 2, 171, 165, 3, 2, 2, 2,
	172, 7, 3, 2, 2, 2, 173, 175, 7, 21, 2, 2, 174, 176, 5, 20, 11, 2, 175,
	174, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 3, 2, 2, 2, 177, 179,
	7, 3, 2, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 9, 3, 2, 2,
	2, 180, 181, 7, 22, 2, 2, 181, 182, 7, 61, 2, 2, 182, 183, 7, 22, 2, 2,
	183, 11, 3, 2, 2, 2, 184, 189, 7, 61, 2, 2, 185, 186, 7, 23, 2, 2, 186,
	188, 7, 61, 2, 2, 187, 185, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189, 187,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 194, 3, 2, 2, 2, 191, 189, 3, 2,
	2, 2, 192, 193, 7, 24, 2, 2, 193, 195, 7, 61, 2, 2, 194, 192, 3, 2, 2,
	2, 194, 195, 3, 2, 2, 2, 195, 13, 3, 2, 2, 2, 196, 201, 5, 32, 17, 2, 197,
	198, 7, 17, 2, 2, 198, 200, 5, 32, 17, 2, 199, 197, 3, 2, 2, 2, 200, 203,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 15, 3, 2,
	2, 2, 203, 201, 3, 2, 2, 2, 204, 209, 5, 34, 18, 2, 205, 206, 7, 17, 2,
	2, 206, 208, 5, 34, 18, 2, 207, 205, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2,
	209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 222, 3, 2, 2, 2, 211,
	209, 3, 2, 2, 2, 212, 213, 5, 24, 13, 2, 213, 218, 5, 32, 17, 2, 214, 215,
	7, 17, 2, 2, 215, 217, 5, 32, 17, 2, 216, 214, 3, 2, 2, 2, 217, 220, 3,
	2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 222, 3, 2, 2,
	2, 220, 218, 3, 2, 2, 2, 221, 204, 3, 2, 2, 2, 221, 212, 3, 2, 2, 2, 222,
	17, 3, 2, 2, 2, 223, 228, 7, 61, 2, 2, 224, 225, 7, 17, 2, 2, 225, 227,
	7, 61, 2, 2, 226, 224, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2,
	2, 2, 228, 229, 3, 2, 2, 2, 229, 19, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2,
	231, 236, 5, 22, 12, 2, 232, 233, 7, 17, 2, 2, 233, 235, 5, 22, 12, 2,
	234, 232, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236,
	237, 3, 2, 2, 2, 237, 21, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 240, 8,
	12, 1, 2, 240, 253, 7, 25, 2, 2, 241, 253, 7, 26, 2, 2, 242, 253, 7, 27,
	2, 2, 243, 253, 5, 74, 38, 2, 244, 253, 5, 76, 39, 2, 245, 253, 7, 28,
	2, 2, 246, 253, 5, 42, 22, 2, 247, 253, 5, 26, 14, 2, 248, 253, 5, 48,
	25, 2, 249, 250, 5, 70, 36, 2, 250, 251, 5, 22, 12, 10, 251, 253, 3, 2,
	2, 2, 252, 239, 3, 2, 2, 2, 252, 241, 3, 2, 2, 2, 252, 242, 3, 2, 2, 2,
	252, 243, 3, 2, 2, 2, 252, 244, 3, 2, 2, 2, 252, 245, 3, 2, 2, 2, 252,
	246, 3, 2, 2, 2, 252, 247, 3, 2, 2, 2, 252, 248, 3, 2, 2, 2, 252, 249,
	3, 2, 2, 2, 253, 288, 3, 2, 2, 2, 254, 255, 12, 11, 2, 2, 255, 256, 5,
	72, 37, 2, 256, 257, 5, 22, 12, 11, 257, 287, 3, 2, 2, 2, 258, 259, 12,
	9, 2, 2, 259, 260, 5, 66, 34, 2, 260, 261, 5, 22, 12, 10, 261, 287, 3,
	2, 2, 2, 262, 263, 12, 8, 2, 2, 263, 264, 5, 64, 33, 2, 264, 265, 5, 22,
	12, 9, 265, 287, 3, 2, 2, 2, 266, 267, 12, 7, 2, 2, 267, 268, 5, 62, 32,
	2, 268, 269, 5, 22, 12, 7, 269, 287, 3, 2, 2, 2, 270, 271, 12, 6, 2, 2,
	271, 272, 5, 60, 31, 2, 272, 273, 5, 22, 12, 7, 273, 287, 3, 2, 2, 2, 274,
	275, 12, 5, 2, 2, 275, 276, 5, 58, 30, 2, 276, 277, 5, 22, 12, 6, 277,
	287, 3, 2, 2, 2, 278, 279, 12, 4, 2, 2, 279, 280, 5, 56, 29, 2, 280, 281,
	5, 22, 12, 5, 281, 287, 3, 2, 2, 2, 282, 283, 12, 3, 2, 2, 283, 284, 5,
	68, 35, 2, 284, 285, 5, 22, 12, 4, 285, 287, 3, 2, 2, 2, 286, 254, 3, 2,
	2, 2, 286, 258, 3, 2, 2, 2, 286, 262, 3, 2, 2, 2, 286, 266, 3, 2, 2, 2,
	286, 270, 3, 2, 2, 2, 286, 274, 3, 2, 2, 2, 286, 278, 3, 2, 2, 2, 286,
	282, 3, 2, 2, 2, 287, 290, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289,
	3, 2, 2, 2, 289, 23, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 291, 292, 9, 2,
	2, 2, 292, 25, 3, 2, 2, 2, 293, 297, 5, 30, 16, 2, 294, 296, 5, 38, 20,
	2, 295, 294, 3, 2, 2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297,
	298, 3, 2, 2, 2, 298, 27, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 302, 5,
	30, 16, 2, 301, 303, 5, 38, 20, 2, 302, 301, 3, 2, 2, 2, 303, 304, 3, 2,
	2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 29, 3, 2, 2, 2,
	306, 312, 5, 32, 17, 2, 307, 308, 7, 32, 2, 2, 308, 309, 5, 22, 12, 2,
	309, 310, 7, 33, 2, 2, 310, 312, 3, 2, 2, 2, 311, 306, 3, 2, 2, 2, 311,
	307, 3, 2, 2, 2, 312, 31, 3, 2, 2, 2, 313, 320, 7, 61, 2, 2, 314, 315,
	7, 32, 2, 2, 315, 316, 5, 22, 12, 2, 316, 317, 7, 33, 2, 2, 317, 318, 5,
	36, 19, 2, 318, 320, 3, 2, 2, 2, 319, 313, 3, 2, 2, 2, 319, 314, 3, 2,
	2, 2, 320, 324, 3, 2, 2, 2, 321, 323, 5, 36, 19, 2, 322, 321, 3, 2, 2,
	2, 323, 326, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325,
	33, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 327, 328, 5, 24, 13, 2, 328, 329,
	5, 32, 17, 2, 329, 35, 3, 2, 2, 2, 330, 332, 5, 38, 20, 2, 331, 330, 3,
	2, 2, 2, 332, 335, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2, 333, 334, 3, 2, 2,
	2, 334, 342, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2, 336, 337, 7, 34, 2, 2, 337,
	338, 5, 22, 12, 2, 338, 339, 7, 35, 2, 2, 339, 343, 3, 2, 2, 2, 340, 341,
	7, 23, 2, 2, 341, 343, 7, 61, 2, 2, 342, 336, 3, 2, 2, 2, 342, 340, 3,
	2, 2, 2, 343, 37, 3, 2, 2, 2, 344, 345, 7, 24, 2, 2, 345, 347, 7, 61, 2,
	2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348,
	349, 5, 40, 21, 2, 349, 39, 3, 2, 2, 2, 350, 352, 7, 32, 2, 2, 351, 353,
	5, 20, 11, 2, 352, 351, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 354, 3,
	2, 2, 2, 354, 358, 7, 33, 2, 2, 355, 358, 5, 48, 25, 2, 356, 358, 5, 76,
	39, 2, 357, 350, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 357, 356, 3, 2, 2, 2,
	358, 41, 3, 2, 2, 2, 359, 360, 7, 19, 2, 2, 360, 361, 5, 44, 23, 2, 361,
	43, 3, 2, 2, 2, 362, 364, 7, 32, 2, 2, 363, 365, 5, 46, 24, 2, 364, 363,
	3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 7, 33,
	2, 2, 367, 368, 5, 4, 3, 2, 368, 369, 7, 8, 2, 2, 369, 45, 3, 2, 2, 2,
	370, 373, 5, 18, 10, 2, 371, 372, 7, 17, 2, 2, 372, 374, 7, 28, 2, 2, 373,
	371, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 377, 3, 2, 2, 2, 375, 377,
	7, 28, 2, 2, 376, 370, 3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 47, 3, 2,
	2, 2, 378, 380, 7, 36, 2, 2, 379, 381, 5, 50, 26, 2, 380, 379, 3, 2, 2,
	2, 380, 381, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 7, 37, 2, 2, 383,
	49, 3, 2, 2, 2, 384, 390, 5, 52, 27, 2, 385, 386, 5, 54, 28, 2, 386, 387,
	5, 52, 27, 2, 387, 389, 3, 2, 2, 2, 388, 385, 3, 2, 2, 2, 389, 392, 3,
	2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 394, 3, 2, 2,
	2, 392, 390, 3, 2, 2, 2, 393, 395, 5, 54, 28, 2, 394, 393, 3, 2, 2, 2,
	394, 395, 3, 2, 2, 2, 395, 51, 3, 2, 2, 2, 396, 397, 7, 34, 2, 2, 397,
	398, 5, 22, 12, 2, 398, 399, 7, 35, 2, 2, 399, 400, 7, 4, 2, 2, 400, 401,
	5, 22, 12, 2, 401, 407, 3, 2, 2, 2, 402, 403, 7, 61, 2, 2, 403, 404, 7,
	4, 2, 2, 404, 407, 5, 22, 12, 2, 405, 407, 5, 22, 12, 2, 406, 396, 3, 2,
	2, 2, 406, 402, 3, 2, 2, 2, 406, 405, 3, 2, 2, 2, 407, 53, 3, 2, 2, 2,
	408, 409, 9, 3, 2, 2, 409, 55, 3, 2, 2, 2, 410, 411, 7, 38, 2, 2, 411,
	57, 3, 2, 2, 2, 412, 413, 7, 39, 2, 2, 413, 59, 3, 2, 2, 2, 414, 415, 9,
	4, 2, 2, 415, 61, 3, 2, 2, 2, 416, 417, 7, 46, 2, 2, 417, 63, 3, 2, 2,
	2, 418, 419, 9, 5, 2, 2, 419, 65, 3, 2, 2, 2, 420, 421, 9, 6, 2, 2, 421,
	67, 3, 2, 2, 2, 422, 423, 9, 7, 2, 2, 423, 69, 3, 2, 2, 2, 424, 425, 9,
	8, 2, 2, 425, 71, 3, 2, 2, 2, 426, 427, 7, 60, 2, 2, 427, 73, 3, 2, 2,
	2, 428, 429, 9, 9, 2, 2, 429, 75, 3, 2, 2, 2, 430, 431, 9, 10, 2, 2, 431,
	77, 3, 2, 2, 2, 39, 84, 88, 126, 131, 143, 169, 171, 175, 178, 189, 194,
	201, 209, 218, 221, 228, 236, 252, 286, 288, 297, 304, 311, 319, 324, 333,
	342, 346, 352, 357, 364, 373, 376, 380, 390, 394, 406,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'='", "'break'", "'goto'", "'do'", "'end'", "'while'", "'repeat'",
	"'until'", "'if'", "'then'", "'elseif'", "'else'", "'for'", "','", "'in'",
	"'function'", "'local'", "'return'", "'::'", "'.'", "':'", "'nil'", "'false'",
	"'true'", "'...'", "'int'", "'float'", "'string'", "'('", "')'", "'['",
	"']'", "'{'", "'}'", "'or'", "'and'", "'<'", "'>'", "'<='", "'>='", "'~='",
	"'=='", "'..'", "'+'", "'-'", "'*'", "'/'", "'%'", "'//'", "'&'", "'|'",
	"'~'", "'<<'", "'>>'", "'not'", "'#'", "'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "NAME", "NORMALSTRING", "CHARSTRING", "LONGSTRING",
	"INT", "HEX", "FLOAT", "HEX_FLOAT", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"chunk", "block", "stat", "retstat", "label", "funcname", "varlist", "typedvarlist",
	"namelist", "explist", "exp", "typeLiteral", "prefixexp", "functioncall",
	"varOrExp", "varId", "typedvar", "varSuffix", "nameAndArgs", "args", "functiondef",
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
	LuaParserRULE_varlist            = 6
	LuaParserRULE_typedvarlist       = 7
	LuaParserRULE_namelist           = 8
	LuaParserRULE_explist            = 9
	LuaParserRULE_exp                = 10
	LuaParserRULE_typeLiteral        = 11
	LuaParserRULE_prefixexp          = 12
	LuaParserRULE_functioncall       = 13
	LuaParserRULE_varOrExp           = 14
	LuaParserRULE_varId              = 15
	LuaParserRULE_typedvar           = 16
	LuaParserRULE_varSuffix          = 17
	LuaParserRULE_nameAndArgs        = 18
	LuaParserRULE_args               = 19
	LuaParserRULE_functiondef        = 20
	LuaParserRULE_funcbody           = 21
	LuaParserRULE_parlist            = 22
	LuaParserRULE_tableconstructor   = 23
	LuaParserRULE_fieldlist          = 24
	LuaParserRULE_field              = 25
	LuaParserRULE_fieldsep           = 26
	LuaParserRULE_operatorOr         = 27
	LuaParserRULE_operatorAnd        = 28
	LuaParserRULE_operatorComparison = 29
	LuaParserRULE_operatorStrcat     = 30
	LuaParserRULE_operatorAddSub     = 31
	LuaParserRULE_operatorMulDivMod  = 32
	LuaParserRULE_operatorBitwise    = 33
	LuaParserRULE_operatorUnary      = 34
	LuaParserRULE_operatorPower      = 35
	LuaParserRULE_numberLiteral      = 36
	LuaParserRULE_stringLiteral      = 37
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

func (s *ChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterChunk(s)
	}
}

func (s *ChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitChunk(s)
	}
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
		p.SetState(76)
		p.Block()
	}
	{
		p.SetState(77)
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

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitBlock(s)
	}
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
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__0)|(1<<LuaParserT__2)|(1<<LuaParserT__3)|(1<<LuaParserT__4)|(1<<LuaParserT__6)|(1<<LuaParserT__7)|(1<<LuaParserT__9)|(1<<LuaParserT__13)|(1<<LuaParserT__16)|(1<<LuaParserT__17)|(1<<LuaParserT__19)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28)|(1<<LuaParserT__29))) != 0) || _la == LuaParserNAME {
		{
			p.SetState(79)
			p.Stat()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__18 {
		{
			p.SetState(85)
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

func (s *StatContext) Typedvarlist() ITypedvarlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedvarlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedvarlistContext)
}

func (s *StatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
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

func (s *StatContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *StatContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

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

func (s *StatContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
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

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStat(s)
	}
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(LuaParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Typedvarlist()
		}
		{
			p.SetState(90)
			p.Match(LuaParserT__1)
		}
		{
			p.SetState(91)
			p.Explist()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Functioncall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)
			p.Match(LuaParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(96)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(97)
			p.Match(LuaParserNAME)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.Match(LuaParserT__4)
		}
		{
			p.SetState(99)
			p.Block()
		}
		{
			p.SetState(100)
			p.Match(LuaParserT__5)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.Match(LuaParserT__6)
		}
		{
			p.SetState(103)
			p.exp(0)
		}
		{
			p.SetState(104)
			p.Match(LuaParserT__4)
		}
		{
			p.SetState(105)
			p.Block()
		}
		{
			p.SetState(106)
			p.Match(LuaParserT__5)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(108)
			p.Match(LuaParserT__7)
		}
		{
			p.SetState(109)
			p.Block()
		}
		{
			p.SetState(110)
			p.Match(LuaParserT__8)
		}
		{
			p.SetState(111)
			p.exp(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(113)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(114)
			p.exp(0)
		}
		{
			p.SetState(115)
			p.Match(LuaParserT__10)
		}
		{
			p.SetState(116)
			p.Block()
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__11 {
			{
				p.SetState(117)
				p.Match(LuaParserT__11)
			}
			{
				p.SetState(118)
				p.exp(0)
			}
			{
				p.SetState(119)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(120)
				p.Block()
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__12 {
			{
				p.SetState(127)
				p.Match(LuaParserT__12)
			}
			{
				p.SetState(128)
				p.Block()
			}

		}
		{
			p.SetState(131)
			p.Match(LuaParserT__5)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(133)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(134)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(135)
			p.Match(LuaParserT__1)
		}
		{
			p.SetState(136)
			p.exp(0)
		}
		{
			p.SetState(137)
			p.Match(LuaParserT__14)
		}
		{
			p.SetState(138)
			p.exp(0)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__14 {
			{
				p.SetState(139)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(140)
				p.exp(0)
			}

		}
		{
			p.SetState(143)
			p.Match(LuaParserT__4)
		}
		{
			p.SetState(144)
			p.Block()
		}
		{
			p.SetState(145)
			p.Match(LuaParserT__5)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(147)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(148)
			p.Namelist()
		}
		{
			p.SetState(149)
			p.Match(LuaParserT__15)
		}
		{
			p.SetState(150)
			p.Explist()
		}
		{
			p.SetState(151)
			p.Match(LuaParserT__4)
		}
		{
			p.SetState(152)
			p.Block()
		}
		{
			p.SetState(153)
			p.Match(LuaParserT__5)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(155)
			p.Match(LuaParserT__16)
		}
		{
			p.SetState(156)
			p.Funcname()
		}
		{
			p.SetState(157)
			p.Funcbody()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(159)
			p.Match(LuaParserT__17)
		}
		{
			p.SetState(160)
			p.Match(LuaParserT__16)
		}
		{
			p.SetState(161)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(162)
			p.Funcbody()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(163)
			p.Match(LuaParserT__17)
		}
		{
			p.SetState(164)
			p.Namelist()
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__1 {
			{
				p.SetState(165)
				p.Match(LuaParserT__1)
			}
			{
				p.SetState(166)
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

func (s *RetstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterRetstat(s)
	}
}

func (s *RetstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitRetstat(s)
	}
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
		p.SetState(171)
		p.Match(LuaParserT__18)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(LuaParserT__16-17))|(1<<(LuaParserT__22-17))|(1<<(LuaParserT__23-17))|(1<<(LuaParserT__24-17))|(1<<(LuaParserT__25-17))|(1<<(LuaParserT__29-17))|(1<<(LuaParserT__33-17))|(1<<(LuaParserT__45-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(LuaParserT__52-53))|(1<<(LuaParserT__55-53))|(1<<(LuaParserT__56-53))|(1<<(LuaParserNAME-53))|(1<<(LuaParserNORMALSTRING-53))|(1<<(LuaParserCHARSTRING-53))|(1<<(LuaParserLONGSTRING-53))|(1<<(LuaParserINT-53))|(1<<(LuaParserHEX-53))|(1<<(LuaParserFLOAT-53))|(1<<(LuaParserHEX_FLOAT-53)))) != 0) {
		{
			p.SetState(172)
			p.Explist()
		}

	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 {
		{
			p.SetState(175)
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

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitLabel(s)
	}
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
		p.SetState(178)
		p.Match(LuaParserT__19)
	}
	{
		p.SetState(179)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(180)
		p.Match(LuaParserT__19)
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

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFuncname(s)
	}
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
		p.SetState(182)
		p.Match(LuaParserNAME)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__20 {
		{
			p.SetState(183)
			p.Match(LuaParserT__20)
		}
		{
			p.SetState(184)
			p.Match(LuaParserNAME)
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__21 {
		{
			p.SetState(190)
			p.Match(LuaParserT__21)
		}
		{
			p.SetState(191)
			p.Match(LuaParserNAME)
		}

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

func (s *VarlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarlist(s)
	}
}

func (s *VarlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarlist(s)
	}
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
	p.EnterRule(localctx, 12, LuaParserRULE_varlist)
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
		p.SetState(194)
		p.VarId()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__14 {
		{
			p.SetState(195)
			p.Match(LuaParserT__14)
		}
		{
			p.SetState(196)
			p.VarId()
		}

		p.SetState(201)
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

func (s *TypedvarlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterTypedvarlist(s)
	}
}

func (s *TypedvarlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitTypedvarlist(s)
	}
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
	p.EnterRule(localctx, 14, LuaParserRULE_typedvarlist)
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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Typedvar()
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__14 {
			{
				p.SetState(203)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(204)
				p.Typedvar()
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.TypeLiteral()
		}
		{
			p.SetState(211)
			p.VarId()
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__14 {
			{
				p.SetState(212)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(213)
				p.VarId()
			}

			p.SetState(218)
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

func (s *NamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterNamelist(s)
	}
}

func (s *NamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitNamelist(s)
	}
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
	p.EnterRule(localctx, 16, LuaParserRULE_namelist)

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
		p.SetState(221)
		p.Match(LuaParserNAME)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(222)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(223)
				p.Match(LuaParserNAME)
			}

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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

func (s *ExplistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExplist(s)
	}
}

func (s *ExplistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExplist(s)
	}
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
	p.EnterRule(localctx, 18, LuaParserRULE_explist)
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
		p.SetState(229)
		p.exp(0)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__14 {
		{
			p.SetState(230)
			p.Match(LuaParserT__14)
		}
		{
			p.SetState(231)
			p.exp(0)
		}

		p.SetState(236)
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

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitExp(s)
	}
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, LuaParserRULE_exp, _p)

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
	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__22:
		{
			p.SetState(238)
			p.Match(LuaParserT__22)
		}

	case LuaParserT__23:
		{
			p.SetState(239)
			p.Match(LuaParserT__23)
		}

	case LuaParserT__24:
		{
			p.SetState(240)
			p.Match(LuaParserT__24)
		}

	case LuaParserINT, LuaParserHEX, LuaParserFLOAT, LuaParserHEX_FLOAT:
		{
			p.SetState(241)
			p.NumberLiteral()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		{
			p.SetState(242)
			p.StringLiteral()
		}

	case LuaParserT__25:
		{
			p.SetState(243)
			p.Match(LuaParserT__25)
		}

	case LuaParserT__16:
		{
			p.SetState(244)
			p.Functiondef()
		}

	case LuaParserT__29, LuaParserNAME:
		{
			p.SetState(245)
			p.Prefixexp()
		}

	case LuaParserT__33:
		{
			p.SetState(246)
			p.Tableconstructor()
		}

	case LuaParserT__45, LuaParserT__52, LuaParserT__55, LuaParserT__56:
		{
			p.SetState(247)
			p.OperatorUnary()
		}
		{
			p.SetState(248)
			p.exp(8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(253)
					p.OperatorPower()
				}
				{
					p.SetState(254)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(257)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(258)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(261)
					p.OperatorAddSub()
				}
				{
					p.SetState(262)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(265)
					p.OperatorStrcat()
				}
				{
					p.SetState(266)
					p.exp(5)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(269)
					p.OperatorComparison()
				}
				{
					p.SetState(270)
					p.exp(5)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(273)
					p.OperatorAnd()
				}
				{
					p.SetState(274)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(277)
					p.OperatorOr()
				}
				{
					p.SetState(278)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(281)
					p.OperatorBitwise()
				}
				{
					p.SetState(282)
					p.exp(2)
				}

			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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

func (s *TypeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterTypeLiteral(s)
	}
}

func (s *TypeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitTypeLiteral(s)
	}
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
	p.EnterRule(localctx, 22, LuaParserRULE_typeLiteral)
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
		p.SetState(289)
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

func (s *PrefixexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterPrefixexp(s)
	}
}

func (s *PrefixexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitPrefixexp(s)
	}
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
	p.EnterRule(localctx, 24, LuaParserRULE_prefixexp)

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
		p.SetState(291)
		p.VarOrExp()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(292)
				p.NameAndArgs()
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFunctioncall(s)
	}
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
	p.EnterRule(localctx, 26, LuaParserRULE_functioncall)

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
		p.SetState(298)
		p.VarOrExp()
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(299)
				p.NameAndArgs()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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

func (s *VarOrExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarOrExp(s)
	}
}

func (s *VarOrExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarOrExp(s)
	}
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
	p.EnterRule(localctx, 28, LuaParserRULE_varOrExp)

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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.VarId()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(306)
			p.exp(0)
		}
		{
			p.SetState(307)
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

func (s *VarIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarId(s)
	}
}

func (s *VarIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarId(s)
	}
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
	p.EnterRule(localctx, 30, LuaParserRULE_varId)

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
	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		{
			p.SetState(311)
			p.Match(LuaParserNAME)
		}

	case LuaParserT__29:
		{
			p.SetState(312)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(313)
			p.exp(0)
		}
		{
			p.SetState(314)
			p.Match(LuaParserT__30)
		}
		{
			p.SetState(315)
			p.VarSuffix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(319)
				p.VarSuffix()
			}

		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
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

func (s *TypedvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterTypedvar(s)
	}
}

func (s *TypedvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitTypedvar(s)
	}
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
	p.EnterRule(localctx, 32, LuaParserRULE_typedvar)

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
		p.SetState(325)
		p.TypeLiteral()
	}
	{
		p.SetState(326)
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

func (s *VarSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterVarSuffix(s)
	}
}

func (s *VarSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitVarSuffix(s)
	}
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
	p.EnterRule(localctx, 34, LuaParserRULE_varSuffix)
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
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__21 || _la == LuaParserT__29 || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(LuaParserT__33-34))|(1<<(LuaParserNORMALSTRING-34))|(1<<(LuaParserCHARSTRING-34))|(1<<(LuaParserLONGSTRING-34)))) != 0) {
		{
			p.SetState(328)
			p.NameAndArgs()
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__31:
		{
			p.SetState(334)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(335)
			p.exp(0)
		}
		{
			p.SetState(336)
			p.Match(LuaParserT__32)
		}

	case LuaParserT__20:
		{
			p.SetState(338)
			p.Match(LuaParserT__20)
		}
		{
			p.SetState(339)
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

func (s *NameAndArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterNameAndArgs(s)
	}
}

func (s *NameAndArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitNameAndArgs(s)
	}
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
	p.EnterRule(localctx, 36, LuaParserRULE_nameAndArgs)
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
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__21 {
		{
			p.SetState(342)
			p.Match(LuaParserT__21)
		}
		{
			p.SetState(343)
			p.Match(LuaParserNAME)
		}

	}
	{
		p.SetState(346)
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

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitArgs(s)
	}
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
	p.EnterRule(localctx, 38, LuaParserRULE_args)
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

	p.SetState(355)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.Match(LuaParserT__29)
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(LuaParserT__16-17))|(1<<(LuaParserT__22-17))|(1<<(LuaParserT__23-17))|(1<<(LuaParserT__24-17))|(1<<(LuaParserT__25-17))|(1<<(LuaParserT__29-17))|(1<<(LuaParserT__33-17))|(1<<(LuaParserT__45-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(LuaParserT__52-53))|(1<<(LuaParserT__55-53))|(1<<(LuaParserT__56-53))|(1<<(LuaParserNAME-53))|(1<<(LuaParserNORMALSTRING-53))|(1<<(LuaParserCHARSTRING-53))|(1<<(LuaParserLONGSTRING-53))|(1<<(LuaParserINT-53))|(1<<(LuaParserHEX-53))|(1<<(LuaParserFLOAT-53))|(1<<(LuaParserHEX_FLOAT-53)))) != 0) {
			{
				p.SetState(349)
				p.Explist()
			}

		}
		{
			p.SetState(352)
			p.Match(LuaParserT__30)
		}

	case LuaParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(353)
			p.Tableconstructor()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(354)
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

func (s *FunctiondefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFunctiondef(s)
	}
}

func (s *FunctiondefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFunctiondef(s)
	}
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
	p.EnterRule(localctx, 40, LuaParserRULE_functiondef)

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
		p.SetState(357)
		p.Match(LuaParserT__16)
	}
	{
		p.SetState(358)
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

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFuncbody(s)
	}
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
	p.EnterRule(localctx, 42, LuaParserRULE_funcbody)
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
		p.SetState(360)
		p.Match(LuaParserT__29)
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__25 || _la == LuaParserNAME {
		{
			p.SetState(361)
			p.Parlist()
		}

	}
	{
		p.SetState(364)
		p.Match(LuaParserT__30)
	}
	{
		p.SetState(365)
		p.Block()
	}
	{
		p.SetState(366)
		p.Match(LuaParserT__5)
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

func (s *ParlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterParlist(s)
	}
}

func (s *ParlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitParlist(s)
	}
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
	p.EnterRule(localctx, 44, LuaParserRULE_parlist)
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

	p.SetState(374)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.Namelist()
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__14 {
			{
				p.SetState(369)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(370)
				p.Match(LuaParserT__25)
			}

		}

	case LuaParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Match(LuaParserT__25)
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

func (s *TableconstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterTableconstructor(s)
	}
}

func (s *TableconstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitTableconstructor(s)
	}
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
	p.EnterRule(localctx, 46, LuaParserRULE_tableconstructor)
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
		p.SetState(376)
		p.Match(LuaParserT__33)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(LuaParserT__16-17))|(1<<(LuaParserT__22-17))|(1<<(LuaParserT__23-17))|(1<<(LuaParserT__24-17))|(1<<(LuaParserT__25-17))|(1<<(LuaParserT__29-17))|(1<<(LuaParserT__31-17))|(1<<(LuaParserT__33-17))|(1<<(LuaParserT__45-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(LuaParserT__52-53))|(1<<(LuaParserT__55-53))|(1<<(LuaParserT__56-53))|(1<<(LuaParserNAME-53))|(1<<(LuaParserNORMALSTRING-53))|(1<<(LuaParserCHARSTRING-53))|(1<<(LuaParserLONGSTRING-53))|(1<<(LuaParserINT-53))|(1<<(LuaParserHEX-53))|(1<<(LuaParserFLOAT-53))|(1<<(LuaParserHEX_FLOAT-53)))) != 0) {
		{
			p.SetState(377)
			p.Fieldlist()
		}

	}
	{
		p.SetState(380)
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

func (s *FieldlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFieldlist(s)
	}
}

func (s *FieldlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFieldlist(s)
	}
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
	p.EnterRule(localctx, 48, LuaParserRULE_fieldlist)
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
		p.SetState(382)
		p.Field()
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(383)
				p.Fieldsep()
			}
			{
				p.SetState(384)
				p.Field()
			}

		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 || _la == LuaParserT__14 {
		{
			p.SetState(391)
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

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitField(s)
	}
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
	p.EnterRule(localctx, 50, LuaParserRULE_field)

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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(395)
			p.exp(0)
		}
		{
			p.SetState(396)
			p.Match(LuaParserT__32)
		}
		{
			p.SetState(397)
			p.Match(LuaParserT__1)
		}
		{
			p.SetState(398)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(401)
			p.Match(LuaParserT__1)
		}
		{
			p.SetState(402)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)
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

func (s *FieldsepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterFieldsep(s)
	}
}

func (s *FieldsepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitFieldsep(s)
	}
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
	p.EnterRule(localctx, 52, LuaParserRULE_fieldsep)
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
		p.SetState(406)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__0 || _la == LuaParserT__14) {
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

func (s *OperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorOr(s)
	}
}

func (s *OperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorOr(s)
	}
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
	p.EnterRule(localctx, 54, LuaParserRULE_operatorOr)

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
		p.SetState(408)
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

func (s *OperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorAnd(s)
	}
}

func (s *OperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorAnd(s)
	}
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
	p.EnterRule(localctx, 56, LuaParserRULE_operatorAnd)

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
		p.SetState(410)
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

func (s *OperatorComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorComparison(s)
	}
}

func (s *OperatorComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorComparison(s)
	}
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
	p.EnterRule(localctx, 58, LuaParserRULE_operatorComparison)
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
		p.SetState(412)
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

func (s *OperatorStrcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorStrcat(s)
	}
}

func (s *OperatorStrcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorStrcat(s)
	}
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
	p.EnterRule(localctx, 60, LuaParserRULE_operatorStrcat)

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
		p.SetState(414)
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

func (s *OperatorAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorAddSub(s)
	}
}

func (s *OperatorAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorAddSub(s)
	}
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
	p.EnterRule(localctx, 62, LuaParserRULE_operatorAddSub)
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
		p.SetState(416)
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

func (s *OperatorMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorMulDivMod(s)
	}
}

func (s *OperatorMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorMulDivMod(s)
	}
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
	p.EnterRule(localctx, 64, LuaParserRULE_operatorMulDivMod)
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
		p.SetState(418)
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

func (s *OperatorBitwiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorBitwise(s)
	}
}

func (s *OperatorBitwiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorBitwise(s)
	}
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
	p.EnterRule(localctx, 66, LuaParserRULE_operatorBitwise)
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
		p.SetState(420)
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

func (s *OperatorUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorUnary(s)
	}
}

func (s *OperatorUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorUnary(s)
	}
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
	p.EnterRule(localctx, 68, LuaParserRULE_operatorUnary)
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
		p.SetState(422)
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

func (s *OperatorPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterOperatorPower(s)
	}
}

func (s *OperatorPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitOperatorPower(s)
	}
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
	p.EnterRule(localctx, 70, LuaParserRULE_operatorPower)

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
		p.SetState(424)
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

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
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
	p.EnterRule(localctx, 72, LuaParserRULE_numberLiteral)
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
		p.SetState(426)
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

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaListener); ok {
		listenerT.ExitStringLiteral(s)
	}
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
	p.EnterRule(localctx, 74, LuaParserRULE_stringLiteral)
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
	case 10:
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
