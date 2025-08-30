package practice_arrays

func CycleSort(arr []int) []int {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] != arr[correct] {
			arr[i], arr[correct] = arr[correct], arr[i]
		} else {
			i++
		}
	}
	return arr
}
