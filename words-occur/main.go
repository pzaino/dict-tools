/*************************************************
 *
 *    File: words-occur.go
 * purpose: Calculate the occurrences and percentage
 *          of each word in a text file
 *  author: Paolo Fabio Zaino
 * license: GPLv2
 *
 *************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/timtadh/getopt"
)

// Globals
var (
	verbose bool
)

func display_usage() {
	fmt.Println("Usage: words-occur -h|-i <input_file>|-i <input_file> -o <output_file>")
}

func main() {
	// Set defaults:
	var verbose bool = false
	var filename string = ""
	var outFileName string = ""

	// Register command-line options
	short := "hi:o:vV"
	long := []string{
		"help",
		"input",
		"output",
		"version",
		"verbose",
	}
	args, optargs, err := getopt.GetOpt(os.Args[1:], short, long)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		display_usage()
		os.Exit(1)
	}

	// Check if any arg was passed
	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "Error: No arguments passed")
		display_usage()
		os.Exit(1)
	}

	// Process command-line options
	for _, oa := range optargs {
		switch oa.Opt() {
		case "-h", "--help":
			display_usage()
		case "-i", "--input":
			filename = oa.Arg()
		case "-o", "--output":
			outFileName = oa.Arg()
			fmt.Println("Output file: ", outFileName)
		case "-v", "--verbose":
			verbose = true
			fmt.Println("Verbose mode enabled: ", verbose)
		case "-V", "--version":
			fmt.Println("words-occur v1.0.0")
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "Unknown option:", oa.Opt())
			display_usage()
			os.Exit(1)
		}
	}

	// Check if the input file was specified
	if filename == "" {
		fmt.Fprintln(os.Stderr, "Error: Input file not specified")
		display_usage()
		os.Exit(1)
	}

	// Check if the input file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Error: %s does not exist\n", filename)
		os.Exit(1)
	}

	// Read the entire file into memory
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Convert the data to string
	fileContents := string(data)

	// Remove newlines and extra spaces
	fileContents = strings.ReplaceAll(fileContents, "\n", " ")
	fileContents = strings.TrimSpace(fileContents)

	// Split the content into words
	words := strings.Fields(fileContents)

	// Count the occurrences of each substring
	substringCount := make(map[string]int)
	totalWords := len(words)

	for _, word := range words {
		substringCount[word]++
	}

	// Calculate the percentage of each substring
	substringPercentage := make(map[string]float64)

	for substring, count := range substringCount {
		substringPercentage[substring] = float64(count) / float64(totalWords) * 100
	}

	// Print the results
	fmt.Printf("Occurrences and Percentage of Substrings in %s:\n", filename)
	for substring, count := range substringCount {
		percentage := substringPercentage[substring]
		fmt.Printf("%s: Occurrences: %d, Percentage: %.2f%%\n", substring, count, percentage)
	}
}
