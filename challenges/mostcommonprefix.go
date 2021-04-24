package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	sort.Strings(strs)
	cp := bytes.NewBufferString("")
	fArr := []rune(strs[0])
	sArr := []rune(strs[len(strs)-1])
	for k, v := range fArr {
		if k == len(sArr) {
			break
		}
		fmt.Println(v, " ", sArr[k])
		if v == sArr[k] {
			cp.WriteRune(v)
		}
		if v != sArr[k] {
			break
		}
	}
	pref := cp.String()
	for _, v := range strs {
		if !strings.HasPrefix(v, pref) {
			return ""
		}
	}
	return pref
}
