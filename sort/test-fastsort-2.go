package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func sort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	var lt, gt, i = lo, hi, lo + 1
	v := a[lo]
	for i <= gt {
		if a[i] < v { // [lo, lt) < v
			a[i], a[lt] = a[lt], a[i]
			lt++
			i++
		} else if a[i] > v { // (gt, hi] > v
			a[gt], a[i] = a[i], a[gt]
			gt--
		} else {
			i++
		}
	}
	sort(a, lo, lt-1)
	sort(a, gt+1, hi)
}

func fastSort(a []int) {
	sort(a, 0, len(a)-1)
}

func main() {
	a := []int{3, 3, 3, 3, 3, 7, 1, 2, 5}
	fastSort(a)
	debug("a: %v\n", a)
}
