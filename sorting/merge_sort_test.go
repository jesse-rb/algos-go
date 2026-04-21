package sorting

import (
	"cmp"
	"testing"
)

func AssertSliceOrder[E cmp.Ordered](t *testing.T, actual []E, expected []E) {
	for i, _ := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	t.Run("sort slice of ints", func(t *testing.T) {
		test := []int{23, 14, 5, 0}
		actual := MergeSort(test)
		expected := []int{0, 5, 14, 23}

		AssertSliceOrder(t, actual, expected)
	})
	t.Run("sort slice of strings", func(t *testing.T) {
		test := []string{"Y", "A", "Z", "X", "C", "B", "U"}
		actual := MergeSort(test)
		expected := []string{"A", "B", "C", "U", "X", "Y", "Z"}
		AssertSliceOrder(t, actual, expected)
	})
	t.Run("sort slice of even length", func(t *testing.T) {
		test := []int{3, 1, 1, 4}
		actual := MergeSort(test)
		expected := []int{1, 1, 3, 4}
		AssertSliceOrder(t, actual, expected)
	})
	t.Run("sort slice of odd length", func(t *testing.T) {
		test := []int{3, 1, 4}
		actual := MergeSort(test)
		expected := []int{1, 3, 4}
		AssertSliceOrder(t, actual, expected)
	})
	t.Run("sort slice of length: 1", func(t *testing.T) {
		test := []int{83}
		actual := MergeSort(test)
		expected := []int{83}
		AssertSliceOrder(t, actual, expected)
	})
	t.Run("sort slice of length: 2", func(t *testing.T) {
		test := []int{112, 83}
		actual := MergeSort(test)
		expected := []int{83, 112}
		AssertSliceOrder(t, actual, expected)
	})
	t.Run("sort empty slice", func(t *testing.T) {
		test := []int{}
		actual := MergeSort(test)
		expected := []int{}
		AssertSliceOrder(t, actual, expected)
	})
}
