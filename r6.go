package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	n := 0
	mult := 0
	fmt.Println("digite um número para ver sua tabuada:")
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		mult = n * i
		fmt.Println("%dX %d=%d/n", n, i, mult)
	}

}
