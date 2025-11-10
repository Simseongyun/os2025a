package main

import "fmt"

func main() {
	subjects := []string{"Go", "", "Python"}

	for _, subjects := range subjects {
		fmt.Println(subjects)
	}
}
