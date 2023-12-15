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

type lens struct {
	length   int
	position int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	text := scanner.Text()
	s := strings.Split(text, ",")
	boxes := make(map[int]map[string]lens)
	for _, v := range s {
		last := v[len(v)-1]
		if last == '-' {
			l := v[:len(v)-1]
			h := runHash(l)
			box, ok := boxes[h]
			if ok {
				lense, ok := box[l]
				if ok {
					o := lense.position
					delete(box, l)
					for k, v := range box {
						if v.position > o {
							v.position--
							box[k] = v
						}
					}
				}
			}

		} else {
			l := v[:len(v)-2]
			h := runHash(l)
			box, ok := boxes[h]
			if ok {
				lense, ok := box[l]
				if ok {
					box[l] = lens{position: lense.position, length: int(last - '0')}
				} else {
					box[l] = lens{position: len(box), length: int(last - '0')}
				}
			} else {
				lenses := make(map[string]lens)
				lenses[l] = lens{position: 0, length: int(last - '0')}
				boxes[h] = lenses
			}
		}
	}
	sum := 0
	for k, v := range boxes {
		for _, v := range v {
			sum += (k + 1) * v.length * (v.position + 1)
		}
	}
	fmt.Println(sum)
}
