package polyfit

import (
	"fmt"
	"github.com/AnthonyHewins/vandermonde"
	"gonum.org/v1/gonum/mat"
)

func PolynomialRegression(x []float64, y []float64, maxDeg int) (coef []float64, err error) {
	n := len(x)
	if n == 0 { return []float64{}, nil }

	if m := len(y); n != m {
		return nil, fmt.Errorf("length mismatch: len(x)=%v is not equal to len(y)=%v", n, m)
	}

	if maxDeg > n {
		return nil, fmt.Errorf("degree must be at least equal to the number of samples")
	}

	X, err := vandermonde.VandermondeWindow(x, 0, maxDeg + 1, 0)
	if err != nil { return nil, err }

	var Theta mat.Dense
	var X_2 mat.Dense
	var X_2_inv mat.Dense

	X_t := X.T()
	Y := mat.NewVecDense(n, y)

	X_2.Mul(X_t, X)
	if err := X_2_inv.Inverse(&X_2); err != nil {
		return nil, err
	}

	Theta.Product(&X_2_inv, X_t, Y)

	return Theta.RawMatrix().Data, nil
}

