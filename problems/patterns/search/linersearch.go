package main

func linerSearching(str []string, search string) int {
	for k, v := range str {
		if v == search {
			return k
		}
	}
	return -1
}
