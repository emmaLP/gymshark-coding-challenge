package main

var originalItemsOrdered int

func calculatePacksFromRequestedItems(itemsOrdered int, packs []int, res []int) map[int]int {
	if contains(packs, itemsOrdered) {
		return mapResults(append(res, itemsOrdered))
	} else {
		remaining := itemsOrdered
		if len(res) == 0 {
			originalItemsOrdered = itemsOrdered
		} else {
			temp := make([]int, len(packs))
			copy(temp, packs)
			lowValue, highValue := findLowAndHighNearestToValue(orderAsc(temp), remaining)
			if remaining >= lowValue && remaining <= highValue {
				res = append(res, highValue)
				remaining = remaining - highValue
			} else if remaining >= lowValue {
				res = append(res, lowValue)
				remaining = remaining - lowValue
			}
		}
		if remaining > 0 {
			for _, pack := range packs {
				if remaining >= pack {
					res = append(res, pack)
					remaining = remaining - pack
					return calculatePacksFromRequestedItems(remaining, packs, res)
				}
			}
			var smallestPack = packs[len(packs)-1]
			if remaining < smallestPack {
				res = append(res, smallestPack)
			}

		}

		//Perform validation to ensure we are sending out the smallest number of packs
		sumResults := sum(res)
		if contains(packs, sumResults) {
			res = []int{sumResults}
		} else if nearest := getNearestValue(originalItemsOrdered, packs); nearest < sumResults && nearest > originalItemsOrdered {
			res = []int{nearest}
		}
	}
	return mapResults(res)
}

func mapResults(arr []int) map[int]int {
	result := make(map[int]int)
	for _, pack := range arr {
		result[pack] += 1
	}

	return result
}
