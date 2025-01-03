package letterfreq

import (
	"sort"
	"unicode"
)

func doCount(t string) map[rune]int {
	count := make(map[rune]int)
	runes := []rune(t)
	for _, r := range runes {
		if unicode.IsSpace(r) {
			continue
		}
		count[r]++
	}
	return count
}

type result struct {
	count map[rune]int
	id    int
}

func LetterFreq(texts []string) []map[rune]int {
	resChan := make(chan result, len(texts))
	for id, t := range texts {
		go func() {
			partialCount := doCount(t)
			resChan <- result{id: id, count: partialCount}
		}()
	}
	res := make([]result, 0, len(texts))
	for range texts {
		res = append(res, <-resChan)
	}
	close(resChan)
	// Return sorted results for easier testing
	sort.Slice(res, func(i, j int) bool {
		return res[i].id < res[j].id
	})
	out := make([]map[rune]int, 0)
	for _, v := range res {
		out = append(out, v.count)
	}
	return out

}
