// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

package stack

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	lock sync.Mutex
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, value)
}

func (s *Stack[T]) Peek() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	length := len(s.data)
	if length == 0 {
		return *new(T), errors.New("trying to peak into empty stack")
	}
	return s.data[length-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	length := len(s.data)
	if length == 0 {
		return *new(T), errors.New("trying to pop from empty stack")
	}
	value := s.data[length-1]
	s.data = s.data[:length-1]
	return value, nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Drain() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = make([]T, 0)
}
