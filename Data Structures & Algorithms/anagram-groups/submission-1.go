func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int)bool {
			return runes[i] < runes[j]
		})
		key := string(runes)
		anagrams[key] = append(anagrams[key], str)
	}
	var result [][]string
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}
