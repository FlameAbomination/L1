package main

import (
	"fmt"
	"os"
	"reflect"
)

func PrintType(value interface{}) {
	fmt.Fprintf(os.Stdout, "%T\n", value)
}

func PrintType2(value interface{}) {
	fmt.Fprintln(os.Stdout, reflect.TypeOf(value))
}

func PrintType3(value interface{}) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int:
		fmt.Fprintln(os.Stdout, "int")
	case reflect.Bool:
		fmt.Fprintln(os.Stdout, "bool")
	case reflect.String:
		fmt.Fprintln(os.Stdout, "string")
	case reflect.Chan:
		fmt.Fprintln(os.Stdout, "chan")
	default:
		fmt.Fprintln(os.Stdout, "unknown")
	}
}

func task14() {
	values := make(chan int)
	values2 := make(chan bool)
	PrintType(42)
	PrintType2(42)
	PrintType3(42)
	PrintType3(values)
	PrintType3(values2)
}
