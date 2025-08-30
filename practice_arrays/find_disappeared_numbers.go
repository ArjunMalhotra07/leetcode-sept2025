package practice_arrays

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if i != correct && nums[correct] != nums[i] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}
	ans := []int{}
	for i, num := range nums {
		if i+1 != num {
			ans = append(ans, i+1)
		}
	}
	return ans
}
