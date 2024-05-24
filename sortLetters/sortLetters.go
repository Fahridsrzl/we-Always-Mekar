package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type LetterCount struct {
	letter rune
	count  int
}

type ByCount []LetterCount

func (a ByCount) Len() int      { return len(a) }
func (a ByCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool {
	if a[i].count == a[j].count {
		if unicode.IsUpper(a[i].letter) && unicode.IsLower(a[j].letter) {
			return true
		}
		if unicode.IsLower(a[i].letter) && unicode.IsUpper(a[j].letter) {
			return false
		}
		return a[i].letter < a[j].letter
	}
	return a[i].count > a[j].count
}

func countAndSortLetters(words []string) string {
	letterCount := make(map[rune]int)

	for _, word := range words {
		for _, char := range word {
			letterCount[char]++
		}
	}

	var counts []LetterCount
	for char, count := range letterCount {
		counts = append(counts, LetterCount{letter: char, count: count})
	}

	sort.Sort(ByCount(counts))

	var result strings.Builder
	for _, lc := range counts {
		result.WriteRune(lc.letter)
	}

	return result.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string

	fmt.Println("Masukkan string (masukkan baris kosong untuk selesai):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Kesalahan membaca input:", err)
		return
	}

	result := countAndSortLetters(inputs)
	fmt.Println("Hasil pengurutan huruf:", result)
}
