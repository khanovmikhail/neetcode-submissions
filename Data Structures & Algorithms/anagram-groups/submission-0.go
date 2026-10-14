func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		rune := []rune(str)
		sort.Slice(rune, func(i, j int)bool {
			return rune[i] < rune[j]
		})
		key := string(rune)
		anagrams[key] = append(anagrams[key], str)
	}
	var result [][]string
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}
