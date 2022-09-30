package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 3, 5} // a[1]==3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
