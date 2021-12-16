package main

import "fmt"

type Queue struct {
	queue []int
}

func (Q *Queue) Size() int {
	return len(Q.queue)
}

func (Q *Queue) Empty() bool {
	return len(Q.queue) == 0
}

func (Q *Queue) Push(data int) {
	Q.queue = append(Q.queue, data)
}

func (Q *Queue) Pop() (int, error) {
	size := len(Q.queue)
	if size > 0 {
		data := Q.queue[0]
		Q.queue = Q.queue[1:]
		return data, nil
	} else {
		return -1, fmt.Errorf("Queue is empty can't pop")
	}
}

func (Q *Queue) Front() (int, error) {
	if len(Q.queue) > 0 {
		return Q.queue[0], nil
	} else {
		return -1, fmt.Errorf("Queue is Empty")
	}
}

func (Q *Queue) Rear() (int, error) {
	if len(Q.queue) > 0 {
		return Q.queue[len(Q.queue)-1], nil
	} else {
		return -1, fmt.Errorf("Queue is Empty")
	}
}

func main() {
	Q := new(Queue) //create Queue

	// add element to Queue
	Q.Push(10)
	Q.Push(20)
	Q.Push(30)

	//print size of Queue
	fmt.Println("Size of the Queue: ", Q.Size())

	//check Queue empty or not
	fmt.Println("Is Queue empty: ", Q.Empty())

	// print Front of the Queue
	front, _ := Q.Front()
	fmt.Println("Front of the Queue: ", front)

	// remove element from Queue
	front, _ = Q.Pop()
	fmt.Println(front)
	front, _ = Q.Pop()
	fmt.Println(front)
	front, _ = Q.Pop()
	fmt.Println(front)
	front, ok := Q.Pop()
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(front)
	}

}
