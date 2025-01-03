package series

func All(n int, s string) []string {
	var subs []string

	if n >= len(s) {
		return []string{s}
	}

	for i := 0; i < len(s); i++ {
		upper := min(len(s), i+n)
		if upper-i < n {
			return subs
		}
		subs = append(subs, s[i:upper])
	}
	return subs
}
