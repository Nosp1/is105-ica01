package main

import (
	"./sum"
	"fmt"
)

func main (){

	v := sum.SumUint32(5,6)
	t := sum.SumInt8(8,8)
	fmt.Println(v,t)
	var i int8
	var j int8
	fmt.Scan(&i)
	fmt.Scan(&j)
	fmt.Println(sum.SumInt8(i,j))
}
