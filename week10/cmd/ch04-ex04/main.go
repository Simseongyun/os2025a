package main

import (
	"week10/pkg/greeting"
	"week10/pkg/greeting/deutsch"
	"week10/pkg/greeting/deutsch/korean"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	korean.Anyung()
	deutsch.GutenTag()
}
