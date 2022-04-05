package main

import "fmt"

func findMaxSubArray(nums []int, left int,right int) int{

	// create a base case to finish to recursion
	if left == right {
		return nums[left]
	} else {
	// perform recurssion for the three segments
	mid := (left + right) / 2
		maxLeft := findMaxSubArray(nums, left, right, mid)
		maxRight := findMaxSubArray(nums, mid+1, right)

		// cross fucntion for a max sub arraybetween the left, mid and right
		cross := maxCross(nums, left, mid, right)

		// this will return the actual max sub array of the three segments
		return max(maxLeft, maxRight, middleCross)
	}
}

func main(){

}