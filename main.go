package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	file, err := os.Open("file.txt") // this is my file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wordFreq := make(map[string]int)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		words := strings.Fields(line)
		for _, word := range words {
			punctuation := ".,!?\"'():;"
			cleaned := strings.ToLower(strings.Trim(word, punctuation))
			if cleaned != "" {
				wordFreq[cleaned]++
			}
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var freqList []WordCount
	for word, count := range wordFreq {
		freqList = append(freqList, WordCount{word, count})
	}

	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i].Count > freqList[j].Count
	})

	fmt.Println("Top 5 most frequent words:")
	for i := 0; i < 5 && i < len(freqList); i++ {
		fmt.Printf("%s: %d\n", freqList[i].Word, freqList[i].Count)
	}
}
