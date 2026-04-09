package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[[TDIWEF][RBNRT][CGFNRT]\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
    re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    for _, line := range lines {
        if re.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
    output := make([]string, 0, len(lines))
	re := regexp.MustCompile(`User\s+(\w+)`)
    for _, line := range lines {
        userMatch := re.FindStringSubmatch(line)
        if userMatch != nil {
            newLine := "[USR] " + userMatch[1] + " " + line
            output = append(output,newLine)
        } else {
            output = append(output,line)
        }
    }
    return output
}
