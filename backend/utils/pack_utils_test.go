package utils

import (
	"reflect"
	"testing"
)

var array = []int{100, 500, 250, 2000, 5000, 10}

func TestOrder(t *testing.T) {
	var expectedResult = []int{5000, 2000, 500, 250, 100, 10}
	var actualResult = OrderPacksDesc(array)
	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Errorf("Ordering of array failed, expected %v, got %v", expectedResult, actualResult)

	}
}
