package main

import "fmt"
func main(){
	xs := []float64{98,93,77,82,83}
  	fmt.Println(average(xs))


}
/*
a function that takes an array of floats
 and returns the average of these numbers
*/
func average(x []float64) float64{
	total := 0.0
	for _, v := range x{
		total += v
	}
	return total / float64(len(x))
}