package main

import _ "fmt"

func main() {
	for i := 1; i <= 100; i += 1 {
		if (i%3 == 0) && (i%5 == 0) {
			print("fizzbuzz,")
		} else if i%5 == 0 {
			print("Buzz,")
		} else if i%3 == 0 {
			print("fizz,")
		} else {
			print(i, " ")
		}
	}
}
