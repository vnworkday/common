package set

type Set[T comparable] interface {
	// Add adds a new element to the set. Returns true if the element was added, false otherwise.
	Add(v T) bool
	// Remove removes an element from the set. Returns true if the element was removed, false otherwise.
	Remove(v T) bool
	// Contains checks if an element is in the set.
	Contains(v T) bool
	// Len returns the number of elements in the set.
	Len() int
	// ToSlice returns a slice of all elements in the set.
	ToSlice() []T
	// Clone returns a new set that is a copy of the original set.
	Clone() Set[T]
	// Clear removes all elements from the set.
	Clear()
}

// New creates a new set.
//
// Parameters:
//   - isThreadSafe: indicates whether the set should be thread-safe.
func New[T comparable](isThreadSafe ...bool) Set[T] {
	var threadSafe bool

	if len(isThreadSafe) > 0 {
		threadSafe = isThreadSafe[0]
	}

	if threadSafe {
		return &safeSet[T]{
			data: make(map[T]struct{}),
		}
	}

	return make(unsafeSet[T])
}

// Merge merges two sets into a new set.
//
// Parameters:
//   - first: the first set.
//   - second: the second set.
//   - threadSafe: indicates whether the new set should be thread-safe.
func Merge[T comparable](first, second Set[T], threadSafe ...bool) Set[T] {
	result := New[T](threadSafe...)

	for _, v := range first.ToSlice() {
		result.Add(v)
	}

	for _, v := range second.ToSlice() {
		result.Add(v)
	}

	return result
}

// Intersect returns a new set that contains the intersection of two sets.
//
// Parameters:
//   - first: the first set.
//   - second: the second set.
//   - threadSafe: indicates whether the new set should be thread-safe.
func Intersect[T comparable](first, second Set[T], threadSafe ...bool) Set[T] {
	result := New[T](threadSafe...)

	for _, v := range first.ToSlice() {
		if second.Contains(v) {
			result.Add(v)
		}
	}

	return result
}

// Difference returns a new set that contains the difference of two sets.
//
// Parameters:
//   - first: the first set.
//   - second: the second set.
//   - threadSafe: indicates whether the new set should be thread-safe.
func Difference[T comparable](first, second Set[T], threadSafe ...bool) (Set[T], Set[T]) {
	resultA := New[T](threadSafe...)
	resultB := New[T](threadSafe...)

	for _, v := range first.ToSlice() {
		if !second.Contains(v) {
			resultA.Add(v)
		}
	}

	for _, v := range second.ToSlice() {
		if !first.Contains(v) {
			resultB.Add(v)
		}
	}

	return resultA, resultB
}

// IsSubset returns true if first is a subset of second.
//
// Parameters:
//   - first: the first set.
//   - second: the second set
func IsSubset[T comparable](first, second Set[T]) bool {
	for _, v := range first.ToSlice() {
		if !second.Contains(v) {
			return false
		}
	}

	return true
}

// IsSuperset returns true if first is a superset of second.
//
// Parameters:
//   - first: the first set.
//   - second: the second set
func IsSuperset[T comparable](first, second Set[T]) bool {
	return IsSubset(second, first)
}
