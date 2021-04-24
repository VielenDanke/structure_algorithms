package main

func fibSequence(n int) int {
	if n <= 2 {
		return 1
	}
	return fibSequence(n-1) + fibSequence(n-2)
}

func fibGo(n int) int {
	var x, y = 1, 1
	for i := 0; i < n; i++ {
		x, y = x+y, x
	}
	return x
}

func fibSequenceD(n int, m map[int]int) int {
	if n <= 2 {
		return 1
	}
	fv, fOk := m[n-1]
	sv, sOk := m[n-2]
	var fRes, sRes int
	if fOk {
		fRes = fv
	} else {
		fRes = fibSequenceD(n-1, m)
		m[n-1] = fRes
	}
	if sOk {
		sRes = sv
	} else {
		sRes = fibSequenceD(n-2, m)
		m[n-2] = sRes
	}
	return fRes + sRes
}

func fibSequenceB(n int, m map[int]int) int {
	num, ok := m[n]
	if ok {
		return num
	}
	if n <= 2 {
		return 1
	}
	res := fibSequenceB(n-1, m) + fibSequenceB(n-2, m)
	m[n] = res
	return res
}

func fibSequenceC(n int) int {
	if n <= 2 {
		return 1
	}
	fibNums := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		fibNums = append(fibNums, fibNums[i-1]+fibNums[i-2])
	}
	return fibNums[n]
}
