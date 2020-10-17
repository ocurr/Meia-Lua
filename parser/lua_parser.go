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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 506,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3,
	99, 10, 3, 12, 3, 14, 3, 102, 11, 3, 3, 3, 5, 3, 105, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 130, 10, 4,
	3, 5, 3, 5, 5, 5, 134, 10, 5, 3, 5, 5, 5, 137, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 146, 10, 7, 12, 7, 14, 7, 149, 11, 7, 3,
	7, 3, 7, 5, 7, 153, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 160, 10,
	8, 12, 8, 14, 8, 163, 11, 8, 3, 8, 5, 8, 166, 10, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 186, 10, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 195, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 213, 10, 13, 5, 13, 215, 10, 13, 3, 14, 3, 14, 3, 14,
	7, 14, 220, 10, 14, 12, 14, 14, 14, 223, 11, 14, 3, 15, 3, 15, 3, 15, 7,
	15, 228, 10, 15, 12, 15, 14, 15, 231, 11, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	7, 15, 237, 10, 15, 12, 15, 14, 15, 240, 11, 15, 5, 15, 242, 10, 15, 3,
	16, 3, 16, 3, 16, 7, 16, 247, 10, 16, 12, 16, 14, 16, 250, 11, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 258, 10, 17, 12, 17, 14, 17,
	261, 11, 17, 3, 18, 3, 18, 3, 18, 7, 18, 266, 10, 18, 12, 18, 14, 18, 269,
	11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 283, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 317, 10, 19, 12, 19,
	14, 19, 320, 11, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 7, 21, 333, 10, 21, 12, 21, 14, 21, 336, 11, 21,
	5, 21, 338, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 344, 10, 21, 12,
	21, 14, 21, 347, 11, 21, 3, 21, 3, 21, 5, 21, 351, 10, 21, 3, 22, 3, 22,
	7, 22, 355, 10, 22, 12, 22, 14, 22, 358, 11, 22, 3, 23, 3, 23, 6, 23, 362,
	10, 23, 13, 23, 14, 23, 363, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24,
	371, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 379, 10,
	25, 3, 25, 7, 25, 382, 10, 25, 12, 25, 14, 25, 385, 11, 25, 3, 26, 3, 26,
	3, 26, 3, 27, 7, 27, 391, 10, 27, 12, 27, 14, 27, 394, 11, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 402, 10, 27, 3, 28, 3, 28, 5, 28,
	406, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 412, 10, 29, 3, 29, 3,
	29, 3, 29, 5, 29, 417, 10, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31,
	424, 10, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 432, 10,
	31, 12, 31, 14, 31, 435, 11, 31, 3, 31, 3, 31, 5, 31, 439, 10, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 5, 32, 447, 10, 32, 3, 32, 5, 32, 450,
	10, 32, 3, 33, 3, 33, 5, 33, 454, 10, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 7, 34, 462, 10, 34, 12, 34, 14, 34, 465, 11, 34, 3, 34, 5, 34,
	468, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 5, 35, 480, 10, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 2, 3, 36, 48,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
	76, 78, 80, 82, 84, 86, 88, 90, 92, 2, 11, 3, 2, 26, 27, 4, 2, 3, 3, 22,
	22, 3, 2, 40, 45, 3, 2, 47, 48, 3, 2, 49, 52, 3, 2, 53, 57, 5, 2, 48, 48,
	55, 55, 58, 59, 3, 2, 65, 68, 3, 2, 62, 64, 2, 534, 2, 94, 3, 2, 2, 2,
	4, 100, 3, 2, 2, 2, 6, 129, 3, 2, 2, 2, 8, 131, 3, 2, 2, 2, 10, 138, 3,
	2, 2, 2, 12, 142, 3, 2, 2, 2, 14, 154, 3, 2, 2, 2, 16, 169, 3, 2, 2, 2,
	18, 174, 3, 2, 2, 2, 20, 177, 3, 2, 2, 2, 22, 183, 3, 2, 2, 2, 24, 214,
	3, 2, 2, 2, 26, 216, 3, 2, 2, 2, 28, 241, 3, 2, 2, 2, 30, 243, 3, 2, 2,
	2, 32, 251, 3, 2, 2, 2, 34, 262, 3, 2, 2, 2, 36, 282, 3, 2, 2, 2, 38, 321,
	3, 2, 2, 2, 40, 350, 3, 2, 2, 2, 42, 352, 3, 2, 2, 2, 44, 359, 3, 2, 2,
	2, 46, 370, 3, 2, 2, 2, 48, 378, 3, 2, 2, 2, 50, 386, 3, 2, 2, 2, 52, 392,
	3, 2, 2, 2, 54, 405, 3, 2, 2, 2, 56, 416, 3, 2, 2, 2, 58, 418, 3, 2, 2,
	2, 60, 421, 3, 2, 2, 2, 62, 449, 3, 2, 2, 2, 64, 451, 3, 2, 2, 2, 66, 457,
	3, 2, 2, 2, 68, 479, 3, 2, 2, 2, 70, 481, 3, 2, 2, 2, 72, 483, 3, 2, 2,
	2, 74, 485, 3, 2, 2, 2, 76, 487, 3, 2, 2, 2, 78, 489, 3, 2, 2, 2, 80, 491,
	3, 2, 2, 2, 82, 493, 3, 2, 2, 2, 84, 495, 3, 2, 2, 2, 86, 497, 3, 2, 2,
	2, 88, 499, 3, 2, 2, 2, 90, 501, 3, 2, 2, 2, 92, 503, 3, 2, 2, 2, 94, 95,
	5, 4, 3, 2, 95, 96, 7, 2, 2, 3, 96, 3, 3, 2, 2, 2, 97, 99, 5, 6, 4, 2,
	98, 97, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101,
	3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 5, 8,
	5, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 5, 3, 2, 2, 2, 106,
	130, 7, 3, 2, 2, 107, 130, 5, 24, 13, 2, 108, 130, 5, 44, 23, 2, 109, 130,
	5, 10, 6, 2, 110, 130, 7, 4, 2, 2, 111, 112, 7, 5, 2, 2, 112, 130, 7, 61,
	2, 2, 113, 114, 7, 6, 2, 2, 114, 115, 5, 4, 3, 2, 115, 116, 7, 7, 2, 2,
	116, 130, 3, 2, 2, 2, 117, 130, 5, 20, 11, 2, 118, 130, 5, 22, 12, 2, 119,
	120, 7, 8, 2, 2, 120, 121, 5, 4, 3, 2, 121, 122, 7, 9, 2, 2, 122, 123,
	5, 36, 19, 2, 123, 130, 3, 2, 2, 2, 124, 130, 5, 14, 8, 2, 125, 126, 7,
	10, 2, 2, 126, 127, 5, 12, 7, 2, 127, 128, 5, 60, 31, 2, 128, 130, 3, 2,
	2, 2, 129, 106, 3, 2, 2, 2, 129, 107, 3, 2, 2, 2, 129, 108, 3, 2, 2, 2,
	129, 109, 3, 2, 2, 2, 129, 110, 3, 2, 2, 2, 129, 111, 3, 2, 2, 2, 129,
	113, 3, 2, 2, 2, 129, 117, 3, 2, 2, 2, 129, 118, 3, 2, 2, 2, 129, 119,
	3, 2, 2, 2, 129, 124, 3, 2, 2, 2, 129, 125, 3, 2, 2, 2, 130, 7, 3, 2, 2,
	2, 131, 133, 7, 11, 2, 2, 132, 134, 5, 34, 18, 2, 133, 132, 3, 2, 2, 2,
	133, 134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 137, 7, 3, 2, 2, 136,
	135, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 9, 3, 2, 2, 2, 138, 139, 7,
	12, 2, 2, 139, 140, 7, 61, 2, 2, 140, 141, 7, 12, 2, 2, 141, 11, 3, 2,
	2, 2, 142, 147, 7, 61, 2, 2, 143, 144, 7, 13, 2, 2, 144, 146, 7, 61, 2,
	2, 145, 143, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 152, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 151,
	7, 14, 2, 2, 151, 153, 7, 61, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3,
	2, 2, 2, 153, 13, 3, 2, 2, 2, 154, 155, 7, 15, 2, 2, 155, 156, 5, 36, 19,
	2, 156, 157, 7, 16, 2, 2, 157, 161, 5, 4, 3, 2, 158, 160, 5, 16, 9, 2,
	159, 158, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161,
	162, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 164, 166,
	5, 18, 10, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 3,
	2, 2, 2, 167, 168, 7, 7, 2, 2, 168, 15, 3, 2, 2, 2, 169, 170, 7, 17, 2,
	2, 170, 171, 5, 36, 19, 2, 171, 172, 7, 16, 2, 2, 172, 173, 5, 4, 3, 2,
	173, 17, 3, 2, 2, 2, 174, 175, 7, 18, 2, 2, 175, 176, 5, 4, 3, 2, 176,
	19, 3, 2, 2, 2, 177, 178, 7, 19, 2, 2, 178, 179, 5, 36, 19, 2, 179, 180,
	7, 6, 2, 2, 180, 181, 5, 4, 3, 2, 181, 182, 7, 7, 2, 2, 182, 21, 3, 2,
	2, 2, 183, 185, 7, 20, 2, 2, 184, 186, 5, 40, 21, 2, 185, 184, 3, 2, 2,
	2, 185, 186, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188, 7, 61, 2, 2, 188,
	189, 7, 21, 2, 2, 189, 190, 5, 36, 19, 2, 190, 191, 7, 22, 2, 2, 191, 194,
	5, 36, 19, 2, 192, 193, 7, 22, 2, 2, 193, 195, 5, 36, 19, 2, 194, 192,
	3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 7, 6,
	2, 2, 197, 198, 5, 4, 3, 2, 198, 199, 7, 7, 2, 2, 199, 23, 3, 2, 2, 2,
	200, 201, 5, 26, 14, 2, 201, 202, 7, 21, 2, 2, 202, 203, 5, 34, 18, 2,
	203, 215, 3, 2, 2, 2, 204, 205, 5, 28, 15, 2, 205, 206, 7, 21, 2, 2, 206,
	207, 5, 34, 18, 2, 207, 215, 3, 2, 2, 2, 208, 209, 7, 23, 2, 2, 209, 212,
	5, 28, 15, 2, 210, 211, 7, 21, 2, 2, 211, 213, 5, 34, 18, 2, 212, 210,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215, 3, 2, 2, 2, 214, 200, 3, 2,
	2, 2, 214, 204, 3, 2, 2, 2, 214, 208, 3, 2, 2, 2, 215, 25, 3, 2, 2, 2,
	216, 221, 5, 48, 25, 2, 217, 218, 7, 22, 2, 2, 218, 220, 5, 48, 25, 2,
	219, 217, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221,
	222, 3, 2, 2, 2, 222, 27, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 229, 5,
	50, 26, 2, 225, 226, 7, 22, 2, 2, 226, 228, 5, 50, 26, 2, 227, 225, 3,
	2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2,
	2, 230, 242, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232, 233, 5, 40, 21, 2,
	233, 238, 5, 48, 25, 2, 234, 235, 7, 22, 2, 2, 235, 237, 5, 48, 25, 2,
	236, 234, 3, 2, 2, 2, 237, 240, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238,
	239, 3, 2, 2, 2, 239, 242, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 224,
	3, 2, 2, 2, 241, 232, 3, 2, 2, 2, 242, 29, 3, 2, 2, 2, 243, 248, 7, 61,
	2, 2, 244, 245, 7, 22, 2, 2, 245, 247, 7, 61, 2, 2, 246, 244, 3, 2, 2,
	2, 247, 250, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249,
	31, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 251, 252, 5, 40, 21, 2, 252, 259,
	7, 61, 2, 2, 253, 254, 7, 22, 2, 2, 254, 255, 5, 40, 21, 2, 255, 256, 7,
	61, 2, 2, 256, 258, 3, 2, 2, 2, 257, 253, 3, 2, 2, 2, 258, 261, 3, 2, 2,
	2, 259, 257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 33, 3, 2, 2, 2, 261,
	259, 3, 2, 2, 2, 262, 267, 5, 36, 19, 2, 263, 264, 7, 22, 2, 2, 264, 266,
	5, 36, 19, 2, 265, 263, 3, 2, 2, 2, 266, 269, 3, 2, 2, 2, 267, 265, 3,
	2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 35, 3, 2, 2, 2, 269, 267, 3, 2, 2,
	2, 270, 271, 8, 19, 1, 2, 271, 283, 7, 24, 2, 2, 272, 283, 5, 38, 20, 2,
	273, 283, 5, 90, 46, 2, 274, 283, 5, 92, 47, 2, 275, 283, 7, 25, 2, 2,
	276, 283, 5, 58, 30, 2, 277, 283, 5, 42, 22, 2, 278, 283, 5, 64, 33, 2,
	279, 280, 5, 86, 44, 2, 280, 281, 5, 36, 19, 10, 281, 283, 3, 2, 2, 2,
	282, 270, 3, 2, 2, 2, 282, 272, 3, 2, 2, 2, 282, 273, 3, 2, 2, 2, 282,
	274, 3, 2, 2, 2, 282, 275, 3, 2, 2, 2, 282, 276, 3, 2, 2, 2, 282, 277,
	3, 2, 2, 2, 282, 278, 3, 2, 2, 2, 282, 279, 3, 2, 2, 2, 283, 318, 3, 2,
	2, 2, 284, 285, 12, 11, 2, 2, 285, 286, 5, 88, 45, 2, 286, 287, 5, 36,
	19, 11, 287, 317, 3, 2, 2, 2, 288, 289, 12, 9, 2, 2, 289, 290, 5, 82, 42,
	2, 290, 291, 5, 36, 19, 10, 291, 317, 3, 2, 2, 2, 292, 293, 12, 8, 2, 2,
	293, 294, 5, 80, 41, 2, 294, 295, 5, 36, 19, 9, 295, 317, 3, 2, 2, 2, 296,
	297, 12, 7, 2, 2, 297, 298, 5, 78, 40, 2, 298, 299, 5, 36, 19, 7, 299,
	317, 3, 2, 2, 2, 300, 301, 12, 6, 2, 2, 301, 302, 5, 76, 39, 2, 302, 303,
	5, 36, 19, 7, 303, 317, 3, 2, 2, 2, 304, 305, 12, 5, 2, 2, 305, 306, 5,
	74, 38, 2, 306, 307, 5, 36, 19, 6, 307, 317, 3, 2, 2, 2, 308, 309, 12,
	4, 2, 2, 309, 310, 5, 72, 37, 2, 310, 311, 5, 36, 19, 5, 311, 317, 3, 2,
	2, 2, 312, 313, 12, 3, 2, 2, 313, 314, 5, 84, 43, 2, 314, 315, 5, 36, 19,
	4, 315, 317, 3, 2, 2, 2, 316, 284, 3, 2, 2, 2, 316, 288, 3, 2, 2, 2, 316,
	292, 3, 2, 2, 2, 316, 296, 3, 2, 2, 2, 316, 300, 3, 2, 2, 2, 316, 304,
	3, 2, 2, 2, 316, 308, 3, 2, 2, 2, 316, 312, 3, 2, 2, 2, 317, 320, 3, 2,
	2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 37, 3, 2, 2, 2,
	320, 318, 3, 2, 2, 2, 321, 322, 9, 2, 2, 2, 322, 39, 3, 2, 2, 2, 323, 351,
	7, 28, 2, 2, 324, 351, 7, 29, 2, 2, 325, 351, 7, 30, 2, 2, 326, 351, 7,
	31, 2, 2, 327, 328, 7, 10, 2, 2, 328, 337, 7, 32, 2, 2, 329, 334, 5, 40,
	21, 2, 330, 331, 7, 22, 2, 2, 331, 333, 5, 40, 21, 2, 332, 330, 3, 2, 2,
	2, 333, 336, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335,
	338, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 337, 329, 3, 2, 2, 2, 337, 338,
	3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 340, 7, 14, 2, 2, 340, 345, 5, 40,
	21, 2, 341, 342, 7, 22, 2, 2, 342, 344, 5, 40, 21, 2, 343, 341, 3, 2, 2,
	2, 344, 347, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346,
	348, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348, 349, 7, 33, 2, 2, 349, 351,
	3, 2, 2, 2, 350, 323, 3, 2, 2, 2, 350, 324, 3, 2, 2, 2, 350, 325, 3, 2,
	2, 2, 350, 326, 3, 2, 2, 2, 350, 327, 3, 2, 2, 2, 351, 41, 3, 2, 2, 2,
	352, 356, 5, 46, 24, 2, 353, 355, 5, 54, 28, 2, 354, 353, 3, 2, 2, 2, 355,
	358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 43, 3,
	2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 361, 5, 46, 24, 2, 360, 362, 5, 54,
	28, 2, 361, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 361, 3, 2, 2, 2,
	363, 364, 3, 2, 2, 2, 364, 45, 3, 2, 2, 2, 365, 371, 5, 48, 25, 2, 366,
	367, 7, 32, 2, 2, 367, 368, 5, 36, 19, 2, 368, 369, 7, 33, 2, 2, 369, 371,
	3, 2, 2, 2, 370, 365, 3, 2, 2, 2, 370, 366, 3, 2, 2, 2, 371, 47, 3, 2,
	2, 2, 372, 379, 7, 61, 2, 2, 373, 374, 7, 32, 2, 2, 374, 375, 5, 36, 19,
	2, 375, 376, 7, 33, 2, 2, 376, 377, 5, 52, 27, 2, 377, 379, 3, 2, 2, 2,
	378, 372, 3, 2, 2, 2, 378, 373, 3, 2, 2, 2, 379, 383, 3, 2, 2, 2, 380,
	382, 5, 52, 27, 2, 381, 380, 3, 2, 2, 2, 382, 385, 3, 2, 2, 2, 383, 381,
	3, 2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 49, 3, 2, 2, 2, 385, 383, 3, 2,
	2, 2, 386, 387, 5, 40, 21, 2, 387, 388, 5, 48, 25, 2, 388, 51, 3, 2, 2,
	2, 389, 391, 5, 54, 28, 2, 390, 389, 3, 2, 2, 2, 391, 394, 3, 2, 2, 2,
	392, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 401, 3, 2, 2, 2, 394,
	392, 3, 2, 2, 2, 395, 396, 7, 34, 2, 2, 396, 397, 5, 36, 19, 2, 397, 398,
	7, 35, 2, 2, 398, 402, 3, 2, 2, 2, 399, 400, 7, 13, 2, 2, 400, 402, 7,
	61, 2, 2, 401, 395, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 402, 53, 3, 2, 2,
	2, 403, 404, 7, 14, 2, 2, 404, 406, 7, 61, 2, 2, 405, 403, 3, 2, 2, 2,
	405, 406, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 408, 5, 56, 29, 2, 408,
	55, 3, 2, 2, 2, 409, 411, 7, 32, 2, 2, 410, 412, 5, 34, 18, 2, 411, 410,
	3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 417, 7, 33,
	2, 2, 414, 417, 5, 64, 33, 2, 415, 417, 5, 92, 47, 2, 416, 409, 3, 2, 2,
	2, 416, 414, 3, 2, 2, 2, 416, 415, 3, 2, 2, 2, 417, 57, 3, 2, 2, 2, 418,
	419, 7, 10, 2, 2, 419, 420, 5, 60, 31, 2, 420, 59, 3, 2, 2, 2, 421, 423,
	7, 32, 2, 2, 422, 424, 5, 62, 32, 2, 423, 422, 3, 2, 2, 2, 423, 424, 3,
	2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 438, 7, 33, 2, 2, 426, 427, 7, 14,
	2, 2, 427, 428, 7, 32, 2, 2, 428, 433, 5, 40, 21, 2, 429, 430, 7, 22, 2,
	2, 430, 432, 5, 40, 21, 2, 431, 429, 3, 2, 2, 2, 432, 435, 3, 2, 2, 2,
	433, 431, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 436, 3, 2, 2, 2, 435,
	433, 3, 2, 2, 2, 436, 437, 7, 33, 2, 2, 437, 439, 3, 2, 2, 2, 438, 426,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 441, 5, 4,
	3, 2, 441, 442, 7, 7, 2, 2, 442, 61, 3, 2, 2, 2, 443, 446, 5, 32, 17, 2,
	444, 445, 7, 22, 2, 2, 445, 447, 7, 25, 2, 2, 446, 444, 3, 2, 2, 2, 446,
	447, 3, 2, 2, 2, 447, 450, 3, 2, 2, 2, 448, 450, 7, 25, 2, 2, 449, 443,
	3, 2, 2, 2, 449, 448, 3, 2, 2, 2, 450, 63, 3, 2, 2, 2, 451, 453, 7, 36,
	2, 2, 452, 454, 5, 66, 34, 2, 453, 452, 3, 2, 2, 2, 453, 454, 3, 2, 2,
	2, 454, 455, 3, 2, 2, 2, 455, 456, 7, 37, 2, 2, 456, 65, 3, 2, 2, 2, 457,
	463, 5, 68, 35, 2, 458, 459, 5, 70, 36, 2, 459, 460, 5, 68, 35, 2, 460,
	462, 3, 2, 2, 2, 461, 458, 3, 2, 2, 2, 462, 465, 3, 2, 2, 2, 463, 461,
	3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 467, 3, 2, 2, 2, 465, 463, 3, 2,
	2, 2, 466, 468, 5, 70, 36, 2, 467, 466, 3, 2, 2, 2, 467, 468, 3, 2, 2,
	2, 468, 67, 3, 2, 2, 2, 469, 470, 7, 34, 2, 2, 470, 471, 5, 36, 19, 2,
	471, 472, 7, 35, 2, 2, 472, 473, 7, 21, 2, 2, 473, 474, 5, 36, 19, 2, 474,
	480, 3, 2, 2, 2, 475, 476, 7, 61, 2, 2, 476, 477, 7, 21, 2, 2, 477, 480,
	5, 36, 19, 2, 478, 480, 5, 36, 19, 2, 479, 469, 3, 2, 2, 2, 479, 475, 3,
	2, 2, 2, 479, 478, 3, 2, 2, 2, 480, 69, 3, 2, 2, 2, 481, 482, 9, 3, 2,
	2, 482, 71, 3, 2, 2, 2, 483, 484, 7, 38, 2, 2, 484, 73, 3, 2, 2, 2, 485,
	486, 7, 39, 2, 2, 486, 75, 3, 2, 2, 2, 487, 488, 9, 4, 2, 2, 488, 77, 3,
	2, 2, 2, 489, 490, 7, 46, 2, 2, 490, 79, 3, 2, 2, 2, 491, 492, 9, 5, 2,
	2, 492, 81, 3, 2, 2, 2, 493, 494, 9, 6, 2, 2, 494, 83, 3, 2, 2, 2, 495,
	496, 9, 7, 2, 2, 496, 85, 3, 2, 2, 2, 497, 498, 9, 8, 2, 2, 498, 87, 3,
	2, 2, 2, 499, 500, 7, 60, 2, 2, 500, 89, 3, 2, 2, 2, 501, 502, 9, 9, 2,
	2, 502, 91, 3, 2, 2, 2, 503, 504, 9, 10, 2, 2, 504, 93, 3, 2, 2, 2, 48,
	100, 104, 129, 133, 136, 147, 152, 161, 165, 185, 194, 212, 214, 221, 229,
	238, 241, 248, 259, 267, 282, 316, 318, 334, 337, 345, 350, 356, 363, 370,
	378, 383, 392, 401, 405, 411, 416, 423, 433, 438, 446, 449, 453, 463, 467,
	479,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'break'", "'goto'", "'do'", "'end'", "'repeat'", "'until'",
	"'function'", "'return'", "'::'", "'.'", "':'", "'if'", "'then'", "'elseif'",
	"'else'", "'while'", "'for'", "'='", "','", "'local'", "'nil'", "'...'",
	"'true'", "'false'", "'int'", "'float'", "'string'", "'bool'", "'('", "')'",
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
	"elsestat", "whilestat", "forstat", "assign", "varlist", "typedvarlist",
	"namelist", "typednamelist", "explist", "exp", "boolLiteral", "typeLiteral",
	"prefixexp", "functioncall", "varOrExp", "varId", "typedvar", "varSuffix",
	"nameAndArgs", "args", "functiondef", "funcbody", "parlist", "tableconstructor",
	"fieldlist", "field", "fieldsep", "operatorOr", "operatorAnd", "operatorComparison",
	"operatorStrcat", "operatorAddSub", "operatorMulDivMod", "operatorBitwise",
	"operatorUnary", "operatorPower", "numberLiteral", "stringLiteral",
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
	LuaParserRULE_whilestat          = 9
	LuaParserRULE_forstat            = 10
	LuaParserRULE_assign             = 11
	LuaParserRULE_varlist            = 12
	LuaParserRULE_typedvarlist       = 13
	LuaParserRULE_namelist           = 14
	LuaParserRULE_typednamelist      = 15
	LuaParserRULE_explist            = 16
	LuaParserRULE_exp                = 17
	LuaParserRULE_boolLiteral        = 18
	LuaParserRULE_typeLiteral        = 19
	LuaParserRULE_prefixexp          = 20
	LuaParserRULE_functioncall       = 21
	LuaParserRULE_varOrExp           = 22
	LuaParserRULE_varId              = 23
	LuaParserRULE_typedvar           = 24
	LuaParserRULE_varSuffix          = 25
	LuaParserRULE_nameAndArgs        = 26
	LuaParserRULE_args               = 27
	LuaParserRULE_functiondef        = 28
	LuaParserRULE_funcbody           = 29
	LuaParserRULE_parlist            = 30
	LuaParserRULE_tableconstructor   = 31
	LuaParserRULE_fieldlist          = 32
	LuaParserRULE_field              = 33
	LuaParserRULE_fieldsep           = 34
	LuaParserRULE_operatorOr         = 35
	LuaParserRULE_operatorAnd        = 36
	LuaParserRULE_operatorComparison = 37
	LuaParserRULE_operatorStrcat     = 38
	LuaParserRULE_operatorAddSub     = 39
	LuaParserRULE_operatorMulDivMod  = 40
	LuaParserRULE_operatorBitwise    = 41
	LuaParserRULE_operatorUnary      = 42
	LuaParserRULE_operatorPower      = 43
	LuaParserRULE_numberLiteral      = 44
	LuaParserRULE_stringLiteral      = 45
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
		p.SetState(92)
		p.Block()
	}
	{
		p.SetState(93)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__0)|(1<<LuaParserT__1)|(1<<LuaParserT__2)|(1<<LuaParserT__3)|(1<<LuaParserT__5)|(1<<LuaParserT__7)|(1<<LuaParserT__9)|(1<<LuaParserT__12)|(1<<LuaParserT__16)|(1<<LuaParserT__17)|(1<<LuaParserT__20)|(1<<LuaParserT__25)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28)|(1<<LuaParserT__29))) != 0) || _la == LuaParserNAME {
		{
			p.SetState(95)
			p.Stat()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__8 {
		{
			p.SetState(101)
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

func (s *StatContext) Whilestat() IWhilestatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhilestatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhilestatContext)
}

func (s *StatContext) Forstat() IForstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForstatContext)
}

