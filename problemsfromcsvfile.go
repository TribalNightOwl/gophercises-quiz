package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// readProblemsFromFile reads problems from a CSV file and returns
// a slice of problems
func readProblemsFromCsvFile(filename string) []problem {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	fmt.Println(lines)
	problems := parseLines(lines)
	fmt.Println(problems)
	return problems
}

// parseLines receives multiple CSV lines and return a slice of problems
func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}
