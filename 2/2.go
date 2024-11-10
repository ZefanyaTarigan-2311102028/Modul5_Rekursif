package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)
	printStar(n)
}

func printStar(n int) {
	if n == 0 {
		return
	}
	printStar(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
