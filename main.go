package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("index.chad")
	result := parse(string(file))

	fmt.Println(len(result))
	fmt.Println(result)
}

func parse(code string) []string {
	chunks := strings.Split(strings.TrimSpace(code), "chad\r\n")
	return chunks
}
