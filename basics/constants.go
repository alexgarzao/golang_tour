package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}