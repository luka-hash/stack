# stack

A thread-safe generic stack implementation for Go (version 1.18+). It is
designed for simplicity.

## Performance

I didn't do any real performance testing, but in my experience, it is fast
enough for most day-to-day cases. One obvious way to speed it up is to allocate
memory for data in chunks (or do some other optimization around underlying
memory for data), but I doubt there is a way to do that in less than 60 lines
of code.

## API

- types:
  - Stack[T any]
- functions:
  - func NewStack[T any]() *Stack[T]
- methods:
  - func (s *Stack[T]) Push(value T)
  - func (s *Stack[T]) Peek() (T, error)
  - func (s *Stack[T]) Pop() (T, error)
  - func (s *Stack[T]) Size() int
  - func (s *Stack[T]) Drain()
 
## Is it any good?

Yes.

## Licence

This code is licensed under MIT licence (see LICENCE for details).
