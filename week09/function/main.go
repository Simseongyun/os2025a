package main

import (
	"fmt"
)

func swap(first int, second int) {
	temp := 0
	temp = first
	first = second
	second = first
	fmt.Println(fisrt, second)
}

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
	swap(a, b)
	fmt.Println(a, b)
}
