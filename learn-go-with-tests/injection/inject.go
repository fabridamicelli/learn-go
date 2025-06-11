package main

import (
	"fmt"
	"io"
)

func Greet(name string, writer io.Writer) {
	fmt.Fprintf(writer, "Hello, %s", name)

}
