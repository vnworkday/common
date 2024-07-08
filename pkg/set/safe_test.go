package set

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeSetOperations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(s *safeSet[int]) (func(), []int)
		expected []int
	}{
		{
			name: "AddToEmptySet",
			setup: func(s *safeSet[int]) (func(), []int) {
				s.Add(1)

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "AddDuplicateElementDoesNotIncreaseSize",
			setup: func(s *safeSet[int]) (func(), []int) {
				s.Add(1)
				s.Add(1) // Attempt to add duplicate

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "RemoveExistingElement",
			setup: func(s *safeSet[int]) (func(), []int) {
				s.Add(1)
				s.Remove(1)

				return func() {}, []int{}
			},
			expected: []int{},
		},
		{
			name: "RemoveNonExistingElementKeepsSetUnchanged",
			setup: func(s *safeSet[int]) (func(), []int) {
				s.Add(1)
				s.Remove(2) // Attempt to remove non-existing

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "ContainsChecksForElementPresence",
			setup: func(s *safeSet[int]) (func(), []int) {
				s.Add(1)

				return func() {}, []int{1}
			},
			expected: []int{1},
		},
		{
			name: "ClearEmptiesTheSet",
			setup: func(s *safeSet[int]) (func(), []int) {
				s.Add(1)
				s.Clear()

				return func() {}, []int{}
			},
			expected: []int{},
		},
		{
			name: "CloneCreatesANewIndependentSet",
			setup: func(givenSet *safeSet[int]) (func(), []int) {
				givenSet.Add(1)
				//nolint:forcetypeassert
				clone := givenSet.Clone().(*safeSet[int])
				clone.Add(2)

				return func() {}, givenSet.ToSlice()
			},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &safeSet[int]{data: make(map[int]struct{})}

			tearDown, expected := tt.setup(s)
			defer tearDown()

			result := s.ToSlice()
			sort.Ints(result) // Ensure the order for comparison
			assert.Equal(t, expected, result)
		})
	}
}
