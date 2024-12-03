package main

import "testing"

// Empty lists return 0 distance
func TestEmptyLists(t *testing.T) {

	// Note that this is an initialized but empty slice; it's distinct from
	// something like []int, which is declared but not initialized ("a nil slice")
	list1 := []int{}
	list2 := []int{}

	result := CalculateDistance(list1, list2)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}


}

// Nil lists resturn 0 distance
func TestNilLists(t *testing.T) {
	var list1 []int
	var list2 []int

	result := CalculateDistance(list1, list2)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// Single element lists with same value
func TestSingleItemSameValue(t *testing.T) {
	list1 := []int{1}
	list2 := []int{1}

	result := CalculateDistance(list1, list2)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// Single element lists with different values
func TestSingleItemDifferentValue(t *testing.T) {
	list1 := []int{1}
	list2 := []int{5}

	result := CalculateDistance(list1, list2)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

// Multi element lists with different values
func TestMutlipleItemDifferentValue(t *testing.T) {
	list1 := []int{1,2,3}
	list2 := []int{5,6,7}

	result := CalculateDistance(list1, list2)
	if result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}
}


// Multi element lists with unsorted values that sum to 0
func TestMutlipleItemUnsorted1(t *testing.T) {
	list1 := []int{1,5,9}
	list2 := []int{1,9,5}

	result := CalculateDistance(list1, list2)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// Multi element lists with unsorted values that sum to nonzero
func TestMutlipleItemUnsorted2(t *testing.T) {
	list1 := []int{1,5,9,2}
	list2 := []int{1,9,5,4}

	result := CalculateDistance(list1, list2)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

// Negative values may be used in the lists
func TestNegativeValue(t *testing.T) {
	list1 := []int{-1,5,9}
	list2 := []int{1,5,9}

	result := CalculateDistance(list1, list2)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

