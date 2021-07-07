package main

import "testing"
import "github.com/stretchr/testify/assert"

var packs = []int{100, 250, 500, 1000, 2000, 5000}

func TestCalcPack_Quantity1(t *testing.T) {
	var res map[int]int
	expectedResultKey := 100

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}
