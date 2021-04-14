package main

import "bytes"

func convert(s string, numRows int) string {
	length := len(s)
	if length < 1 || length > 1000 {
		return ""
	}
	if numRows < 1 || numRows > 1000 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	buffArr := make([]*bytes.Buffer, 0)
	for i := 0; i < calculateMin(numRows, len(s)); i++ {
		buffArr = append(buffArr, bytes.NewBufferString(""))
	}
	isGoingDown := false
	currentRow := 0
	for _, v := range s {
		buffArr[currentRow].WriteRune(v)
		if currentRow == 0 || currentRow == numRows - 1 {
			isGoingDown = !isGoingDown
		}
		if isGoingDown {
			currentRow += 1
		} else {
			currentRow += -1
		}
	}
	resBuff := bytes.NewBufferString("")
	for _, v := range buffArr {
		resBuff.Write(v.Bytes())
	}
	return resBuff.String()
}

func calculateMin(f, s int) (min int) {
	if f > s {
		min = s
		return
	} else {
		min = f
		return
	}
}