package main

import (
	"fmt"
)

const miles2kilometes = 1.6

func main() {
	i := 10
goolakh:
	for i = 0; i <= 10; i++ {
		continue
		switch i {
		case 2:
			fmt.Println("42")
			break goolakh
		case 43:
			fmt.Println("43")

		}
	}
	for i = 2; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("hello")
}
