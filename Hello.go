package main

import (
	"fmt"
)

const s string = "const"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(1 + 1)
	fmt.Println(true && false)
	fmt.Println(7.0 / 3.0)

	var a, b int = 1, 2
	fmt.Println(a, b)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)

	const d = 3e20
	fmt.Println(d)

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for n := range 4 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
