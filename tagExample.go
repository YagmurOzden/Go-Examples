package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `requried max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
