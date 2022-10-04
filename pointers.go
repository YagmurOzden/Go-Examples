package main

import "fmt"

func main() {
	a := 42
	b := a
	fmt.Println(a, b)
	//OUTPUTS: 42 42

	a = 42
	var c *int = &a
	fmt.Println(a, c)
	//OUTPUTS: 42 0xc0000b2008

	fmt.Println(&a, c)
	//OUTPUTS: 0xc000014090 0xc000014090

	fmt.Println(a, *c)
	//OUTPUTS: 42 42

	d := [3]int{1, 2, 3}
	k := &d[0]
	t := &d[1]
	fmt.Printf("%v %p %p \n", d, k, t)
	//OUTPUTS: [1 2 3] 0xc00011c000 0xc00011c008

	var ms1 myStuct
	ms1 = myStuct{foo: 42}
	fmt.Println(ms1)
	//OUTPUT: {42}

	var ms2 *myStuct
	ms2 = &myStuct{foo: 42}
	fmt.Println(ms2)
	//OUTPUT: &{42}

	var ms3 *myStuct
	fmt.Println(ms3)
	//OUTPUT <nil>
	ms3 = new(myStuct)
	fmt.Println(ms3)
	//OUTPUT: &{0}

	var ms4 *myStuct
	ms4 = new(myStuct)
	(*ms4).foo = 42 //dereferancing variable
	fmt.Println((*ms4).foo)
	//OUTPUT: 42

	//clean version
	var ms5 *myStuct
	ms5 = new(myStuct)
	ms5.foo = 42 //dereferancing variable
	fmt.Println(ms5.foo)
	//OUTPUT: 42

	//for arrays
	a1 := [3]int{1, 2, 3}
	b1 := a1
	fmt.Println(a1, b1)
	//OUTPUTS:[1 2 3] [1 2 3]
	a1[1] = 42
	fmt.Println(a1, b1)
	//OUTPUTS:[1 42 3] [1 2 3]

	//slices acts diffrent
	a2 := []int{1, 2, 3}
	b2 := a2
	fmt.Println(a2, b2)
	//OUTPUTS [1 2 3] [1 2 3]
	a2[1] = 42
	fmt.Println(a2, b2)
	//OUTPUTS [1 42 3] [1 42 3]

	//for MAPS
	a_map := map[string]string{"foo": "bar", "baz": "buz"}
	b_map := a_map
	fmt.Println(a_map, b_map)
	//OPUTPUT: map[baz:buz foo:bar] map[baz:buz foo:bar]
	a_map["foo"] = "qux"
	fmt.Println(a_map, b_map)
	//OUTPUT: map[baz:buz foo:qux] map[baz:buz foo:qux]
}

type myStuct struct {
	foo int
}
