package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func search1(a []int, target int) int {
	lo, hi := 0, len(a) // [lo, hi)
	for lo < hi {
		mi := lo + (hi-lo)/2 // overflow protection
		//debug("lo: %v, hi: %v, mi: %v\n", lo, hi, mi)
		if a[mi] < target {
			lo = mi + 1 // (..., lo) < target
		} else if target < a[mi] {
			hi = mi // [hi, ...) > target
		} else {
			hi = mi // [hi, ...) >= target
		}
	}
	if lo < len(a) && a[lo] == target {
		return lo
	}

	return -1
}

func search2(a []int, target int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mi := lo + (hi-lo)/2
		//debug("lo: %v, hi: %v, mi: %v\n", lo, hi, mi)
		if a[mi] < target {
			lo = mi + 1 // (..., lo) < target
		} else if target < a[mi] {
			hi = mi // [hi, ...) > target
		} else {
			lo = mi + 1 // (..., lo) <= target
		}
	}
	if 0 <= lo-1 && a[lo-1] == target {
		return lo - 1
	}

	return -1
}

func main() {
	a := []int{1, 2, 3, 3, 4, 4}
	ret1 := search1(a, 4)
	debug("ret1: %v\n", ret1)
	ret2 := search2(a, 4)
	debug("ret2: %v\n", ret2)

	ret3 := search1(a, 5)
	debug("ret3: %v\n", ret3)
	ret4 := search1(a, 0)
	debug("ret4: %v\n", ret4)

	ret5 := search2(a, 5)
	debug("ret5: %v\n", ret5)
	ret6 := search2(a, 0)
	debug("ret6: %v\n", ret6)
}
