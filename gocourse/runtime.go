package gocourse

type Run struct {
	Time   int
	Result string
	Failed bool
}

func averageRuntimeInSeconds(runs []Run) float64 {
	var totalTime int
	for _, run := range runs {
		totalTime += run.Time
	}
	return float64(totalTime / len(runs))
}

func BinarySearch(dataList []int, target int) int {
	low := 0
	high := len(dataList) - 1
	for low <= high {
		mid := (low + high) / 2
		if dataList[mid] == target {
			return mid
		} else if dataList[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
