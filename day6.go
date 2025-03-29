//Length of Last Word
func lengthOfLastWord(s string) int {
	
	s = strings.TrimSpace(s)

	lastSpaceIndex := strings.LastIndex(s, " ")

	
	return len(s) - lastSpaceIndex - 1
}