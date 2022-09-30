package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 3, 5} // a[1]==3
	var d = [3]int{1, 2, 3}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// OUTPUT:
	// [1 2 3]
	// [1 2 3]
	// [1 3 5]
	// [1 2 3]

	// You can access a specific array element by referring to the index number.
	// In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.
	numbers := [3]int{11, 22, 33}

	fmt.Println(numbers[0])
	fmt.Println(numbers[2])
	//output :
	// 11
	// 33

	numbers[2] = 50
	fmt.Println(numbers)
	//OUTPUT: [11 22 50]

	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{0, 1, 2}       //partially initialized
	arr3 := [5]int{0, 1, 2, 3, 4} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	// OUTPUT:
	// [0 0 0 0 0]
	// [0 1 2 0 0]
	// [0 1 2 3 4]
}

//Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
