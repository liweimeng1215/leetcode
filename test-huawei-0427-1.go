package main

import "fmt"

func searchLow(a int, b []int) int {
	lo := 0
	hi := len(b)
	for lo < hi {
		mi := (lo + hi) / 2
		if b[mi] > a {

		} else if b[mi] < a {

		} else if b[mi] == a {

		}
	}
}

func searchTop(a int, b []int) int {
	lo := 0
	hi := len(b)
	for lo < hi {
		mi := (lo + hi) / 2
		if b[mi] > a {
			hi = mi
		} else if b[mi] < a {
			lo = mi + 1
		} else if b[mi] == a {
			return mi
		}
	}
	return lo
}

func solution(a []int, b []int, r int) {
	for _, c := range a {
		p := searchTop(c, b)
		if p >= len(b) {
			break
		}
		if b[p]-c <= r {
			fmt.Printf("%d %d\n", c, b[p])
		}
		b = b[p:]
	}
}

func main() {
	var m, n, R int
	fmt.Scanf("%d %d %d\n", &m, &n, &R)
	var a, b []int
	for i := 0; i < m; i++ {
		x := 0
		fmt.Scan(&x)
		a = append(a, x)
	}
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		b = append(b, x)
	}
	solution(a, b, R)
}
