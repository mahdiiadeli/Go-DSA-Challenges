package main

import (
	"fmt"
	"hello_algo/arrays"
)

func main()  {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	arrays.Merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)
}