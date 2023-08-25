package main

import (
	"fmt"
	"os"

	"main.go/pkg"
)

func main() {
	var str string
	fmt.Println("Введите строку, содержащую скобки (),[] или {}")
	fmt.Scan(&str)
	isBalanced, err := pkg.IsStrBalanced(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isBalanced {
		fmt.Println("Скобки сбалансированы")
	} else {
		fmt.Println("Скобки несбалансированы")
	}
}
