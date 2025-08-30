package practice_arrays

// ! https://leetcode.com/problems/find-the-duplicate-number/
func FindDuplicate(nums []int) int {
	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if correct != i && nums[correct] != nums[i] {
			nums[correct], nums[i] = nums[i], nums[correct]
		} else {
			i += 1
		}
	}
	for i, val := range nums {
		if i+1 != val {
			return val
		}
	}
	return -1
}
