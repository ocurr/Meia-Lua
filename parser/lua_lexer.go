// Code generated from Lua.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 72, 628,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50,
	3, 50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3,
	55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58,
	3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 7, 60, 372, 10, 60, 12, 60, 14, 60,
	375, 11, 60, 3, 61, 3, 61, 3, 61, 7, 61, 380, 10, 61, 12, 61, 14, 61, 383,
	11, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 7, 62, 390, 10, 62, 12, 62,
	14, 62, 393, 11, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3,
	64, 3, 64, 3, 64, 3, 64, 3, 64, 7, 64, 407, 10, 64, 12, 64, 14, 64, 410,
	11, 64, 3, 64, 5, 64, 413, 10, 64, 3, 65, 6, 65, 416, 10, 65, 13, 65, 14,
	65, 417, 3, 66, 3, 66, 3, 66, 6, 66, 423, 10, 66, 13, 66, 14, 66, 424,
	3, 67, 6, 67, 428, 10, 67, 13, 67, 14, 67, 429, 3, 67, 3, 67, 7, 67, 434,
	10, 67, 12, 67, 14, 67, 437, 11, 67, 3, 67, 5, 67, 440, 10, 67, 3, 67,
	3, 67, 6, 67, 444, 10, 67, 13, 67, 14, 67, 445, 3, 67, 5, 67, 449, 10,
	67, 3, 67, 6, 67, 452, 10, 67, 13, 67, 14, 67, 453, 3, 67, 3, 67, 5, 67,
	458, 10, 67, 3, 68, 3, 68, 3, 68, 6, 68, 463, 10, 68, 13, 68, 14, 68, 464,
	3, 68, 3, 68, 7, 68, 469, 10, 68, 12, 68, 14, 68, 472, 11, 68, 3, 68, 5,
	68, 475, 10, 68, 3, 68, 3, 68, 3, 68, 3, 68, 6, 68, 481, 10, 68, 13, 68,
	14, 68, 482, 3, 68, 5, 68, 486, 10, 68, 3, 68, 3, 68, 3, 68, 6, 68, 491,
	10, 68, 13, 68, 14, 68, 492, 3, 68, 3, 68, 5, 68, 497, 10, 68, 3, 69, 3,
	69, 5, 69, 501, 10, 69, 3, 69, 6, 69, 504, 10, 69, 13, 69, 14, 69, 505,
	3, 70, 3, 70, 5, 70, 510, 10, 70, 3, 70, 6, 70, 513, 10, 70, 13, 70, 14,
	70, 514, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 521, 10, 71, 3, 71, 3, 71,
	3, 71, 3, 71, 5, 71, 527, 10, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3,
	72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 5, 72, 540, 10, 72, 3, 73, 3, 73,
	3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 6, 74, 552, 10,
	74, 13, 74, 14, 74, 553, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77,
	3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3,
	78, 3, 78, 3, 78, 3, 78, 7, 78, 577, 10, 78, 12, 78, 14, 78, 580, 11, 78,
	3, 78, 3, 78, 7, 78, 584, 10, 78, 12, 78, 14, 78, 587, 11, 78, 3, 78, 3,
	78, 7, 78, 591, 10, 78, 12, 78, 14, 78, 594, 11, 78, 3, 78, 3, 78, 7, 78,
	598, 10, 78, 12, 78, 14, 78, 601, 11, 78, 5, 78, 603, 10, 78, 3, 78, 3,
	78, 3, 78, 5, 78, 608, 10, 78, 3, 78, 3, 78, 3, 79, 6, 79, 613, 10, 79,
	13, 79, 14, 79, 614, 3, 79, 3, 79, 3, 80, 3, 80, 3, 80, 7, 80, 622, 10,
	80, 12, 80, 14, 80, 625, 11, 80, 3, 80, 3, 80, 3, 408, 2, 81, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 2, 129,
	65, 131, 66, 133, 67, 135, 68, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2,
	147, 2, 149, 2, 151, 2, 153, 69, 155, 70, 157, 71, 159, 72, 3, 2, 19, 5,
	2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2,
	36, 36, 94, 94, 4, 2, 41, 41, 94, 94, 4, 2, 90, 90, 122, 122, 4, 2, 71,
	71, 103, 103, 4, 2, 45, 45, 47, 47, 4, 2, 82, 82, 114, 114, 12, 2, 36,
	36, 41, 41, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116, 118, 118, 120,
	120, 124, 124, 3, 2, 50, 52, 3, 2, 50, 59, 5, 2, 50, 59, 67, 72, 99, 104,
	6, 2, 12, 12, 15, 15, 63, 63, 93, 93, 4, 2, 12, 12, 15, 15, 5, 2, 12, 12,
	15, 15, 93, 93, 4, 3, 12, 12, 15, 15, 5, 2, 11, 12, 14, 15, 34, 34, 2,
	665, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3,
	2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71,
	3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2,
	79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2,
	2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2,
	2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3,
	2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2,
	109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2,
	2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123,
	3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2,
	2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3,
	2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 3, 161, 3, 2, 2, 2, 5,
	163, 3, 2, 2, 2, 7, 169, 3, 2, 2, 2, 9, 174, 3, 2, 2, 2, 11, 177, 3, 2,
	2, 2, 13, 181, 3, 2, 2, 2, 15, 188, 3, 2, 2, 2, 17, 194, 3, 2, 2, 2, 19,
	203, 3, 2, 2, 2, 21, 210, 3, 2, 2, 2, 23, 213, 3, 2, 2, 2, 25, 215, 3,
	2, 2, 2, 27, 217, 3, 2, 2, 2, 29, 220, 3, 2, 2, 2, 31, 225, 3, 2, 2, 2,
	33, 232, 3, 2, 2, 2, 35, 237, 3, 2, 2, 2, 37, 243, 3, 2, 2, 2, 39, 247,
	3, 2, 2, 2, 41, 249, 3, 2, 2, 2, 43, 251, 3, 2, 2, 2, 45, 257, 3, 2, 2,
	2, 47, 261, 3, 2, 2, 2, 49, 265, 3, 2, 2, 2, 51, 270, 3, 2, 2, 2, 53, 276,
	3, 2, 2, 2, 55, 280, 3, 2, 2, 2, 57, 286, 3, 2, 2, 2, 59, 293, 3, 2, 2,
	2, 61, 298, 3, 2, 2, 2, 63, 300, 3, 2, 2, 2, 65, 302, 3, 2, 2, 2, 67, 304,
	3, 2, 2, 2, 69, 306, 3, 2, 2, 2, 71, 308, 3, 2, 2, 2, 73, 310, 3, 2, 2,
	2, 75, 313, 3, 2, 2, 2, 77, 317, 3, 2, 2, 2, 79, 319, 3, 2, 2, 2, 81, 321,
	3, 2, 2, 2, 83, 324, 3, 2, 2, 2, 85, 327, 3, 2, 2, 2, 87, 330, 3, 2, 2,
	2, 89, 333, 3, 2, 2, 2, 91, 336, 3, 2, 2, 2, 93, 338, 3, 2, 2, 2, 95, 340,
	3, 2, 2, 2, 97, 342, 3, 2, 2, 2, 99, 344, 3, 2, 2, 2, 101, 346, 3, 2, 2,
	2, 103, 349, 3, 2, 2, 2, 105, 351, 3, 2, 2, 2, 107, 353, 3, 2, 2, 2, 109,
	355, 3, 2, 2, 2, 111, 358, 3, 2, 2, 2, 113, 361, 3, 2, 2, 2, 115, 365,
	3, 2, 2, 2, 117, 367, 3, 2, 2, 2, 119, 369, 3, 2, 2, 2, 121, 376, 3, 2,
	2, 2, 123, 386, 3, 2, 2, 2, 125, 396, 3, 2, 2, 2, 127, 412, 3, 2, 2, 2,
	129, 415, 3, 2, 2, 2, 131, 419, 3, 2, 2, 2, 133, 457, 3, 2, 2, 2, 135,
	496, 3, 2, 2, 2, 137, 498, 3, 2, 2, 2, 139, 507, 3, 2, 2, 2, 141, 526,
	3, 2, 2, 2, 143, 539, 3, 2, 2, 2, 145, 541, 3, 2, 2, 2, 147, 546, 3, 2,
	2, 2, 149, 557, 3, 2, 2, 2, 151, 559, 3, 2, 2, 2, 153, 561, 3, 2, 2, 2,
	155, 570, 3, 2, 2, 2, 157, 612, 3, 2, 2, 2, 159, 618, 3, 2, 2, 2, 161,
	162, 7, 61, 2, 2, 162, 4, 3, 2, 2, 2, 163, 164, 7, 100, 2, 2, 164, 165,
	7, 116, 2, 2, 165, 166, 7, 103, 2, 2, 166, 167, 7, 99, 2, 2, 167, 168,
	7, 109, 2, 2, 168, 6, 3, 2, 2, 2, 169, 170, 7, 105, 2, 2, 170, 171, 7,
	113, 2, 2, 171, 172, 7, 118, 2, 2, 172, 173, 7, 113, 2, 2, 173, 8, 3, 2,
	2, 2, 174, 175, 7, 102, 2, 2, 175, 176, 7, 113, 2, 2, 176, 10, 3, 2, 2,
	2, 177, 178, 7, 103, 2, 2, 178, 179, 7, 112, 2, 2, 179, 180, 7, 102, 2,
	2, 180, 12, 3, 2, 2, 2, 181, 182, 7, 116, 2, 2, 182, 183, 7, 103, 2, 2,
	183, 184, 7, 114, 2, 2, 184, 185, 7, 103, 2, 2, 185, 186, 7, 99, 2, 2,
	186, 187, 7, 118, 2, 2, 187, 14, 3, 2, 2, 2, 188, 189, 7, 119, 2, 2, 189,
	190, 7, 112, 2, 2, 190, 191, 7, 118, 2, 2, 191, 192, 7, 107, 2, 2, 192,
	193, 7, 110, 2, 2, 193, 16, 3, 2, 2, 2, 194, 195, 7, 104, 2, 2, 195, 196,
	7, 119, 2, 2, 196, 197, 7, 112, 2, 2, 197, 198, 7, 101, 2, 2, 198, 199,
	7, 118, 2, 2, 199, 200, 7, 107, 2, 2, 200, 201, 7, 113, 2, 2, 201, 202,
	7, 112, 2, 2, 202, 18, 3, 2, 2, 2, 203, 204, 7, 116, 2, 2, 204, 205, 7,
	103, 2, 2, 205, 206, 7, 118, 2, 2, 206, 207, 7, 119, 2, 2, 207, 208, 7,
	116, 2, 2, 208, 209, 7, 112, 2, 2, 209, 20, 3, 2, 2, 2, 210, 211, 7, 60,
	2, 2, 211, 212, 7, 60, 2, 2, 212, 22, 3, 2, 2, 2, 213, 214, 7, 48, 2, 2,
	214, 24, 3, 2, 2, 2, 215, 216, 7, 60, 2, 2, 216, 26, 3, 2, 2, 2, 217, 218,
	7, 107, 2, 2, 218, 219, 7, 104, 2, 2, 219, 28, 3, 2, 2, 2, 220, 221, 7,
	118, 2, 2, 221, 222, 7, 106, 2, 2, 222, 223, 7, 103, 2, 2, 223, 224, 7,
	112, 2, 2, 224, 30, 3, 2, 2, 2, 225, 226, 7, 103, 2, 2, 226, 227, 7, 110,
	2, 2, 227, 228, 7, 117, 2, 2, 228, 229, 7, 103, 2, 2, 229, 230, 7, 107,
	2, 2, 230, 231, 7, 104, 2, 2, 231, 32, 3, 2, 2, 2, 232, 233, 7, 103, 2,
	2, 233, 234, 7, 110, 2, 2, 234, 235, 7, 117, 2, 2, 235, 236, 7, 103, 2,
	2, 236, 34, 3, 2, 2, 2, 237, 238, 7, 121, 2, 2, 238, 239, 7, 106, 2, 2,
	239, 240, 7, 107, 2, 2, 240, 241, 7, 110, 2, 2, 241, 242, 7, 103, 2, 2,
	242, 36, 3, 2, 2, 2, 243, 244, 7, 104, 2, 2, 244, 245, 7, 113, 2, 2, 245,
	246, 7, 116, 2, 2, 246, 38, 3, 2, 2, 2, 247, 248, 7, 63, 2, 2, 248, 40,
	3, 2, 2, 2, 249, 250, 7, 46, 2, 2, 250, 42, 3, 2, 2, 2, 251, 252, 7, 110,
	2, 2, 252, 253, 7, 113, 2, 2, 253, 254, 7, 101, 2, 2, 254, 255, 7, 99,
	2, 2, 255, 256, 7, 110, 2, 2, 256, 44, 3, 2, 2, 2, 257, 258, 7, 112, 2,
	2, 258, 259, 7, 107, 2, 2, 259, 260, 7, 110, 2, 2, 260, 46, 3, 2, 2, 2,
	261, 262, 7, 48, 2, 2, 262, 263, 7, 48, 2, 2, 263, 264, 7, 48, 2, 2, 264,
	48, 3, 2, 2, 2, 265, 266, 7, 118, 2, 2, 266, 267, 7, 116, 2, 2, 267, 268,
	7, 119, 2, 2, 268, 269, 7, 103, 2, 2, 269, 50, 3, 2, 2, 2, 270, 271, 7,
	104, 2, 2, 271, 272, 7, 99, 2, 2, 272, 273, 7, 110, 2, 2, 273, 274, 7,
	117, 2, 2, 274, 275, 7, 103, 2, 2, 275, 52, 3, 2, 2, 2, 276, 277, 7, 107,
	2, 2, 277, 278, 7, 112, 2, 2, 278, 279, 7, 118, 2, 2, 279, 54, 3, 2, 2,
	2, 280, 281, 7, 104, 2, 2, 281, 282, 7, 110, 2, 2, 282, 283, 7, 113, 2,
	2, 283, 284, 7, 99, 2, 2, 284, 285, 7, 118, 2, 2, 285, 56, 3, 2, 2, 2,
	286, 287, 7, 117, 2, 2, 287, 288, 7, 118, 2, 2, 288, 289, 7, 116, 2, 2,
	289, 290, 7, 107, 2, 2, 290, 291, 7, 112, 2, 2, 291, 292, 7, 105, 2, 2,
	292, 58, 3, 2, 2, 2, 293, 294, 7, 100, 2, 2, 294, 295, 7, 113, 2, 2, 295,
	296, 7, 113, 2, 2, 296, 297, 7, 110, 2, 2, 297, 60, 3, 2, 2, 2, 298, 299,
	7, 42, 2, 2, 299, 62, 3, 2, 2, 2, 300, 301, 7, 43, 2, 2, 301, 64, 3, 2,
	2, 2, 302, 303, 7, 93, 2, 2, 303, 66, 3, 2, 2, 2, 304, 305, 7, 95, 2, 2,
	305, 68, 3, 2, 2, 2, 306, 307, 7, 125, 2, 2, 307, 70, 3, 2, 2, 2, 308,
	309, 7, 127, 2, 2, 309, 72, 3, 2, 2, 2, 310, 311, 7, 113, 2, 2, 311, 312,
	7, 116, 2, 2, 312, 74, 3, 2, 2, 2, 313, 314, 7, 99, 2, 2, 314, 315, 7,
	112, 2, 2, 315, 316, 7, 102, 2, 2, 316, 76, 3, 2, 2, 2, 317, 318, 7, 62,
	2, 2, 318, 78, 3, 2, 2, 2, 319, 320, 7, 64, 2, 2, 320, 80, 3, 2, 2, 2,
	321, 322, 7, 62, 2, 2, 322, 323, 7, 63, 2, 2, 323, 82, 3, 2, 2, 2, 324,
	325, 7, 64, 2, 2, 325, 326, 7, 63, 2, 2, 326, 84, 3, 2, 2, 2, 327, 328,
	7, 128, 2, 2, 328, 329, 7, 63, 2, 2, 329, 86, 3, 2, 2, 2, 330, 331, 7,
	63, 2, 2, 331, 332, 7, 63, 2, 2, 332, 88, 3, 2, 2, 2, 333, 334, 7, 48,
	2, 2, 334, 335, 7, 48, 2, 2, 335, 90, 3, 2, 2, 2, 336, 337, 7, 45, 2, 2,
	337, 92, 3, 2, 2, 2, 338, 339, 7, 47, 2, 2, 339, 94, 3, 2, 2, 2, 340, 341,
	7, 44, 2, 2, 341, 96, 3, 2, 2, 2, 342, 343, 7, 49, 2, 2, 343, 98, 3, 2,
	2, 2, 344, 345, 7, 39, 2, 2, 345, 100, 3, 2, 2, 2, 346, 347, 7, 49, 2,
	2, 347, 348, 7, 49, 2, 2, 348, 102, 3, 2, 2, 2, 349, 350, 7, 40, 2, 2,
	350, 104, 3, 2, 2, 2, 351, 352, 7, 126, 2, 2, 352, 106, 3, 2, 2, 2, 353,
	354, 7, 128, 2, 2, 354, 108, 3, 2, 2, 2, 355, 356, 7, 62, 2, 2, 356, 357,
	7, 62, 2, 2, 357, 110, 3, 2, 2, 2, 358, 359, 7, 64, 2, 2, 359, 360, 7,
	64, 2, 2, 360, 112, 3, 2, 2, 2, 361, 362, 7, 112, 2, 2, 362, 363, 7, 113,
	2, 2, 363, 364, 7, 118, 2, 2, 364, 114, 3, 2, 2, 2, 365, 366, 7, 37, 2,
	2, 366, 116, 3, 2, 2, 2, 367, 368, 7, 96, 2, 2, 368, 118, 3, 2, 2, 2, 369,
	373, 9, 2, 2, 2, 370, 372, 9, 3, 2, 2, 371, 370, 3, 2, 2, 2, 372, 375,
	3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 120, 3, 2,
	2, 2, 375, 373, 3, 2, 2, 2, 376, 381, 7, 36, 2, 2, 377, 380, 5, 141, 71,
	2, 378, 380, 10, 4, 2, 2, 379, 377, 3, 2, 2, 2, 379, 378, 3, 2, 2, 2, 380,
	383, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 384,
	3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 384, 385, 7, 36, 2, 2, 385, 122, 3, 2,
	2, 2, 386, 391, 7, 41, 2, 2, 387, 390, 5, 141, 71, 2, 388, 390, 10, 5,
	2, 2, 389, 387, 3, 2, 2, 2, 389, 388, 3, 2, 2, 2, 390, 393, 3, 2, 2, 2,
	391, 389, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 394, 3, 2, 2, 2, 393,
	391, 3, 2, 2, 2, 394, 395, 7, 41, 2, 2, 395, 124, 3, 2, 2, 2, 396, 397,
	7, 93, 2, 2, 397, 398, 5, 127, 64, 2, 398, 399, 7, 95, 2, 2, 399, 126,
	3, 2, 2, 2, 400, 401, 7, 63, 2, 2, 401, 402, 5, 127, 64, 2, 402, 403, 7,
	63, 2, 2, 403, 413, 3, 2, 2, 2, 404, 408, 7, 93, 2, 2, 405, 407, 11, 2,
	2, 2, 406, 405, 3, 2, 2, 2, 407, 410, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2,
	408, 406, 3, 2, 2, 2, 409, 411, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 411,
	413, 7, 95, 2, 2, 412, 400, 3, 2, 2, 2, 412, 404, 3, 2, 2, 2, 413, 128,
	3, 2, 2, 2, 414, 416, 5, 149, 75, 2, 415, 414, 3, 2, 2, 2, 416, 417, 3,
	2, 2, 2, 417, 415, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 130, 3, 2, 2,
	2, 419, 420, 7, 50, 2, 2, 420, 422, 9, 6, 2, 2, 421, 423, 5, 151, 76, 2,
	422, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 424,
	425, 3, 2, 2, 2, 425, 132, 3, 2, 2, 2, 426, 428, 5, 149, 75, 2, 427, 426,
	3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 427, 3, 2, 2, 2, 429, 430, 3, 2,
	2, 2, 430, 431, 3, 2, 2, 2, 431, 435, 7, 48, 2, 2, 432, 434, 5, 149, 75,
	2, 433, 432, 3, 2, 2, 2, 434, 437, 3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 435,
	436, 3, 2, 2, 2, 436, 439, 3, 2, 2, 2, 437, 435, 3, 2, 2, 2, 438, 440,
	5, 137, 69, 2, 439, 438, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 458, 3,
	2, 2, 2, 441, 443, 7, 48, 2, 2, 442, 444, 5, 149, 75, 2, 443, 442, 3, 2,
	2, 2, 444, 445, 3, 2, 2, 2, 445, 443, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2,
	446, 448, 3, 2, 2, 2, 447, 449, 5, 137, 69, 2, 448, 447, 3, 2, 2, 2, 448,
	449, 3, 2, 2, 2, 449, 458, 3, 2, 2, 2, 450, 452, 5, 149, 75, 2, 451, 450,
	3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 453, 454, 3, 2,
	2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 5, 137, 69, 2, 456, 458, 3, 2, 2,
	2, 457, 427, 3, 2, 2, 2, 457, 441, 3, 2, 2, 2, 457, 451, 3, 2, 2, 2, 458,
	134, 3, 2, 2, 2, 459, 460, 7, 50, 2, 2, 460, 462, 9, 6, 2, 2, 461, 463,
	5, 151, 76, 2, 462, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 462, 3,
	2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 470, 7, 48, 2,
	2, 467, 469, 5, 151, 76, 2, 468, 467, 3, 2, 2, 2, 469, 472, 3, 2, 2, 2,
	470, 468, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 474, 3, 2, 2, 2, 472,
	470, 3, 2, 2, 2, 473, 475, 5, 139, 70, 2, 474, 473, 3, 2, 2, 2, 474, 475,
	3, 2, 2, 2, 475, 497, 3, 2, 2, 2, 476, 477, 7, 50, 2, 2, 477, 478, 9, 6,
	2, 2, 478, 480, 7, 48, 2, 2, 479, 481, 5, 151, 76, 2, 480, 479, 3, 2, 2,
	2, 481, 482, 3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483,
	485, 3, 2, 2, 2, 484, 486, 5, 139, 70, 2, 485, 484, 3, 2, 2, 2, 485, 486,
	3, 2, 2, 2, 486, 497, 3, 2, 2, 2, 487, 488, 7, 50, 2, 2, 488, 490, 9, 6,
	2, 2, 489, 491, 5, 151, 76, 2, 490, 489, 3, 2, 2, 2, 491, 492, 3, 2, 2,
	2, 492, 490, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494,
	495, 5, 139, 70, 2, 495, 497, 3, 2, 2, 2, 496, 459, 3, 2, 2, 2, 496, 476,
	3, 2, 2, 2, 496, 487, 3, 2, 2, 2, 497, 136, 3, 2, 2, 2, 498, 500, 9, 7,
	2, 2, 499, 501, 9, 8, 2, 2, 500, 499, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2,
	501, 503, 3, 2, 2, 2, 502, 504, 5, 149, 75, 2, 503, 502, 3, 2, 2, 2, 504,
	505, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 138,
	3, 2, 2, 2, 507, 509, 9, 9, 2, 2, 508, 510, 9, 8, 2, 2, 509, 508, 3, 2,
	2, 2, 509, 510, 3, 2, 2, 2, 510, 512, 3, 2, 2, 2, 511, 513, 5, 149, 75,
	2, 512, 511, 3, 2, 2, 2, 513, 514, 3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 514,
	515, 3, 2, 2, 2, 515, 140, 3, 2, 2, 2, 516, 517, 7, 94, 2, 2, 517, 527,
	9, 10, 2, 2, 518, 520, 7, 94, 2, 2, 519, 521, 7, 15, 2, 2, 520, 519, 3,
	2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522, 527, 7, 12, 2,
	2, 523, 527, 5, 143, 72, 2, 524, 527, 5, 145, 73, 2, 525, 527, 5, 147,
	74, 2, 526, 516, 3, 2, 2, 2, 526, 518, 3, 2, 2, 2, 526, 523, 3, 2, 2, 2,
	526, 524, 3, 2, 2, 2, 526, 525, 3, 2, 2, 2, 527, 142, 3, 2, 2, 2, 528,
	529, 7, 94, 2, 2, 529, 540, 5, 149, 75, 2, 530, 531, 7, 94, 2, 2, 531,
	532, 5, 149, 75, 2, 532, 533, 5, 149, 75, 2, 533, 540, 3, 2, 2, 2, 534,
	535, 7, 94, 2, 2, 535, 536, 9, 11, 2, 2, 536, 537, 5, 149, 75, 2, 537,
	538, 5, 149, 75, 2, 538, 540, 3, 2, 2, 2, 539, 528, 3, 2, 2, 2, 539, 530,
	3, 2, 2, 2, 539, 534, 3, 2, 2, 2, 540, 144, 3, 2, 2, 2, 541, 542, 7, 94,
	2, 2, 542, 543, 7, 122, 2, 2, 543, 544, 5, 151, 76, 2, 544, 545, 5, 151,
	76, 2, 545, 146, 3, 2, 2, 2, 546, 547, 7, 94, 2, 2, 547, 548, 7, 119, 2,
	2, 548, 549, 7, 125, 2, 2, 549, 551, 3, 2, 2, 2, 550, 552, 5, 151, 76,
	2, 551, 550, 3, 2, 2, 2, 552, 553, 3, 2, 2, 2, 553, 551, 3, 2, 2, 2, 553,
	554, 3, 2, 2, 2, 554, 555, 3, 2, 2, 2, 555, 556, 7, 127, 2, 2, 556, 148,
	3, 2, 2, 2, 557, 558, 9, 12, 2, 2, 558, 150, 3, 2, 2, 2, 559, 560, 9, 13,
	2, 2, 560, 152, 3, 2, 2, 2, 561, 562, 7, 47, 2, 2, 562, 563, 7, 47, 2,
	2, 563, 564, 7, 93, 2, 2, 564, 565, 3, 2, 2, 2, 565, 566, 5, 127, 64, 2,
	566, 567, 7, 95, 2, 2, 567, 568, 3, 2, 2, 2, 568, 569, 8, 77, 2, 2, 569,
	154, 3, 2, 2, 2, 570, 571, 7, 47, 2, 2, 571, 572, 7, 47, 2, 2, 572, 602,
	3, 2, 2, 2, 573, 603, 3, 2, 2, 2, 574, 578, 7, 93, 2, 2, 575, 577, 7, 63,
	2, 2, 576, 575, 3, 2, 2, 2, 577, 580, 3, 2, 2, 2, 578, 576, 3, 2, 2, 2,
	578, 579, 3, 2, 2, 2, 579, 603, 3, 2, 2, 2, 580, 578, 3, 2, 2, 2, 581,
	585, 7, 93, 2, 2, 582, 584, 7, 63, 2, 2, 583, 582, 3, 2, 2, 2, 584, 587,
	3, 2, 2, 2, 585, 583, 3, 2, 2, 2, 585, 586, 3, 2, 2, 2, 586, 588, 3, 2,
	2, 2, 587, 585, 3, 2, 2, 2, 588, 592, 10, 14, 2, 2, 589, 591, 10, 15, 2,
	2, 590, 589, 3, 2, 2, 2, 591, 594, 3, 2, 2, 2, 592, 590, 3, 2, 2, 2, 592,
	593, 3, 2, 2, 2, 593, 603, 3, 2, 2, 2, 594, 592, 3, 2, 2, 2, 595, 599,
	10, 16, 2, 2, 596, 598, 10, 15, 2, 2, 597, 596, 3, 2, 2, 2, 598, 601, 3,
	2, 2, 2, 599, 597, 3, 2, 2, 2, 599, 600, 3, 2, 2, 2, 600, 603, 3, 2, 2,
	2, 601, 599, 3, 2, 2, 2, 602, 573, 3, 2, 2, 2, 602, 574, 3, 2, 2, 2, 602,
	581, 3, 2, 2, 2, 602, 595, 3, 2, 2, 2, 603, 607, 3, 2, 2, 2, 604, 605,
	7, 15, 2, 2, 605, 608, 7, 12, 2, 2, 606, 608, 9, 17, 2, 2, 607, 604, 3,
	2, 2, 2, 607, 606, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609, 610, 8, 78, 2,
	2, 610, 156, 3, 2, 2, 2, 611, 613, 9, 18, 2, 2, 612, 611, 3, 2, 2, 2, 613,
	614, 3, 2, 2, 2, 614, 612, 3, 2, 2, 2, 614, 615, 3, 2, 2, 2, 615, 616,
	3, 2, 2, 2, 616, 617, 8, 79, 3, 2, 617, 158, 3, 2, 2, 2, 618, 619, 7, 37,
	2, 2, 619, 623, 7, 35, 2, 2, 620, 622, 10, 15, 2, 2, 621, 620, 3, 2, 2,
	2, 622, 625, 3, 2, 2, 2, 623, 621, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 624,
	626, 3, 2, 2, 2, 625, 623, 3, 2, 2, 2, 626, 627, 8, 80, 2, 2, 627, 160,
	3, 2, 2, 2, 42, 2, 373, 379, 381, 389, 391, 408, 412, 417, 424, 429, 435,
	439, 445, 448, 453, 457, 464, 470, 474, 482, 485, 492, 496, 500, 505, 509,
	514, 520, 526, 539, 553, 578, 585, 592, 599, 602, 607, 614, 623, 4, 2,
	3, 2, 8, 2, 2,
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
	"", "';'", "'break'", "'goto'", "'do'", "'end'", "'repeat'", "'until'",
	"'function'", "'return'", "'::'", "'.'", "':'", "'if'", "'then'", "'elseif'",
	"'else'", "'while'", "'for'", "'='", "','", "'local'", "'nil'", "'...'",
	"'true'", "'false'", "'int'", "'float'", "'string'", "'bool'", "'('", "')'",
	"'['", "']'", "'{'", "'}'", "'or'", "'and'", "'<'", "'>'", "'<='", "'>='",
	"'~='", "'=='", "'..'", "'+'", "'-'", "'*'", "'/'", "'%'", "'//'", "'&'",
	"'|'", "'~'", "'<<'", "'>>'", "'not'", "'#'", "'^'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "NAME", "NORMALSTRING", "CHARSTRING", "LONGSTRING",
	"INT", "HEX", "FLOAT", "HEX_FLOAT", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
	"T__57", "NAME", "NORMALSTRING", "CHARSTRING", "LONGSTRING", "NESTED_STR",
	"INT", "HEX", "FLOAT", "HEX_FLOAT", "ExponentPart", "HexExponentPart",
	"EscapeSequence", "DecimalEscape", "HexEscape", "UtfEscape", "Digit", "HexDigit",
	"COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

