package pangram

import "strings"

func IsPangram(input string) bool {
	var seen [26]bool

	count := 0
	for _, r := range strings.ToLower(input) {
		if r < 'a' || r > 'z' {
			continue
		}
		idx := r - 'a'
		if !seen[idx] {
			seen[idx] = true
			count++
		}
		if count == 26 {
			return true
		}
	}
	return false
}
