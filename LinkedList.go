package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (LL *LinkedList) Size() int {
	return LL.size
}

func (LL *LinkedList) InsertFront(data int) {
	newNode := new(Node)
	newNode.data = data
	newNode.next = LL.head
	LL.head = newNode
	LL.size++
}

func (LL *LinkedList) InsertBack(data int) {
	newNode := new(Node)
	newNode.data = data
	newNode.next = nil
	if LL.head != nil {
		temp := LL.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	} else {
		LL.head = newNode
	}
	LL.size++
}

func (LL *LinkedList) DeleteFront() (int, error) {
	if LL.size == 0 {
		return -1, fmt.Errorf("Linked List is Empty!!")
	} else {
		data := LL.head.data
		LL.head = LL.head.next
		LL.size--
		return data, nil
	}
}

func (LL *LinkedList) DeleteBack() (int, error) {
	if LL.size == 0 {
		return -1, fmt.Errorf("Linked List is Empty!!")
	} else {
		var data int
		if LL.size == 1 {
			data = LL.head.data
			LL.head = nil
		} else {
			temp := LL.head
			for temp.next.next != nil {
				temp = temp.next
			}
			data = temp.next.data
			temp.next = nil
		}
		LL.size--
		return data, nil
	}
}

func (LL *LinkedList) Print() {
	temp := LL.head
	for temp != nil {
		fmt.Print(temp.data)
		if temp.next != nil {
			fmt.Print(" -> ")
		}
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	LL := new(LinkedList) //Create Linked List

	// Add elements to the end of Linked List

	LL.InsertBack(10)
	LL.InsertBack(20)
	LL.InsertBack(30)
	LL.InsertBack(40)

	// Print size of the Linked List
	fmt.Println("Size of Linked List: ", LL.size)

	// Print Linked List
	fmt.Print("Linked List : ")
	LL.Print()

	//Delete element from back
	data, _ := LL.DeleteBack()
	fmt.Println("Delete element: ", data)

	//Delete element from front
	data, _ = LL.DeleteFront()
	fmt.Println("Delete element: ", data)

	fmt.Print("Linked List : ")
	LL.Print()

	// Insert element at start
	LL.InsertFront(10)

	fmt.Print("Linked List : ")
	LL.Print()
}
