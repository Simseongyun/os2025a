package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var now time.Time = time.Now()
	//month := now.Month()
	var month time.Month = now.Month()
	fmt.Println(month)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	log.Fatal(err) // report the error and exit program
	fmt.Println(i)
}
