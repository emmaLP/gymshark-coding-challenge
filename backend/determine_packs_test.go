package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeterminePacks(t *testing.T) {
	data := &Data{}
	data.ItemsOrdered = 250
	data.Packs = []int{250, 500, 1000, 2000, 5000}
	jsonDataBytes := determinePacks(data)
	assert.Equal(t, "{\"Packs\":{\"250\":1}}", string(jsonDataBytes[:]))

}
