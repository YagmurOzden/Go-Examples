package main

import "fmt"

func main() {
	//You can create variables with diffrent ways
	var a string = "Yagmur"
	fmt.Println(a)
	//output: Yagmur

	var b string
	b = "Yagmur"
	fmt.Println(b)
	//output: Yagmur

	var c string
	c = "Yagmur"
	fmt.Println(c)
	//output: Yagmur

	d := "Yagmur"
	fmt.Println(d)
	//output: Yagmur

	var e, f, g, t int = 1, 3, 5, 7

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(t)
	//output:1
	//output:3
	//output:5
	//output:7

	var s, q = 6, "Hello"
	r, y := 7, "World!"

	fmt.Println(s)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(y)
	//output:6
	//output:Hello
	//output:7
	//output:Werld!
}

// Go Variable Naming Rules
// A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

// Go variable naming rules:
// A variable name must start with a letter or an underscore character (_)
// A variable name cannot start with a digit
// A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
// Variable names are case-sensitive (age, Age and AGE are three different variables)
// There is no limit on the length of the variable name
// A variable name cannot contain spaces
// The variable name cannot be any Go keywords
