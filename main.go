package main

import "fmt"

// Node представляет узел односвязного списка
type Node struct {
	data int
	next *Node
}

// LinkedList представляет односвязный список
type LinkedList struct {
	head *Node
}

// Добавление элемента в конец списка
func (ll *LinkedList) append(data int) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Вывод содержимого списка
func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}

	ll.append(1)
	ll.append(2)
	ll.append(3)

	ll.printList()
}
