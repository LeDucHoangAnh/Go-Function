package main

import "fmt"

// func func_name(params) return type{ //code}
func Hello() {
	fmt.Println("Hello")
}

func HelloWithParam(name string) {
	fmt.Println("Hello", name)
}

func main() {
	Hello()
	HelloWithParam("World")
}
