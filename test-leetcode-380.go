package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type ScaleArray struct {
	arr       []int
	lastIndex int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (sa *ScaleArray) Get(i int) (int, bool) {
	if i < 0 || i > sa.lastIndex {
		return 0, false
	} else {
		return sa.arr[i], true
	}
}

func (sa *ScaleArray) Remove(i int) bool {
	if i > sa.lastIndex || i < 0 {
		return false
	} else if i == sa.lastIndex {
		sa.lastIndex--
		return true
	} else {
		sa.arr[i], sa.arr[sa.lastIndex] = sa.arr[sa.lastIndex], sa.arr[i]
		sa.lastIndex--
		return true
	}
}

func (sa *ScaleArray) Add(v int) int {
	if len(sa.arr)-1 > sa.lastIndex {
		sa.lastIndex++
		sa.arr[sa.lastIndex] = v
	} else if len(sa.arr)-1 < sa.lastIndex {
		log.Fatal("SA arr length is wrong! length: %d, lastIndex: %d", len(sa.arr)-1, sa.lastIndex)
	} else {
		sa.arr = append(sa.arr, v)
		sa.lastIndex++
	}
	return sa.lastIndex
}

func (sa *ScaleArray) Random() int {
	i := rand.Int() % (sa.lastIndex + 1)
	return sa.arr[i]
}

type RandomizedSet struct {
	SA  ScaleArray
	Map map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{SA: ScaleArray{lastIndex: -1, arr: make([]int, 0)}, Map: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if v, ok := this.Map[val]; ok == true && v != -1 { // has occupied
		return false
	} else {
		this.Map[val] = this.SA.Add(val)
		log.Infof("after insert %v, SA: %v, Map: %v", val, this.SA, this.Map)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.Map[val]; ok == false || v == -1 {
		return false
	} else {
		index := this.Map[val]
		this.SA.Remove(index)
		this.Map[val] = -1
		if v, ok := this.SA.Get(index); ok {
			this.Map[v] = index
		}
		log.Infof("after remove %v, SA: %v, Map: %v", val, this.SA, this.Map)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.SA.Random()
}

func main() {
	R := Constructor()
	fmt.Printf("res: %v\n", R.Insert(0))
	fmt.Printf("res: %v\n", R.Insert(1))
	fmt.Printf("res: %v\n", R.Remove(0))
	fmt.Printf("res: %v\n", R.Insert(2))
	fmt.Printf("res: %v\n", R.Remove(1))
	fmt.Printf("res: %v\n", R.GetRandom())
}
