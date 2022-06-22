package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3} //початковий слайс
	var result []int
	var inArr bool //показчик чи є елемент в result

	for i := 0; i < len(arr); i++ {
		inArr = false
		for j := 0; j < len(result); j++ {

			if arr[i] == result[j] {
				inArr = true
			}
		}
		if inArr == false {
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}
