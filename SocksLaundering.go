package main

func Solution11(K int, C []int, D []int) int {

	groupedCleanSocks := make(map[int]int)
	// Remove clean sock pairs, count pairs
	for _, sockNum := range C {
		groupedCleanSocks[sockNum]++
	}

	var finalPairs int
	var finalCleanSocks []int
	for sockNum, sockCount := range groupedCleanSocks {
		if sockCount%2 != 0 {
			finalCleanSocks = append(finalCleanSocks, sockNum)
		}
		finalPairs += sockCount / 2
	}

	// Count to K through D, see if match exist in C
	var washes int
	for dirtySockIndex, dirtySockNum := range D {
		for cleanSockIndex, cleanSockNum := range finalCleanSocks {
			if cleanSockNum == dirtySockNum {
				washes++
				if washes > K {
					return finalPairs
				}
				finalPairs++
				finalCleanSocks[cleanSockIndex] = -1
				D[dirtySockIndex] = -1
				break
			}
		}
	}

	remainderSockWashes := K - washes
	if remainderSockWashes <= 1 {
		return finalPairs
	}

	groupedDirtySocks := make(map[int]int)
	for _, dirtySockNum := range D {
		if dirtySockNum != -1 {
			groupedDirtySocks[dirtySockNum]++
		}
	}

	for _, sockCount := range groupedDirtySocks {
		for {
			if remainderSockWashes <= 1 {
				return finalPairs
			}
			if sockCount <= 1 {
				break
			}
			sockCount -= 2
			finalPairs++
			remainderSockWashes -= 2
		}
	}
	return finalPairs
}
