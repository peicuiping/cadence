// Code generated from parser/Cadence.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 88, 633,
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
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60,
	3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63,
	3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3,
	66, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68,
	3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3,
	70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72,
	3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3,
	74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75,
	3, 76, 3, 76, 7, 76, 492, 10, 76, 12, 76, 14, 76, 495, 11, 76, 3, 77, 5,
	77, 498, 10, 77, 3, 78, 3, 78, 5, 78, 502, 10, 78, 3, 79, 3, 79, 7, 79,
	506, 10, 79, 12, 79, 14, 79, 509, 11, 79, 3, 79, 5, 79, 512, 10, 79, 3,
	79, 3, 79, 3, 79, 7, 79, 517, 10, 79, 12, 79, 14, 79, 520, 11, 79, 3, 79,
	5, 79, 523, 10, 79, 3, 80, 3, 80, 7, 80, 527, 10, 80, 12, 80, 14, 80, 530,
	11, 80, 3, 81, 3, 81, 3, 81, 3, 81, 6, 81, 536, 10, 81, 13, 81, 14, 81,
	537, 3, 82, 3, 82, 3, 82, 3, 82, 6, 82, 544, 10, 82, 13, 82, 14, 82, 545,
	3, 83, 3, 83, 3, 83, 3, 83, 6, 83, 552, 10, 83, 13, 83, 14, 83, 553, 3,
	84, 3, 84, 3, 84, 7, 84, 559, 10, 84, 12, 84, 14, 84, 562, 11, 84, 3, 85,
	3, 85, 7, 85, 566, 10, 85, 12, 85, 14, 85, 569, 11, 85, 3, 85, 3, 85, 3,
	86, 3, 86, 5, 86, 575, 10, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87,
	3, 87, 6, 87, 584, 10, 87, 13, 87, 14, 87, 585, 3, 87, 3, 87, 5, 87, 590,
	10, 87, 3, 88, 3, 88, 3, 89, 6, 89, 595, 10, 89, 13, 89, 14, 89, 596, 3,
	89, 3, 89, 3, 90, 6, 90, 602, 10, 90, 13, 90, 14, 90, 603, 3, 90, 3, 90,
	3, 91, 3, 91, 3, 91, 3, 91, 3, 91, 7, 91, 613, 10, 91, 12, 91, 14, 91,
	616, 11, 91, 3, 91, 3, 91, 3, 91, 3, 91, 3, 91, 3, 92, 3, 92, 3, 92, 3,
	92, 7, 92, 627, 10, 92, 12, 92, 14, 92, 630, 11, 92, 3, 92, 3, 92, 3, 614,
	2, 93, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75,
	39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93,
	48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56,
	111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64,
	127, 65, 129, 66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71, 141, 72,
	143, 73, 145, 74, 147, 75, 149, 76, 151, 77, 153, 2, 155, 2, 157, 78, 159,
	79, 161, 80, 163, 81, 165, 82, 167, 83, 169, 84, 171, 2, 173, 2, 175, 2,
	177, 85, 179, 86, 181, 87, 183, 88, 3, 2, 16, 5, 2, 67, 92, 97, 97, 99,
	124, 3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 4, 2, 50, 51, 97, 97, 4, 2, 50,
	57, 97, 97, 6, 2, 50, 59, 67, 72, 97, 97, 99, 104, 4, 2, 67, 92, 99, 124,
	6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 6, 2, 12, 12, 15, 15, 36, 36, 94,
	94, 9, 2, 36, 36, 41, 41, 50, 50, 94, 94, 112, 112, 116, 116, 118, 118,
	5, 2, 50, 59, 67, 72, 99, 104, 6, 2, 2, 2, 11, 11, 13, 14, 34, 34, 5, 2,
	12, 12, 15, 15, 8234, 8235, 4, 2, 12, 12, 15, 15, 2, 647, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2,
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2,
	2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2,
	2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3,
	2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81,
	3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2,
	89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2,
	2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2,
	2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111,
	3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2,
	2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3,
	2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2,
	133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2,
	2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147,
	3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2,
	2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3,
	2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2,
	179, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 3, 185, 3, 2,
	2, 2, 5, 187, 3, 2, 2, 2, 7, 189, 3, 2, 2, 2, 9, 191, 3, 2, 2, 2, 11, 193,
	3, 2, 2, 2, 13, 195, 3, 2, 2, 2, 15, 197, 3, 2, 2, 2, 17, 199, 3, 2, 2,
	2, 19, 201, 3, 2, 2, 2, 21, 205, 3, 2, 2, 2, 23, 207, 3, 2, 2, 2, 25, 210,
	3, 2, 2, 2, 27, 213, 3, 2, 2, 2, 29, 215, 3, 2, 2, 2, 31, 217, 3, 2, 2,
	2, 33, 220, 3, 2, 2, 2, 35, 223, 3, 2, 2, 2, 37, 225, 3, 2, 2, 2, 39, 227,
	3, 2, 2, 2, 41, 230, 3, 2, 2, 2, 43, 233, 3, 2, 2, 2, 45, 236, 3, 2, 2,
	2, 47, 239, 3, 2, 2, 2, 49, 241, 3, 2, 2, 2, 51, 243, 3, 2, 2, 2, 53, 245,
	3, 2, 2, 2, 55, 247, 3, 2, 2, 2, 57, 249, 3, 2, 2, 2, 59, 254, 3, 2, 2,
	2, 61, 256, 3, 2, 2, 2, 63, 258, 3, 2, 2, 2, 65, 261, 3, 2, 2, 2, 67, 265,
	3, 2, 2, 2, 69, 267, 3, 2, 2, 2, 71, 271, 3, 2, 2, 2, 73, 274, 3, 2, 2,
	2, 75, 278, 3, 2, 2, 2, 77, 282, 3, 2, 2, 2, 79, 284, 3, 2, 2, 2, 81, 286,
	3, 2, 2, 2, 83, 288, 3, 2, 2, 2, 85, 300, 3, 2, 2, 2, 87, 307, 3, 2, 2,
	2, 89, 316, 3, 2, 2, 2, 91, 325, 3, 2, 2, 2, 93, 335, 3, 2, 2, 2, 95, 339,
	3, 2, 2, 2, 97, 345, 3, 2, 2, 2, 99, 350, 3, 2, 2, 2, 101, 354, 3, 2, 2,
	2, 103, 359, 3, 2, 2, 2, 105, 364, 3, 2, 2, 2, 107, 368, 3, 2, 2, 2, 109,
	372, 3, 2, 2, 2, 111, 379, 3, 2, 2, 2, 113, 383, 3, 2, 2, 2, 115, 388,
	3, 2, 2, 2, 117, 396, 3, 2, 2, 2, 119, 403, 3, 2, 2, 2, 121, 409, 3, 2,
	2, 2, 123, 418, 3, 2, 2, 2, 125, 422, 3, 2, 2, 2, 127, 426, 3, 2, 2, 2,
	129, 429, 3, 2, 2, 2, 131, 434, 3, 2, 2, 2, 133, 440, 3, 2, 2, 2, 135,
	444, 3, 2, 2, 2, 137, 447, 3, 2, 2, 2, 139, 452, 3, 2, 2, 2, 141, 458,
	3, 2, 2, 2, 143, 462, 3, 2, 2, 2, 145, 469, 3, 2, 2, 2, 147, 474, 3, 2,
	2, 2, 149, 481, 3, 2, 2, 2, 151, 489, 3, 2, 2, 2, 153, 497, 3, 2, 2, 2,
	155, 501, 3, 2, 2, 2, 157, 503, 3, 2, 2, 2, 159, 524, 3, 2, 2, 2, 161,
	531, 3, 2, 2, 2, 163, 539, 3, 2, 2, 2, 165, 547, 3, 2, 2, 2, 167, 555,
	3, 2, 2, 2, 169, 563, 3, 2, 2, 2, 171, 574, 3, 2, 2, 2, 173, 589, 3, 2,
	2, 2, 175, 591, 3, 2, 2, 2, 177, 594, 3, 2, 2, 2, 179, 601, 3, 2, 2, 2,
	181, 607, 3, 2, 2, 2, 183, 622, 3, 2, 2, 2, 185, 186, 7, 61, 2, 2, 186,
	4, 3, 2, 2, 2, 187, 188, 7, 125, 2, 2, 188, 6, 3, 2, 2, 2, 189, 190, 7,
	127, 2, 2, 190, 8, 3, 2, 2, 2, 191, 192, 7, 46, 2, 2, 192, 10, 3, 2, 2,
	2, 193, 194, 7, 60, 2, 2, 194, 12, 3, 2, 2, 2, 195, 196, 7, 48, 2, 2, 196,
	14, 3, 2, 2, 2, 197, 198, 7, 93, 2, 2, 198, 16, 3, 2, 2, 2, 199, 200, 7,
	95, 2, 2, 200, 18, 3, 2, 2, 2, 201, 202, 7, 62, 2, 2, 202, 203, 7, 47,
	2, 2, 203, 204, 7, 64, 2, 2, 204, 20, 3, 2, 2, 2, 205, 206, 7, 63, 2, 2,
	206, 22, 3, 2, 2, 2, 207, 208, 7, 126, 2, 2, 208, 209, 7, 126, 2, 2, 209,
	24, 3, 2, 2, 2, 210, 211, 7, 40, 2, 2, 211, 212, 7, 40, 2, 2, 212, 26,
	3, 2, 2, 2, 213, 214, 7, 126, 2, 2, 214, 28, 3, 2, 2, 2, 215, 216, 7, 96,
	2, 2, 216, 30, 3, 2, 2, 2, 217, 218, 7, 63, 2, 2, 218, 219, 7, 63, 2, 2,
	219, 32, 3, 2, 2, 2, 220, 221, 7, 35, 2, 2, 221, 222, 7, 63, 2, 2, 222,
	34, 3, 2, 2, 2, 223, 224, 7, 62, 2, 2, 224, 36, 3, 2, 2, 2, 225, 226, 7,
	64, 2, 2, 226, 38, 3, 2, 2, 2, 227, 228, 7, 62, 2, 2, 228, 229, 7, 63,
	2, 2, 229, 40, 3, 2, 2, 2, 230, 231, 7, 64, 2, 2, 231, 232, 7, 63, 2, 2,
	232, 42, 3, 2, 2, 2, 233, 234, 7, 62, 2, 2, 234, 235, 7, 62, 2, 2, 235,
	44, 3, 2, 2, 2, 236, 237, 7, 64, 2, 2, 237, 238, 7, 64, 2, 2, 238, 46,
	3, 2, 2, 2, 239, 240, 7, 45, 2, 2, 240, 48, 3, 2, 2, 2, 241, 242, 7, 47,
	2, 2, 242, 50, 3, 2, 2, 2, 243, 244, 7, 44, 2, 2, 244, 52, 3, 2, 2, 2,
	245, 246, 7, 49, 2, 2, 246, 54, 3, 2, 2, 2, 247, 248, 7, 39, 2, 2, 248,
	56, 3, 2, 2, 2, 249, 250, 7, 99, 2, 2, 250, 251, 7, 119, 2, 2, 251, 252,
	7, 118, 2, 2, 252, 253, 7, 106, 2, 2, 253, 58, 3, 2, 2, 2, 254, 255, 7,
	40, 2, 2, 255, 60, 3, 2, 2, 2, 256, 257, 7, 35, 2, 2, 257, 62, 3, 2, 2,
	2, 258, 259, 7, 62, 2, 2, 259, 260, 7, 47, 2, 2, 260, 64, 3, 2, 2, 2, 261,
	262, 7, 62, 2, 2, 262, 263, 7, 47, 2, 2, 263, 264, 7, 35, 2, 2, 264, 66,
	3, 2, 2, 2, 265, 266, 7, 65, 2, 2, 266, 68, 3, 2, 2, 2, 267, 268, 5, 177,
	89, 2, 268, 269, 7, 65, 2, 2, 269, 270, 7, 65, 2, 2, 270, 70, 3, 2, 2,
	2, 271, 272, 7, 99, 2, 2, 272, 273, 7, 117, 2, 2, 273, 72, 3, 2, 2, 2,
	274, 275, 7, 99, 2, 2, 275, 276, 7, 117, 2, 2, 276, 277, 7, 65, 2, 2, 277,
	74, 3, 2, 2, 2, 278, 279, 7, 99, 2, 2, 279, 280, 7, 117, 2, 2, 280, 281,
	7, 35, 2, 2, 281, 76, 3, 2, 2, 2, 282, 283, 7, 66, 2, 2, 283, 78, 3, 2,
	2, 2, 284, 285, 7, 42, 2, 2, 285, 80, 3, 2, 2, 2, 286, 287, 7, 43, 2, 2,
	287, 82, 3, 2, 2, 2, 288, 289, 7, 118, 2, 2, 289, 290, 7, 116, 2, 2, 290,
	291, 7, 99, 2, 2, 291, 292, 7, 112, 2, 2, 292, 293, 7, 117, 2, 2, 293,
	294, 7, 99, 2, 2, 294, 295, 7, 101, 2, 2, 295, 296, 7, 118, 2, 2, 296,
	297, 7, 107, 2, 2, 297, 298, 7, 113, 2, 2, 298, 299, 7, 112, 2, 2, 299,
	84, 3, 2, 2, 2, 300, 301, 7, 117, 2, 2, 301, 302, 7, 118, 2, 2, 302, 303,
	7, 116, 2, 2, 303, 304, 7, 119, 2, 2, 304, 305, 7, 101, 2, 2, 305, 306,
	7, 118, 2, 2, 306, 86, 3, 2, 2, 2, 307, 308, 7, 116, 2, 2, 308, 309, 7,
	103, 2, 2, 309, 310, 7, 117, 2, 2, 310, 311, 7, 113, 2, 2, 311, 312, 7,
	119, 2, 2, 312, 313, 7, 116, 2, 2, 313, 314, 7, 101, 2, 2, 314, 315, 7,
	103, 2, 2, 315, 88, 3, 2, 2, 2, 316, 317, 7, 101, 2, 2, 317, 318, 7, 113,
	2, 2, 318, 319, 7, 112, 2, 2, 319, 320, 7, 118, 2, 2, 320, 321, 7, 116,
	2, 2, 321, 322, 7, 99, 2, 2, 322, 323, 7, 101, 2, 2, 323, 324, 7, 118,
	2, 2, 324, 90, 3, 2, 2, 2, 325, 326, 7, 107, 2, 2, 326, 327, 7, 112, 2,
	2, 327, 328, 7, 118, 2, 2, 328, 329, 7, 103, 2, 2, 329, 330, 7, 116, 2,
	2, 330, 331, 7, 104, 2, 2, 331, 332, 7, 99, 2, 2, 332, 333, 7, 101, 2,
	2, 333, 334, 7, 103, 2, 2, 334, 92, 3, 2, 2, 2, 335, 336, 7, 104, 2, 2,
	336, 337, 7, 119, 2, 2, 337, 338, 7, 112, 2, 2, 338, 94, 3, 2, 2, 2, 339,
	340, 7, 103, 2, 2, 340, 341, 7, 120, 2, 2, 341, 342, 7, 103, 2, 2, 342,
	343, 7, 112, 2, 2, 343, 344, 7, 118, 2, 2, 344, 96, 3, 2, 2, 2, 345, 346,
	7, 103, 2, 2, 346, 347, 7, 111, 2, 2, 347, 348, 7, 107, 2, 2, 348, 349,
	7, 118, 2, 2, 349, 98, 3, 2, 2, 2, 350, 351, 7, 114, 2, 2, 351, 352, 7,
	116, 2, 2, 352, 353, 7, 103, 2, 2, 353, 100, 3, 2, 2, 2, 354, 355, 7, 114,
	2, 2, 355, 356, 7, 113, 2, 2, 356, 357, 7, 117, 2, 2, 357, 358, 7, 118,
	2, 2, 358, 102, 3, 2, 2, 2, 359, 360, 7, 114, 2, 2, 360, 361, 7, 116, 2,
	2, 361, 362, 7, 107, 2, 2, 362, 363, 7, 120, 2, 2, 363, 104, 3, 2, 2, 2,
	364, 365, 7, 114, 2, 2, 365, 366, 7, 119, 2, 2, 366, 367, 7, 100, 2, 2,
	367, 106, 3, 2, 2, 2, 368, 369, 7, 117, 2, 2, 369, 370, 7, 103, 2, 2, 370,
	371, 7, 118, 2, 2, 371, 108, 3, 2, 2, 2, 372, 373, 7, 99, 2, 2, 373, 374,
	7, 101, 2, 2, 374, 375, 7, 101, 2, 2, 375, 376, 7, 103, 2, 2, 376, 377,
	7, 117, 2, 2, 377, 378, 7, 117, 2, 2, 378, 110, 3, 2, 2, 2, 379, 380, 7,
	99, 2, 2, 380, 381, 7, 110, 2, 2, 381, 382, 7, 110, 2, 2, 382, 112, 3,
	2, 2, 2, 383, 384, 7, 117, 2, 2, 384, 385, 7, 103, 2, 2, 385, 386, 7, 110,
	2, 2, 386, 387, 7, 104, 2, 2, 387, 114, 3, 2, 2, 2, 388, 389, 7, 99, 2,
	2, 389, 390, 7, 101, 2, 2, 390, 391, 7, 101, 2, 2, 391, 392, 7, 113, 2,
	2, 392, 393, 7, 119, 2, 2, 393, 394, 7, 112, 2, 2, 394, 395, 7, 118, 2,
	2, 395, 116, 3, 2, 2, 2, 396, 397, 7, 116, 2, 2, 397, 398, 7, 103, 2, 2,
	398, 399, 7, 118, 2, 2, 399, 400, 7, 119, 2, 2, 400, 401, 7, 116, 2, 2,
	401, 402, 7, 112, 2, 2, 402, 118, 3, 2, 2, 2, 403, 404, 7, 100, 2, 2, 404,
	405, 7, 116, 2, 2, 405, 406, 7, 103, 2, 2, 406, 407, 7, 99, 2, 2, 407,
	408, 7, 109, 2, 2, 408, 120, 3, 2, 2, 2, 409, 410, 7, 101, 2, 2, 410, 411,
	7, 113, 2, 2, 411, 412, 7, 112, 2, 2, 412, 413, 7, 118, 2, 2, 413, 414,
	7, 107, 2, 2, 414, 415, 7, 112, 2, 2, 415, 416, 7, 119, 2, 2, 416, 417,
	7, 103, 2, 2, 417, 122, 3, 2, 2, 2, 418, 419, 7, 110, 2, 2, 419, 420, 7,
	103, 2, 2, 420, 421, 7, 118, 2, 2, 421, 124, 3, 2, 2, 2, 422, 423, 7, 120,
	2, 2, 423, 424, 7, 99, 2, 2, 424, 425, 7, 116, 2, 2, 425, 126, 3, 2, 2,
	2, 426, 427, 7, 107, 2, 2, 427, 428, 7, 104, 2, 2, 428, 128, 3, 2, 2, 2,
	429, 430, 7, 103, 2, 2, 430, 431, 7, 110, 2, 2, 431, 432, 7, 117, 2, 2,
	432, 433, 7, 103, 2, 2, 433, 130, 3, 2, 2, 2, 434, 435, 7, 121, 2, 2, 435,
	436, 7, 106, 2, 2, 436, 437, 7, 107, 2, 2, 437, 438, 7, 110, 2, 2, 438,
	439, 7, 103, 2, 2, 439, 132, 3, 2, 2, 2, 440, 441, 7, 104, 2, 2, 441, 442,
	7, 113, 2, 2, 442, 443, 7, 116, 2, 2, 443, 134, 3, 2, 2, 2, 444, 445, 7,
	107, 2, 2, 445, 446, 7, 112, 2, 2, 446, 136, 3, 2, 2, 2, 447, 448, 7, 118,
	2, 2, 448, 449, 7, 116, 2, 2, 449, 450, 7, 119, 2, 2, 450, 451, 7, 103,
	2, 2, 451, 138, 3, 2, 2, 2, 452, 453, 7, 104, 2, 2, 453, 454, 7, 99, 2,
	2, 454, 455, 7, 110, 2, 2, 455, 456, 7, 117, 2, 2, 456, 457, 7, 103, 2,
	2, 457, 140, 3, 2, 2, 2, 458, 459, 7, 112, 2, 2, 459, 460, 7, 107, 2, 2,
	460, 461, 7, 110, 2, 2, 461, 142, 3, 2, 2, 2, 462, 463, 7, 107, 2, 2, 463,
	464, 7, 111, 2, 2, 464, 465, 7, 114, 2, 2, 465, 466, 7, 113, 2, 2, 466,
	467, 7, 116, 2, 2, 467, 468, 7, 118, 2, 2, 468, 144, 3, 2, 2, 2, 469, 470,
	7, 104, 2, 2, 470, 471, 7, 116, 2, 2, 471, 472, 7, 113, 2, 2, 472, 473,
	7, 111, 2, 2, 473, 146, 3, 2, 2, 2, 474, 475, 7, 101, 2, 2, 475, 476, 7,
	116, 2, 2, 476, 477, 7, 103, 2, 2, 477, 478, 7, 99, 2, 2, 478, 479, 7,
	118, 2, 2, 479, 480, 7, 103, 2, 2, 480, 148, 3, 2, 2, 2, 481, 482, 7, 102,
	2, 2, 482, 483, 7, 103, 2, 2, 483, 484, 7, 117, 2, 2, 484, 485, 7, 118,
	2, 2, 485, 486, 7, 116, 2, 2, 486, 487, 7, 113, 2, 2, 487, 488, 7, 123,
	2, 2, 488, 150, 3, 2, 2, 2, 489, 493, 5, 153, 77, 2, 490, 492, 5, 155,
	78, 2, 491, 490, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493, 491, 3, 2, 2, 2,
	493, 494, 3, 2, 2, 2, 494, 152, 3, 2, 2, 2, 495, 493, 3, 2, 2, 2, 496,
	498, 9, 2, 2, 2, 497, 496, 3, 2, 2, 2, 498, 154, 3, 2, 2, 2, 499, 502,
	9, 3, 2, 2, 500, 502, 5, 153, 77, 2, 501, 499, 3, 2, 2, 2, 501, 500, 3,
	2, 2, 2, 502, 156, 3, 2, 2, 2, 503, 511, 9, 3, 2, 2, 504, 506, 9, 4, 2,
	2, 505, 504, 3, 2, 2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507,
	508, 3, 2, 2, 2, 508, 510, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 512,
	9, 3, 2, 2, 511, 507, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 513, 3, 2,
	2, 2, 513, 514, 7, 48, 2, 2, 514, 522, 9, 3, 2, 2, 515, 517, 9, 4, 2, 2,
	516, 515, 3, 2, 2, 2, 517, 520, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2, 518,
	519, 3, 2, 2, 2, 519, 521, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 521, 523,
	9, 3, 2, 2, 522, 518, 3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523, 158, 3, 2,
	2, 2, 524, 528, 9, 3, 2, 2, 525, 527, 9, 4, 2, 2, 526, 525, 3, 2, 2, 2,
	527, 530, 3, 2, 2, 2, 528, 526, 3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529,
	160, 3, 2, 2, 2, 530, 528, 3, 2, 2, 2, 531, 532, 7, 50, 2, 2, 532, 533,
	7, 100, 2, 2, 533, 535, 3, 2, 2, 2, 534, 536, 9, 5, 2, 2, 535, 534, 3,
	2, 2, 2, 536, 537, 3, 2, 2, 2, 537, 535, 3, 2, 2, 2, 537, 538, 3, 2, 2,
	2, 538, 162, 3, 2, 2, 2, 539, 540, 7, 50, 2, 2, 540, 541, 7, 113, 2, 2,
	541, 543, 3, 2, 2, 2, 542, 544, 9, 6, 2, 2, 543, 542, 3, 2, 2, 2, 544,
	545, 3, 2, 2, 2, 545, 543, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546, 164,
	3, 2, 2, 2, 547, 548, 7, 50, 2, 2, 548, 549, 7, 122, 2, 2, 549, 551, 3,
	2, 2, 2, 550, 552, 9, 7, 2, 2, 551, 550, 3, 2, 2, 2, 552, 553, 3, 2, 2,
	2, 553, 551, 3, 2, 2, 2, 553, 554, 3, 2, 2, 2, 554, 166, 3, 2, 2, 2, 555,
	556, 7, 50, 2, 2, 556, 560, 9, 8, 2, 2, 557, 559, 9, 9, 2, 2, 558, 557,
	3, 2, 2, 2, 559, 562, 3, 2, 2, 2, 560, 558, 3, 2, 2, 2, 560, 561, 3, 2,
	2, 2, 561, 168, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 563, 567, 7, 36, 2, 2,
	564, 566, 5, 171, 86, 2, 565, 564, 3, 2, 2, 2, 566, 569, 3, 2, 2, 2, 567,
	565, 3, 2, 2, 2, 567, 568, 3, 2, 2, 2, 568, 570, 3, 2, 2, 2, 569, 567,
	3, 2, 2, 2, 570, 571, 7, 36, 2, 2, 571, 170, 3, 2, 2, 2, 572, 575, 5, 173,
	87, 2, 573, 575, 10, 10, 2, 2, 574, 572, 3, 2, 2, 2, 574, 573, 3, 2, 2,
	2, 575, 172, 3, 2, 2, 2, 576, 577, 7, 94, 2, 2, 577, 590, 9, 11, 2, 2,
	578, 579, 7, 94, 2, 2, 579, 580, 7, 119, 2, 2, 580, 581, 3, 2, 2, 2, 581,
	583, 7, 125, 2, 2, 582, 584, 5, 175, 88, 2, 583, 582, 3, 2, 2, 2, 584,
	585, 3, 2, 2, 2, 585, 583, 3, 2, 2, 2, 585, 586, 3, 2, 2, 2, 586, 587,
	3, 2, 2, 2, 587, 588, 7, 127, 2, 2, 588, 590, 3, 2, 2, 2, 589, 576, 3,
	2, 2, 2, 589, 578, 3, 2, 2, 2, 590, 174, 3, 2, 2, 2, 591, 592, 9, 12, 2,
	2, 592, 176, 3, 2, 2, 2, 593, 595, 9, 13, 2, 2, 594, 593, 3, 2, 2, 2, 595,
	596, 3, 2, 2, 2, 596, 594, 3, 2, 2, 2, 596, 597, 3, 2, 2, 2, 597, 598,
	3, 2, 2, 2, 598, 599, 8, 89, 2, 2, 599, 178, 3, 2, 2, 2, 600, 602, 9, 14,
	2, 2, 601, 600, 3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603, 601, 3, 2, 2, 2,
	603, 604, 3, 2, 2, 2, 604, 605, 3, 2, 2, 2, 605, 606, 8, 90, 2, 2, 606,
	180, 3, 2, 2, 2, 607, 608, 7, 49, 2, 2, 608, 609, 7, 44, 2, 2, 609, 614,
	3, 2, 2, 2, 610, 613, 5, 181, 91, 2, 611, 613, 11, 2, 2, 2, 612, 610, 3,
	2, 2, 2, 612, 611, 3, 2, 2, 2, 613, 616, 3, 2, 2, 2, 614, 615, 3, 2, 2,
	2, 614, 612, 3, 2, 2, 2, 615, 617, 3, 2, 2, 2, 616, 614, 3, 2, 2, 2, 617,
	618, 7, 44, 2, 2, 618, 619, 7, 49, 2, 2, 619, 620, 3, 2, 2, 2, 620, 621,
	8, 91, 2, 2, 621, 182, 3, 2, 2, 2, 622, 623, 7, 49, 2, 2, 623, 624, 7,
	49, 2, 2, 624, 628, 3, 2, 2, 2, 625, 627, 10, 15, 2, 2, 626, 625, 3, 2,
	2, 2, 627, 630, 3, 2, 2, 2, 628, 626, 3, 2, 2, 2, 628, 629, 3, 2, 2, 2,
	629, 631, 3, 2, 2, 2, 630, 628, 3, 2, 2, 2, 631, 632, 8, 92, 2, 2, 632,
	184, 3, 2, 2, 2, 24, 2, 493, 497, 501, 507, 511, 518, 522, 528, 537, 545,
	553, 560, 567, 574, 585, 589, 596, 603, 612, 614, 628, 3, 2, 3, 2,
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
	"", "';'", "'{'", "'}'", "','", "':'", "'.'", "'['", "']'", "'<->'", "'='",
	"'||'", "'&&'", "'|'", "'^'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
	"'<<'", "'>>'", "'+'", "'-'", "'*'", "'/'", "'%'", "'auth'", "'&'", "'!'",
	"'<-'", "'<-!'", "'?'", "", "'as'", "'as?'", "'as!'", "'@'", "'('", "')'",
	"'transaction'", "'struct'", "'resource'", "'contract'", "'interface'",
	"'fun'", "'event'", "'emit'", "'pre'", "'post'", "'priv'", "'pub'", "'set'",
	"'access'", "'all'", "'self'", "'account'", "'return'", "'break'", "'continue'",
	"'let'", "'var'", "'if'", "'else'", "'while'", "'for'", "'in'", "'true'",
	"'false'", "'nil'", "'import'", "'from'", "'create'", "'destroy'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "ShiftLeft", "ShiftRight",
	"Plus", "Minus", "Mul", "Div", "Mod", "Auth", "Ampersand", "Negate", "Move",
	"MoveForced", "Optional", "NilCoalescing", "Casting", "FailableCasting",
	"ForceCasting", "ResourceAnnotation", "OpenParen", "CloseParen", "Transaction",
	"Struct", "Resource", "Contract", "Interface", "Fun", "Event", "Emit",
	"Pre", "Post", "Priv", "Pub", "Set", "Access", "All", "Self", "Account",
	"Return", "Break", "Continue", "Let", "Var", "If", "Else", "While", "For",
	"In", "True", "False", "Nil", "Import", "From", "Create", "Destroy", "Identifier",
	"PositiveFixedPointLiteral", "DecimalLiteral", "BinaryLiteral", "OctalLiteral",
	"HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral", "WS", "Terminator",
	"BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "Equal", "Unequal", "Less",
	"Greater", "LessEqual", "GreaterEqual", "ShiftLeft", "ShiftRight", "Plus",
	"Minus", "Mul", "Div", "Mod", "Auth", "Ampersand", "Negate", "Move", "MoveForced",
	"Optional", "NilCoalescing", "Casting", "FailableCasting", "ForceCasting",
	"ResourceAnnotation", "OpenParen", "CloseParen", "Transaction", "Struct",
	"Resource", "Contract", "Interface", "Fun", "Event", "Emit", "Pre", "Post",
	"Priv", "Pub", "Set", "Access", "All", "Self", "Account", "Return", "Break",
	"Continue", "Let", "Var", "If", "Else", "While", "For", "In", "True", "False",
	"Nil", "Import", "From", "Create", "Destroy", "Identifier", "IdentifierHead",
	"IdentifierCharacter", "PositiveFixedPointLiteral", "DecimalLiteral", "BinaryLiteral",
	"OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral",
	"QuotedText", "EscapedCharacter", "HexadecimalDigit", "WS", "Terminator",
	"BlockComment", "LineComment",
}

