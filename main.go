// Command gophercises-quiz will quiz the user
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// problem should have only one question and one answer
type problem struct {
	question string
	answer   string
}

func main() {
	cvsFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	problems := readProblemsFromCsvFile(*cvsFilename)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// exit will print a message and exit with code 1
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
