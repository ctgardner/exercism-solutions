package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validLine      = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	separator      = regexp.MustCompile(`<[~*=-]*>`)
	quotedPassword = regexp.MustCompile(`".*(?i:password).*"`)
	endOfLine      = regexp.MustCompile(`end-of-line\d+`)
	username       = regexp.MustCompile(`User\s+(\w+)`)
)

func IsValidLine(text string) bool {
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, s := range lines {
		if quotedPassword.MatchString(s) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLine.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, 0, len(lines))
	for _, s := range lines {
		if match := username.FindStringSubmatch(s); match != nil {
			s = fmt.Sprintf("[USR] %s %s", match[1], s)
		}
		taggedLines = append(taggedLines, s)
	}
	return taggedLines
}
