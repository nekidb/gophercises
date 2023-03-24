package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestGetQuizesFromFile(t *testing.T) {
	t.Run("one quiz in file", func(t *testing.T) {
		fileData := "5+5,10"

		got := GetQuizesFromFile(strings.NewReader(fileData))
		want := []Quiz{
			{"5+5", "10"},
		}

		assertParsedQuizes(t, got, want)
	})

	t.Run("five quizes in file", func(t *testing.T) {
		fileData := `7+3,10
1+1,2
8+3,11
1+2,3
8+6,14`

		got := GetQuizesFromFile(strings.NewReader(fileData))
		want := []Quiz{
			{"7+3", "10"},
			{"1+1", "2"},
			{"8+3", "11"},
			{"1+2", "3"},
			{"8+6", "14"},
		}

		assertParsedQuizes(t, got, want)
	})

	t.Run("question with commas", func(t *testing.T) {
		fileData := `"what 2+2, sir?",4`

		got := GetQuizesFromFile(strings.NewReader(fileData))
		want := []Quiz{
			{"what 2+2, sir?", "4"},
		}

		assertParsedQuizes(t, got, want)
	})
}

func TestAskOneQuiz(t *testing.T) {
	cases := []struct {
		testName string
		quiz     Quiz
		answer   string
		result   bool
	}{
		{"correct answer", Quiz{Question: "2+2", Answer: "4"}, "4", true},
		{"incorrect answer", Quiz{Question: "3+3", Answer: "6"}, "7", false},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			questionWriter := &bytes.Buffer{}

			answerReader := &bytes.Buffer{}
			answerReader.WriteString(c.answer)

			got := AskOneQuiz(c.quiz, questionWriter, answerReader)

			if c.quiz.Question != questionWriter.String() {
				t.Errorf("wanted %v output, but got %v", c.quiz.Question, questionWriter.String())
			}
			if got != c.result {
				t.Errorf("wanted %v, but got %v", c.result, got)
			}
		})
	}
}

// Dont understand how to actually test this function
// func TestWaitAndExecute(t *testing.T) {
// 	t.Run("run function when timer expired", func(t *testing.T) {
// 		isExecuted := false
// 		var toExecute = func() {
// 			isExecuted = true
// 		}
//
// 		WaitAndExecute(1*time.Second, time.Timer{}, toExecute)
//
// 		if isExecuted != true {
// 			t.Errorf("want %v, got %v", true, isExecuted)
// 		}
// 	})
//
// 	t.Run("testing waiting", func(t *testing.T) {
// 		durationToWait := 30 * time.Second
// 		spyTimer := &SpyTimer{}
//
// 		WaitAndExecute(durationToWait, spyTimer, func() {})
//
// 		if spyTimer.DurationWaited != durationToWait {
// 			t.Errorf("should have wait for %v, but waited for %v", durationToWaint, spyTimer.DurationWaited)
// 		}
// 	})
// }

func assertParsedQuizes(t *testing.T, got, want []Quiz) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, but got %v", want, got)
	}
}

// type SpyTimer struct {
// 	DurationWaited time.Duration
// }
//
// func (s *SpyTimer) NewTimer(duration time.Duration) {
// 	s.DurationWaited = duration
// }
