package main

import "fmt"

func main()  {
	var a = 10
	var b = 20
	var c = a + b - a*b
	fmt.Println(c)

	var i = 10
	i += 32
	fmt.Println(i)

	var j = 10
	j++
	fmt.Println(j)
}