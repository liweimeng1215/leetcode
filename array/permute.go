// LC46. 全排列
// 用深度优先进行搜索，记录走过的路径
// 回溯时撤销之前走过的记录
package main

type Solution struct {
	Visited map[int]struct{}
	Ans     [][]int
	Path    []int
}

func (s *Solution) Dfs(nums []int) {
	N := len(nums)
	if N == 0 {
		return
	}

	if len(s.Path) == N {
		path := make([]int, N)
		copy(path, s.Path)
		s.Ans = append(s.Ans, path)
	}

	for i := range nums {
		if _, ok := s.Visited[nums[i]]; ok {
			continue
		} else {
			s.Visited[nums[i]] = struct{}{}
			s.Path = append(s.Path, nums[i])
			s.Dfs(nums)
			delete(s.Visited, nums[i])
			s.Path = s.Path[0 : len(s.Path)-1]
		}
	}
}

func permute(nums []int) [][]int {
	s := Solution{Visited: map[int]struct{}{}, Ans: [][]int{}, Path: []int{}}
	s.Dfs(nums)
	return s.Ans
}
