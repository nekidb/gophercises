package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func GetQuizesFromFile(fileData io.Reader) (result []Quiz) {
	r := csv.NewReader(fileData)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(record) != 2 {
			panic(errors.New(fmt.Sprintf("unsufficient length of file record: %v", record)))
		}
		question := record[0]
		answer := record[1]
		if err != nil {
			panic(err)
		}
		quiz := Quiz{question, answer}

		result = append(result, quiz)
	}

	return result
}

func AskOneQuiz(quiz Quiz, questionWriter io.Writer, answerReader io.Reader) bool {
	fmt.Fprintf(questionWriter, quiz.Question)

	var answer string
	fmt.Fscan(answerReader, &answer)
	if quiz.Answer != answer {
		return false
	}
	return true
}

// Dont understand how to test this function, so will not write it
// func WaitAndExecute(duration time.Duration, timer Timer, toExecute func()) {
// 	toExecute()
// }

type ScoreCounter struct {
	Score    uint
	MaxScore uint
}

func (sc *ScoreCounter) Print() {
	fmt.Printf("Your score: %d/%d", sc.Score, sc.MaxScore)
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timerDuration := flag.Int("timer", 30, "timer duration in seconds")

	flag.Parse()

	fileData, err := os.ReadFile(*csvFilename)
	if err != nil {
		panic(err)
	}

	quizes := GetQuizesFromFile(bytes.NewReader(fileData))

	scoreCounter := ScoreCounter{Score: 0, MaxScore: uint(len(quizes))}

	timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)
	go func() {
		<-timer.C
		scoreCounter.Print()
		os.Exit(1)
	}()

	for _, quiz := range quizes {
		result := AskOneQuiz(quiz, os.Stdout, os.Stdin)
		if result {
			scoreCounter.Score++
		}
	}

	scoreCounter.Print()
}
