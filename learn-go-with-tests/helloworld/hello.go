package main

import (
	"fmt"
)

const (
	enHelloPrefix = "Hello, "
	deHelloPrefix = "Hallo, "
	esHelloPrefix = "Hola, "
)

var validLangs = []string{"en", "es", "de", ""}

func checkLang(lang string) (error, string) {
	for _, l := range validLangs {
		if lang == l {
			return nil, lang
		}
	}
	return fmt.Errorf("Invalid language %s", lang), ""
}

func Hello(name string, lang string) string {
	err, lang := checkLang(lang)
	if err != nil {
		panic(err)
	}
	if name == "" {
		name = "world"
	}

	var prefix string
	switch lang {
	case "es":
		prefix = esHelloPrefix
	case "de":
		prefix = deHelloPrefix
	case "en":
		prefix = enHelloPrefix
	case "":
		prefix = enHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Fab", ""))
}
