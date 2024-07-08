package set

type unsafeSet[T comparable] map[T]struct{}

func (s unsafeSet[T]) Add(v T) bool {
	if _, ok := s[v]; ok {
		return false
	}

	s[v] = struct{}{}

	return true
}

func (s unsafeSet[T]) Remove(v T) bool {
	if _, ok := s[v]; !ok {
		return false
	}

	delete(s, v)

	return true
}

func (s unsafeSet[T]) Contains(v T) bool {
	_, ok := s[v]

	return ok
}

func (s unsafeSet[T]) Len() int {
	return len(s)
}

func (s unsafeSet[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}

	return result
}

func (s unsafeSet[T]) Clone() Set[T] {
	result := make(unsafeSet[T])

	for k := range s {
		result[k] = struct{}{}
	}

	return result
}

func (s unsafeSet[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}
