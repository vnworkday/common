package set

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsafeSetOperations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(s unsafeSet[int]) (func(), []int)
		expected []int
	}{
		{
			name: "AddToEmptySet",
			setup: func(s unsafeSet[int]) (func(), []int) {
				s.Add(1)

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "AddDuplicateElement",
			setup: func(s unsafeSet[int]) (func(), []int) {
				s.Add(1)
				s.Add(1) // Attempt to add duplicate

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "RemoveExistingElement",
			setup: func(s unsafeSet[int]) (func(), []int) {
				s.Add(1)
				s.Remove(1)

				return func() {}, []int{}
			},
			expected: []int{},
		},
		{
			name: "RemoveNonExistingElement",
			setup: func(s unsafeSet[int]) (func(), []int) {
				s.Add(1)
				s.Remove(2) // Attempt to remove non-existing

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "ContainsExistingElement",
			setup: func(s unsafeSet[int]) (func(), []int) {
				s.Add(1)

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "ClearSet",
			setup: func(s unsafeSet[int]) (func(), []int) {
				s.Add(1)
				s.Clear()

				return func() {}, []int{}
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := make(unsafeSet[int])

			tearDown, expected := tt.setup(s)
			defer tearDown()

			result := s.ToSlice()
			sort.Ints(result) // Ensure the order for comparison
			assert.Equal(t, expected, result)
		})
	}
}
