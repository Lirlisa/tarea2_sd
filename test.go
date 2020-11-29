package main

import "fmt"

func main() {
	a := [3](int){4, 5, 6}
	for i, val := range a {
		fmt.Printf("%d: %d\n", i, val)
	}
}
