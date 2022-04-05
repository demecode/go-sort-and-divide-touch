package main

import "ftm"

func findMaxSubArray(nums []int, left int right int) int{

	// this will return the actual max sub array of the three segments
	return max(maxLeft, maxRight, middleCross)
}

func main(){

}