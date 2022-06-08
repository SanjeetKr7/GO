package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	list []int
}

func (q *Queue) isEmpty() error {
	if len(q.list) == 0 {
		return errors.New("Queue is Empty")
	}
	return nil
}

func (q *Queue) Enqueue(x int) {
	q.list = append(q.list, x)
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() != nil {
		return -1, errors.New("Queue is Empty")
	}
	top := q.list[0]
	q.list = q.list[1:]
	return top, nil
}

func (q *Queue) Peek() {

}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Queue value ", q)

	val, err := q.Dequeue()
	if err != nil {
		fmt.Println("Queue is Empty")
	} else {
		fmt.Println("Peek Value is ", val)
	}

	fmt.Println("After dequeue the queue value is ", q.list)

}
