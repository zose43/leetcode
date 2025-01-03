package medium

type listNode struct {
	Val  int
	Next *listNode
}

const maxBuf = 100 * 2

func addTwoNumbers(l *listNode, ll *listNode) *listNode {
	buf := make([][]int, maxBuf)
	var i int
	for ; ; i++ {
		if l == nil && ll == nil {
			break
		} else {
			buf[i] = make([]int, 0, 2)
			if l != nil {
				buf[i] = append(buf[i], l.Val)
				l = l.Next
			}
			if ll != nil {
				buf[i] = append(buf[i], ll.Val)
				ll = ll.Next
			}
		}
	}
	result := new(listNode)
	var parent *listNode
	for j := 0; j <= i-1; j++ {
		sum := calc(buf[j])
		sum = processUnlimitedFirstOrLast(sum, &i, &buf[j+1], j+1 <= i-1)
		if j == 0 {
			result.Val = sum
			parent = result
		} else {
			sum = processUnlimitedFirstOrLast(sum, &i, &buf[j+1], false)
			r := &listNode{
				Val: sum,
			}
			parent.Next = r
			parent = r
		}
	}
	return result
}

func processUnlimitedFirstOrLast(sum int, i *int, arr *[]int, needSizedBuff bool) int {
	if sum > 9 {
		sum = sum % 10
		if needSizedBuff {
			(*arr)[0]++
		} else {
			*arr = append(*arr, 1)
			*i++
		}
	}
	return sum
}

func calc(sum []int) int {
	var val int
	for _, v := range sum {
		val += v
	}
	return val
}
