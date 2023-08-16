// Copyright (c) 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack[int]()
	// should not error
	s.Push(42)
	value, err := s.Peek()
	if err != nil {
		t.Errorf("[In TestPush()] s.Push(42), s.Peek() - expected nil error, but got %s\n", err)
	}
	if value != 42 {
		t.Errorf("[In TestPush()] s.Push(42), s.Peek() - expected 42, but got %d\n", value)
	}
	// generics work
	// s.Push("Hello, World!")
}

func TestPop(t *testing.T) {
	s := NewStack[int]()
	// should error
	_, err := s.Pop()
	if err == nil {
		t.Errorf("[In TestPop()] s.Pop() - expected \"trying to pop from empty stack\" error, but got nil\n")
	}
	// add value
	s.Push(42)
	// should not error
	value, err := s.Pop()
	if err != nil {
		t.Errorf("[In TestPop()] s.Pop() - expected nil error, but got %s\n", err)
	}
	if value != 42 {
		t.Errorf("[In TestPop()] s.Pop() - expected 42, but got %d\n", value)
	}
	// add some more values
	s.Push(21)
	s.Push(84)
	value, err = s.Pop()
	if err != nil {
		t.Errorf("[In TestPop()] s.Pop() - expected nil error but got %s\n", err)
	}
	if value != 84 {
		t.Errorf("[In TestPop()] s.Pop() - expected 84, but got %d\n", value)
	}
	value, err = s.Pop()
	if err != nil {
		t.Errorf("[In TestPop()] s.Pop() - expected nil error but got %s\n", err)
	}
	if value != 21 {
		t.Errorf("[In TestPop()] s.Pop() - expected 21, but got %d\n", value)
	}
}

func TestPeek(t *testing.T) {
	s := NewStack[int]()
	// should error
	_, err := s.Peek()
	if err == nil {
		t.Errorf("[In TestPeek()] s.Peek() - expected \"trying to peak into empty stack\" error, but got nil\n")
	}
	// add value
	s.Push(42)
	value, err := s.Peek()
	if err != nil {
		t.Errorf("[In TestPeek()] s.Peek() - expected nil error but got %s\n", err)
	}
	if value != 42 {
		t.Errorf("[In TestPeek()] s.Peek() - expected 42, but got %d\n", value)
	}
	// add some more values
	s.Push(21)
	s.Push(84)
	value, err = s.Peek()
	if err != nil {
		t.Errorf("[In TestPeek()] s.Peek() - expected nil error but got %s\n", err)
	}
	if value != 84 {
		t.Errorf("[In TestPeek()] s.Peek() - expected 84, but got %d\n", value)
	}
}

func TestSize(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(5)
	sz := s.Size()
	if sz != 4 {
		t.Errorf("[In TestSize()] s.Size() - expected 4, but got %d\n", sz)
	}
}

func TestDrain(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(5)
	s.Drain()
	sz := s.Size()
	if sz != 0 {
		t.Errorf("[In TestDrain()] s.Drain() - expected 0, but got %d\n", sz)
	}
}
