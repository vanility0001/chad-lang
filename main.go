package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		File string
	}

	arg.MustParse(&args)

	file, _ := ioutil.ReadFile(args.File)
	result := parse(string(file))

	fmt.Println(len(result))
	fmt.Println(result)
}

func parse(code string) []string {
	chunks := strings.Split(strings.TrimSpace(code), "chad\r\n")
	return chunks
}
