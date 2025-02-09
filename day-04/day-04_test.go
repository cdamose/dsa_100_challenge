package day04

import "testing"

func TestFindOnceApperanceUisngHash(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 4},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 5},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 2},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 3},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 4},
	}

	for _, test := range tests {
		result := FindOnceApperanceUisngHash(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}

func TestFindOnceApperanceUisngXOR(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 4},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 5},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 2},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 3},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 4},
	}

	for _, test := range tests {
		result := FindOnceApperanceUisngXOR(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}

func TestFindMajorityElementNaive(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 0},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 0},
	}

	for _, test := range tests {
		result := FindMajorityElementNaive(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}

func TestFindMajorityElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 0},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 0},
	}

	for _, test := range tests {
		result := FindMajorityElement(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestFindMajorityElementUsingSorting(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 0},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 0},
	}

	for _, test := range tests {
		result := FindMajorityElementUsingSorting(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestFindMajorityElementUsingMooreVoting(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 0},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 0},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 0},
	}

	for _, test := range tests {
		result := FindMajorityElementUsingMooreVoting(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

// Helper function to compare two slices for equality
func equalSlices(a, b []int) bool {

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestFindFirstandLastOccurence(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 2, []int{0, 1}},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 2, []int{0, 1}},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 3, []int{1, 2}},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 4, []int{3, 4}},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 5, []int{5, 6}},
	}

	for _, test := range tests {
		first, last := FindFirstandLastOccurence(test.input, test.target)
		result := []int{first, last}
		if !equalSlices(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}

func TestFindFirstandLastOccurenceUsingBinarySearch(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 2, []int{0, 1}},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 2, []int{0, 1}},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 3, []int{1, 2}},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 4, []int{3, 4}},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 5, []int{5, 6}},
	}

	for _, test := range tests {
		first, last := FindFirstandLastOccurenceUsingBinarySearch(test.input, test.target)
		result := []int{first, last}
		if !equalSlices(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}
