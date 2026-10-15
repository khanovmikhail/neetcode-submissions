func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	buckets := make([][]int, len(nums) + 1)
	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}
	res := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}
	return res
}
