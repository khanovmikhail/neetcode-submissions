func hasDuplicate(nums []int) bool {
    counts := make(map[int]bool)
	for _, v := range nums {
		if counts[v] {
			return true
		}
		counts[v] = true
	}
	return false
}
