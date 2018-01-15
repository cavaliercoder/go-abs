/*
Package abs provides multiple implementations of the abs function, to compute
the absolute value of a signed, 64-bit integer.

This package complements the following article, which compares each
implementation:

http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
*/
package abs

import "math"

// WithBranch uses control structures to return the absolute value of an
// integer.
func WithBranch(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// WithStdLib uses the standard library's math package to compute the
// absolute value on an integer.
//
// We expect test for correctness to fail on large numbers that overflow
// float64.
func WithStdLib(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

// WithTwosComplement uses a trick from Henry S. Warren's incredible book,
// Hacker's Delight. It utilizes Two's Complement arithmetic to compute the
// absolute value of an integer.
func WithTwosComplement(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// WithASM uses the Two's Complement trick, but implemented in Assembly to
// compute the absolute value of a signed integer.
func WithASM(n int64) int64
