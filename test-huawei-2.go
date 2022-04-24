/*
  给定一个射击比赛成绩单
  包含多个选手若干次射击的成绩分数
  请对每个选手按其最高三个分数之和进行降序排名
  输出降序排名后的选手id序列
  条件如下
    1. 一个选手可以有多个射击成绩的分数，且次序不固定
    2. 如果一个选手成绩少于3个，则认为选手的所有成绩无效，排名忽略该选手
    3. 如果选手的成绩之和相等，则相等的选手按照其id降序排列

   输入描述:
     输入第一行
         一个整数N
         表示该场比赛总共进行了N次射击
         产生N个成绩分数
         2<=N<=100

     输入第二行
       一个长度为N整数序列
       表示参与每次射击的选手id
       0<=id<=99

     输入第三行
        一个长度为N整数序列
        表示参与每次射击选手对应的成绩
        0<=成绩<=100

   输出描述:
      符合题设条件的降序排名后的选手ID序列

   示例一
      输入:
        13
        3,3,7,4,4,4,4,7,7,3,5,5,5
        53,80,68,24,39,76,66,16,100,55,53,80,55
      输出:
        5,3,7,4
      说明:
        该场射击比赛进行了13次
        参赛的选手为{3,4,5,7}
        3号选手成绩53,80,55 最高三个成绩的和为188
        4号选手成绩24,39,76,66  最高三个成绩的和为181
        5号选手成绩53,80,55  最高三个成绩的和为188
        7号选手成绩68,16,100  最高三个成绩的和为184
        比较各个选手最高3个成绩的和
        有3号=5号>7号>4号
        由于3号和5号成绩相等  且id 5>3
        所以输出5,3,7,4
*/
package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Id    int
	Score int
}

func makePlayer(id int, scores []int) Player {
	p := Player{}
	p.Id = id
	if len(scores) < 3 {
		p.Score = 0
		return p
	}

	top3 := make([]int, 3)
	for _, s := range scores {
		if top3[2] > s {
			continue
		} else {
			top3[2] = s
		}
		for i := 2; i > 0; i-- {
			if top3[i] > top3[i-1] {
				top3[i], top3[i-1] = top3[i-1], top3[i]
			}
		}
	}
	p.Score = top3[0] + top3[1] + top3[2]
	return p
}

type RankPlayer []Player

func (r RankPlayer) Swap(i int, j int) { r[i], r[j] = r[j], r[i] }
func (r RankPlayer) Len() int          { return len(r) }
func (r RankPlayer) Less(i int, j int) bool {
	if r[i].Score < r[j].Score {
		return true
	} else if r[i].Score == r[j].Score {
		return r[i].Id < r[j].Id
	} else {
		return false
	}
}
func (r RankPlayer) Print() {
	i := r.Len() - 1
	fmt.Printf("%d", r[i].Id)
	for k := i - 1; k >= 0; k-- {
		fmt.Printf(",%d", r[k].Id)
	}
	fmt.Printf("\n")
}

func main() {
	//var n int
	//var ids []int
	//var scores []int
	//fmt.Scanln(&n)
	//fmt.Scanln(&ids)
	//fmt.Scanln(&scores)
	n := 13
	ids := []int{3, 3, 7, 4, 4, 4, 4, 7, 7, 3, 5, 5, 5}
	scores := []int{53, 80, 68, 24, 39, 76, 66, 16, 100, 55, 53, 80, 55}
	m := map[int][]int{}
	for i := 0; i < n; i++ {
		id, score := ids[i], scores[i]
		m[id] = append(m[id], score)
	}
	var players []Player
	for k, v := range m {
		players = append(players, makePlayer(k, v))
	}
	sort.Sort(RankPlayer(players))
	RankPlayer(players).Print()
}
