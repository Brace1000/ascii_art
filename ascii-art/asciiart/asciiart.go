package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// readFile reads the content of a file and returns an array of strings representing each line
func readFile(fileName string) []string {
	wortArt, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Split content into lines
	arrWordArt := strings.Split(string(wortArt), "\n")
	if fileName == "thinkertoy.txt" {
		arrWordArt = strings.Split(string(wortArt), "\r\n")
	}

	return arrWordArt
}

// formatString formats a string by replacing "\n" with "\\n" and splitting it into an array of strings
// If the parameter is "\n", return a newline character
func formatString(param string) []string {
	if param == "\\n" {
		return []string{"\n"}
	}

	// Replace occurrences of "\n" with "\\n"
	// Split the formatted string into an array of strings based on "\\n"
	strWord := strings.ReplaceAll(param, "\n", "\\n")
	strWord = strings.ReplaceAll(strWord, "\\t", "    ")
	words := strings.Split(strWord, "\\n")

	return words
}

// CreateStringart generates ASCII art from a given array of words and word art
// If there's only one word which is a newline character, return a newline
// If there's only one word which is empty, return an empty string
func CreateStringart(words, arrWordArt []string) string {
	if len(words) == 1 && words[0] == "\n" {
		return "\n"
	}
	if len(words) == 1 && words[0] == "" {
		return ""
	}

	// Initialize lineCount to determine the number of lines to print for each word
	// Initialize the ASCII art string
	// Iterate through each line of the word art
	// Iterate through each character of the word
	// Calculate the index of the character in the word art
	// Check if the index is within the range of the word art array
	// Check if the character is printable
	lineCount := 0
	strArt := ""

	for _, word := range words {
		if word == "" {
			lineCount = 1
		} else {
			lineCount = 8
		}
		for i := 1; i <= lineCount; i++ {

			for _, v := range word {
				if v > 126 {
					fmt.Println("The parameter includes a non-printable character")
					os.Exit(0)
				}
				index := 9*(v-32) + rune(i)
				if int(index) > len(arrWordArt)-1 {
					fmt.Println("Some character(s) cannot be found in the banner file")
					return ""
				}
				strArt += arrWordArt[index]
			}
			strArt += "\n"

		}
	}
	strArt = trimString(strArt)
	return strArt
}
// trimString removes any leading or trailing newline characters from a string
/// Iterate through each character and corresponding index in the input string
// Initialize an empty string to store the trimmed result
 // If the current character is not a newline character
 // After iterating through the string, reverse the loop direction
 // Append the character to the trimmed string
func trimString(str string) string {
	strTrim := ""

	for i, ch := range str {
		if ch != '\n' {
			return str
		}
		if i != len(str)-1 {
			strTrim += string(ch)
		}
	}
	return strTrim
}

// Asciiart is the main function of the program that generates ASCII art from command-line arguments
// If three arguments are provided, use the third argument as the file name
// Generate ASCII art using the formatted input string and word art
// Read the content of the banner file
// Format the input string (handle newlines)
func Asciiart() string {
	param := os.Args
	if len(param) < 2 || len(os.Args) > 3 {
		fmt.Println("The program expects two and not more than three arguments")
		os.Exit(0)
	}

	fileName := "standard.txt"
	if len(param) == 3 {
		fileName = param[2]
	}
	arrWordArt := readFile(fileName)
	words := formatString(param[1])
	strArt := CreateStringart(words, arrWordArt)

	return strArt
}
