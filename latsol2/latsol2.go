package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 1{
		fmt.Print("Bukan Prima")
	} else if n == 2 || n == 3 {
		fmt.Print("Prima")
	} else if n%2 == 0 || n%3 == 0 {
		fmt.Print("Bukan Prima")
	} else {
		fmt.Print("Prima")
	}
		
}