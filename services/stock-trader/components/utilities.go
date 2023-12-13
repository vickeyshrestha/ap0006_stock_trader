package stocktrader

func enqueue(queue []TimeSeries, element TimeSeries) []TimeSeries {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []TimeSeries) (TimeSeries, []TimeSeries) {
	element := queue[0]
	if len(queue) == 1 {
		var tmp []TimeSeries
		return element, tmp

	}
	return element, queue[1:]
}

func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}
