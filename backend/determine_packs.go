package main

import (
	"encoding/json"
	"fmt"
)

func determinePacks(data *Data) []byte {
	orderDesc(data.Packs)
	fmt.Println("Packs: ", data.Packs)

	result := calculatePacksFromRequestedItems(data.ItemsOrdered, data.Packs, []int{})

	response := Response{Packs: result}
	var jsonData []byte

	jsonData, _ = json.Marshal(response)
	return jsonData
}
