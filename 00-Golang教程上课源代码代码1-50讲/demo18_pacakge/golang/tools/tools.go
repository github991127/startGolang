package tools

import (
	"fmt"
	"golang/calc"
)

func init() {
	fmt.Println("tools init..")
}

func PrintInfo() {
	calc.Add(10, 20)
	fmt.Println("tools print..")
}
