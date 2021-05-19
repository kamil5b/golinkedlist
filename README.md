# golinkedlist
I created my own linked list using Go. Even though Go had their own library for linked list named 'list'

I'm using doubly linear linked list because its practical and easy to use.

# Features
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
