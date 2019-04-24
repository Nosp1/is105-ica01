package main

import (
	"./sum"
	"fmt"
	"os"
	"strconv"
)

func main (){

	var d1 int
	var d2 int

	args := os.Args

	if len(args) < 3 {
		panic("Not enough parameters to sum, need 2")
	}

	d1, _ = strconv.Atoi(args[1])
	d2, _ = strconv.Atoi(args[2])

	fmt.Println(sum.SumInt8(d1, d2))
}
