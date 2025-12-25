package main

import (
	"fmt"
)

func main() {
	printDiamond(4)
}


func printDiamond(n int) {
   fmt.Println("Мой бриллиант:")

    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            fmt.Print(" ")
        }
        for k := 0; k < 2*i+1; k++ {
            if k == 0 || k == 2*i { 
                fmt.Print("#")
            } else { 
                fmt.Print(" ")
            }
        }
        fmt.Println() 
    }

    for i := n - 2; i >= 0; i-- {
        for j := 0; j < n-i-1; j++ {
            fmt.Print(" ")
        }
        for k := 0; k < 2*i+1; k++ {
            if k == 0 || k == 2*i { 
                fmt.Print("#")
            } else { 
                fmt.Print(" ")
            }
        }
        fmt.Println() 
    }
}