func (s *StatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(LuaParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.Functioncall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(107)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(108)
			p.Match(LuaParserT__1)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(109)
			p.Match(LuaParserT__2)
		}
		{
			p.SetState(110)
			p.Match(LuaParserNAME)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
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

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(115)
			p.Whilestat()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(116)
			p.Forstat()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(117)
			p.Match(LuaParserT__5)
		}
		{
			p.SetState(118)
			p.Block()
		}
		{
			p.SetState(119)
			p.Match(LuaParserT__6)
		}
		{
			p.SetState(120)
			p.exp(0)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(122)
			p.Ifstat()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(123)
			p.Match(LuaParserT__7)
		}
		{
			p.SetState(124)
			p.Funcname()
		}
		{
			p.SetState(125)
			p.Funcbody()
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
		p.SetState(129)
		p.Match(LuaParserT__8)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-8)&-(0x1f+1)) == 0 && ((1<<uint((_la-8)))&((1<<(LuaParserT__7-8))|(1<<(LuaParserT__21-8))|(1<<(LuaParserT__22-8))|(1<<(LuaParserT__23-8))|(1<<(LuaParserT__24-8))|(1<<(LuaParserT__29-8))|(1<<(LuaParserT__33-8)))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46))|(1<<(LuaParserNAME-46))|(1<<(LuaParserNORMALSTRING-46))|(1<<(LuaParserCHARSTRING-46))|(1<<(LuaParserLONGSTRING-46))|(1<<(LuaParserINT-46))|(1<<(LuaParserHEX-46))|(1<<(LuaParserFLOAT-46))|(1<<(LuaParserHEX_FLOAT-46)))) != 0) {
		{
			p.SetState(130)
			p.Explist()
		}

	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 {
		{
			p.SetState(133)
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
		p.SetState(136)
		p.Match(LuaParserT__9)
	}
	{
		p.SetState(137)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(138)
		p.Match(LuaParserT__9)
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
		p.SetState(140)
		p.Match(LuaParserNAME)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__10 {
		{
			p.SetState(141)
			p.Match(LuaParserT__10)
		}
		{
			p.SetState(142)
			p.Match(LuaParserNAME)
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__11 {
		{
			p.SetState(148)
			p.Match(LuaParserT__11)
		}
		{
			p.SetState(149)
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
		p.SetState(152)
		p.Match(LuaParserT__12)
	}
	{
		p.SetState(153)
		p.exp(0)
	}
	{
		p.SetState(154)
		p.Match(LuaParserT__13)
	}
	{
		p.SetState(155)
		p.Block()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__14 {
		{
			p.SetState(156)
			p.Elseifstat()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__15 {
		{
			p.SetState(162)
			p.Elsestat()
		}

	}
	{
		p.SetState(165)
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
		p.SetState(167)
		p.Match(LuaParserT__14)
	}
	{
		p.SetState(168)
		p.exp(0)
	}
	{
		p.SetState(169)
		p.Match(LuaParserT__13)
	}
	{
		p.SetState(170)
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
		p.SetState(172)
		p.Match(LuaParserT__15)
	}
	{
		p.SetState(173)
		p.Block()
	}

	return localctx
}

// IWhilestatContext is an interface to support dynamic dispatch.
type IWhilestatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhilestatContext differentiates from other interfaces.
	IsWhilestatContext()
}

type WhilestatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestatContext() *WhilestatContext {
	var p = new(WhilestatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_whilestat
	return p
}

func (*WhilestatContext) IsWhilestatContext() {}

func NewWhilestatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestatContext {
	var p = new(WhilestatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_whilestat

	return p
}

func (s *WhilestatContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *WhilestatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitWhilestat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Whilestat() (localctx IWhilestatContext) {
	localctx = NewWhilestatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LuaParserRULE_whilestat)

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
		p.SetState(175)
		p.Match(LuaParserT__16)
	}
	{
		p.SetState(176)
		p.exp(0)
	}
	{
		p.SetState(177)
		p.Match(LuaParserT__3)
	}
	{
		p.SetState(178)
		p.Block()
	}
	{
		p.SetState(179)
		p.Match(LuaParserT__4)
	}

	return localctx
}

// IForstatContext is an interface to support dynamic dispatch.
type IForstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForstatContext differentiates from other interfaces.
	IsForstatContext()
}

type ForstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForstatContext() *ForstatContext {
	var p = new(ForstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_forstat
	return p
}

func (*ForstatContext) IsForstatContext() {}

func NewForstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstatContext {
	var p = new(ForstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_forstat

	return p
}

func (s *ForstatContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstatContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *ForstatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ForstatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ForstatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForstatContext) TypeLiteral() ITypeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLiteralContext)
}

func (s *ForstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitForstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Forstat() (localctx IForstatContext) {
	localctx = NewForstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LuaParserRULE_forstat)
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
		p.SetState(181)
		p.Match(LuaParserT__17)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__7)|(1<<LuaParserT__25)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28))) != 0 {
		{
			p.SetState(182)
			p.TypeLiteral()
		}

	}
	{
		p.SetState(185)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(186)
		p.Match(LuaParserT__18)
	}
	{
		p.SetState(187)
		p.exp(0)
	}
	{
		p.SetState(188)
		p.Match(LuaParserT__19)
	}
	{
		p.SetState(189)
		p.exp(0)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__19 {
		{
			p.SetState(190)
			p.Match(LuaParserT__19)
		}
		{
			p.SetState(191)
			p.exp(0)
		}

	}
	{
		p.SetState(194)
		p.Match(LuaParserT__3)
	}
	{
		p.SetState(195)
		p.Block()
	}
	{
		p.SetState(196)
		p.Match(LuaParserT__4)
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
	p.EnterRule(localctx, 22, LuaParserRULE_assign)
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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29, LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Varlist()
		}
		{
			p.SetState(199)
			p.Match(LuaParserT__18)
		}
		{
			p.SetState(200)
			p.Explist()
		}

	case LuaParserT__7, LuaParserT__25, LuaParserT__26, LuaParserT__27, LuaParserT__28:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Typedvarlist()
		}
		{
			p.SetState(203)
			p.Match(LuaParserT__18)
		}
		{
			p.SetState(204)
			p.Explist()
		}

	case LuaParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
			p.Match(LuaParserT__20)
		}
		{
			p.SetState(207)
			p.Typedvarlist()
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__18 {
			{
				p.SetState(208)
				p.Match(LuaParserT__18)
			}
			{
				p.SetState(209)
				p.Explist()
			}

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
	p.EnterRule(localctx, 24, LuaParserRULE_varlist)
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
		p.SetState(214)
		p.VarId()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__19 {
		{
			p.SetState(215)
			p.Match(LuaParserT__19)
		}
		{
			p.SetState(216)
			p.VarId()
		}

		p.SetState(221)
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
	p.EnterRule(localctx, 26, LuaParserRULE_typedvarlist)
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

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Typedvar()
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__19 {
			{
				p.SetState(223)
				p.Match(LuaParserT__19)
			}
			{
				p.SetState(224)
				p.Typedvar()
			}

			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.TypeLiteral()
		}
		{
			p.SetState(231)
			p.VarId()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__19 {
			{
				p.SetState(232)
				p.Match(LuaParserT__19)
			}
			{
				p.SetState(233)
				p.VarId()
			}

			p.SetState(238)
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
	p.EnterRule(localctx, 28, LuaParserRULE_namelist)
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
		p.SetState(241)
		p.Match(LuaParserNAME)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__19 {
		{
			p.SetState(242)
			p.Match(LuaParserT__19)
		}
		{
			p.SetState(243)
			p.Match(LuaParserNAME)
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypednamelistContext is an interface to support dynamic dispatch.
type ITypednamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypednamelistContext differentiates from other interfaces.
	IsTypednamelistContext()
}

type TypednamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypednamelistContext() *TypednamelistContext {
	var p = new(TypednamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_typednamelist
	return p
}

func (*TypednamelistContext) IsTypednamelistContext() {}

func NewTypednamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypednamelistContext {
	var p = new(TypednamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_typednamelist

	return p
}

func (s *TypednamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *TypednamelistContext) AllTypeLiteral() []ITypeLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem())
	var tst = make([]ITypeLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeLiteralContext)
		}
	}

	return tst
}

func (s *TypednamelistContext) TypeLiteral(i int) ITypeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeLiteralContext)
}

func (s *TypednamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *TypednamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *TypednamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypednamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypednamelistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTypednamelist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Typednamelist() (localctx ITypednamelistContext) {
	localctx = NewTypednamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LuaParserRULE_typednamelist)

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
		p.SetState(249)
		p.TypeLiteral()
	}
	{
		p.SetState(250)
		p.Match(LuaParserNAME)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(251)
				p.Match(LuaParserT__19)
			}
			{
				p.SetState(252)
				p.TypeLiteral()
			}
			{
				p.SetState(253)
				p.Match(LuaParserNAME)
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 32, LuaParserRULE_explist)
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
		p.SetState(260)
		p.exp(0)
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__19 {
		{
			p.SetState(261)
			p.Match(LuaParserT__19)
		}
		{
			p.SetState(262)
			p.exp(0)
		}

		p.SetState(267)
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, LuaParserRULE_exp, _p)

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
	p.SetState(280)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__21:
		{
			p.SetState(269)
			p.Match(LuaParserT__21)
		}

	case LuaParserT__23, LuaParserT__24:
		{
			p.SetState(270)
			p.BoolLiteral()
		}

	case LuaParserINT, LuaParserHEX, LuaParserFLOAT, LuaParserHEX_FLOAT:
		{
			p.SetState(271)
			p.NumberLiteral()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		{
			p.SetState(272)
			p.StringLiteral()
		}

	case LuaParserT__22:
		{
			p.SetState(273)
			p.Match(LuaParserT__22)
		}

	case LuaParserT__7:
		{
			p.SetState(274)
			p.Functiondef()
		}

	case LuaParserT__29, LuaParserNAME:
		{
			p.SetState(275)
			p.Prefixexp()
		}

	case LuaParserT__33:
		{
			p.SetState(276)
			p.Tableconstructor()
		}

	case LuaParserT__45, LuaParserT__52, LuaParserT__55, LuaParserT__56:
		{
			p.SetState(277)
			p.OperatorUnary()
		}
		{
			p.SetState(278)
			p.exp(8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(283)
					p.OperatorPower()
				}
				{
					p.SetState(284)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(287)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(288)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(291)
					p.OperatorAddSub()
				}
				{
					p.SetState(292)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(294)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(295)
					p.OperatorStrcat()
				}
				{
					p.SetState(296)
					p.exp(5)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(299)
					p.OperatorComparison()
				}
				{
					p.SetState(300)
					p.exp(5)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(303)
					p.OperatorAnd()
				}
				{
					p.SetState(304)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(306)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(307)
					p.OperatorOr()
				}
				{
					p.SetState(308)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(311)
					p.OperatorBitwise()
				}
				{
					p.SetState(312)
					p.exp(2)
				}

			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, LuaParserRULE_boolLiteral)
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
		p.SetState(319)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__23 || _la == LuaParserT__24) {
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

func (s *TypeLiteralContext) AllTypeLiteral() []ITypeLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem())
	var tst = make([]ITypeLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeLiteralContext)
		}
	}

	return tst
}

func (s *TypeLiteralContext) TypeLiteral(i int) ITypeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeLiteralContext)
}

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
	p.EnterRule(localctx, 38, LuaParserRULE_typeLiteral)
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

	p.SetState(348)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(LuaParserT__25)
		}

	case LuaParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(LuaParserT__26)
		}

	case LuaParserT__27:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(323)
			p.Match(LuaParserT__27)
		}

	case LuaParserT__28:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(324)
			p.Match(LuaParserT__28)
		}

	case LuaParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(325)
			p.Match(LuaParserT__7)
		}
		{
			p.SetState(326)
			p.Match(LuaParserT__29)
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__7)|(1<<LuaParserT__25)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28))) != 0 {
			{
				p.SetState(327)
				p.TypeLiteral()
			}
			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == LuaParserT__19 {
				{
					p.SetState(328)
					p.Match(LuaParserT__19)
				}
				{
					p.SetState(329)
					p.TypeLiteral()
				}

				p.SetState(334)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(337)
			p.Match(LuaParserT__11)
		}

		{
			p.SetState(338)
			p.TypeLiteral()
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__19 {
			{
				p.SetState(339)
				p.Match(LuaParserT__19)
			}
			{
				p.SetState(340)
				p.TypeLiteral()
			}

			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		{
			p.SetState(346)
			p.Match(LuaParserT__30)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 40, LuaParserRULE_prefixexp)

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
		p.SetState(350)
		p.VarOrExp()
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(351)
				p.NameAndArgs()
			}

		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 42, LuaParserRULE_functioncall)

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
		p.SetState(357)
		p.VarOrExp()
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(358)
				p.NameAndArgs()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 44, LuaParserRULE_varOrExp)

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

	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.VarId()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(364)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(365)
			p.exp(0)
		}
		{
			p.SetState(366)
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
	p.EnterRule(localctx, 46, LuaParserRULE_varId)

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
	p.SetState(376)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		{
			p.SetState(370)
			p.Match(LuaParserNAME)
		}

	case LuaParserT__29:
		{
			p.SetState(371)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(372)
			p.exp(0)
		}
		{
			p.SetState(373)
			p.Match(LuaParserT__30)
		}
		{
			p.SetState(374)
			p.VarSuffix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(378)
				p.VarSuffix()
			}

		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, LuaParserRULE_typedvar)

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
		p.SetState(384)
		p.TypeLiteral()
	}
	{
		p.SetState(385)
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
	p.EnterRule(localctx, 50, LuaParserRULE_varSuffix)
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
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__11 || _la == LuaParserT__29 || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(LuaParserT__33-34))|(1<<(LuaParserNORMALSTRING-34))|(1<<(LuaParserCHARSTRING-34))|(1<<(LuaParserLONGSTRING-34)))) != 0) {
		{
			p.SetState(387)
			p.NameAndArgs()
		}

		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__31:
		{
			p.SetState(393)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(394)
			p.exp(0)
		}
		{
			p.SetState(395)
			p.Match(LuaParserT__32)
		}

	case LuaParserT__10:
		{
			p.SetState(397)
			p.Match(LuaParserT__10)
		}
		{
			p.SetState(398)
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
	p.EnterRule(localctx, 52, LuaParserRULE_nameAndArgs)
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
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__11 {
		{
			p.SetState(401)
			p.Match(LuaParserT__11)
		}
		{
			p.SetState(402)
			p.Match(LuaParserNAME)
		}

	}
	{
		p.SetState(405)
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
	p.EnterRule(localctx, 54, LuaParserRULE_args)
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

	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(407)
			p.Match(LuaParserT__29)
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-8)&-(0x1f+1)) == 0 && ((1<<uint((_la-8)))&((1<<(LuaParserT__7-8))|(1<<(LuaParserT__21-8))|(1<<(LuaParserT__22-8))|(1<<(LuaParserT__23-8))|(1<<(LuaParserT__24-8))|(1<<(LuaParserT__29-8))|(1<<(LuaParserT__33-8)))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46))|(1<<(LuaParserNAME-46))|(1<<(LuaParserNORMALSTRING-46))|(1<<(LuaParserCHARSTRING-46))|(1<<(LuaParserLONGSTRING-46))|(1<<(LuaParserINT-46))|(1<<(LuaParserHEX-46))|(1<<(LuaParserFLOAT-46))|(1<<(LuaParserHEX_FLOAT-46)))) != 0) {
			{
				p.SetState(408)
				p.Explist()
			}

		}
		{
			p.SetState(411)
			p.Match(LuaParserT__30)
		}

	case LuaParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.Tableconstructor()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(413)
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
	p.EnterRule(localctx, 56, LuaParserRULE_functiondef)

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
		p.Match(LuaParserT__7)
	}
	{
		p.SetState(417)
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

func (s *FuncbodyContext) AllTypeLiteral() []ITypeLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem())
	var tst = make([]ITypeLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeLiteralContext)
		}
	}

	return tst
}

