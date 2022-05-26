package main

import (
	"errors"
	"fmt"
)

var d = true

func Debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

type Queue struct {
	a          []int
	start, end int
	cap        int
}

func (q Queue) String() string {
	return fmt.Sprintf("a: %v, cap: %v, [%v, %v)", q.a, q.cap, q.start, q.end)
}

func New(c int) Queue {
	t := Queue{a: make([]int, c), cap: c}
	return t
}

func (q Queue) IsEmpty() bool {
	return q.start == q.end
}

func (q Queue) IsFull() bool {
	return q.start+q.cap == q.end
}

func (q *Queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	i := q.start % q.cap
	q.start++
	return q.a[i], nil
}

func (q *Queue) PushQueue(v int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	i := (q.end) % q.cap
	q.a[i] = v
	q.end++
	return nil
}

func main() {
	q := New(3)
	Debug("q: {%v}\n", q)
	q.PushQueue(1)
	Debug("q: %v\n", q)
	q.PushQueue(2)
	q.PushQueue(3)
	q.PushQueue(4) // full
	Debug("q: %v\n", q)
	q.DeQueue()
	q.DeQueue()
	q.DeQueue() // empty
	Debug("q: %v\n", q)
	q.PushQueue(5)
	q.PushQueue(6)
	q.PushQueue(7)
	q.PushQueue(8)
	Debug("q: %v\n", q)
}
