package main

import "fmt"

func twoSum(arr []int, target int) [] int {

	hash := make(map[int]int)

	for i, v := range arr{
		diff  := target- v

		if idx, found := hash[diff]; found{
			return [] int{idx, i}
		}
		hash[v] = i
	}
	return nil
}

func main(){

	arr:= []int{12,653,1245,7543,1223,-1234,54}
	target := -11

	result := twoSum(arr, target)
	fmt.Println("Indices", result)
}





