// LC54. 螺旋矩阵
// bottom top left right 作为下上左右的边界，超出边界则无法继续遍历
package main

func spiralOrder(matrix [][]int) []int {
	N := len(matrix)
	M := len(matrix[0])
	ans := []int{}

	bottom, top := 0, N-1
	left, right := 0, M-1
	for {
		// 向右
		if left > right {
			break
		}
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[bottom][i])
		}
		bottom++
		// 向上
		if bottom > top {
			break
		}
		for i := bottom; i <= top; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		// 向左
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[top][i])
		}
		top--
		// 向下
		if top < bottom {
			break
		}
		for i := top; i >= bottom; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++

	}
	return ans
}
