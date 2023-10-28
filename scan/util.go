package scan

func isAlpha(r rune) bool {
	return r >= runeALower && r <= runeZLower ||
		r >= runeAUpper && r <= runeZUpper ||
		r == runeUnderscore
}
func isDigit(r rune) bool {
	return r >= runeZero && r <= runeNine
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isDigit(r)
}
