package main

import (
	"math"
	"math/rand"
	"sort"
)

// 去除重复项
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
	//nums := []int{1, 2, 3}
	//k := removeDuplicates(nums)
	//for i := 0; i < k; i++ {
	//	fmt.Println(nums[i])
	//}

	//arr := []int{1, 2, 3, 4, 5}
	//rotate(arr, 3)
	//fmt.Println(arr)

}

// 轮转数组
func rotate(nums []int, k int) {
	// 翻转3次数组
	// nums = "----->-->"; k =3
	// result = "-->----->";
	//
	// reverse "----->-->" we can get "<--<-----"
	// reverse "<--" we can get "--><-----"
	// reverse "<-----" we can get "-->----->"

	k %= len(nums)

	reverseArray(nums, 0, len(nums)-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)
}

// reverse array
func reverseArray(arr []int, s int, e int) {
	for i, j := s, e; i < j; i, j = i+1, j-1 {
		(arr)[i], (arr)[j] = (arr)[j], (arr)[i]
	}
}

// 股票最佳时机
func maxProfit(prices []int) int {
	// 先找最低，找到最低就往右边找最高
	minPrice := math.MaxInt
	maxPrice := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxPrice {
			maxPrice = price - minPrice
		}
	}
	return maxPrice
}

// 跳跃游戏
func canJump(nums []int) bool {
	// 贪心

	mostRight := 0
	for i := 0; i <= mostRight; i++ {
		mostRight = max(mostRight, i+nums[i])
		if mostRight >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 跳跃游戏2
func jump(nums []int) int {
	// 贪心

	// 「上一跳的落点」=「下一跳的起点」
	// 除了第一跳的起点固定为0（等效于单值区间[0,0]），其余每次落点和起点都是一个范围，所以每一跳都是从一个区间跳入另一个区间
	// 下一个区间范围取决于上一个区间范围的元素(「下一个区间的右端点」=「遍历上一个区间元素所得到的最右能够跳到的位置」)
	// 不用关心区间的左端点，只需关心区间的右端点是否到达最后一个元素
	// 「计算跳跃次数(steps)」=「计算区间个数(count) - 1」

	rangeRight := 0                           // 区间右端点
	mostRight := 0                            // 上个区间内能跳到的最右位置
	count := 1                                // 区间个数，[0,0]固定为第1个区间，所以计数从1开始
	lastIndex := len(nums) - 1                // 最后一个元素的位置
	for i := 0; rangeRight < lastIndex; i++ { // 当 rangeRight > lastIndex 时，该区间已到达最后一个元素，结束循环
		mostRight = max(mostRight, i+nums[i])
		if i == rangeRight { // 当遍历到当前区间的右端点时，将区间的右端点更新为下一个区间的右端点
			rangeRight = mostRight
			count++ // 区间数+1
		}
	}
	return count - 1
}

// H 指数
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j] // 倒序排序
	})

	h := 0
	for i := 0; i < len(citations) && citations[i] > h; i++ { // i 左边的数都满足 > h，因此每次只需最右端这个元素是否大于 h （即>=当前以满足该条件的元素的个数）即可
		h++
	}
	return h
}

// RandomizedSet O(1) 时间插入、删除、获取随机元素
type RandomizedSet struct {
	values       []int       // 存储元素的值
	valueToIndex map[int]int // 记录每个元素对应在 nums 中的索引
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values:       make([]int, 0),
		valueToIndex: make(map[int]int),
	}
}

// Insert 插入到尾部
func (this *RandomizedSet) Insert(val int) bool {
	// 如果已存在
	if _, ok := this.valueToIndex[val]; ok {
		return false
	}

	// 插入到尾部
	this.values = append(this.values, val)
	this.valueToIndex[val] = len(this.values) - 1
	return true
}

// Remove 交换到尾部，移除尾部
func (this *RandomizedSet) Remove(val int) bool {
	// 如果不存在
	if _, ok := this.valueToIndex[val]; !ok {
		return false
	}

	// 将尾部元素替换到该位置，然后移除尾部元素
	valueIndex := this.valueToIndex[val]
	lastElement := this.values[len(this.values)-1]
	this.values[valueIndex] = lastElement
	this.valueToIndex[lastElement] = valueIndex
	this.values = this.values[:len(this.values)-1]
	delete(this.valueToIndex, val)
	return true
}

// GetRandom 根据索引随机获取
func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// 238. 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	// 构建左右乘积数组
	// 空间复杂度 O(n)

	n := len(nums)

	L, R, answer := make([]int, n), make([]int, n), make([]int, n)

	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	for i := 0; i < n; i++ {
		answer[i] = L[i] * R[i]
	}

	return answer
}

// 238. 除自身以外数组的乘积
func productExceptSelf2(nums []int) []int {
	// 用结果数组承载左侧乘积，最后计算结果的同时计算右侧乘积
	// 空间复杂度 O(1)

	n := len(nums)

	answer := make([]int, n)

	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	R := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i] // 右侧乘积
	}

	return answer
}

// 134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	// 图像法
	// 能否跑完环路：「总加油量 - 总耗油量」是否 >= 0
	// 找出发站（如果能跑完环路）：将每一站的「加油量-耗油量」累加差值作图，找到最低点，最低点就是出发点（最低点即油箱油量最少的一站，将这一站作为出发站，其余每一站的油箱油量只多不少）

	// 油箱油量、油箱最低油量、出发站
	sum, minSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ { // 无论从哪一站出发，最低点都是确定的一站，因为每站的相对关系不变，这里假设从0站出发进行计算
		sum += gas[i] - cost[i]
		if sum < minSum {
			// 经过第 i 站，油箱油量达到新低
			minSum = sum
			// 第 i+1 站，就是最低点（起点）
			start = i + 1
		}
	}

	if sum < 0 {
		return -1
	}

	if start == len(gas) {
		return 0
	}

	return start
}

// 135. 分发糖果
func candy(ratings []int) int {
	// 分别从左、右按照「大了就加」的规则得出每人的糖果数，最后取两者中更大的数即可

	n := len(ratings)
	left := make([]int, n)
	count := 0

	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		count += max(left[i], right)
	}

	return count
}

// 42. 接雨水
func trap(height []int) int {
	// 双指针

	// 1. 初始化两个指针 left = 0 和 right = height。length - 1。
	// 2． 当 left ＜ right 时向中间移动两个指针：
	// 	- 如果 height[left] < height[right] 说明储水量依赖于 height[left] 的高度（可能构成低洼的右边界很大）
	//		- 如果 height[left] > left_max 说明没有或超出左边边界，不构成低洼，left_max = height[left] 。
	// 		- 如果 height[left] <= left_max 说明构成低洼，往答案中累加积水量。ans += left_max - height[left]
	//		- 前进 left。left++
	//	- 如果 height[left] >= height[right] 说明储水量依赖于 height[right] 的高度（可能构成低洼的左边界很大）
	//		- 如果 height[right] > right_max 说明没有或超出右边边界，不构成低洼，right_max = height[right]
	// 		- 如果 height[right] <= right_max 说明构成低洼，往答案中累加积水量。ans += right_max - height[right]
	//		- 前进 right。right--

	ans, l, lMax, r, rMax := 0, 0, 0, len(height)-1, 0

	for l < r {
		if height[l] < height[r] {
			lMax = max(lMax, height[l])
			ans += lMax - height[l]
			l++
		} else {
			rMax = max(rMax, height[r])
			ans += rMax - height[r]
			r--
		}
	}

	return ans
}
