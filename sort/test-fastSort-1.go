package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func partition(a []int, lo, hi int) int {
	var i, j = lo, hi
	v := a[lo]
	for {
		for a[i] <= v { // 从左往右遍历 [lo,i)<=v
			i++
			if i == hi {
				break
			}
		}
		for a[j] >= v { // 从右往左遍历 (j, hi]>=v
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func sort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}

func fastSort(a []int) {
	sort(a, 0, len(a)-1)
}

func main() {
	a := []int{3, 3, 3, 3, 3, 7, 1, 2, 5}
	fastSort(a)
	debug("a: %v\n", a)
}
