package main

import (
	"fmt"
	"sort"
)

type Lesser interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		string | float32 | float64
}

// T - type parameter, comparable - constraint

type Set[T Lesser] map[T]struct{}

//type Set[T comparable] map[T]struct{}

func NewSet[T Lesser]() Set[T] {
	return map[T]struct{}{}
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) AddValues(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) ToSlice() []T {
	res := make([]T, 0, len(s))
	for v := range s {
		res = append(res, v)
	}
	return res
}

func (s Set[T]) ToSortedSlice() []T {
	arr := s.ToSlice()
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func NewFromValues[T Lesser](values ...T) Set[T] {
	//func NewFromValues[T comparable](values ...T) Set[T] {
	res := make(map[T]struct{}, len(values))
	for _, v := range values {
		res[v] = struct{}{}
	}
	return res
}

func main() {
	s := NewSet[int]()
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Remove(4)
	fmt.Println(s.ToSlice())

	setString := NewSet[string]()
	setString.Add("a")
	setString.Add("b")
	setString.Add("c")
	fmt.Println(setString.ToSlice())
	setString.AddValues("f", "g")
	fmt.Println(setString.ToSlice())

	in := NewFromValues(1, 2, 3, 4, 167).ToSortedSlice()
	fmt.Println(in)
}

//func Printf(template string, values ...any) (string, error) {
//	return "", nil
//}
