package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runHash(text string) int {
	cur := 0
	for _, v := range text {
		cur += int(v)
		cur *= 17
		cur = cur % 256
	}
	return cur
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	text := scanner.Text()
	s := strings.Split(text, ",")
	sum := 0
	for _, v := range s {
		sum += runHash(v)
	}
	fmt.Println(sum)
}
