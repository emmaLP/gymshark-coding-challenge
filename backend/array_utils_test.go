package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var array = []int{100, 500, 250, 2000, 5000, 10}

func TestOrderDesc(t *testing.T) {
	var expectedResult = []int{5000, 2000, 500, 250, 100, 10}
	var actualResult = orderDesc(array)
	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("Ordering of array failed, expected %v, got %v", expectedResult, actualResult)

	}
}

func TestContainsFalse(t *testing.T) {
	packContains := contains(array, 123)
	assert.False(t, packContains)
}

func TestContainsTrue(t *testing.T) {
	packContains := contains(array, 5000)
	assert.True(t, packContains)
}

func TestFindLowAndHighForValue160(t *testing.T) {
	var array = []int{50, 100, 500, 2000, 5000, 10}
	low, high := findLowAndHighNearestToValue(array, 160)
	assert.Equal(t, 100, low)
	assert.Equal(t, 500, high)
}

func TestFindLowAndHighForValue1000(t *testing.T) {
	var array = []int{50, 100, 500, 2000, 5000, 10}
	low, high := findLowAndHighNearestToValue(array, 1000)
	assert.Equal(t, 500, low)
	assert.Equal(t, 2000, high)
}

func TestGetNearestValue10(t *testing.T) {
	var array = []int{50, 100, 500, 2000, 5000, 10}
	nearest := getNearestValue(10, array)
	assert.Equal(t, 10, nearest)
}

func TestGetNearestValue150(t *testing.T) {
	var array = []int{50, 100, 500, 2000, 5000, 10}
	nearest := getNearestValue(150, array)
	assert.Equal(t, 100, nearest)
}
