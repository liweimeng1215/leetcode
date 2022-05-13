/**
1.给一个数组和目标值，求数组中和为目标值的所有组合，数组中每个值可以重复使用。

比如：数组【2，3，6，7】，目标值7

所有组合：【2，2，3】，【7】
**/
package main

import "fmt"

type Solution struct {
	Path []int
	Ans  [][]int
}

func (s Solution) pathSum() int {
	sum := 0
	for i := range s.Path {
		sum += s.Path[i]
	}
	return sum
}

func (s *Solution) Dfs(candidate []int, target, i int) { // i作为起始搜索点，避免重复
	if s.pathSum() == target {
		path := make([]int, len(s.Path))
		copy(path, s.Path)
		s.Ans = append(s.Ans, path)
		return
	} else if s.pathSum() > target {
		return
	}

	for i := i; i < len(candidate); i++ {
		s.Path = append(s.Path, candidate[i])
		s.Dfs(candidate, target, i)
		s.Path = s.Path[0 : len(s.Path)-1]
	}
}

func main() {
	candidate := []int{2, 3, 6, 7}
	s := Solution{Path: []int{}, Ans: [][]int{}}
	s.Dfs(candidate, 7, 0)
	fmt.Printf("ans: %v", s.Ans)
}
