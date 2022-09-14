package num

import "math"

func Sq(x float64) float64 { return x * x }

func Dist(X1, Y1, X2, Y2 float64) float64 {
	return math.Sqrt(Sq(X1-X2) + Sq(Y1-Y2))
}

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func Clamp(value, min, max float64) float64 {
	r := value
	if r < min {
		r = min
	} else if r > max {
		r = max
	}
	return r
}
