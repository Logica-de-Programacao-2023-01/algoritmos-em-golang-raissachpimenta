package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	n := 0
	fmt.Println("digite um número para ver seus divisores")
	fmt.Scan(&n)
	fmt.Println("os dividores do núemro são:")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			print(i, "")
		}
	}
}
