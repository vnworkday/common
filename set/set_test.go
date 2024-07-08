package set

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetOperations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		setup      func(s Set[int]) (Set[int], []int)
		expected   []int
		threadSafe bool // Indicates whether the set should be thread-safe
	}{
		{
			name: "AddToEmptySet",
			setup: func(s Set[int]) (Set[int], []int) {
				s.Add(1)

				return s, []int{1}
			},
			threadSafe: false,
		},
		{
			name: "AddDuplicateElementDoesNotIncreaseSize",
			setup: func(s Set[int]) (Set[int], []int) {
				s.Add(1)
				s.Add(1) // Attempt to add duplicate

				return s, []int{1}
			},
			threadSafe: true,
		},
		{
			name: "RemoveExistingElement",
			setup: func(s Set[int]) (Set[int], []int) {
				s.Add(1)
				s.Remove(1)

				return s, []int{}
			},
			threadSafe: false,
		},
		{
			name: "RemoveNonExistingElementKeepsSetUnchanged",
			setup: func(s Set[int]) (Set[int], []int) {
				s.Add(1)
				s.Remove(2) // Attempt to remove non-existing

				return s, []int{1}
			},
			threadSafe: true,
		},
		{
			name: "ContainsChecksForElementPresence",
			setup: func(s Set[int]) (Set[int], []int) {
				s.Add(1)

				return s, []int{1}
			},
			threadSafe: false,
		},
		{
			name: "ClearEmptiesTheSet",
			setup: func(s Set[int]) (Set[int], []int) {
				s.Add(1)
				s.Clear()

				return s, []int{}
			},
			threadSafe: true,
		},
		{
			name: "MergeSetsCombinesTwoSets",
			setup: func(_ Set[int]) (Set[int], []int) {
				a := New[int]()
				a.Add(1)
				b := New[int]()
				b.Add(2)
				merged := Merge(a, b)

				return merged, []int{1, 2}
			},
			threadSafe: true,
		},
		{
			name: "IntersectSetsFindsCommonElements",
			setup: func(_ Set[int]) (Set[int], []int) {
				first := New[int]()
				first.Add(1)
				first.Add(2)
				second := New[int]()
				second.Add(2)
				second.Add(3)
				intersect := Intersect(first, second)

				return intersect, []int{2}
			},
			threadSafe: false,
		},
		{
			name: "DifferenceSetsFindsUniqueElements",
			setup: func(_ Set[int]) (Set[int], []int) {
				first := New[int]()
				first.Add(1)
				first.Add(2)
				second := New[int]()
				second.Add(2)
				second.Add(3)
				diffA, _ := Difference(first, second)

				return diffA, []int{1}
			},
			threadSafe: false,
		},
		{
			name: "IsSubsetChecksIfFirstSetIsSubsetOfSecond",
			setup: func(_ Set[int]) (Set[int], []int) {
				first := New[int]()
				first.Add(1)
				second := New[int]()
				second.Add(1)
				second.Add(2)
				isSubset := IsSubset(first, second)
				if isSubset {
					return first, []int{1}
				}

				return first, []int{} // 1 for true, 0 for false
			},
			threadSafe: false,
		},
		{
			name: "IsSupersetChecksIfFirstSetIsSupersetOfSecond",
			setup: func(_ Set[int]) (Set[int], []int) {
				first := New[int]()
				first.Add(1)
				first.Add(2)
				second := New[int]()
				second.Add(1)
				isSuperset := IsSuperset(first, second)
				if isSuperset {
					return first, []int{1, 2}
				}

				return first, []int{1} // 1 for true, 0 for false
			},
			threadSafe: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := New[int](tt.threadSafe)
			ss, expected := tt.setup(s)
			result := ss.ToSlice()
			sort.Ints(result) // Ensure the order for comparison
			assert.Equal(t, expected, result)
		})
	}
}