type LuaLexer struct {
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

func NewLuaLexer(input antlr.CharStream) *LuaLexer {

	l := new(LuaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Lua.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LuaLexer tokens.
const (
	LuaLexerT__0         = 1
	LuaLexerT__1         = 2
	LuaLexerT__2         = 3
	LuaLexerT__3         = 4
	LuaLexerT__4         = 5
	LuaLexerT__5         = 6
	LuaLexerT__6         = 7
	LuaLexerT__7         = 8
	LuaLexerT__8         = 9
	LuaLexerT__9         = 10
	LuaLexerT__10        = 11
	LuaLexerT__11        = 12
	LuaLexerT__12        = 13
	LuaLexerT__13        = 14
	LuaLexerT__14        = 15
	LuaLexerT__15        = 16
	LuaLexerT__16        = 17
	LuaLexerT__17        = 18
	LuaLexerT__18        = 19
	LuaLexerT__19        = 20
	LuaLexerT__20        = 21
	LuaLexerT__21        = 22
	LuaLexerT__22        = 23
	LuaLexerT__23        = 24
	LuaLexerT__24        = 25
	LuaLexerT__25        = 26
	LuaLexerT__26        = 27
	LuaLexerT__27        = 28
	LuaLexerT__28        = 29
	LuaLexerT__29        = 30
	LuaLexerT__30        = 31
	LuaLexerT__31        = 32
	LuaLexerT__32        = 33
	LuaLexerT__33        = 34
	LuaLexerT__34        = 35
	LuaLexerT__35        = 36
	LuaLexerT__36        = 37
	LuaLexerT__37        = 38
	LuaLexerT__38        = 39
	LuaLexerT__39        = 40
	LuaLexerT__40        = 41
	LuaLexerT__41        = 42
	LuaLexerT__42        = 43
	LuaLexerT__43        = 44
	LuaLexerT__44        = 45
	LuaLexerT__45        = 46
	LuaLexerT__46        = 47
	LuaLexerT__47        = 48
	LuaLexerT__48        = 49
	LuaLexerT__49        = 50
	LuaLexerT__50        = 51
	LuaLexerT__51        = 52
	LuaLexerT__52        = 53
	LuaLexerT__53        = 54
	LuaLexerT__54        = 55
	LuaLexerT__55        = 56
	LuaLexerT__56        = 57
	LuaLexerT__57        = 58
	LuaLexerNAME         = 59
	LuaLexerNORMALSTRING = 60
	LuaLexerCHARSTRING   = 61
	LuaLexerLONGSTRING   = 62
	LuaLexerINT          = 63
	LuaLexerHEX          = 64
	LuaLexerFLOAT        = 65
	LuaLexerHEX_FLOAT    = 66
	LuaLexerCOMMENT      = 67
	LuaLexerLINE_COMMENT = 68
	LuaLexerWS           = 69
	LuaLexerSHEBANG      = 70
)