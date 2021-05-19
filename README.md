# golinkedlist
I created my own linked list using Go. Even though Go had their own library for linked list named 'list'

I'm using doubly linear linked list because its practical and easy to use.

Suitable for learning data structure with Go Language, its totally open source you just have to clone or download it

# Features
## Headers
```Go
type Node struct {
	value      interface{}
	next, prev *Node
}

type LinkedList struct {
	first, last *Node
	Size        int
}
```
## Create List
```Go
func CreateList() LinkedList
```
This feature creating the list
There is a sub featur for emptying the list
```Go
func (*LinkedList) EmptyingList()
```
## Add
```Go 
func (*LinkedList) Add(interface{})
```
This feature adds an object to the list. It can be anything like functions, integer, etc.
There are sub feature for Add feature.
```Go
func (*LinkedList) AddFirst(*Node)
```
```Go
func (*LinkedList) AddLast(*Node)
```
## Get
```Go
func (*LinkedList) Get(int) (interface{}, error)
```
This feature return the value that user gave index in the list. It can occur error if the input parameter goes beyond the size of the list or if the list is empty
```Go
if l.Size == 0 {
		return nil, errors.New("list empty")
	}
if i >= l.Size {
		return nil, errors.New("out of bound")
}
```
## Remove
```Go

```

