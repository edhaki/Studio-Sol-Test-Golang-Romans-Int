package romanos

import (
	"regexp"
	"strings"
)

var MapRomans = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	total := 0
	for i := range s {
		value := MapRomans[s[i]]

		if i > 0 && MapRomans[s[i-1]] < value {
			total -= MapRomans[s[i-1]] * 2
		}
		total += value
	}

	return total
}

func RomanNumerals(text string) (string, int) {
	var romans = findRomans(text)
	var max_roman = ""
	var max = 0

	for _, roman := range romans {
		value := RomanToInt(roman)
		if value > max {
			max_roman = roman
			max = value
		}
	}
	return max_roman, max
}

func findRomans(text string) []string {
	if text == "" {
		return []string{}
	}
	romans_list := []string{}

	roman := strings.ToUpper(text)
	r := regexp.MustCompile("M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})")
	regexGroups := r.FindAllStringSubmatch(roman, -1)

	for i := 1; i < len(regexGroups); i++ {
		romans_list = append(romans_list, regexGroups[i][0])
	}
	return romans_list
}
