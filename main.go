package main

import "fmt"

func main() {
	printTable(4)
}
func printTable(num int) {

	for i := 1; i <= num; i++ {
		for a := 1; a <= num; a++ {
			if a == num {
				fmt.Printf("%d x %d = %d\n", i, a, i*a)
				continue
			}
			fmt.Printf("%d x %d = %d\t", i, a, i*a)
		}
	}
}
