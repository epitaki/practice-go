package main

import "fmt"

func SomeSums(N int, A int, B int) int {
	var o5 int
	var o4 int
	var o3 int
	var o2 int
	var o1 int
	var Sum int
	var result int
	for i := 1; i <= N; i++ {
		o5 = (i / 10000) % 10
		o4 = (i / 1000) % 10
		o3 = (i / 100) % 10
		o2 = (i / 10) % 10
		o1 = i % 10
		Sum = o5 + o4 + o3 + o2 + o1
		if A <= Sum && Sum <= B {
			result += i
		}
	}
	return result
}

func main() {
	var N int
	var A int
	var B int

	fmt.Scan(&N)
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Println(SomeSums(N, A, B))
}
