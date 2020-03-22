package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	var s1 int
	var s2 int
	var s3 int

	s1, _ = strconv.Atoi(string(s[0]))
	s2, _ = strconv.Atoi(string(s[1]))
	s3, _ = strconv.Atoi(string(s[2]))
	fmt.Println(s1 + s2 + s3)

}
