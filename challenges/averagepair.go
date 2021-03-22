package main

func averagePair(arr []int, average float64) bool {
	f := 0
	s := len(arr) - 1
	for f < s {
		res := (float64(arr[f]) + float64(arr[s])) / 2
		if res == average {
			return true
		}
		if res < average {
			f++
		}
		if res > average {
			s++
		}
	}
	return false
}
