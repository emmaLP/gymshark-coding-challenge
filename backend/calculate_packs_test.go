package main

import (
	"reflect"
	"testing"
)
import "github.com/stretchr/testify/assert"

var packs = []int{250, 500, 1000, 2000, 5000}

func TestCalcPackQuantity1(t *testing.T) {
	var res map[int]int
	expectedResultKey := 250

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity250(t *testing.T) {
	var res map[int]int
	expectedResultKey := 250

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity251(t *testing.T) {
	var res map[int]int
	expectedResultKey := 500

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity252(t *testing.T) {
	var res map[int]int
	expectedResultKey := 500

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity501(t *testing.T) {
	var res map[int]int
	expectedResult := make(map[int]int)
	expectedResult[500] = 1
	expectedResult[250] = 1

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.True(t, reflect.DeepEqual(expectedResult, actualResult))
}

func TestCalcPackQuantity12001(t *testing.T) {
	var res map[int]int
	expectedResult := make(map[int]int)
	expectedResult[5000] = 2
	expectedResult[2000] = 1
	expectedResult[250] = 1

	actualResult := calculatePacksFromRequestedItems(1, packs, res)
	assert.True(t, reflect.DeepEqual(expectedResult, actualResult))
}
