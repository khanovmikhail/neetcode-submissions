import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))
	for _, str := range strs {
		bytes := []byte(str)
		slices.Sort(bytes)
		key := string(bytes)
		groups[key] = append(groups[key], str)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
