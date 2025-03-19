package main

import (
	"fmt"
	"sort"
)

// func arrays() {
// 	// Declaration with var keyword and type
// 	var arr1 [5]int

// 	// Declaration with initialization
// 	var arr2 [3]string = [3]string{"Go", "Rust", "Python"}

// 	// Short declaration with initialization
// 	arr3 := [4]float64{1.1, 2.2, 3.3, 4.4}

// 	// Using ... to let the compiler count the elements
// 	arr4 := [...]int{10, 20, 30, 40, 50}

// 	// Sparse array (specific indexes)
// 	arr5 := [5]int{0: 100, 4: 400}

// }

// func slices() {
// 	// Empty slice declaration
// 	var slice1 []int

// 	// Slice with initialization
// 	var slice2 []string = []string{"Go", "Rust", "Python"}

// 	// Short declaration with initialization
// 	slice3 := []float64{1.1, 2.2, 3.3, 4.4}

// 	// Make function with length
// 	slice4 := make([]int, 5)

// 	// Make function with length and capacity
// 	slice5 := make([]int, 3, 10)

// 	// Slice created from an array
// 	arr := [5]int{10, 20, 30, 40, 50}
// 	slice6 := arr[1:4] // Contains [20, 30, 40]

// 	// Slice of a slice
// 	slice7 := slice6[1:2] // Contains [30]
// }

// func main() {
// 	// arr := [3]int{1, 2, 3}
// 	// copiedArr := arr
// 	// fmt.Printf("Address of arr: %p\n", &arr)
// 	// fmt.Printf("Address of copied arr: %p\n", &copiedArr)

// 	slice := []int{1, 2, 3}
// 	copiedSlice := slice
// 	fmt.Printf("Address of slice header: %p\n", &slice)
// 	fmt.Printf("Address of copiedSlice header: %p\n", &copiedSlice)
// 	fmt.Printf("Address of slice: %p\n", &slice[0])
// 	fmt.Printf("Address of copiedSlice: %p\n", &copiedSlice[0])
// }

// func main() {
// 	// For arrays
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	fmt.Printf("MAIN: Address of array: %p\n", &arr)

// 	// For slices
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Printf("MAIN: Address of slice header: %p\n", &slice)
// 	fmt.Printf("MAIN: Address of slice's underlying array: %p\n", &slice[0])
// 	fmt.Printf("MAIN BEFORE blah: B[0]: %d\n", slice[0])
// 	blah(arr, slice)
// 	fmt.Printf("MAIN: arr[0]: %d\n", arr[0])

// 	fmt.Printf("MAIN: slice[0]: %d\n", slice[0])
// }

// func blah(a [5]int, b []int) {
// 	fmt.Printf("FUNC: Address of array: %p\n", &a)
// 	a[0] = 1000
// 	fmt.Printf("FUNC: a[0]: %d\n", a[0])

// 	fmt.Printf("FUNC: Address of slice header: %p\n", &b)
// 	fmt.Printf("FUNC: Address of slice: %p\n", b)
// 	b[0] = 100
// 	fmt.Printf("FUNC: B[0]: %d\n", b[0])
// }

func main() {
	scores := []int{}
	fmt.Printf("Initial: scores = %v, length = %d\n", scores, len(scores))

	scores = append(scores, 85)
	scores = append(scores, 70, 65)
	fmt.Printf("After append: scores = %v, length = %d\n", scores, len(scores))

	moreScore := []int{1, 29}
	scores = append(scores, moreScore...)
	fmt.Printf("After append slice: scores = %v, length = %d\n", scores, len(scores))

	names := []string{"Tim", "Bob", "Jess", "Zac"}
	fmt.Printf("Unsorted: names = %v\n", names)

	sort.Strings(names)
	fmt.Printf("Sorted: names = %v\n", names)

	sort.Sort(sort.Reverse((sort.IntSlice(scores))))
	fmt.Printf("Scores sorted (descending): scores = %v\n", scores)

}
