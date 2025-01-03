package letterfreq

import (
	"sort"
	"sync"
	"unicode"
)

func doCount(t string, wg *sync.WaitGroup) map[rune]int {
	defer wg.Done()
	count := make(map[rune]int)
	runes := []rune(t)
	for _, r := range runes {
		if unicode.IsSpace(r) {
			continue
		}
		if _, ok := count[r]; ok {
			count[r]++
		} else {
			count[r] = 1
		}
	}
	return count
}

func arrayHasRune(r rune, arr []rune) bool {
	for _, v := range arr {
		if v == r {
			return true
		}
	}
	return false
}

type Result struct {
	count map[rune]int
	id    int
}

func LetterFreq(texts []string) []map[rune]int {

	var wg sync.WaitGroup
	wg.Add(len(texts))
	res := make([]Result, 0)
	for id, t := range texts {
		go func() {
			partialCount := doCount(t, &wg)
			res = append(res, Result{id: id, count: partialCount})
		}()
	}
	wg.Wait()
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
