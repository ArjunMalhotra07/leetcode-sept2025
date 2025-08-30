package practice_arrays

func missingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i]
		if i != correctIndex && nums[i] != len(nums){
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
		} else {
			i++
		}
	}
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	return len(nums)
}
