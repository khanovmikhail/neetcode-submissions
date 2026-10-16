func longestConsecutive(nums []int) int {
	elements := make(map[int]bool)
	for _, v := range nums {
		elements[v] = true
	}
	max := 0
	for k, _ := range elements {
		if !elements[k - 1] {
			length := 1
			for i := k + 1; elements[i]; i++ {
				length++
			}
			if length > max {
				max = length
			}
		}
	}
	return max
}
