package main

import "fmt"

// func func_name(params) return type{ //code}
func Hello() {
	fmt.Println("Hello")
}

func HelloWithParam(name string) {
	fmt.Println("Hello", name)
}

func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}
func main() {
	Hello()
	HelloWithParam("World")
	result := greeting("Kai")
	fmt.Println(result)
}
