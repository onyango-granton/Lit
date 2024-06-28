package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// firstFilter := []int{}
	res := []int{}
	// for _, ch := range nums {
	// 	if ch > target {
	// 		continue
	// 	}
	// 	firstFilter = append(firstFilter, ch)
	// }

	for i,ch := range nums{
		diff := target - ch
		for v, d := range nums{
			if diff == d && i != v{
				res = append(res, i)
				res = append(res, v)
			}
			if len(res) == 2{
				return res
			}
		}
	}

	return res
}

func main() {
	numArr := []int{3,3}
	target := 6
	fmt.Println(twoSum(numArr, target))
}
