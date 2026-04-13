package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]struct{})

	for _, c := range strings.ToLower(word) {
		if c == '-' || c == ' ' {
			continue
		}

		if _, ok := seen[c]; ok {
            return false
        }

		seen[c] = struct{}{}
	}
	return true
}