func (s *FuncbodyContext) TypeLiteral(i int) ITypeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeLiteralContext)
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
	p.EnterRule(localctx, 58, LuaParserRULE_funcbody)
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
		p.SetState(419)
		p.Match(LuaParserT__29)
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__7)|(1<<LuaParserT__22)|(1<<LuaParserT__25)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28))) != 0 {
		{
			p.SetState(420)
			p.Parlist()
		}

	}
	{
		p.SetState(423)
		p.Match(LuaParserT__30)
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__11 {
		{
			p.SetState(424)
			p.Match(LuaParserT__11)
		}
		{
			p.SetState(425)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(426)
			p.TypeLiteral()
		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__19 {
			{
				p.SetState(427)
				p.Match(LuaParserT__19)
			}
			{
				p.SetState(428)
				p.TypeLiteral()
			}

			p.SetState(433)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(434)
			p.Match(LuaParserT__30)
		}

	}
	{
		p.SetState(438)
		p.Block()
	}
	{
		p.SetState(439)
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

func (s *ParlistContext) Typednamelist() ITypednamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypednamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypednamelistContext)
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
	p.EnterRule(localctx, 60, LuaParserRULE_parlist)
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

	p.SetState(447)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__7, LuaParserT__25, LuaParserT__26, LuaParserT__27, LuaParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)
			p.Typednamelist()
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__19 {
			{
				p.SetState(442)
				p.Match(LuaParserT__19)
			}
			{
				p.SetState(443)
				p.Match(LuaParserT__22)
			}

		}

	case LuaParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Match(LuaParserT__22)
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
	p.EnterRule(localctx, 62, LuaParserRULE_tableconstructor)
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
		p.SetState(449)
		p.Match(LuaParserT__33)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-8)&-(0x1f+1)) == 0 && ((1<<uint((_la-8)))&((1<<(LuaParserT__7-8))|(1<<(LuaParserT__21-8))|(1<<(LuaParserT__22-8))|(1<<(LuaParserT__23-8))|(1<<(LuaParserT__24-8))|(1<<(LuaParserT__29-8))|(1<<(LuaParserT__31-8))|(1<<(LuaParserT__33-8)))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(LuaParserT__45-46))|(1<<(LuaParserT__52-46))|(1<<(LuaParserT__55-46))|(1<<(LuaParserT__56-46))|(1<<(LuaParserNAME-46))|(1<<(LuaParserNORMALSTRING-46))|(1<<(LuaParserCHARSTRING-46))|(1<<(LuaParserLONGSTRING-46))|(1<<(LuaParserINT-46))|(1<<(LuaParserHEX-46))|(1<<(LuaParserFLOAT-46))|(1<<(LuaParserHEX_FLOAT-46)))) != 0) {
		{
			p.SetState(450)
			p.Fieldlist()
		}

	}
	{
		p.SetState(453)
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
	p.EnterRule(localctx, 64, LuaParserRULE_fieldlist)
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
		p.SetState(455)
		p.Field()
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(456)
				p.Fieldsep()
			}
			{
				p.SetState(457)
				p.Field()
			}

		}
		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 || _la == LuaParserT__19 {
		{
			p.SetState(464)
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
	p.EnterRule(localctx, 66, LuaParserRULE_field)

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

	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(467)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(468)
			p.exp(0)
		}
		{
			p.SetState(469)
			p.Match(LuaParserT__32)
		}
		{
			p.SetState(470)
			p.Match(LuaParserT__18)
		}
		{
			p.SetState(471)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(473)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(474)
			p.Match(LuaParserT__18)
		}
		{
			p.SetState(475)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(476)
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
	p.EnterRule(localctx, 68, LuaParserRULE_fieldsep)
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
		p.SetState(479)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserT__0 || _la == LuaParserT__19) {
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
	p.EnterRule(localctx, 70, LuaParserRULE_operatorOr)

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
		p.SetState(481)
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
	p.EnterRule(localctx, 72, LuaParserRULE_operatorAnd)

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
		p.SetState(483)
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
	p.EnterRule(localctx, 74, LuaParserRULE_operatorComparison)
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
		p.SetState(485)
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
	p.EnterRule(localctx, 76, LuaParserRULE_operatorStrcat)

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
		p.SetState(487)
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
	p.EnterRule(localctx, 78, LuaParserRULE_operatorAddSub)
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
		p.SetState(489)
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
	p.EnterRule(localctx, 80, LuaParserRULE_operatorMulDivMod)
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
		p.SetState(491)
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
	p.EnterRule(localctx, 82, LuaParserRULE_operatorBitwise)
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
		p.SetState(493)
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
	p.EnterRule(localctx, 84, LuaParserRULE_operatorUnary)
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
		p.SetState(495)
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
	p.EnterRule(localctx, 86, LuaParserRULE_operatorPower)

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
		p.SetState(497)
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
	p.EnterRule(localctx, 88, LuaParserRULE_numberLiteral)
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
		p.SetState(499)
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
	p.EnterRule(localctx, 90, LuaParserRULE_stringLiteral)
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
		p.SetState(501)
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
	case 17:
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
