package main

import "fmt"

type Stack struct {
	stack []int
}

func (S *Stack) Empty() bool {
	return len(S.stack) == 0
}

func (S *Stack) Size() int {
	return len(S.stack)
}

func (S *Stack) Top() (int, error) {
	if len(S.stack) > 0 {
		return S.stack[len(S.stack)-1], nil
	} else {
		return -1, fmt.Errorf("Stack is empty")
	}
}

func (S *Stack) Push(data int) {
	S.stack = append(S.stack, data)
}

func (S *Stack) Pop() (int, error) {
	size := len(S.stack)
	if size > 0 {
		data := S.stack[size-1]
		S.stack = S.stack[:size-1]
		return data, nil
	} else {
		return -1, fmt.Errorf("Stack is empty can't pop")
	}
}

func main() {
	S := new(Stack) //create stack

	// add element to stack
	S.Push(10)
	S.Push(20)
	S.Push(30)

	//print size of stack
	fmt.Println("Size of the stack: ", S.Size())

	//check stack empty or not
	fmt.Println("Is stack empty: ", S.Empty())

	// print top of the stack
	top, _ := S.Top()
	fmt.Println("Top of the stack: ", top)

	// remove element from stack
	top, _ = S.Pop()
	fmt.Println(top)
	top, _ = S.Pop()
	fmt.Println(top)
	top, _ = S.Pop()
	fmt.Println(top)
	top, ok := S.Pop()
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(top)
	}
}
