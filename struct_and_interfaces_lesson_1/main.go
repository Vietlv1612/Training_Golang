package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func AnalyzeStruct(p interface{}) {
	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)

	if t.Kind() != reflect.Struct {
		fmt.Println("Person is not a struct")
		return
	}

	fmt.Println("Analyzing struct:", t.Name())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s, Value: %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
}

func main() {
	p := Person{
		Name:   "John Doe",
		Age:    30,
		Gender: "Male",
	}

	AnalyzeStruct(p)
}
