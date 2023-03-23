package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{-2, 3, 15, 4}
	arr3 := []int{0}

	expected1 := 55
	expected2 := 20
	expected3 := 0

	actual1 := ConcurrentSum(arr1)
	actual2 := ConcurrentSum(arr2)
	actual3 := ConcurrentSum(arr3)

	if actual1 != expected1 {
		t.Errorf("expected %d, was %d", expected1, actual1)
	}

	if actual2 != expected2 {
		t.Errorf("expected %d, was %d", expected2, actual2)
	}

	if actual3 != expected3 {
		t.Errorf("expected %d, was %d", expected3, actual3)
	}
}
