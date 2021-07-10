package main

import (
	"math"
	"sort"
)

func orderDesc(arr []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

func orderAsc(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func contains(arr []int, value int) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

//Find the nearest numbers either side of the specified value in an array
func findLowAndHighNearestToValue(arr []int, value int) (int, int) {
	arr = orderAsc(arr)
	var low = 0
	var high = len(arr) - 1
	var mid int
	for high-low > 1 {
		mid = int(math.Floor(float64((low + high) / 2)))
		if arr[mid] < value {
			low = mid
		} else {
			high = mid
		}
	}

	return arr[low], arr[high]
}

//Get the nearest value in an array to the number specified
func getNearestValue(number int, arr []int) int {
	var nearest = arr[0]
	for _, pack := range arr {
		a := math.Abs(float64(number - pack))
		b := math.Abs(float64(number - nearest))

		if a < b {
			nearest = pack
		}
	}

	return nearest
}
