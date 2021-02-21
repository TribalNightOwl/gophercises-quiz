// Package getproblems implements function to get problems from different sources
package getproblems

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

// ReadProblemsFromFile returns a struct
func ReadProblemsFromFile(filename string) []Problem {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	problems := parseLines(lines)

	return problems
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
