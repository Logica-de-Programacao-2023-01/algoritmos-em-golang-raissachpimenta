package main

import "fmt"

func main() {
	n := 0.0
	resp := ""
	count := 0.0
	soma := 0.0
	med := 0.0
	for true {
		fmt.Println("digite um número")
		fmt.Scan(&n)
		count++
		soma += n
		fmt.Println("deseja continuar[S/N]?")
		fmt.Scan(&resp)
		if resp == "N" || resp == "n" {
			break
		}
	}
	med = soma / count
	fmt.Println("a média dos números é %.2f, med")
}
