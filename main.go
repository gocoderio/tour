package main

import "runtime"

func main() {
	println("Start by clicking Hello, 世界 in the window on the left")

	println("(", runtime.Version(), "- runs entirely in your web browser!)")
}
