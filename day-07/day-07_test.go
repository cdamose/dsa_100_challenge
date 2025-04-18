package day07

import "testing"

func TestRearrangePositiveAndNegative(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "All positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: []int{-1, -2, -3, -4, -5},
		},
		{
			name:     "Mixed positive and negative numbers",
			input:    []int{1, -2, 3, -4, 5},
			expected: []int{1, -2, 3, -4, 5},
		},

		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single positive number",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Single negative number",
			input:    []int{-1},
			expected: []int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RearrangePositiveAndNegative(tt.input)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRearrangePositiveAndNegativeUsingTwoPointers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "All positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: []int{-1, -2, -3, -4, -5},
		},
		{
			name:     "Mixed positive and negative numbers",
			input:    []int{1, -2, 3, -4, 5},
			expected: []int{-4, -2, 3, 1, 5},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single positive number",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Single negative number",
			input:    []int{-1},
			expected: []int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RearrangePositiveAndNegativeUsingTwoPointers(tt.input)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDutchNationalFlagProblem(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "All zeros",
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "All ones",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			name:     "All twos",
			input:    []int{2, 2, 2, 2, 2},
			expected: []int{2, 2, 2, 2, 2},
		},
		{
			name:     "Mixed zeros, ones, and twos",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "Single one",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Single two",
			input:    []int{2},
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DutchNationalFlagProblem(tt.input)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDutchNationalFlagProblemUsingHashing(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "All zeros",
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "All ones",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			name:     "All twos",
			input:    []int{2, 2, 2, 2, 2},
			expected: []int{2, 2, 2, 2, 2},
		},
		{
			name:     "Mixed zeros, ones, and twos",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "Single one",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Single two",
			input:    []int{2},
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DutchNationalFlagProblemUsingHashing(tt.input)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
