package main

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//返回滑动窗口中的最大值。

//给定一个数组和一个窗口为 K 的窗口，当窗口从数组的左边滑动到数组右边的时候，输出每次移动窗口以后，在窗口内的最大值。
//这道题最暴力的方法就是 2 层循环，时间复杂度 O(n * K)。
//另一种思路是用优先队列，每次窗口以后的时候都向优先队列里面新增一个节点，并删除一个节点。时间复杂度是 O(n * log n)
//最优的解法是用双端队列，队列的一边永远都存的是窗口的最大值，队列的另外一个边存的是比最大值小的值。队列中最大值左边的所有值都出队。
//在保证了双端队列的一边即是最大值以后，时间复杂度是 O(n)，空间复杂度是 O(K)

//暴力解法
func maxSlidingWindow1(nums []int, k int) []int {
	res := make([]int, 0, k)
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	for i := 0; i <= n-k; i++ {
		max := nums[i]
		for j := 1; j < k; j++ {
			if max < nums[i+j] {
				max = nums[i+j]
			}
		}
		res = append(res, max)
	}
	return res
}

//双端队列，滑动窗口
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:len(window)]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}

func main() {

}
