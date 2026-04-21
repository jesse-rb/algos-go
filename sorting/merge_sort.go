package sorting

import (
	"cmp"
)

func mergeSortMerge[E cmp.Ordered](left []E, right []E, toSort []E) {
	l := 0
	r := 0
	i := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			toSort[i] = left[l]
			l++
		} else {
			toSort[i] = right[r]
			r++
		}
		i++
	}

	for l < len(left) {
		toSort[i] = left[l]
		l++
		i++
	}
	for r < len(right) {
		toSort[i] = right[r]
		r++
		i++
	}
}

// MergeSort sorts s using the merge sort algorithm O(n log n)
func MergeSort[E cmp.Ordered](s []E) []E {
	sorted := make([]E, len(s))
	if len(s) == 0 {
		return sorted
	}
	if len(s) == 1 {
		sorted[0] = s[0]
		return s
	}

	m := len(s) / 2
	left := MergeSort(s[:m])
	right := MergeSort(s[m:])
	mergeSortMerge(left, right, sorted)

	return sorted
}
