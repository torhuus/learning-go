package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	validPattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`

	re := regexp.MustCompile(validPattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {

	if text == "" {
		return []string{""}
	}

	validPattern := `<[~*=-]*>`

	re := regexp.MustCompile(validPattern)

	sections := re.Split(text, -1)

	var result []string

	for _, section := range sections {
		trimmed := strings.TrimSpace(section)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

func CountQuotedPasswords(lines []string) int {
	passwords := 0

	validPattern := `"(?i)([^"]*?\bpassword\b[^"]*?)"`

	re := regexp.MustCompile(validPattern)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match != "" {
				passwords++
			}
		}
	}

	return passwords
}

func RemoveEndOfLineText(text string) string {
	validPattern := `end-of-line[0-9]*`

	re := regexp.MustCompile(validPattern)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	validPattern := `User \s*\w*`
	prefix := `[USR]`

	re := regexp.MustCompile(validPattern)

	for i, line := range lines {
		match := re.FindString(line)
		if match != "" {
			username := strings.Fields(match)
			lines[i] = prefix + username[1] + " " + line
			lines[i] = fmt.Sprintf("%s %s %s", prefix, username[1], line)

		}
	}
	return lines
}
