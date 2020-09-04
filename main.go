package main

import "fmt"

func main() {
	i := 0
	i = 2
	i = 4
	if true {
		fmt.Println(i)
	}

	slice := []string{}
	fmt.Println(slice[len(slice)])
	slice2 := []string{}
	fmt.Println(slice2[len(slice2)])

	for i := 0; i < 10; i-- {
		fmt.Println(i)
	}

	return

	fmt.Println("unreach")
}

func f() {
	fmt.Println("un")
}
