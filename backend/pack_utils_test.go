package main

import (
	"reflect"
	"testing"
)
import "github.com/stretchr/testify/assert"

var array = []int{100, 500, 250, 2000, 5000, 10}

func TestOrderDesc(t *testing.T) {
	var expectedResult = []int{5000, 2000, 500, 250, 100, 10}
	var actualResult = orderPacksDesc(array)
	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("Ordering of array failed, expected %v, got %v", expectedResult, actualResult)

	}
}

func TestPackContainsFalse(t *testing.T) {
	packContains := packContains(array, 123)
	assert.False(t, packContains)
}

func TestPackContainsTrue(t *testing.T) {
	packContains := packContains(array, 5000)
	assert.True(t, packContains)
}
