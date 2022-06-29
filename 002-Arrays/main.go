package main

import (
	"fmt"
)

func main() {
/* ARRAYS */
	var x [5]int
	x[4] = 100
	fmt.Println(x)
//prints: [0 0 0 0 100]
	var f [5]float64
	var num float64 = 77
	fmt.Println(num)
	for j := 0; j < 5; j++{
		f[j] = num /5
		num++
	}
	fmt.Println(f)
	//prints: [15.4 15.6 15.8 16 16.2]

/* SLICES */
	var slice1 []float64

	slice2 := make([]float64, 5)
	fmt.Println(slice1)
	//prints: []
	slice2[3] = 13.5
	fmt.Println(slice2) 
	//prints: [0 0 0 13.5 0]

	slice3 := []int{1,2,3}
	slice4 := append(slice3, 4, 5)

	fmt.Println(slice3, slice4)
	//prints: [1 2 3] [1 2 3 4 5]

/* MAPS */
	m := make(map[string]int)
	m["first"] = 10

	fmt.Println(m)
	//prints: map[first:10]

	//elements map
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	//to check if an element is in the map or not
	if name, ok := elements["He"]; ok{
		fmt.Println(name, ok)
		//prints: Helium true
	}

}