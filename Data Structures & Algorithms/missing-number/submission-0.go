func missingNumber(nums []int) int {
	res := 0
	for i := 0; i <= len(nums); i++ {
		res ^= i
	}
	for _, v := range nums {
		res ^= v
	}
	return res
}
