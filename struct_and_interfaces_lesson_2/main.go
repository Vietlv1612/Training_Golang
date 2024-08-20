package main

import (
	"fmt"
	"reflect"
)

// Create a new basic ORM use Reflection
type User struct {
	ID      int
	Name    string
	Age     int
	Address string
	Email   string
}

func InsertQuery(tableName string, s interface{}) string {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	fields := ""
	values := ""

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i).Name
		value := v.Field(i)

		fields += field + ","
		if value.Kind() == reflect.String {
			values += fmt.Sprintf("'%v'", value)
		} else {
			values += fmt.Sprintf("%v", value)
		}

		if i != v.NumField()-1 {
			fields += ","
			values += ","
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, fields, values)
	return query
}

func main() {
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	query := InsertQuery("users", user)
	fmt.Println(query)
}
