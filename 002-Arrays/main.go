package main

import "fmt"

func main() {
// Arrays with ints
	var x [5]int
	x[4] = 100
	fmt.Println(x)
//prints: [0 0 0 0 100]
	var f [5]float64
	var num float64 = 77
	println(num)
	for j := 0; j < 5; j++{
		f[j] = num /5
		num++
	}
	fmt.Println(f)
	//prints: [15.4 15.6 15.8 16 16.2]

}