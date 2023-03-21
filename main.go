package main

import "fmt"

func main() {
	fmt.Println("Inception")
	inception()
}

func inception() {
	fmt.Println("Level 1: Hello from inside Inception")
	level2()
}

func level2() {
	fmt.Println("Level 2: Hello from inside Level 2")
	level3()
}

func level3() {
	fmt.Println("Level 3: Hello from inside Level 3")
}
