package SplitSlice

func SplitStringSlice(data []string, step int) (result [][]string) {
	var StartIndex = 0
	var EndIndex = 0
	Length := len(data)
	for {
		EndIndex = StartIndex + step
		if EndIndex >= Length {
			result = append(result, data[StartIndex:])
			return
		}
		result = append(result, data[StartIndex:EndIndex])
		StartIndex += step
	}
}

func SplitInt64Slice(data []int64, step int) (result [][]int64) {
	var StartIndex = 0
	var EndIndex = 0
	Length := len(data)
	for {
		EndIndex = StartIndex + step
		if EndIndex >= Length {
			result = append(result, data[StartIndex:])
			return
		}
		result = append(result, data[StartIndex:EndIndex])
		StartIndex += step
	}
}

func SplitFloat64Slice(data []float64, step int) (result [][]float64) {
	var StartIndex = 0
	var EndIndex = 0
	Length := len(data)
	for {
		EndIndex = StartIndex + step
		if EndIndex >= Length {
			result = append(result, data[StartIndex:])
			return
		}
		result = append(result, data[StartIndex:EndIndex])
		StartIndex += step
	}
}
