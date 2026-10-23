func isValid(s string) bool {
    pairs := map[rune]rune {
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune
	for _, r := range s {
		if len(stack) == 0 {
			stack = append(stack, r)
			continue
		}
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
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
