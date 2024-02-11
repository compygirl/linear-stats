package functions

import (
	"math"
	"sort"
)

func Average(inp []int) float64 {
	sum := 0
	for _, el := range inp {
		sum += (el)
	}
	return float64(sum) / float64(len(inp))
}

func Median(inp []int) float64 {
	sort.Ints(inp)
	if len(inp)%2 == 0 {
		ind2 := len(inp) / 2
		ind1 := ind2 - 1
		return float64(inp[ind1]+inp[ind2]) / 2
	} else {
		return float64(inp[len(inp)/2])
	}
}

func Variance(inp []int) float64 {
	avg := Average(inp)
	sum := 0.0
	for _, el := range inp {
		sum += math.Pow(float64(el)-avg, 2)
	}
	return sum / float64(len(inp))
}

func StdDev(inp []int) float64 {
	return math.Sqrt(Variance(inp))
}

func LinearRegLine(inp []int) (float64, float64) {
	mean_y := Average(inp)
	mean_x := 0.0
	for i := 0; i < len(inp); i++ {
		mean_x += float64(i)
	}
	mean_x = mean_x / float64(len(inp))
	e_xy := 0.0
	esqr_x := 0.0
	for i := 0; i < len(inp); i++ {
		e_xy += (float64(inp[i]) - mean_y) * (float64(i+1) - mean_x)
		esqr_x += (float64(i) - mean_x) * (float64(i) - mean_x)
	}
	coef_k := e_xy / esqr_x
	coef_b := mean_y - coef_k*mean_x
	return coef_k, coef_b
}

func PearsonCoef(inp []int) float64 {
	mean_y := Average(inp)
	mean_x := 0.0
	for i := 0; i < len(inp); i++ {
		mean_x += float64(i)
	}
	mean_x = mean_x / float64(len(inp))
	var_x := 0.0
	for i := 0; i < len(inp); i++ {
		var_x += (mean_x - float64(i)) * (mean_x - float64(i))
	}
	var_x = var_x / float64(len(inp))
	std_x := math.Sqrt(var_x)
	std_y := StdDev(inp)
	e_xy := 0.0
	m_std := std_x * std_y
	for i := 0; i < len(inp); i++ {
		e_xy += (float64(inp[i]) - mean_y) * (float64(i) - mean_x)
	}
	cov := e_xy / float64(len(inp))
	r := cov / m_std
	return r
}
