package day01

// FindSecondSmallestAndSecondLargest finds the second smallest and second largest elements in an array.
// Approach 1: Using two auxiliary variables.

// Initial Array:
// -------------------------------------------------
// Index:      0    1    2    3    4    5
// Array:     [ 3,   1,   5,   0,   4,   2 ]
// -------------------------------------------------
// Initial State:
//     min1 = 3      max1 = 3
//     min2 = 3      max2 = 3
// -------------------------------------------------

// Iteration over each element:

// 1. Index 0, Element = 3
//    - Compare 3 with min1 (3):  3 < 3? → False.
//    - Compare 3 with max1 (3):  3 > 3? → False.
//    State remains:
//        min1 = 3,   min2 = 3,   max1 = 3,   max2 = 3

// -------------------------------------------------

// 2. Index 1, Element = 1
//    - Compare 1 with min1 (3):  1 < 3? → True.
//          => Update: min2 = min1 (3), then min1 = 1.
//    - Compare 1 with max1 (3):  1 > 3? → False.
//    Updated State:
//        min1 = 1,   min2 = 3,   max1 = 3,   max2 = 3

// -------------------------------------------------

// 3. Index 2, Element = 5
//    - Compare 5 with min1 (1):  5 < 1? → False.
//    - Compare 5 with max1 (3):  5 > 3? → True.
//          => Update: max2 = max1 (3), then max1 = 5.
//    Updated State:
//        min1 = 1,   min2 = 3,   max1 = 5,   max2 = 3

// -------------------------------------------------

// 4. Index 3, Element = 0
//    - Compare 0 with min1 (1):  0 < 1? → True.
//          => Update: min2 = min1 (1), then min1 = 0.
//    - Compare 0 with max1 (5):  0 > 5? → False.
//    Updated State:
//        min1 = 0,   min2 = 1,   max1 = 5,   max2 = 3

// -------------------------------------------------

// 5. Index 4, Element = 4
//    - Compare 4 with min1 (0):  4 < 0? → False.
//    - Compare 4 with max1 (5):  4 > 5? → False.
//    State remains:
//        min1 = 0,   min2 = 1,   max1 = 5,   max2 = 3

// -------------------------------------------------

// 6. Index 5, Element = 2
//    - Compare 2 with min1 (0):  2 < 0? → False.
//    - Compare 2 with max1 (5):  2 > 5? → False.
//    Final State remains:
//        min1 = 0,   min2 = 1,   max1 = 5,   max2 = 3

// -------------------------------------------------
// Final Result:
//     Second Smallest (min2) = 1
//     Second Largest  (max2) = 3
// -------------------------------------------------

func FindSecondSmallestAndSecondLargest(arr []int) (int, int) {
	min1, max1 := arr[0], arr[0]
	min2, max2 := arr[0], arr[0]
	for _, v := range arr {
		if v < min1 {
			min2 = min1
			min1 = v
		}

		if v > max1 {
			max2 = max1
			max1 = v
		}

	}
	return min2, max2
}
