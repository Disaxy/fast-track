package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

const WordRegex = `[a-zA-Zа-яА-ЯёЁ0-9]+`

var (
	ErrFilePathNotProvided = errors.New("file path argument not provided")
	ErrFilePathIsEmpty     = errors.New("file path is empty")
	ErrOpenFile            = errors.New("error opening file")
	ErrReadFile            = errors.New("error reading file")
)

type kv struct {
	key   string
	value int
}

func getFilePath() (string, error) {
	args := os.Args

	if len(args) < 2 {
		return "", ErrFilePathNotProvided
	}

	filePath := args[1]

	if strings.TrimSpace(filePath) == "" {
		return "", ErrFilePathIsEmpty
	}

	return filePath, nil
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, ErrOpenFile
	}

	return file, nil
}

func readFile(file *os.File) (map[string]int, error) {
	wordCount := make(map[string]int)
	wordRegex := regexp.MustCompile(WordRegex)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := wordRegex.FindAllString(strings.ToLower(line), -1)
		for _, w := range words {
			wordCount[w]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, ErrReadFile
	}

	return wordCount, nil
}

func sortWordCount(wordCount map[string]int) []kv {
	var sorted []kv

	for k, v := range wordCount {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].value > sorted[j].value
	})

	return sorted
}

func printTop10(sortedWordCount []kv) {
	fmt.Println("TOP-10 words:")
	for i, kv := range sortedWordCount {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s — %d\n", i+1, kv.key, kv.value)
	}
}

func main() {
	filePath, err := getFilePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := openFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	wordCount, err := readFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sorted := sortWordCount(wordCount)
	printTop10(sorted)
}
