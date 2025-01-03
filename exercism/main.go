package main

import (
	"fmt"
)

func main() {
	// fmt.Println(runlength.RunningLengthDecode("10a11b12C1o"))
	// fmt.Println(runlength.RunningLengthEncode("a"))
	// fmt.Println(len("⌣"))
	// fmt.Println([]byte("⌣"))
	// fmt.Println([]rune("⌣"))
	// fmt.Println(len([]rune("⌣")))
	// fmt.Println(len(string([]rune("⌣"))))
	// fmt.Println([]byte("⌣"))
	//
	// var c rune = '\u7684'
	// fmt.Println(c)
	// fmt.Println(string(c))
	// s := []byte(string(c))
	// fmt.Println(string(s))
	//
	// // 󰩱
	// point := 985713
	// fmt.Println(string(rune(point)))
	// //fabri
	//
	// fmt.Println(string([]rune{102, 97, 98, 114, 105}))
	// for _, c := range "fabri" {
	// 	fmt.Println(c)
	// }
	// runes := []string{"f3", "b0", "a9", "b1"}
	// all := make([]rune, 0)
	// for _, r := range runes {
	// 	fmt.Println(string(r))
	// 	out, _ := utf8.DecodeRune([]byte(r))
	// 	all = append(all, out)
	// }
	// fmt.Println(string(all))

	// b := "hhello"
	// s := strings.TrimLeft(b, "h")
	// fmt.Println(s)

	m := map[rune]int{'h': 1}
	s, _ := m['h']
	ss, _ := m['g']
	fmt.Println(s)
	fmt.Println(ss)

}
