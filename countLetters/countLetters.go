package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func countLetters(input string) string {
	input = strings.ToLower(input)
	letterCount := make(map[rune]int)

	for _, char := range input {
		if unicode.IsLetter(char) {
			letterCount[char]++
		}
	}

	var result []string
	for char, count := range letterCount {
		result = append(result, fmt.Sprintf("%c=%d", char, count))
	}

	return strings.Join(result, ", ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan string:")
	if scanner.Scan() {
		input := scanner.Text()
		result := countLetters(input)
		fmt.Println("Hasil penghitungan huruf:", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Kesalahan membaca input:", err)
	}
}
