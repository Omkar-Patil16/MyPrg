package main

import "fmt"

func main() {
	fmt.Println("this is the first go code for me")

	mytest()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func mytest() {
	fmt.Println("this is an example of func calling")
}
