package day01

// FindMaximumAndMinimum finds the maximum and minimum elements in an array.

// Assume arr = [ 3, 1, 5, 0, 4 ]

// Initial:
//    min = arr[0] = 3
//    max = arr[0] = 3

// --------------------------------------------------
// Index:   0    1    2    3    4
// Array:  [ 3,   1,   5,   0,   4 ]
//         -------------------------
// Step 0:
//    min = 3, max = 3

// --------------------------------------------------
// Index 0: Element = 3
//    Compare 3 with min (3) → No change
//    Compare 3 with max (3) → No change
//    State: min = 3, max = 3

// --------------------------------------------------
// Index 1: Element = 1
//    Compare 1 with min (3) → 1 < 3, update min = 1
//    Compare 1 with max (3) → 1 is not > 3, so max remains 3
//    State: min = 1, max = 3

// --------------------------------------------------
// Index 2: Element = 5
//    Compare 5 with min (1) → 5 is not < 1, so min remains 1
//    Compare 5 with max (3) → 5 > 3, update max = 5
//    State: min = 1, max = 5

// --------------------------------------------------
// Index 3: Element = 0
//    Compare 0 with min (1) → 0 < 1, update min = 0
//    Compare 0 with max (5) → 0 is not > 5, so max remains 5
//    State: min = 0, max = 5

// --------------------------------------------------
// Index 4: Element = 4
//    Compare 4 with min (0) → 4 is not < 0, so min remains 0
//    Compare 4 with max (5) → 4 is not > 5, so max remains 5
//    State: min = 0, max = 5

// --------------------------------------------------
// Final Result:
//    min = 0
//    max = 5
// --------------------------------------------------

func FindMaximumAndMinimum(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
