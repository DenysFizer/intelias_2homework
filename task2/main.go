package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	_ = result
	var min, max int
	//var num int
	numsinString := strings.Fields(input)
	var numsArr []int

	for _, i := range numsinString {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numsArr = append(numsArr, j)
	}
	max = numsArr[0]
	min = numsArr[0]

	for i := 0; i < len(numsArr); i++ {
		switch {
		case numsArr[i] > max:
			max = numsArr[i]
		case numsArr[i] < min:
			min = numsArr[i]
		}
	}
	if max == min {
		result = result + strconv.Itoa(max)
	} else {
		result = result + strconv.Itoa(max) + " " + strconv.Itoa(min)
	}
	fmt.Println(result)

}
