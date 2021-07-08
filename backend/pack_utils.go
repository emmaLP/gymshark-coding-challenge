package main

import "sort"

func orderPacksDesc(packs []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	return packs
}
func packContains(arr []int, value int) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}
