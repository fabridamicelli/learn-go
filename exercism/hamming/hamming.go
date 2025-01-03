package hamming

import "fmt"

func Hamming(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("Incompatible inputs, len(a)=%d != len(b)=%d", len(a), len(b))
	}
	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
