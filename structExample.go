package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	compansion []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		compansion: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)               //{3 Jon Pertwee [Liz Shaw Jo Grant Sarah Jane Smith]}
	fmt.Println(aDoctor.compansion[1]) //Jo Grant

	// OUTPUT:
	// {3 Jon Pertwee [Liz Shaw Jo Grant Sarah Jane Smith]}
	// Jo Grant
}

// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
// While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
// A struct can be useful for grouping data together to create records.
