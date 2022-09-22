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
func rectInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

//Named return values
func rectInfo1(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}
func main() {
	Hello()

	HelloWithParam("World")

	result := greeting("Kai")
	fmt.Println(result)

	w, h, area := rectInfo(100, 200)
	fmt.Println("Width = ", w)
	fmt.Println("Height = ", h)
	fmt.Println("area = ", area)

	w, h, isSquare := rectInfo1(100, 200)
	if isSquare {
		fmt.Println("'This is square'")
	} else {

		fmt.Println("width = ", w)
		fmt.Println("height = ", h)
	}

}
