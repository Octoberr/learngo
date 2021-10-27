package test

import (
	"fmt"
	"learngo/gocourse"
	"testing"
)

func Test_binarysearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := gocourse.BinarySearch(a, 5)
	if res != 4 {
		t.Error("BinarySearch failed")
	}
	fmt.Println(res)

}
