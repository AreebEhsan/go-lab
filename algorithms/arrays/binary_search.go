package main

import "fmt"

func binarySearch(arr []int, target int) int{
	left :=0
	right := len(arr) - 1

	for left<=right{
		mid := left + (right - left)/2

		if arr[mid] == target{
			return mid
		}

		if arr[mid] < target{
			left = mid + 1

		}else{
			right = mid - 1
		}
	}

	return -1
}

func main(){
	arr:=[]int{34,67,342,6335,75434,8574327,5746355463,}
	fmt.Println("Index:", binarySearch(arr,75434))
}
