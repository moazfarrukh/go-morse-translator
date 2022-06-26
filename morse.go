package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//A translating map which contains morse code and its alphabet equivalent
	morse := map[string]string{".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E", "..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J", "-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O", ".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T", "..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y", "--..": "Z"}

	var input string
	fmt.Println("Enter morse code to be converted:")
	// taking user input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	// slice that stores the result of the conversion
	var result []string

	//convert input to a slice of rune for easier
	str := []rune(input)

	//starting index for each letter in morse code
	prevIdx := 0

	//using the next whitespace a letter in morse code is extracted as a substring then processed
	for i, j := range str {
		if j == ' ' || i == len(str)-1 {
			//substring of morsecode alphabet is translated through the morse map then appended to result
			result = append(result, morse[strings.TrimSpace(string(str[prevIdx:i+1]))])

			// the index after the next whitespace is set as the starting point for next alphabet
			prevIdx = i + 1
		}
		if j == ' ' && str[i+1] == ' ' {
			// when there are two white spaces in a row. a space is added between words
			result = append(result, " ")
		}

	}
	// join and display the slice of alphabets
	fmt.Printf("Translation: %s", strings.Join(result, ""))

}
