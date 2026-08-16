package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"freqchars"
)

func main() {
	topN := 0
	bar := "#"
	scale := 1
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-top":
			if i+1 < len(args) {
				if n, err := strconv.Atoi(args[i+1]); err == nil {
					topN = n
				}
				i++
			}
		case "-bar":
			if i+1 < len(args) {
				bar = args[i+1]
				i++
			}
		case "-scale":
			if i+1 < len(args) {
				if n, err := strconv.Atoi(args[i+1]); err == nil {
					scale = n
				}
				i++
			}
		}
	}
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var sb strings.Builder
	for sc.Scan() {
		sb.WriteString(sc.Text())
	}
	var items []freqchars.Item
	if topN > 0 {
		items = freqchars.Top(sb.String(), topN)
	} else {
		items = freqchars.Count(sb.String())
	}
	fmt.Print(freqchars.Bar(items, bar, scale))
}
