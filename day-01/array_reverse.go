	package day01

	// Original Array:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 1 |  2 |  3 |  4 |  5 |
	//          -------------------------

	// Step 1: Create an empty auxiliary array of the same size.
	// Auxiliary Array (initially):
	// Index:    0    1    2    3    4
	//          -------------------------
	//          |   |    |    |    |   |
	//          -------------------------

	// Step 2: Copy elements in reverse order:
	// - Copy original[4] (which is 5) into aux[0]
	// - Copy original[3] (which is 4) into aux[1]
	// - Copy original[2] (which is 3) into aux[2]
	// - Copy original[1] (which is 2) into aux[3]
	// - Copy original[0] (which is 1) into aux[4]

	// Auxiliary Array after copying:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 5 |  4 |  3 |  2 |  1 |
	//          -------------------------

	// Step 3 (Optional): Copy the auxiliary array back to the original array.
	// Final Reversed Array:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 5 |  4 |  3 |  2 |  1 |
	//          -------------------------

	// Approach 1: Using an auxiliary array  time o(n) space o(n)

	func reverseArray_using_aux_array(arr []int) {
		aux := make([]int, len(arr))
		for i := 0; i < len(arr); i++ {
			aux[i] = arr[len(arr)-1-i]
		}

		for i := 0; i < len(arr); i++ {
			arr[i] = aux[i]
		}
	}

	// Approach 2: Using two pointers  time o(n) space o(1)

	// Original Array:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 1 |  2 |  3 |  4 |  5 |
	//          -------------------------

	// Step 1: Initialize two pointers.
	// Set start = 0 and end = 4

	// Step 2: Swap the elements at indices start and end:
	// - Swap arr[0] (which is 1) with arr[4] (which is 5)
	// Array becomes:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 5 |  2 |  3 |  4 |  1 |
	//          -------------------------

	// Step 3: Move the pointers inward:
	// - Increment start to 1 and decrement end to 3

	// Step 4: Swap the elements at indices start and end:
	// - Swap arr[1] (which is 2) with arr[3] (which is 4)
	// Array becomes:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 5 |  4 |  3 |  2 |  1 |
	//          -------------------------

	// Step 5: Now start equals 2 and end equals 2, so the loop terminates.

	// Final Reversed Array:
	// Index:    0    1    2    3    4
	//          -------------------------
	//          | 5 |  4 |  3 |  2 |  1 |
	//          -------------------------

	func reverseArrayTwoPointer(arr []int) {
		start := 0
		end := len(arr) - 1
		for start < end {
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
			start++
			end--
		}
	}
