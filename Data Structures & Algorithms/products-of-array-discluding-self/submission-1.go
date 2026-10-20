func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	left := 1
	for i := 0; i < length; i++ {
		res[i] = left
		left *= nums[i]
	}
	right := 1
	for i := length -1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
// right= 1
// nums =   [1, 2, 3, 4]
// suffix = [24, 12, 4, 1]
// answer = [24, 12, 8, 6]
// res =    [1, 1, 2, 6]