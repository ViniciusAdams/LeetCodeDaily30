func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if open, exists := bracketMap[char]; exists { // If it's a closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // Pop the top element
		} else { // It's an opening bracket
			stack = append(stack, char)
		}
	}

	return len(stack) == 0 // Stack should be empty if valid
}

func main() {
	fmt.Println(isValid("()"))      // true
	fmt.Println(isValid("()[]{}"))  // true
	fmt.Println(isValid("(]"))      // false
	fmt.Println(isValid("([])"))    // true
	fmt.Println(isValid("{[]}"))    // true
}
