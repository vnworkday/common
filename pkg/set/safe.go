package set

import "sync"

type safeSet[T comparable] struct {
	data map[T]struct{}
	lock sync.Mutex
}

func (s *safeSet[T]) Add(val T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.data[val]; ok {
		return false
	}

	s.data[val] = struct{}{}

	return true
}

func (s *safeSet[T]) Remove(val T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.data[val]; !ok {
		return false
	}

	delete(s.data, val)

	return true
}

func (s *safeSet[T]) Contains(val T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[val]

	return ok
}

func (s *safeSet[T]) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.data)
}

func (s *safeSet[T]) ToSlice() []T {
	s.lock.Lock()
	defer s.lock.Unlock()

	result := make([]T, 0, len(s.data))
	for k := range s.data {
		result = append(result, k)
	}

	return result
}

func (s *safeSet[T]) Clone() Set[T] {
	s.lock.Lock()
	defer s.lock.Unlock()

	result := make(map[T]struct{})

	for k := range s.data {
		result[k] = struct{}{}
	}

	return &safeSet[T]{
		data: result,
	}
}

func (s *safeSet[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for k := range s.data {
		delete(s.data, k)
	}
}
