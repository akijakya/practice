package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	// the time limit is also going to be a(n integer) flag
	timeLimit := flag.Int("limit", 30, "The time limit for the for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	//this takes an io.Reader type, which our 'file' is
	r := csv.NewReader(file)
	// lines is going to be a 2D slice of the whole file, because it wont be that big
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	// a beadott timeLimit beszorozva a time.Second enum-mal adja ki, hogy ennyi másodperc legyen a timer
	// we need to convert the timelimit to a time.Duration type, because time.Second is one 
	// and they need to be the same type so the result would be the same type
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d questions!\n", correct, len(problems))
			return
		default:
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			var answer string
			//Scanf gets rid of every whitespace, so it is not good if you want to input multiple words
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d questions!\n", correct, len(problems))
}

// this is going to take in a lines 2D string slice and a problem slice
func parseLines (lines [][]string) []problem {
	// the return is going to be a slice of problems of the length of the lines input
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem {
			q: line[0],
			// we are trimming it because what if the csv is bad and there are unnecessary whitespaces,
			// and the Scanf for the answer also removes whitespace so here we are also doing that for foolproofing reasons
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// this will represent a question-answer pair
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
