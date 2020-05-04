package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Fika", "Handani"}
	printMessage("Hello", names)
}
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ") // strings.Join adalah sebuah package yang sudah disediakan dalam "strings"
	fmt.Println(message, nameString)
}
