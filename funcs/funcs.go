package funcs

import "math"

const eps = 1e-3

func Dirac(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		if v > -eps && v < eps {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func Comb(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		closestInt := math.Round(v)
		if math.Abs(v-closestInt) < eps {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func Rect(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		if math.Abs(v) <= 0.5 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func Sinc(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		if v == 0 {
			result[i] = 1
		} else {
			result[i] = (math.Sin(math.Pi * v)) / (math.Pi * v)
		}
	}
	return result
}

func Sgn(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		if v < 0 {
			result[i] = -1
		} else if v > 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func Triang(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		if math.Abs(v) < 1 {
			result[i] = 1 - math.Abs(v)
		} else {
			result[i] = 0
		}
	}
	return result
}

func FourierTransform(y []float64) []float64 {
	n := len(y)
	re := make([]float64, n)
	im := make([]float64, n)
	res := make([]float64, n/2)

	for k := 0; k < n; k++ {
		for j := 0; j < n; j++ {
			re[k] += y[j] * math.Cos(-2*math.Pi*float64(k*j)/float64(n))
			im[k] += y[j] * math.Sin(-2*math.Pi*float64(k*j)/float64(n))
		}
	}

	for i := 0; i < n/2; i++ {
		res[i] = math.Sqrt(re[i]*re[i] + im[i]*im[i])
	}

	return res
}

func SinPi(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		result[i] = math.Sin(math.Pi * v)
	}
	return result
}

func CosPi(x []float64) []float64 {
	result := make([]float64, len(x))
	for i, v := range x {
		result[i] = math.Cos(math.Pi * v)
	}
	return result
}
