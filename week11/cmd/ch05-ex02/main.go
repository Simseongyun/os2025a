package main

import (
	"fmt"
	"github.com/headfirstgo/datafile" 
	"log"
)

func main() {
	weights, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	hap := 0.0
	//for _, number := range numbers {
	for i := 0; i < len(weights); i++{
		hap = hap + weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("평균: ", hap/weeks)
}