type CadenceLexer struct {
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

func NewCadenceLexer(input antlr.CharStream) *CadenceLexer {

	l := new(CadenceLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Cadence.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CadenceLexer tokens.
const (
	CadenceLexerT__0                      = 1
	CadenceLexerT__1                      = 2
	CadenceLexerT__2                      = 3
	CadenceLexerT__3                      = 4
	CadenceLexerT__4                      = 5
	CadenceLexerT__5                      = 6
	CadenceLexerT__6                      = 7
	CadenceLexerT__7                      = 8
	CadenceLexerT__8                      = 9
	CadenceLexerT__9                      = 10
	CadenceLexerT__10                     = 11
	CadenceLexerT__11                     = 12
	CadenceLexerT__12                     = 13
	CadenceLexerT__13                     = 14
	CadenceLexerEqual                     = 15
	CadenceLexerUnequal                   = 16
	CadenceLexerLess                      = 17
	CadenceLexerGreater                   = 18
	CadenceLexerLessEqual                 = 19
	CadenceLexerGreaterEqual              = 20
	CadenceLexerShiftLeft                 = 21
	CadenceLexerShiftRight                = 22
	CadenceLexerPlus                      = 23
	CadenceLexerMinus                     = 24
	CadenceLexerMul                       = 25
	CadenceLexerDiv                       = 26
	CadenceLexerMod                       = 27
	CadenceLexerAuth                      = 28
	CadenceLexerAmpersand                 = 29
	CadenceLexerNegate                    = 30
	CadenceLexerMove                      = 31
	CadenceLexerMoveForced                = 32
	CadenceLexerOptional                  = 33
	CadenceLexerNilCoalescing             = 34
	CadenceLexerCasting                   = 35
	CadenceLexerFailableCasting           = 36
	CadenceLexerForceCasting              = 37
	CadenceLexerResourceAnnotation        = 38
	CadenceLexerOpenParen                 = 39
	CadenceLexerCloseParen                = 40
	CadenceLexerTransaction               = 41
	CadenceLexerStruct                    = 42
	CadenceLexerResource                  = 43
	CadenceLexerContract                  = 44
	CadenceLexerInterface                 = 45
	CadenceLexerFun                       = 46
	CadenceLexerEvent                     = 47
	CadenceLexerEmit                      = 48
	CadenceLexerPre                       = 49
	CadenceLexerPost                      = 50
	CadenceLexerPriv                      = 51
	CadenceLexerPub                       = 52
	CadenceLexerSet                       = 53
	CadenceLexerAccess                    = 54
	CadenceLexerAll                       = 55
	CadenceLexerSelf                      = 56
	CadenceLexerAccount                   = 57
	CadenceLexerReturn                    = 58
	CadenceLexerBreak                     = 59
	CadenceLexerContinue                  = 60
	CadenceLexerLet                       = 61
	CadenceLexerVar                       = 62
	CadenceLexerIf                        = 63
	CadenceLexerElse                      = 64
	CadenceLexerWhile                     = 65
	CadenceLexerFor                       = 66
	CadenceLexerIn                        = 67
	CadenceLexerTrue                      = 68
	CadenceLexerFalse                     = 69
	CadenceLexerNil                       = 70
	CadenceLexerImport                    = 71
	CadenceLexerFrom                      = 72
	CadenceLexerCreate                    = 73
	CadenceLexerDestroy                   = 74
	CadenceLexerIdentifier                = 75
	CadenceLexerPositiveFixedPointLiteral = 76
	CadenceLexerDecimalLiteral            = 77
	CadenceLexerBinaryLiteral             = 78
	CadenceLexerOctalLiteral              = 79
	CadenceLexerHexadecimalLiteral        = 80
	CadenceLexerInvalidNumberLiteral      = 81
	CadenceLexerStringLiteral             = 82
	CadenceLexerWS                        = 83
	CadenceLexerTerminator                = 84
	CadenceLexerBlockComment              = 85
	CadenceLexerLineComment               = 86
)
