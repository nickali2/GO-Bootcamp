package main

import "fmt"

const miles2kilometes = 1.6

func main() {
	var miles float64
	fmt.Print("enter miles:")
	_, err := fmt.Scanf("%v", &miles)
	if err != nil {
		fmt.Println(err)
	}
	kilomete := miles2kilometes * miles
	fmt.Println(miles, " is equal to ", kilomete)
}
