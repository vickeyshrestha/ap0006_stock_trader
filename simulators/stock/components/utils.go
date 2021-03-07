package stocks

func ClearMap(m map[string]float64) {
	for k := range m {
		delete(m, k)
	}
}
