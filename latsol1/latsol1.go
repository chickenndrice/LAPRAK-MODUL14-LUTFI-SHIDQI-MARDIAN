package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	y := 0
	for i := 1; i <= x; i++ {
		if i%2 != 0 {
			y++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil\n", y)
}