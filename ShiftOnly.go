package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	slice := strings.Split(s, " ")
	var inputsArray []int
	for i := 0; i < N; i++ {
		bar, _ := strconv.Atoi(slice[i])
		inputsArray = append(inputsArray, bar)
	}
	result := 0
  breakflag := 0
	for {
		for _, s := range inputsArray {
          nino:=int(math.Pow(2, float64(result)))
          if s%nino != 0 {
            	breakflag = 1
				break
			}
		}
      if breakflag==1{
      break
      }
		result += 1
	}
  fmt.Println(result-1)

}

func ea