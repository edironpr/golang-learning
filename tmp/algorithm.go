package main

import "fmt"

func removeDuplicates(nums []int) int {
	// 快慢指针

	// 处理边界
	if len(nums) <= 2 {
		return len(nums)
	}

	// 直接从第二、三个元素开始（由于可保留两个重复项，因此第一个元素和第二个元素是否相等并不重要，不需要判断）
	slow, fast := 1, 2
	for fast < len(nums) {
		// 在相等的情况下，再判断一下 slow 前一个元素的值
		if nums[slow] != nums[fast] || nums[slow-1] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func main() {
	nums := []int{1, 2, 3}
	k := removeDuplicates(nums)
	for i := 0; i < k; i++ {
		fmt.Println(nums[i])
	}
}
