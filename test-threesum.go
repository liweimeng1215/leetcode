// lc 15. 三数之和
// 排序+双指针方法
package main

import "sort"

type MySort []int

func (ms MySort) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }
func (ms MySort) Len() int           { return len(ms) }
func (ms MySort) Less(i, j int) bool { return ms[i] < ms[j] }

func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}

	sort.Sort(MySort(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ans
		}
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] { // 去重，该数值所有可能的结果已记录
			i++
		}
		lo := i + 1
		hi := len(nums) - 1
		for lo < hi {
			if nums[i]+nums[lo]+nums[hi] == 0 {
				ans = append(ans, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--
				for lo < len(nums) && nums[lo-1] == nums[lo] { // 去重，找到下一个改变结果的位置
					lo++
				}
				for hi >= 0 && nums[hi] == nums[hi+1] {
					hi--
				}
			} else if nums[i]+nums[lo]+nums[hi] < 0 {
				lo++
			} else if nums[i]+nums[lo]+nums[hi] > 0 {
				hi--
			}
		}
	}
	return ans
}
