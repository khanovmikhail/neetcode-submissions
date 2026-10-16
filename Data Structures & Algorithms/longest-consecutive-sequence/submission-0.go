func longestConsecutive(nums []int) int {
	elements := make(map[int]bool)
	for _, v := range nums {
		elements[v] = true
	}
	max := 0
	for k, _ := range elements {
		if !elements[k - 1] {
			length, el:= 1, k + 1
			for elements[el] {
				length++
				el++
			}
			if length > max {
				max = length
			}
		}
	}
	return max
}
