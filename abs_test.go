package abs

import (
	"math/rand"
	"testing"
)

const (
	MaxInt int64 = 1<<63 - 1
	MinInt int64 = -1 << 63
)

// An absFunc is a function that returns the absolute value of an integer.
type absFunc func(int64) int64

var (
	testInputs  = []int64{MinInt + 1, MinInt + 2, -1, -0, 1, 2, MaxInt - 1, MaxInt}
	testOutputs = []int64{MaxInt, MaxInt - 1, 1, 0, 1, 2, MaxInt - 1, MaxInt}
	testFuncs   = []struct {
		Name string
		Func absFunc
	}{
		{
			Name: "WithBranch",
			Func: WithBranch,
		},
		{
			// test failure expected on large numbers
			Name: "WithStdLib",
			Func: WithStdLib,
		},
		{
			Name: "WithTwosComplement",
			Func: WithTwosComplement,
		},
		{
			Name: "WithASM",
			Func: WithASM,
		},
	}
)

func TestAbs(t *testing.T) {
	for _, ts := range testFuncs {
		t.Run(ts.Name, func(t *testing.T) {
			for i := 0; i < len(testInputs); i++ {
				actual := ts.Func(testInputs[i])
				if actual != testOutputs[i] {
					t.Errorf("%s(%d)", ts.Name, testInputs[i])
					t.Errorf("	input:		%064b (%d)", uint64(testInputs[i]), testInputs[i])
					t.Errorf("	expected:	%064b (%d)", uint64(testOutputs[i]), testOutputs[i])
					t.Errorf("	actual:		%064b (%d)", uint64(actual), actual)
				}
			}
		})
	}
}

func BenchmarkAbs(b *testing.B) {
	const maxInputs = 10000

	// Use a set of random inputs so that the compiler cannot optimize it away
	benchInputs := make([]int64, maxInputs)

	// Store the outputs so that the compiler cannnot optimize them away
	benchOutputs := make([]int64, maxInputs)

	for i := 0; i < maxInputs; i++ {
		benchInputs[i] = rand.Int63()
		if rand.Float32() > 0.5 {
			benchInputs[i] = -1 * benchInputs[i]
		}
	}

	for _, ts := range testFuncs {
		var i int
		b.Run(ts.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StartTimer()
				benchOutputs[i] = ts.Func(benchInputs[i])
				b.StopTimer()
				if i++; i >= maxInputs {
					i = 0
				}
			}
		})
	}
}
