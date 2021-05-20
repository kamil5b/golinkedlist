package golinkedlist

import (
	"errors"
	"reflect"
)

func lessThan(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false, errors.New("different data type")
	}
	switch a.(type) {
	case int8:
		return a.(int8) < b.(int8), nil
	case int16:
		return a.(int16) < b.(int16), nil
	case int32:
		return a.(int32) < b.(int32), nil
	case int64:
		return a.(int64) < b.(int64), nil
	case uint8:
		return a.(uint8) < b.(uint8), nil
	case uint16:
		return a.(uint16) < b.(uint16), nil
	case uint32:
		return a.(uint32) < b.(uint32), nil
	case uint64:
		return a.(uint64) < b.(uint64), nil
	case int:
		return a.(int) < b.(int), nil
	case uint:
		return a.(uint) < b.(uint), nil
	case uintptr:
		return a.(uintptr) < b.(uintptr), nil
	case string:
		return a.(string) < b.(string), nil
	case float32:
		return a.(float32) < b.(float32), nil
	case float64:
		return a.(float64) < b.(float64), nil
	default:
		return false, errors.New("uncomparable data types or interface")
	}
}

func moreThan(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false, errors.New("different data type")
	}
	switch a.(type) {
	case int8:
		return a.(int8) > b.(int8), nil
	case int16:
		return a.(int16) > b.(int16), nil
	case int32:
		return a.(int32) > b.(int32), nil
	case int64:
		return a.(int64) > b.(int64), nil
	case uint8:
		return a.(uint8) > b.(uint8), nil
	case uint16:
		return a.(uint16) > b.(uint16), nil
	case uint32:
		return a.(uint32) > b.(uint32), nil
	case uint64:
		return a.(uint64) > b.(uint64), nil
	case int:
		return a.(int) > b.(int), nil
	case uint:
		return a.(uint) > b.(uint), nil
	case uintptr:
		return a.(uintptr) > b.(uintptr), nil
	case string:
		return a.(string) > b.(string), nil
	case float32:
		return a.(float32) > b.(float32), nil
	case float64:
		return a.(float64) > b.(float64), nil
	default:
		return false, errors.New("uncomparable data types or interface")
	}
}

func (l *DoublyLinkedList) sortDescending() error {

	for current := l.first; current.next != nil; current = current.next {
		for index := current.next; index != nil; index = index.next {
			if ok, err := lessThan(current.value, index.value); ok {
				temp := current.value
				current.value = index.value
				index.value = temp
			} else if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *DoublyLinkedList) sortAscending() error {

	for current := l.first; current.next != nil; current = current.next {
		for index := current.next; index != nil; index = index.next {
			if ok, err := moreThan(current.value, index.value); ok {
				temp := current.value
				current.value = index.value
				index.value = temp
			} else if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *SinglyLinkedList) sortDescending() error {

	for current := l.first; current.next != nil; current = current.next {
		for index := current.next; index != nil; index = index.next {
			if ok, err := lessThan(current.value, index.value); ok {
				temp := current.value
				current.value = index.value
				index.value = temp
			} else if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *SinglyLinkedList) sortAscending() error {

	for current := l.first; current.next != nil; current = current.next {
		for index := current.next; index != nil; index = index.next {
			if ok, err := moreThan(current.value, index.value); ok {
				temp := current.value
				current.value = index.value
				index.value = temp
			} else if err != nil {
				return err
			}
		}
	}

	return nil
}
