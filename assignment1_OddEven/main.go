package main

import "fmt"

func main() {

	numbers := generateNumbers()

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, "is Even")
		} else {
			fmt.Println(n, "is Odd")
		}
	}
}
