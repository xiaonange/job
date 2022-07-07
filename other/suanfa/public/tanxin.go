package public

func MaxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	var currSum, maxSum = 0, nums[0]
	for _, v := range nums {
		if currSum > 0 {
			currSum += v
		} else {
			currSum = v
		}
		if maxSum < currSum {
			maxSum = currSum
		}
	}
	return maxSum
}
//动态规划
func MaxSubArrayDT(nums []int) int {
	max_sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i - 1] > 0 {
			nums[i] += nums[i - 1]
		}
		if nums[i] > max_sum {
			max_sum = nums[i]
		}
	}
	return max_sum
}
