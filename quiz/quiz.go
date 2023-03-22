package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

const defaultFileName = "problems.csv"

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
	fmt.Fprintln(questionWriter, quiz.Question)

	var answer string
	fmt.Fscan(answerReader, &answer)
	if quiz.Answer != answer {
		return false
	}
	return true
}

func GetFileName() string {
	if len(os.Args) == 1 {
		return defaultFileName
	}

	return os.Args[1]
}

func main() {
	fileName := GetFileName()

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	quizes := GetQuizesFromFile(bytes.NewReader(fileData))

	counter := 0
	for _, quiz := range quizes {
		result := AskOneQuiz(quiz, os.Stdout, os.Stdin)
		if result {
			counter++
		}
	}

	fmt.Printf("\nYour score: %d/%d", counter, len(quizes))
}
