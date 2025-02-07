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
