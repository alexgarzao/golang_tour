package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	res := make(map[string]int)
	
	for _, value := range tokens {
		res[value] ++
	}
	
	return res
}

func main() {
	wc.Test(WordCount)
}
