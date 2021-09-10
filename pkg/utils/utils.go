package utils

import (
	"strings"
)

func FirstLetterToUpper(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

func FirstLetterToLower(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToLower(str[:1]) + str[1:]
}

func FindStrIgnoreCaseInSlice(str string, slice []string) bool {
	for _, e := range slice {
		if strings.EqualFold(str, e) {
			return true
		}
	}

	return false
}

func SliceToStrByDelimiter(slice []string, delimiter string) string {
	res := ""
	for i, v := range slice {
		if i == 0 {
			res += v
		} else {
			res += delimiter + v
		}
	}

	return res
}
