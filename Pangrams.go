package hackerranksubmission

import "strings"

func pangrams(s string) string {
	s = strings.ToLower(s)
	dict := make(map[string]int)

	for _, value := range s {
		dict[string(value)] += 1
	}

	result := len(dict) - 1

	if result == 26 {

		return "pangram"

	} else {

		return "not pangram"
	}
}
