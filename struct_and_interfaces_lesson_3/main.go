package main

import (
	"fmt"
	"reflect"
	plugin1 "struct_and_interfaces_lesson_3/plugin"
)

// Design a structure a Plugin use Reflection and Interfaces to implement the following methods:
type Plugin interface {
	Run(data string) string
}

func LoadPlugin(name string) Plugin {
	pluginType := reflect.TypeOf(plugin1.Plugin1{})
	pluginValue := reflect.New(pluginType)
	return pluginValue.Interface().(Plugin)
}

func main() {
	p := LoadPlugin("plugin1.Plugin1")
	fmt.Println(p.Run("Hello"))
}
