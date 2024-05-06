package rmap

import (
	"fmt"
)

type IntegerSet[T comparable] struct {
	m map[T]struct{}
}

func NewIntegerSet[T comparable]() *IntegerSet[T] {
	return &IntegerSet[T]{m: make(map[T]struct{})}
}

func (s *IntegerSet[T]) Add(item T) {
	s.m[item] = struct{}{}
}

func (s *IntegerSet[T]) AddList(items []T) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *IntegerSet[T]) Size() int {
	return len(s.m)
}

func (s *IntegerSet[T]) List() []T {
	v := make([]T, 0, s.Size())
	for item := range s.m {
		v = append(v, item)
	}
	return v
}

func (s *IntegerSet[T]) In(k T) bool {
	_, ok := s.m[k]
	return ok
}

func (s *IntegerSet[T]) ListIn(ks []T) bool {
	for _, val := range ks {
		_, ok := s.m[val]
		if !ok {
			return false
		}
	}
	return true
}

func (s *IntegerSet[T]) Delete(k T) {
	delete(s.m, k)
}

func (s *IntegerSet[T]) DeleteList(ks []T) {
	for _, k := range ks {
		delete(s.m, k)
	}
}

func (s *IntegerSet[T]) Merge2String(split string) (str string) {
	for item := range s.m {
		tmp := fmt.Sprintf("%v", item)
		if str == "" {
			str = tmp
		} else {
			str = str + split + tmp
		}
	}
	return
}
