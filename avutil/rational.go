// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/rational.h>
//#include <stdlib.h>
import "C"
import (
	"math/big"
	"unsafe"
)

type (
	Rational C.struct_AVRational
)

// Compare rational to another, returns 0 if equal, 1 if r>other and -1 if r<other.
//
// C-Function: av_cmp_q
func (r Rational) Compare(other Rational) int {
	return int(C.av_cmp_q((C.struct_AVRational)(r), (C.struct_AVRational(other))))
}

// Divide rational by another.
//
// C-Function: av_div_q
func (r Rational) Divide(other Rational) Rational {
	return Rational(C.av_div_q((C.struct_AVRational)(r), (C.struct_AVRational)(other)))
}

// Multiply rational by another.
//
// C-Function: av_mul_q
func (r Rational) Multiply(other Rational) Rational {
	return Rational(C.av_mul_q((C.struct_AVRational)(r), (C.struct_AVRational)(other)))
}

// Subtract other rational.
//
// C-Function: av_sub_q
func (r Rational) Subtract(other Rational) Rational {
	return Rational(C.av_sub_q((C.struct_AVRational)(r), (C.struct_AVRational)(other)))
}

// NearestIndex returns the index of the nearest rational in the provided array.
//
// C-Function: av_find_nearest_q_idx
func (r Rational) NearestIndex(arr []Rational) int {
	return int(C.av_find_nearest_q_idx((C.struct_AVRational)(r), (*C.struct_AVRational)(unsafe.Pointer(&arr))))
}

// Nearer determines whether rational is closer to the first or second provided
// rationals.  Returns 1 if first is closer, -1 if second is closer, and 0 if
// they have the same distance.
//
// C-Function: av_nearer_q
func (r Rational) Nearer(first Rational, second Rational) int {
	return int(C.av_nearer_q((C.struct_AVRational)(r), (C.struct_AVRational)(first), (C.struct_AVRational)(second)))
}

// Float converts rational to float64
//
// C-Function: av_q2d
func (r Rational) Float() float64 {
	return float64(C.av_q2d((C.struct_AVRational)(r)))
}

// Rat converts rational to *big.Rat
func (r Rational) Rat() *big.Rat {
	return big.NewRat(int64(r.num), int64(r.den))
}

// FloatToRational converts a float64 to a Rational, with a maximum denominator.
//
// C-Function: av_d2q
func FloatToRational(f float64, max int) Rational {
	return Rational(C.av_d2q((C.double)(f), (C.int)(max)))
}
