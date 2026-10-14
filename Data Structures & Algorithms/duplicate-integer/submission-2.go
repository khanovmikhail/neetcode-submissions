func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
	for _, v := range nums {
		if x := m[v]; x {
			return true
		}
		m[v] = true
	}
	return false
}
