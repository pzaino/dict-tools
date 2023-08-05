/*************************************************
 *
 *    File: extr-alphabet.go
 * purpose: Extracts the alphabet from a text file
 *  author: Paolo Fabio Zaino
 * license: GPLv2
 *
 *************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <src_file> [dst_file]")
		return
	}

	srcFile := os.Args[1]
	dstFile := ""

	if len(os.Args) > 2 {
		dstFile = os.Args[2]
	}

	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println("Error reading the source file:", err)
		return
	}

	contents := string(content)
	words := strings.Split(contents, "")

	// Sort the words and remove duplicates
	sort.Strings(words)
	uniqueWords := unique(words)

	if dstFile != "" {
		// Write the sorted unique words to the destination file
		err = ioutil.WriteFile(dstFile, []byte(strings.Join(uniqueWords, "\n")), 0644)
		if err != nil {
			fmt.Println("Error writing to the destination file:", err)
			return
		}
	} else {
		// Print the sorted unique words to the console
		fmt.Println(strings.Join(uniqueWords, "\n"))
	}
}

func unique(words []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueWords []string

	for _, word := range words {
		if !uniqueMap[word] {
			uniqueMap[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}

	return uniqueWords
}
