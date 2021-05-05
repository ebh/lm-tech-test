package lm_tech_test

import "unicode"

func SearchCaseInsensitive(textToSearch, subtext string) []int {
	return searchCaseSensitive(
		toLower(textToSearch),
		toLower(subtext))
}

func searchCaseSensitive(textToSearch, subtext string) []int {
	if len(textToSearch) == 0 {
		return nil
	}

	if len(subtext) == 0 {
		return nil
	}

	var r []int

	for i := 0; i < len(textToSearch)-len(subtext)+1; i++ {
		if textToSearch[i:i+len(subtext)] == subtext {
			r = append(r, i+1)
		}
	}

	return r
}

func toLower(s string) string {
	rs := []rune(s)
	o := make([]rune, len(s))
	for i, r := range rs {
		o[i] = unicode.ToLower(r)
	}
	return string(o)
}
