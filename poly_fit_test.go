package polyfit

import (
	"math"
	"testing"
)

var (
	RandomX   = []float64{ 70, 69, 107, 1, 42, 89}
	RandomY   = []float64{170, 20, 2.2, 6,  0,  3}
	RandomAns = []float64{
		-3044.458915,
		3230.173225,
		-183.4957863,
		3.815471992,
		-0.0340936794,
		0.0001107360238,
	}
)

func TestPolyFit(t *testing.T) {
	// 0 element x array yields 0 element polynomial, regardless of y
	base_case, _ := PolyFit([]float64{}, []float64{}, 0)
	if len(base_case) != 0 {
		t.Errorf("Should have gotten blank array for base case")
	}

	_, len_mismatch := PolyFit([]float64{1}, []float64{}, 0)
	if len_mismatch == nil {
		t.Errorf("Should have gotten length mismatch for these 2 arrays: len(x) != len(y)")
	}

	_, not_unique := PolyFit([]float64{1,1}, []float64{1,2}, 0)
	if not_unique == nil {
		t.Errorf("Should have gotten uniqueness error for non-unique x values")
	}

	_, too_large_polynomial := PolyFit([]float64{1,2}, []float64{1,2}, 3)
	if too_large_polynomial == nil {
		t.Errorf("polynomial degree should have been too large for sample size, wasn't")
	}

	linear, _ := PolyFit(
		[]float64{1,2,3,4},
		[]float64{1,2,3,4},
		4,
	)
	comp_array("testing linear", t, []float64{0,1,0,0}, linear)

	quadratic, _ := PolyFit(
		[]float64{1,2,3,4},
		[]float64{1,4,9,16},
		4,
	)
	comp_array("testing quadratic", t, []float64{0,0,1,0}, quadratic)

	cubic, _ := PolyFit(
		[]float64{1,2,3,4},
		[]float64{1,8,27,64},
		4,
	)
	comp_array("testing cubic", t, []float64{0,0,0,1}, cubic)

	// Precomputed polyreg problem where the answer is known up to some sigfigs
	random, _ := PolyFit(
		RandomX,
		RandomY,
		5,
	)
	comp_array("testing random", t, RandomAns, random)
}

func comp_array(s string, t *testing.T, expected, actual []float64) {
	n := len(expected)
	m := len(actual)
	if len(expected) != len(actual) {
		t.Errorf("%s:unequal array lengths: expected %v, got %v", s, n, m)
		return
	}

	for i := 0; i < n; i++ {
		exp := expected[i]
		act := actual[i]
		diff := math.Abs(exp - act)
		if diff > 0.005 {
			t.Errorf("%s:element %v should be %v, but is %v. Diff: %v", s, i, exp, act, diff)
		}
	}
}
