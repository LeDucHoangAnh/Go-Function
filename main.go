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

//Multiple return values
func rectInfo(w int, h int) (int, int) {
	return w, h
}

func main() {
	Hello()
	HelloWithParam("World")
	result := greeting("Kai")
	fmt.Println(result)

	w, h := rectInfo(100, 200)
	fmt.Println("Width = ", w)
	fmt.Println("Height = ", h)
}
