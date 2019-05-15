package statistic

func Average(xs []float64) float64 {
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func Sum(xs []float64) float64 {
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total
}
