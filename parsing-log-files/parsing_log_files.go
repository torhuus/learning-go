package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	validPattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`

	re := regexp.MustCompile(validPattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	validPattern := `^<[-=*~]*>$`

	re := regexp.MustCompile(validPattern)
	sections := re.Split(text, -1)

	var result []string
	for _, section := range sections {
		if section != "" {
			result = append(result, section)
		}
	}

	return result
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
