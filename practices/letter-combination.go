package practices

import "strings"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
// medium
func LetterCombinations(digits string) []string {
	letters := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}

	queue := []string{""}
	res := []string{}

	for len(queue) > 0 {
		comb := queue[0]
		queue = queue[1:] // dequeue

		if len(comb) == len(digits) {
			res = append(res, comb)
		} else {
			curDigit := strings.Split(digits, "")[len(comb)]

			for _, char := range letters[curDigit] {
				queue = append(queue, comb+char)
			}
		}

	}

	return res
}
