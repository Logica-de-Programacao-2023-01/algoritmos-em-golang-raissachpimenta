package main

import "fmt"

func main() {
	n := 0
	count := 0
	men := 0
	mai := 0
	for true {
		fmt.Println("digite um número:")
		fmt.Scan(&n)
		count++

		if count == 1 {
			mai = n
			men = n
		}
		if n > mai {
			mai = n
		} else if men > n {
			men = n
		}
		if n == 0 {
			break
		}
	}
	fmt.Println("o maior número digitado foi %d", mai)
}
