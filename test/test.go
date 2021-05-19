package main

import (
	"fmt"

	"github.com/kamil5b/golinkedlist"
)

func main() {
	l := golinkedlist.CreateList()

	l.Add(5)
	l.Add(2)
	l.Add(17)
	l.Add(12)

	for i := 0; i < l.Size; i++ {
		val, _ := l.Get(i)
		fmt.Println(val)
	}
	fmt.Println()

	l.Sort(true)
	for i := 0; i < l.Size; i++ {
		val, _ := l.Get(i)
		fmt.Println(val)
	}
	fmt.Println()

	l.Remove(2)

	for i := 0; i < l.Size; i++ {
		val, _ := l.Get(i)
		fmt.Println(val)
	}

}
