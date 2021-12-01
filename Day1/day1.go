package Day1

func GetDepths(data []int) int{
	counter := 0
	lastMeasurement := data[0]

	for index, measurement := range data {
		if (index > 0)  && (lastMeasurement < measurement){
			counter++
		}
		lastMeasurement = measurement
	}
	return counter
}

func GetDepthsWindow(data []int) int{
	counter := 0
	lastWindowValue := -1

	for index, measurement := range data {
		if index > 1 {
			windowValue := data[index-2] + data[index-1] + measurement
			if lastWindowValue >= 0 && (windowValue > lastWindowValue) {
				counter++
			}
			lastWindowValue = windowValue
		}
	}
	return counter
}
