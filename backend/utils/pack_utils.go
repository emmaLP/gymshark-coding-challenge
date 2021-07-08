package utils

import "sort"

func OrderPacksDesc(packs []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	return packs
}
