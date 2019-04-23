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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 442,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 85, 10, 3, 12, 3, 14, 3,
	88, 11, 3, 3, 3, 5, 3, 91, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 124, 10, 4, 12, 4, 14, 4, 127, 11, 4, 3, 4, 3, 4, 5, 4, 131, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 143,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 169, 10, 4, 5, 4, 171, 10, 4, 3, 5, 3, 5, 5, 5, 175, 10, 5,
	3, 5, 5, 5, 178, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7,
	187, 10, 7, 12, 7, 14, 7, 190, 11, 7, 3, 7, 3, 7, 5, 7, 194, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 204, 10, 8, 3, 9, 3,
	9, 3, 9, 7, 9, 209, 10, 9, 12, 9, 14, 9, 212, 11, 9, 3, 10, 3, 10, 3, 10,
	7, 10, 217, 10, 10, 12, 10, 14, 10, 220, 11, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 226, 10, 10, 12, 10, 14, 10, 229, 11, 10, 5, 10, 231, 10, 10,
	3, 11, 3, 11, 3, 11, 7, 11, 236, 10, 11, 12, 11, 14, 11, 239, 11, 11, 3,
	12, 3, 12, 3, 12, 7, 12, 244, 10, 12, 12, 12, 14, 12, 247, 11, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 262, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 296, 10, 13, 12, 13, 14, 13,
	299, 11, 13, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 305, 10, 15, 12, 15, 14,
	15, 308, 11, 15, 3, 16, 3, 16, 6, 16, 312, 10, 16, 13, 16, 14, 16, 313,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 321, 10, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 5, 18, 329, 10, 18, 3, 18, 7, 18, 332, 10, 18,
	12, 18, 14, 18, 335, 11, 18, 3, 19, 3, 19, 3, 19, 3, 20, 7, 20, 341, 10,
	20, 12, 20, 14, 20, 344, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 5, 20, 352, 10, 20, 3, 21, 3, 21, 5, 21, 356, 10, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 5, 22, 362, 10, 22, 3, 22, 3, 22, 3, 22, 5, 22, 367, 10,
	22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 5, 24, 374, 10, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 383, 10, 25, 3, 25, 5, 25, 386,
	10, 25, 3, 26, 3, 26, 5, 26, 390, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 7, 27, 398, 10, 27, 12, 27, 14, 27, 401, 11, 27, 3, 27, 5, 27,
	404, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 5, 28, 416, 10, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 2, 3, 24, 41,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
	76, 78, 2, 11, 3, 2, 29, 31, 4, 2, 3, 3, 17, 17, 3, 2, 40, 45, 3, 2, 47,
	48, 3, 2, 49, 52, 3, 2, 53, 57, 5, 2, 48, 48, 55, 55, 58, 59, 3, 2, 65,
	68, 3, 2, 62, 64, 2, 469, 2, 80, 3, 2, 2, 2, 4, 86, 3, 2, 2, 2, 6, 170,
	3, 2, 2, 2, 8, 172, 3, 2, 2, 2, 10, 179, 3, 2, 2, 2, 12, 183, 3, 2, 2,
	2, 14, 203, 3, 2, 2, 2, 16, 205, 3, 2, 2, 2, 18, 230, 3, 2, 2, 2, 20, 232,
	3, 2, 2, 2, 22, 240, 3, 2, 2, 2, 24, 261, 3, 2, 2, 2, 26, 300, 3, 2, 2,
	2, 28, 302, 3, 2, 2, 2, 30, 309, 3, 2, 2, 2, 32, 320, 3, 2, 2, 2, 34, 328,
	3, 2, 2, 2, 36, 336, 3, 2, 2, 2, 38, 342, 3, 2, 2, 2, 40, 355, 3, 2, 2,
	2, 42, 366, 3, 2, 2, 2, 44, 368, 3, 2, 2, 2, 46, 371, 3, 2, 2, 2, 48, 385,
	3, 2, 2, 2, 50, 387, 3, 2, 2, 2, 52, 393, 3, 2, 2, 2, 54, 415, 3, 2, 2,
	2, 56, 417, 3, 2, 2, 2, 58, 419, 3, 2, 2, 2, 60, 421, 3, 2, 2, 2, 62, 423,
	3, 2, 2, 2, 64, 425, 3, 2, 2, 2, 66, 427, 3, 2, 2, 2, 68, 429, 3, 2, 2,
	2, 70, 431, 3, 2, 2, 2, 72, 433, 3, 2, 2, 2, 74, 435, 3, 2, 2, 2, 76, 437,
	3, 2, 2, 2, 78, 439, 3, 2, 2, 2, 80, 81, 5, 4, 3, 2, 81, 82, 7, 2, 2, 3,
	82, 3, 3, 2, 2, 2, 83, 85, 5, 6, 4, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2,
	2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86,
	3, 2, 2, 2, 89, 91, 5, 8, 5, 2, 90, 89, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2,
	91, 5, 3, 2, 2, 2, 92, 171, 7, 3, 2, 2, 93, 171, 5, 14, 8, 2, 94, 171,
	5, 30, 16, 2, 95, 171, 5, 10, 6, 2, 96, 171, 7, 4, 2, 2, 97, 98, 7, 5,
	2, 2, 98, 171, 7, 61, 2, 2, 99, 100, 7, 6, 2, 2, 100, 101, 5, 4, 3, 2,
	101, 102, 7, 7, 2, 2, 102, 171, 3, 2, 2, 2, 103, 104, 7, 8, 2, 2, 104,
	105, 5, 24, 13, 2, 105, 106, 7, 6, 2, 2, 106, 107, 5, 4, 3, 2, 107, 108,
	7, 7, 2, 2, 108, 171, 3, 2, 2, 2, 109, 110, 7, 9, 2, 2, 110, 111, 5, 4,
	3, 2, 111, 112, 7, 10, 2, 2, 112, 113, 5, 24, 13, 2, 113, 171, 3, 2, 2,
	2, 114, 115, 7, 11, 2, 2, 115, 116, 5, 24, 13, 2, 116, 117, 7, 12, 2, 2,
	117, 125, 5, 4, 3, 2, 118, 119, 7, 13, 2, 2, 119, 120, 5, 24, 13, 2, 120,
	121, 7, 12, 2, 2, 121, 122, 5, 4, 3, 2, 122, 124, 3, 2, 2, 2, 123, 118,
	3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2,
	2, 2, 126, 130, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 129, 7, 14, 2, 2,
	129, 131, 5, 4, 3, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131,
	132, 3, 2, 2, 2, 132, 133, 7, 7, 2, 2, 133, 171, 3, 2, 2, 2, 134, 135,
	7, 15, 2, 2, 135, 136, 7, 61, 2, 2, 136, 137, 7, 16, 2, 2, 137, 138, 5,
	24, 13, 2, 138, 139, 7, 17, 2, 2, 139, 142, 5, 24, 13, 2, 140, 141, 7,
	17, 2, 2, 141, 143, 5, 24, 13, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2,
	2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 6, 2, 2, 145, 146, 5, 4, 3, 2,
	146, 147, 7, 7, 2, 2, 147, 171, 3, 2, 2, 2, 148, 149, 7, 15, 2, 2, 149,
	150, 5, 20, 11, 2, 150, 151, 7, 18, 2, 2, 151, 152, 5, 22, 12, 2, 152,
	153, 7, 6, 2, 2, 153, 154, 5, 4, 3, 2, 154, 155, 7, 7, 2, 2, 155, 171,
	3, 2, 2, 2, 156, 157, 7, 19, 2, 2, 157, 158, 5, 12, 7, 2, 158, 159, 5,
	46, 24, 2, 159, 171, 3, 2, 2, 2, 160, 161, 7, 20, 2, 2, 161, 162, 7, 19,
	2, 2, 162, 163, 7, 61, 2, 2, 163, 171, 5, 46, 24, 2, 164, 165, 7, 20, 2,
	2, 165, 168, 5, 20, 11, 2, 166, 167, 7, 16, 2, 2, 167, 169, 5, 22, 12,
	2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170,
	92, 3, 2, 2, 2, 170, 93, 3, 2, 2, 2, 170, 94, 3, 2, 2, 2, 170, 95, 3, 2,
	2, 2, 170, 96, 3, 2, 2, 2, 170, 97, 3, 2, 2, 2, 170, 99, 3, 2, 2, 2, 170,
	103, 3, 2, 2, 2, 170, 109, 3, 2, 2, 2, 170, 114, 3, 2, 2, 2, 170, 134,
	3, 2, 2, 2, 170, 148, 3, 2, 2, 2, 170, 156, 3, 2, 2, 2, 170, 160, 3, 2,
	2, 2, 170, 164, 3, 2, 2, 2, 171, 7, 3, 2, 2, 2, 172, 174, 7, 21, 2, 2,
	173, 175, 5, 22, 12, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175,
	177, 3, 2, 2, 2, 176, 178, 7, 3, 2, 2, 177, 176, 3, 2, 2, 2, 177, 178,
	3, 2, 2, 2, 178, 9, 3, 2, 2, 2, 179, 180, 7, 22, 2, 2, 180, 181, 7, 61,
	2, 2, 181, 182, 7, 22, 2, 2, 182, 11, 3, 2, 2, 2, 183, 188, 7, 61, 2, 2,
	184, 185, 7, 23, 2, 2, 185, 187, 7, 61, 2, 2, 186, 184, 3, 2, 2, 2, 187,
	190, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 193,
	3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 191, 192, 7, 24, 2, 2, 192, 194, 7, 61,
	2, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 13, 3, 2, 2, 2,
	195, 196, 5, 16, 9, 2, 196, 197, 7, 16, 2, 2, 197, 198, 5, 22, 12, 2, 198,
	204, 3, 2, 2, 2, 199, 200, 5, 18, 10, 2, 200, 201, 7, 16, 2, 2, 201, 202,
	5, 22, 12, 2, 202, 204, 3, 2, 2, 2, 203, 195, 3, 2, 2, 2, 203, 199, 3,
	2, 2, 2, 204, 15, 3, 2, 2, 2, 205, 210, 5, 34, 18, 2, 206, 207, 7, 17,
	2, 2, 207, 209, 5, 34, 18, 2, 208, 206, 3, 2, 2, 2, 209, 212, 3, 2, 2,
	2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 17, 3, 2, 2, 2, 212,
	210, 3, 2, 2, 2, 213, 218, 5, 36, 19, 2, 214, 215, 7, 17, 2, 2, 215, 217,
	5, 36, 19, 2, 216, 214, 3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3,
	2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 231, 3, 2, 2, 2, 220, 218, 3, 2, 2,
	2, 221, 222, 5, 26, 14, 2, 222, 227, 5, 34, 18, 2, 223, 224, 7, 17, 2,
	2, 224, 226, 5, 34, 18, 2, 225, 223, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2,
	227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229,
	227, 3, 2, 2, 2, 230, 213, 3, 2, 2, 2, 230, 221, 3, 2, 2, 2, 231, 19, 3,
	2, 2, 2, 232, 237, 7, 61, 2, 2, 233, 234, 7, 17, 2, 2, 234, 236, 7, 61,
	2, 2, 235, 233, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2,
	237, 238, 3, 2, 2, 2, 238, 21, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 245,
	5, 24, 13, 2, 241, 242, 7, 17, 2, 2, 242, 244, 5, 24, 13, 2, 243, 241,
	3, 2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2,
	2, 2, 246, 23, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 249, 8, 13, 1, 2,
	249, 262, 7, 25, 2, 2, 250, 262, 7, 26, 2, 2, 251, 262, 7, 27, 2, 2, 252,
	262, 5, 76, 39, 2, 253, 262, 5, 78, 40, 2, 254, 262, 7, 28, 2, 2, 255,
	262, 5, 44, 23, 2, 256, 262, 5, 28, 15, 2, 257, 262, 5, 50, 26, 2, 258,
	259, 5, 72, 37, 2, 259, 260, 5, 24, 13, 10, 260, 262, 3, 2, 2, 2, 261,
	248, 3, 2, 2, 2, 261, 250, 3, 2, 2, 2, 261, 251, 3, 2, 2, 2, 261, 252,
	3, 2, 2, 2, 261, 253, 3, 2, 2, 2, 261, 254, 3, 2, 2, 2, 261, 255, 3, 2,
	2, 2, 261, 256, 3, 2, 2, 2, 261, 257, 3, 2, 2, 2, 261, 258, 3, 2, 2, 2,
	262, 297, 3, 2, 2, 2, 263, 264, 12, 11, 2, 2, 264, 265, 5, 74, 38, 2, 265,
	266, 5, 24, 13, 11, 266, 296, 3, 2, 2, 2, 267, 268, 12, 9, 2, 2, 268, 269,
	5, 68, 35, 2, 269, 270, 5, 24, 13, 10, 270, 296, 3, 2, 2, 2, 271, 272,
	12, 8, 2, 2, 272, 273, 5, 66, 34, 2, 273, 274, 5, 24, 13, 9, 274, 296,
	3, 2, 2, 2, 275, 276, 12, 7, 2, 2, 276, 277, 5, 64, 33, 2, 277, 278, 5,
	24, 13, 7, 278, 296, 3, 2, 2, 2, 279, 280, 12, 6, 2, 2, 280, 281, 5, 62,
	32, 2, 281, 282, 5, 24, 13, 7, 282, 296, 3, 2, 2, 2, 283, 284, 12, 5, 2,
	2, 284, 285, 5, 60, 31, 2, 285, 286, 5, 24, 13, 6, 286, 296, 3, 2, 2, 2,
	287, 288, 12, 4, 2, 2, 288, 289, 5, 58, 30, 2, 289, 290, 5, 24, 13, 5,
	290, 296, 3, 2, 2, 2, 291, 292, 12, 3, 2, 2, 292, 293, 5, 70, 36, 2, 293,
	294, 5, 24, 13, 4, 294, 296, 3, 2, 2, 2, 295, 263, 3, 2, 2, 2, 295, 267,
	3, 2, 2, 2, 295, 271, 3, 2, 2, 2, 295, 275, 3, 2, 2, 2, 295, 279, 3, 2,
	2, 2, 295, 283, 3, 2, 2, 2, 295, 287, 3, 2, 2, 2, 295, 291, 3, 2, 2, 2,
	296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298,
	25, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 9, 2, 2, 2, 301, 27, 3,
	2, 2, 2, 302, 306, 5, 32, 17, 2, 303, 305, 5, 40, 21, 2, 304, 303, 3, 2,
	2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2,
	307, 29, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 311, 5, 32, 17, 2, 310,
	312, 5, 40, 21, 2, 311, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 311,
	3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 31, 3, 2, 2, 2, 315, 321, 5, 34,
	18, 2, 316, 317, 7, 32, 2, 2, 317, 318, 5, 24, 13, 2, 318, 319, 7, 33,
	2, 2, 319, 321, 3, 2, 2, 2, 320, 315, 3, 2, 2, 2, 320, 316, 3, 2, 2, 2,
	321, 33, 3, 2, 2, 2, 322, 329, 7, 61, 2, 2, 323, 324, 7, 32, 2, 2, 324,
	325, 5, 24, 13, 2, 325, 326, 7, 33, 2, 2, 326, 327, 5, 38, 20, 2, 327,
	329, 3, 2, 2, 2, 328, 322, 3, 2, 2, 2, 328, 323, 3, 2, 2, 2, 329, 333,
	3, 2, 2, 2, 330, 332, 5, 38, 20, 2, 331, 330, 3, 2, 2, 2, 332, 335, 3,
	2, 2, 2, 333, 331, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 35, 3, 2, 2,
	2, 335, 333, 3, 2, 2, 2, 336, 337, 5, 26, 14, 2, 337, 338, 5, 34, 18, 2,
	338, 37, 3, 2, 2, 2, 339, 341, 5, 40, 21, 2, 340, 339, 3, 2, 2, 2, 341,
	344, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 351,
	3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 345, 346, 7, 34, 2, 2, 346, 347, 5, 24,
	13, 2, 347, 348, 7, 35, 2, 2, 348, 352, 3, 2, 2, 2, 349, 350, 7, 23, 2,
	2, 350, 352, 7, 61, 2, 2, 351, 345, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 352,
	39, 3, 2, 2, 2, 353, 354, 7, 24, 2, 2, 354, 356, 7, 61, 2, 2, 355, 353,
	3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 5, 42,
	22, 2, 358, 41, 3, 2, 2, 2, 359, 361, 7, 32, 2, 2, 360, 362, 5, 22, 12,
	2, 361, 360, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363,
	367, 7, 33, 2, 2, 364, 367, 5, 50, 26, 2, 365, 367, 5, 78, 40, 2, 366,
	359, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 366, 365, 3, 2, 2, 2, 367, 43, 3,
	2, 2, 2, 368, 369, 7, 19, 2, 2, 369, 370, 5, 46, 24, 2, 370, 45, 3, 2,
	2, 2, 371, 373, 7, 32, 2, 2, 372, 374, 5, 48, 25, 2, 373, 372, 3, 2, 2,
	2, 373, 374, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 376, 7, 33, 2, 2, 376,
	377, 5, 4, 3, 2, 377, 378, 7, 7, 2, 2, 378, 47, 3, 2, 2, 2, 379, 382, 5,
	20, 11, 2, 380, 381, 7, 17, 2, 2, 381, 383, 7, 28, 2, 2, 382, 380, 3, 2,
	2, 2, 382, 383, 3, 2, 2, 2, 383, 386, 3, 2, 2, 2, 384, 386, 7, 28, 2, 2,
	385, 379, 3, 2, 2, 2, 385, 384, 3, 2, 2, 2, 386, 49, 3, 2, 2, 2, 387, 389,
	7, 36, 2, 2, 388, 390, 5, 52, 27, 2, 389, 388, 3, 2, 2, 2, 389, 390, 3,
	2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 392, 7, 37, 2, 2, 392, 51, 3, 2, 2,
	2, 393, 399, 5, 54, 28, 2, 394, 395, 5, 56, 29, 2, 395, 396, 5, 54, 28,
	2, 396, 398, 3, 2, 2, 2, 397, 394, 3, 2, 2, 2, 398, 401, 3, 2, 2, 2, 399,
	397, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 403, 3, 2, 2, 2, 401, 399,
	3, 2, 2, 2, 402, 404, 5, 56, 29, 2, 403, 402, 3, 2, 2, 2, 403, 404, 3,
	2, 2, 2, 404, 53, 3, 2, 2, 2, 405, 406, 7, 34, 2, 2, 406, 407, 5, 24, 13,
	2, 407, 408, 7, 35, 2, 2, 408, 409, 7, 16, 2, 2, 409, 410, 5, 24, 13, 2,
	410, 416, 3, 2, 2, 2, 411, 412, 7, 61, 2, 2, 412, 413, 7, 16, 2, 2, 413,
	416, 5, 24, 13, 2, 414, 416, 5, 24, 13, 2, 415, 405, 3, 2, 2, 2, 415, 411,
	3, 2, 2, 2, 415, 414, 3, 2, 2, 2, 416, 55, 3, 2, 2, 2, 417, 418, 9, 3,
	2, 2, 418, 57, 3, 2, 2, 2, 419, 420, 7, 38, 2, 2, 420, 59, 3, 2, 2, 2,
	421, 422, 7, 39, 2, 2, 422, 61, 3, 2, 2, 2, 423, 424, 9, 4, 2, 2, 424,
	63, 3, 2, 2, 2, 425, 426, 7, 46, 2, 2, 426, 65, 3, 2, 2, 2, 427, 428, 9,
	5, 2, 2, 428, 67, 3, 2, 2, 2, 429, 430, 9, 6, 2, 2, 430, 69, 3, 2, 2, 2,
	431, 432, 9, 7, 2, 2, 432, 71, 3, 2, 2, 2, 433, 434, 9, 8, 2, 2, 434, 73,
	3, 2, 2, 2, 435, 436, 7, 60, 2, 2, 436, 75, 3, 2, 2, 2, 437, 438, 9, 9,
	2, 2, 438, 77, 3, 2, 2, 2, 439, 440, 9, 10, 2, 2, 440, 79, 3, 2, 2, 2,
	40, 86, 90, 125, 130, 142, 168, 170, 174, 177, 188, 193, 203, 210, 218,
	227, 230, 237, 245, 261, 295, 297, 306, 313, 320, 328, 333, 342, 351, 355,
	361, 366, 373, 382, 385, 389, 399, 403, 415,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'break'", "'goto'", "'do'", "'end'", "'while'", "'repeat'",
	"'until'", "'if'", "'then'", "'elseif'", "'else'", "'for'", "'='", "','",
	"'in'", "'function'", "'local'", "'return'", "'::'", "'.'", "':'", "'nil'",
	"'false'", "'true'", "'...'", "'int'", "'float'", "'string'", "'('", "')'",
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
	"chunk", "block", "stat", "retstat", "label", "funcname", "assign", "varlist",
	"typedvarlist", "namelist", "explist", "exp", "typeLiteral", "prefixexp",
	"functioncall", "varOrExp", "varId", "typedvar", "varSuffix", "nameAndArgs",
	"args", "functiondef", "funcbody", "parlist", "tableconstructor", "fieldlist",
	"field", "fieldsep", "operatorOr", "operatorAnd", "operatorComparison",
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
	LuaParserRULE_assign             = 6
	LuaParserRULE_varlist            = 7
	LuaParserRULE_typedvarlist       = 8
	LuaParserRULE_namelist           = 9
	LuaParserRULE_explist            = 10
	LuaParserRULE_exp                = 11
	LuaParserRULE_typeLiteral        = 12
	LuaParserRULE_prefixexp          = 13
	LuaParserRULE_functioncall       = 14
	LuaParserRULE_varOrExp           = 15
	LuaParserRULE_varId              = 16
	LuaParserRULE_typedvar           = 17
	LuaParserRULE_varSuffix          = 18
	LuaParserRULE_nameAndArgs        = 19
	LuaParserRULE_args               = 20
	LuaParserRULE_functiondef        = 21
	LuaParserRULE_funcbody           = 22
	LuaParserRULE_parlist            = 23
	LuaParserRULE_tableconstructor   = 24
	LuaParserRULE_fieldlist          = 25
	LuaParserRULE_field              = 26
	LuaParserRULE_fieldsep           = 27
	LuaParserRULE_operatorOr         = 28
	LuaParserRULE_operatorAnd        = 29
	LuaParserRULE_operatorComparison = 30
	LuaParserRULE_operatorStrcat     = 31
	LuaParserRULE_operatorAddSub     = 32
	LuaParserRULE_operatorMulDivMod  = 33
	LuaParserRULE_operatorBitwise    = 34
	LuaParserRULE_operatorUnary      = 35
	LuaParserRULE_operatorPower      = 36
	LuaParserRULE_numberLiteral      = 37
	LuaParserRULE_stringLiteral      = 38
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
		p.SetState(78)
		p.Block()
	}
	{
		p.SetState(79)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserT__0)|(1<<LuaParserT__1)|(1<<LuaParserT__2)|(1<<LuaParserT__3)|(1<<LuaParserT__5)|(1<<LuaParserT__6)|(1<<LuaParserT__8)|(1<<LuaParserT__12)|(1<<LuaParserT__16)|(1<<LuaParserT__17)|(1<<LuaParserT__19)|(1<<LuaParserT__26)|(1<<LuaParserT__27)|(1<<LuaParserT__28)|(1<<LuaParserT__29))) != 0) || _la == LuaParserNAME {
		{
			p.SetState(81)
			p.Stat()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__18 {
		{
			p.SetState(87)
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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(LuaParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Functioncall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(LuaParserT__1)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(LuaParserT__2)
		}
		{
			p.SetState(96)
			p.Match(LuaParserNAME)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(97)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(98)
			p.Block()
		}
		{
			p.SetState(99)
			p.Match(LuaParserT__4)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(101)
			p.Match(LuaParserT__5)
		}
		{
			p.SetState(102)
			p.exp(0)
		}
		{
			p.SetState(103)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(104)
			p.Block()
		}
		{
			p.SetState(105)
			p.Match(LuaParserT__4)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(107)
			p.Match(LuaParserT__6)
		}
		{
			p.SetState(108)
			p.Block()
		}
		{
			p.SetState(109)
			p.Match(LuaParserT__7)
		}
		{
			p.SetState(110)
			p.exp(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(112)
			p.Match(LuaParserT__8)
		}
		{
			p.SetState(113)
			p.exp(0)
		}
		{
			p.SetState(114)
			p.Match(LuaParserT__9)
		}
		{
			p.SetState(115)
			p.Block()
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__10 {
			{
				p.SetState(116)
				p.Match(LuaParserT__10)
			}
			{
				p.SetState(117)
				p.exp(0)
			}
			{
				p.SetState(118)
				p.Match(LuaParserT__9)
			}
			{
				p.SetState(119)
				p.Block()
			}

			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__11 {
			{
				p.SetState(126)
				p.Match(LuaParserT__11)
			}
			{
				p.SetState(127)
				p.Block()
			}

		}
		{
			p.SetState(130)
			p.Match(LuaParserT__4)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(132)
			p.Match(LuaParserT__12)
		}
		{
			p.SetState(133)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(134)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(135)
			p.exp(0)
		}
		{
			p.SetState(136)
			p.Match(LuaParserT__14)
		}
		{
			p.SetState(137)
			p.exp(0)
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__14 {
			{
				p.SetState(138)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(139)
				p.exp(0)
			}

		}
		{
			p.SetState(142)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(143)
			p.Block()
		}
		{
			p.SetState(144)
			p.Match(LuaParserT__4)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(146)
			p.Match(LuaParserT__12)
		}
		{
			p.SetState(147)
			p.Namelist()
		}
		{
			p.SetState(148)
			p.Match(LuaParserT__15)
		}
		{
			p.SetState(149)
			p.Explist()
		}
		{
			p.SetState(150)
			p.Match(LuaParserT__3)
		}
		{
			p.SetState(151)
			p.Block()
		}
		{
			p.SetState(152)
			p.Match(LuaParserT__4)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(154)
			p.Match(LuaParserT__16)
		}
		{
			p.SetState(155)
			p.Funcname()
		}
		{
			p.SetState(156)
			p.Funcbody()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(158)
			p.Match(LuaParserT__17)
		}
		{
			p.SetState(159)
			p.Match(LuaParserT__16)
		}
		{
			p.SetState(160)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(161)
			p.Funcbody()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(162)
			p.Match(LuaParserT__17)
		}
		{
			p.SetState(163)
			p.Namelist()
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__13 {
			{
				p.SetState(164)
				p.Match(LuaParserT__13)
			}
			{
				p.SetState(165)
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
		p.SetState(170)
		p.Match(LuaParserT__18)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(LuaParserT__16-17))|(1<<(LuaParserT__22-17))|(1<<(LuaParserT__23-17))|(1<<(LuaParserT__24-17))|(1<<(LuaParserT__25-17))|(1<<(LuaParserT__29-17))|(1<<(LuaParserT__33-17))|(1<<(LuaParserT__45-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(LuaParserT__52-53))|(1<<(LuaParserT__55-53))|(1<<(LuaParserT__56-53))|(1<<(LuaParserNAME-53))|(1<<(LuaParserNORMALSTRING-53))|(1<<(LuaParserCHARSTRING-53))|(1<<(LuaParserLONGSTRING-53))|(1<<(LuaParserINT-53))|(1<<(LuaParserHEX-53))|(1<<(LuaParserFLOAT-53))|(1<<(LuaParserHEX_FLOAT-53)))) != 0) {
		{
			p.SetState(171)
			p.Explist()
		}

	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 {
		{
			p.SetState(174)
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
		p.SetState(177)
		p.Match(LuaParserT__19)
	}
	{
		p.SetState(178)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(179)
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
		p.SetState(181)
		p.Match(LuaParserNAME)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__20 {
		{
			p.SetState(182)
			p.Match(LuaParserT__20)
		}
		{
			p.SetState(183)
			p.Match(LuaParserNAME)
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__21 {
		{
			p.SetState(189)
			p.Match(LuaParserT__21)
		}
		{
			p.SetState(190)
			p.Match(LuaParserNAME)
		}

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
	p.EnterRule(localctx, 12, LuaParserRULE_assign)

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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29, LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Varlist()
		}
		{
			p.SetState(194)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(195)
			p.Explist()
		}

	case LuaParserT__26, LuaParserT__27, LuaParserT__28:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Typedvarlist()
		}
		{
			p.SetState(198)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(199)
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
	p.EnterRule(localctx, 14, LuaParserRULE_varlist)
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
		p.SetState(203)
		p.VarId()
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__14 {
		{
			p.SetState(204)
			p.Match(LuaParserT__14)
		}
		{
			p.SetState(205)
			p.VarId()
		}

		p.SetState(210)
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
	p.EnterRule(localctx, 16, LuaParserRULE_typedvarlist)
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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Typedvar()
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
				p.Typedvar()
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.TypeLiteral()
		}
		{
			p.SetState(220)
			p.VarId()
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LuaParserT__14 {
			{
				p.SetState(221)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(222)
				p.VarId()
			}

			p.SetState(227)
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
	p.EnterRule(localctx, 18, LuaParserRULE_namelist)

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
		p.SetState(230)
		p.Match(LuaParserNAME)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(231)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(232)
				p.Match(LuaParserNAME)
			}

		}
		p.SetState(237)
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
	p.EnterRule(localctx, 20, LuaParserRULE_explist)
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
		p.SetState(238)
		p.exp(0)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__14 {
		{
			p.SetState(239)
			p.Match(LuaParserT__14)
		}
		{
			p.SetState(240)
			p.exp(0)
		}

		p.SetState(245)
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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, LuaParserRULE_exp, _p)

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
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__22:
		{
			p.SetState(247)
			p.Match(LuaParserT__22)
		}

	case LuaParserT__23:
		{
			p.SetState(248)
			p.Match(LuaParserT__23)
		}

	case LuaParserT__24:
		{
			p.SetState(249)
			p.Match(LuaParserT__24)
		}

	case LuaParserINT, LuaParserHEX, LuaParserFLOAT, LuaParserHEX_FLOAT:
		{
			p.SetState(250)
			p.NumberLiteral()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		{
			p.SetState(251)
			p.StringLiteral()
		}

	case LuaParserT__25:
		{
			p.SetState(252)
			p.Match(LuaParserT__25)
		}

	case LuaParserT__16:
		{
			p.SetState(253)
			p.Functiondef()
		}

	case LuaParserT__29, LuaParserNAME:
		{
			p.SetState(254)
			p.Prefixexp()
		}

	case LuaParserT__33:
		{
			p.SetState(255)
			p.Tableconstructor()
		}

	case LuaParserT__45, LuaParserT__52, LuaParserT__55, LuaParserT__56:
		{
			p.SetState(256)
			p.OperatorUnary()
		}
		{
			p.SetState(257)
			p.exp(8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(262)
					p.OperatorPower()
				}
				{
					p.SetState(263)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(266)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(267)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(270)
					p.OperatorAddSub()
				}
				{
					p.SetState(271)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(273)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(274)
					p.OperatorStrcat()
				}
				{
					p.SetState(275)
					p.exp(5)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(278)
					p.OperatorComparison()
				}
				{
					p.SetState(279)
					p.exp(5)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(282)
					p.OperatorAnd()
				}
				{
					p.SetState(283)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(286)
					p.OperatorOr()
				}
				{
					p.SetState(287)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(290)
					p.OperatorBitwise()
				}
				{
					p.SetState(291)
					p.exp(2)
				}

			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 24, LuaParserRULE_typeLiteral)
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
		p.SetState(298)
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
	p.EnterRule(localctx, 26, LuaParserRULE_prefixexp)

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
		p.SetState(300)
		p.VarOrExp()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(301)
				p.NameAndArgs()
			}

		}
		p.SetState(306)
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
	p.EnterRule(localctx, 28, LuaParserRULE_functioncall)

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
		p.SetState(307)
		p.VarOrExp()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(308)
				p.NameAndArgs()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(311)
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
	p.EnterRule(localctx, 30, LuaParserRULE_varOrExp)

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

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.VarId()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(315)
			p.exp(0)
		}
		{
			p.SetState(316)
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
	p.EnterRule(localctx, 32, LuaParserRULE_varId)

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
	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		{
			p.SetState(320)
			p.Match(LuaParserNAME)
		}

	case LuaParserT__29:
		{
			p.SetState(321)
			p.Match(LuaParserT__29)
		}
		{
			p.SetState(322)
			p.exp(0)
		}
		{
			p.SetState(323)
			p.Match(LuaParserT__30)
		}
		{
			p.SetState(324)
			p.VarSuffix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(328)
				p.VarSuffix()
			}

		}
		p.SetState(333)
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
	p.EnterRule(localctx, 34, LuaParserRULE_typedvar)

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
		p.SetState(334)
		p.TypeLiteral()
	}
	{
		p.SetState(335)
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
	p.EnterRule(localctx, 36, LuaParserRULE_varSuffix)
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
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserT__21 || _la == LuaParserT__29 || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(LuaParserT__33-34))|(1<<(LuaParserNORMALSTRING-34))|(1<<(LuaParserCHARSTRING-34))|(1<<(LuaParserLONGSTRING-34)))) != 0) {
		{
			p.SetState(337)
			p.NameAndArgs()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__31:
		{
			p.SetState(343)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(344)
			p.exp(0)
		}
		{
			p.SetState(345)
			p.Match(LuaParserT__32)
		}

	case LuaParserT__20:
		{
			p.SetState(347)
			p.Match(LuaParserT__20)
		}
		{
			p.SetState(348)
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
	p.EnterRule(localctx, 38, LuaParserRULE_nameAndArgs)
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

	if _la == LuaParserT__21 {
		{
			p.SetState(351)
			p.Match(LuaParserT__21)
		}
		{
			p.SetState(352)
			p.Match(LuaParserNAME)
		}

	}
	{
		p.SetState(355)
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
	p.EnterRule(localctx, 40, LuaParserRULE_args)
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

	p.SetState(364)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.Match(LuaParserT__29)
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(LuaParserT__16-17))|(1<<(LuaParserT__22-17))|(1<<(LuaParserT__23-17))|(1<<(LuaParserT__24-17))|(1<<(LuaParserT__25-17))|(1<<(LuaParserT__29-17))|(1<<(LuaParserT__33-17))|(1<<(LuaParserT__45-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(LuaParserT__52-53))|(1<<(LuaParserT__55-53))|(1<<(LuaParserT__56-53))|(1<<(LuaParserNAME-53))|(1<<(LuaParserNORMALSTRING-53))|(1<<(LuaParserCHARSTRING-53))|(1<<(LuaParserLONGSTRING-53))|(1<<(LuaParserINT-53))|(1<<(LuaParserHEX-53))|(1<<(LuaParserFLOAT-53))|(1<<(LuaParserHEX_FLOAT-53)))) != 0) {
			{
				p.SetState(358)
				p.Explist()
			}

		}
		{
			p.SetState(361)
			p.Match(LuaParserT__30)
		}

	case LuaParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Tableconstructor()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(363)
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
	p.EnterRule(localctx, 42, LuaParserRULE_functiondef)

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
		p.SetState(366)
		p.Match(LuaParserT__16)
	}
	{
		p.SetState(367)
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
	p.EnterRule(localctx, 44, LuaParserRULE_funcbody)
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
		p.SetState(369)
		p.Match(LuaParserT__29)
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__25 || _la == LuaParserNAME {
		{
			p.SetState(370)
			p.Parlist()
		}

	}
	{
		p.SetState(373)
		p.Match(LuaParserT__30)
	}
	{
		p.SetState(374)
		p.Block()
	}
	{
		p.SetState(375)
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
	p.EnterRule(localctx, 46, LuaParserRULE_parlist)
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

	p.SetState(383)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)
			p.Namelist()
		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserT__14 {
			{
				p.SetState(378)
				p.Match(LuaParserT__14)
			}
			{
				p.SetState(379)
				p.Match(LuaParserT__25)
			}

		}

	case LuaParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
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
	p.EnterRule(localctx, 48, LuaParserRULE_tableconstructor)
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
		p.SetState(385)
		p.Match(LuaParserT__33)
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(LuaParserT__16-17))|(1<<(LuaParserT__22-17))|(1<<(LuaParserT__23-17))|(1<<(LuaParserT__24-17))|(1<<(LuaParserT__25-17))|(1<<(LuaParserT__29-17))|(1<<(LuaParserT__31-17))|(1<<(LuaParserT__33-17))|(1<<(LuaParserT__45-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(LuaParserT__52-53))|(1<<(LuaParserT__55-53))|(1<<(LuaParserT__56-53))|(1<<(LuaParserNAME-53))|(1<<(LuaParserNORMALSTRING-53))|(1<<(LuaParserCHARSTRING-53))|(1<<(LuaParserLONGSTRING-53))|(1<<(LuaParserINT-53))|(1<<(LuaParserHEX-53))|(1<<(LuaParserFLOAT-53))|(1<<(LuaParserHEX_FLOAT-53)))) != 0) {
		{
			p.SetState(386)
			p.Fieldlist()
		}

	}
	{
		p.SetState(389)
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
	p.EnterRule(localctx, 50, LuaParserRULE_fieldlist)
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
		p.SetState(391)
		p.Field()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(392)
				p.Fieldsep()
			}
			{
				p.SetState(393)
				p.Field()
			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserT__0 || _la == LuaParserT__14 {
		{
			p.SetState(400)
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
	p.EnterRule(localctx, 52, LuaParserRULE_field)

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

	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(403)
			p.Match(LuaParserT__31)
		}
		{
			p.SetState(404)
			p.exp(0)
		}
		{
			p.SetState(405)
			p.Match(LuaParserT__32)
		}
		{
			p.SetState(406)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(407)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(410)
			p.Match(LuaParserT__13)
		}
		{
			p.SetState(411)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)
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
	p.EnterRule(localctx, 54, LuaParserRULE_fieldsep)
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
		p.SetState(415)
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
	p.EnterRule(localctx, 56, LuaParserRULE_operatorOr)

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
		p.SetState(417)
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
	p.EnterRule(localctx, 58, LuaParserRULE_operatorAnd)

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
	p.EnterRule(localctx, 60, LuaParserRULE_operatorComparison)
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
		p.SetState(421)
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
	p.EnterRule(localctx, 62, LuaParserRULE_operatorStrcat)

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
		p.SetState(423)
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
	p.EnterRule(localctx, 64, LuaParserRULE_operatorAddSub)
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
		p.SetState(425)
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
	p.EnterRule(localctx, 66, LuaParserRULE_operatorMulDivMod)
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
		p.SetState(427)
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
	p.EnterRule(localctx, 68, LuaParserRULE_operatorBitwise)
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
		p.SetState(429)
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
	p.EnterRule(localctx, 70, LuaParserRULE_operatorUnary)
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
		p.SetState(431)
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
	p.EnterRule(localctx, 72, LuaParserRULE_operatorPower)

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
		p.SetState(433)
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
	p.EnterRule(localctx, 74, LuaParserRULE_numberLiteral)
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
		p.SetState(435)
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
	p.EnterRule(localctx, 76, LuaParserRULE_stringLiteral)
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
		p.SetState(437)
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
	case 11:
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
