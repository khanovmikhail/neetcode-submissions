func isValid(s string) bool {
    pairs := map[rune]rune {
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			l := len(stack) - 1
			pop := stack[l]
			stack = stack[:l]
			if pairs[pop] != r {
				return false
			}
		}
	}
	return len(stack) == 0;
}
