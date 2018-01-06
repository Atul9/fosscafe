package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Println("Hello\tworld")
	fmt.Println("value of area :", area)
	max_number := max(5, 10)
	min_number := min(5, 10)
	fmt.Println("Max out of 5 & 10 :", max_number)
	fmt.Println("Min out of 5 & 10 :", min_number)
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func min(num1, num2 int) int {
	var result int
	if num1 < num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
