package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var packs = []int{250, 500, 1000, 5000, 2000}

func TestCalcPackQuantity1(t *testing.T) {
	expectedResultKey := 250

	actualResult := calculatePacksFromRequestedItems(1, orderDesc(packs), []int{})
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity160(t *testing.T) {
	packs := orderDesc([]int{50, 100, 500, 1000, 5000, 2000})
	expectedResult := make(map[int]int)
	expectedResult[100] = 2

	actualResult := calculatePacksFromRequestedItems(160, orderDesc(packs), []int{})
	fmt.Println(actualResult)
	assert.True(t, reflect.DeepEqual(expectedResult, actualResult))
}

func TestCalcPackQuantity239(t *testing.T) {
	packs := orderDesc([]int{100, 250, 500, 1000, 5000, 2000})
	expectedResult := make(map[int]int)
	expectedResult[250] = 1

	actualResult := calculatePacksFromRequestedItems(239, orderDesc(packs), []int{})
	fmt.Println(actualResult)
	assert.True(t, reflect.DeepEqual(expectedResult, actualResult))
}

func TestCalcPackQuantity250(t *testing.T) {
	expectedResultKey := 250

	actualResult := calculatePacksFromRequestedItems(250, orderDesc(packs), []int{})
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity251(t *testing.T) {
	expectedResultKey := 500

	actualResult := calculatePacksFromRequestedItems(251, orderDesc(packs), []int{})
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity252(t *testing.T) {
	expectedResultKey := 500

	actualResult := calculatePacksFromRequestedItems(252, orderDesc(packs), []int{})
	assert.Contains(t, actualResult, expectedResultKey)
	quantity := actualResult[expectedResultKey]
	assert.Equal(t, 1, quantity)
}

func TestCalcPackQuantity501(t *testing.T) {
	expectedResult := make(map[int]int)
	expectedResult[500] = 1
	expectedResult[250] = 1

	actualResult := calculatePacksFromRequestedItems(501, orderDesc(packs), []int{})
	assert.True(t, reflect.DeepEqual(expectedResult, actualResult))
}

func TestCalcPackQuantity12001(t *testing.T) {
	expectedResult := make(map[int]int)
	expectedResult[5000] = 2
	expectedResult[2000] = 1
	expectedResult[250] = 1

	actualResult := calculatePacksFromRequestedItems(12001, orderDesc(packs), []int{})
	fmt.Println("actual result", actualResult)
	assert.True(t, reflect.DeepEqual(expectedResult, actualResult))
}
