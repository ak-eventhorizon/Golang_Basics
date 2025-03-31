package main

import (
	"fmt"
)

// ********* Linked List (Duplex)*********
// Пример реализации структуры двустороннего связанного списка
//
// Push() -->                    <-prev next->
// 		nil <- (head) el1 <-> el2 <-> el3 <-> el4 (tail) -> nil
// Pop() <--

// Элемент связаного списка
type Element struct {
	prev  *Element
	next  *Element
	value any
}

// Связанный список
type LinkedList struct {
	head *Element
	tail *Element
}

// Добавления элемента
func (ll *LinkedList) Push(value any) {
	newElem := &Element{
		next:  ll.head,
		value: value,
	}

	if ll.head != nil {
		ll.head.prev = newElem
	}
	ll.head = newElem

	el := ll.head
	for el.next != nil {
		el = el.next
	}
	ll.tail = el
}

// Извлечение элемента
func (ll *LinkedList) Pop() (value any) {

	if ll.head != nil {
		value = ll.head.value
		ll.head = ll.head.next
	} else {
		value = nil
	}

	return value
}

// Вывод связанного списка
func (ll *LinkedList) Display() {
	elem := ll.head

	for elem != nil {
		if elem == ll.head && elem == ll.tail {
			fmt.Printf("<-(head) %+v (tail)->", elem.value)
		} else if elem == ll.head {
			fmt.Printf("<-(head) %+v <-> ", elem.value)
		} else if elem == ll.tail {
			fmt.Printf("%+v (tail)->", elem.value)
		} else {
			fmt.Printf("%+v <-> ", elem.value)
		}

		elem = elem.next
	}

	fmt.Println()
}

// Разворачивание списка
func (ll *LinkedList) Reverse() {
	currentElem := ll.head
	var nextElem *Element
	var previousElem *Element

	ll.tail = ll.head

	for currentElem != nil {
		nextElem, currentElem.next = currentElem.next, previousElem
		previousElem, currentElem = currentElem, nextElem
	}
	ll.head = previousElem
}

func main() {

	myList := new(LinkedList)
	myList.Push("1")
	myList.Push("2")
	myList.Push("3")
	myList.Push("4")
	myList.Push("5")
	myList.Push("6")
	myList.Display()

	fmt.Println(myList.Pop())
	myList.Display()

	myList.Reverse()
	myList.Display()

}
