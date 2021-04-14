package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstArr := findDeepest(l1, make([]int, 0))
	secondArr := findDeepest(l2, make([]int, 0))

	resArr := calculateSum(firstArr, secondArr)

	return appendNode(resArr)
}

func calculateSum(first, second []int) []int {
	var result []int
	min := findMin(first, second)
	counter := 0
	toAdd := 0
	for counter < min {
		sum := first[counter] + second[counter] + toAdd
		result, toAdd = appendDependsOnSum(sum, result, toAdd)
		counter++
	}
	for counter < len(first) {
		sum := first[counter] + toAdd
		result, toAdd = appendDependsOnSum(sum, result, toAdd)
		counter++
	}
	for counter < len(second) {
		sum := second[counter] + toAdd
		result, toAdd = appendDependsOnSum(sum, result, toAdd)
		counter++
	}
	if toAdd != 0 {
		result = append(result, toAdd)
	}
	return result
}

func findMin(first, second []int) int {
	if len(first) > len(second) {
		return len(second)
	} else {
		return len(first)
	}
}

func appendDependsOnSum(sum int, result []int, toAdd int) ([]int, int) {
	if sum > 9 {
		result = append(result, sum-10)
		toAdd = 1
	} else {
		result = append(result, sum)
		toAdd = 0
	}
	return result, toAdd
}

func appendNode(arr []int) *ListNode {
	if len(arr) == 1 {
		return &ListNode{Val: arr[0]}
	}
	val := arr[:1][0]
	l := &ListNode{Val: val}
	l.Next = appendNode(arr[1:])
	return l
}

func findDeepest(l *ListNode, arr []int) []int {
	if l == nil {
		return arr
	}
	arr = append(arr, l.Val)
	return findDeepest(l.Next, arr)
}

func addTwoNumbersBetter(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	p, q, curr := l1, l2, dummy
	carry := 0
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if q != nil {
			q = q.Next
		}
		if p != nil {
			p = p.Next
		}
		if carry > 0 {
			curr.Next = &ListNode{Val: carry}
		}
	}
	return dummy.Next
}