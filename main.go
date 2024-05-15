package main

import (
	"fmt"
	"os"
	"zephyr-go/lexer"
)

func main() {
	var contents = "var a = 2;"

	var data, err = lexer.Lex(contents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(data)
}
