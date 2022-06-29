/*
practice program Golang Loops
*/
package main

func main() {
// for loop
	x := [10]int{}
	for i := 0; i < 10; i++{
		x[i] = i+3
	}
// range loop
	for  l, value:= range x{ //use _ to tell go you dont need this 
		println(value)
		println(l)
	} 
}