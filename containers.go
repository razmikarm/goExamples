package main

import "fmt"

func arrays() {
	var nums [3]int

	nums[0] = 2
	nums[1] = 4
	nums[2] = 7

	fmt.Println(nums[2])

	var numbers = [4]int{3, 5, 8, 11}

	fmt.Println(numbers)

	for _, val := range numbers {
		fmt.Println(val)
	}

}

func withSlice() {

	numSlice := []int{0, 1, 2, 3, 4, 5, 6, 7}

	numSlice2 := numSlice[6:8]

	fmt.Println(numSlice2)

	numSlice3 := make([]int, 8, 10)

	copy(numSlice3, numSlice)

	numSlice[1] = 8

	numSlice3 = append(numSlice3, 8, 9)

	fmt.Println(numSlice)
	fmt.Println(numSlice3)

}

func withMaps() {
	pet := map[string]string{
		"name": "Rex",
		"type": "dog",
	}
	fmt.Println(pet)
	pairs := make(map[int][]int)
	pair := []int{0, 1, 2}
	pairs[42] = append(pairs[0], pair...)
	pairs[6] = append(pairs[1], 3, 4)
	fmt.Println(pairs)
	delete(pairs, 5)
	delete(pairs, 6)
	fmt.Println(pairs)
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func switch2Vals(a int, b int) (int, int) {
	return b, a
}
