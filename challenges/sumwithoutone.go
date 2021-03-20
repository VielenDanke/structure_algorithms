package main

func sumWithoutOne(arr []int) []int {
	// difine the exclusion index
	sumArr := make([]int, 0)
	exclusion := 0
	tempSum := 0
	counter := 0

	// loop over and sum
	for {
		if exclusion == len(arr) {
			break
		}
		if counter == len(arr) {
			sumArr = append(sumArr, tempSum)
			tempSum = 0
			counter = 0
			exclusion++
		}
		if counter == exclusion {
			counter++
			continue
		}
		tempSum += arr[counter]
		counter++
	}

	// return new arr
	return sumArr
}
