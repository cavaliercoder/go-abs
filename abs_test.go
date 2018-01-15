package abs

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	MaxInt int64 = 1<<63 - 1
	MinInt int64 = -1 << 63
)

// An absFunc is a function that returns the absolute value of an integer.
type absFunc func(int64) int64

func funcName(v interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	return s[strings.LastIndex(s, ".")+1:]
}

func TestAbs(t *testing.T) {
	inputs := []int64{MinInt + 1, MinInt + 2, -1, -0, 1, 2, MaxInt - 1, MaxInt}
	outputs := []int64{MaxInt, MaxInt - 1, 1, 0, 1, 2, MaxInt - 1, MaxInt}
	testFuncs := []absFunc{
		WithBranch,
		// WithStdLib, // test failure expected on large numbers
		WithTwosComplement,
		WithASM,
	}
	for _, f := range testFuncs {
		testName := funcName(f)
		t.Run(testName, func(t *testing.T) {
			for i := 0; i < len(inputs); i++ {
				actual := f(inputs[i])
				if actual != outputs[i] {
					t.Errorf("%s(%d)", testName, inputs[i])
					t.Errorf("	input:		%064b (%d)", uint64(inputs[i]), inputs[i])
					t.Errorf("	expected:	%064b (%d)", uint64(outputs[i]), outputs[i])
					t.Errorf("	actual:		%064b (%d)", uint64(actual), actual)
				}
			}
		})
	}
}

var r uint64 = 0xdeadbeef

// Pseudo-random number generator adapted from
// https://github.com/dgryski/trifles/blob/master/fastrand/fastrand.go
func Rand() int64 {
	r ^= r >> 12 // a
	r ^= r << 25 // b
	r ^= r >> 27 // c
	return int64(r * 2685821657736338717)
}

// sink is used to prevent the compiler from dropping function calls where the
// returned value is not used within benchmarks.
var sink int64

func BenchmarkWithBranch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sink += WithBranch(Rand())
	}
}

func BenchmarkWithStdLib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sink += WithStdLib(Rand())
	}
}

func BenchmarkWithTwosComplement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sink += WithTwosComplement(Rand())
	}
}

func BenchmarkWithASM(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sink += WithASM(Rand())
	}
}
