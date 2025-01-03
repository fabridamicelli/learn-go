package runlength

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func appendPart(all string, n int, s string) string {
	if n == 1 {
		return fmt.Sprintf("%s%s", all, s)
	}
	return fmt.Sprintf("%s%d%s", all, n, s)
}

func formatOutput(numbers []int, letters []rune) (string, error) {

	if len(letters) != len(numbers) {
		return "", fmt.Errorf("Wrong number of letters(%d) and numbers(%d)", len(letters), len(numbers))
	}

	out := make([]string, 0)
	for i := 0; i < len(letters); i++ {
		n := numbers[i]
		rep := 1
		if n != -1 {
			rep = n
		}
		news := strings.Repeat(string(letters[i]), rep)
		out = append(out, news)
	}

	return strings.Join(out, ""), nil
}

func RunningLengthEncode(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}

	out := ""
	runes := []rune(s)
	current := runes[0]
	count := 1
	for _, c := range runes[1:] {
		if c == current {
			count++
		} else {
			out = appendPart(out, count, string(current))
			count = 1
			current = c
		}
	}

	return appendPart(out, count, string(current))
}

func RunningLengthDecode(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	numberRunes := make([][]rune, 0)
	letters := make([]rune, 0)
	runes := []rune(s)

	currentNumber := make([]rune, 0)
	sawNumber := false
	for _, c := range runes {
		if unicode.IsNumber(c) {
			currentNumber = append(currentNumber, c)
			sawNumber = true
		} else if unicode.IsLetter(c) {
			letters = append(letters, c)
			if sawNumber {
				numberRunes = append(numberRunes, currentNumber)
				currentNumber = make([]rune, 0)
				sawNumber = false
			} else {
				numberRunes = append(numberRunes, nil)
				currentNumber = make([]rune, 0)
			}
		}
	}
	numbers := make([]int, 0)
	for _, numrunes := range numberRunes {
		if numrunes == nil {
			numbers = append(numbers, -1)
		} else {
			n, err := strconv.Atoi(string(numrunes))
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, n)
		}
	}
	out, err := formatOutput(numbers, letters)
	if err != nil {
		return "", err
	}

	return out, nil

}
