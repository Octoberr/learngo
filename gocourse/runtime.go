package gocourse

import "fmt"

type Run struct {
	Time   int
	Result string
	Failed bool
}

type Set map[string]struct{}

func RecognizeKey() {
	set := make(Set)
	for _, key := range []string{"a", "b", "c", "d"} {
		set[key] = struct{}{}
	}
	fmt.Println(len(set))
	_, ok := set["a"]
	if ok {
		fmt.Println("a is in the set")
	}
	delete(set, "a")
	fmt.Println(set)

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
