package main

import (
	"flag"
	"fmt"
	"os"
	"quizgame/service"
	"strings"
	"time"
)

func main() {
	if !promptToStart() {
		os.Exit(0)
	}

	var score = 0

	filePath := flag.String("path", "file/problems.csv", "a csv file path in the format of queston,answer")
	programDuration := flag.Int("duration", 5, "program duration in seconds")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file, %s", err)
		os.Exit(1)
	}
	defer file.Close()

	questionsAndResults, err := service.GetQuestionsAndResult(file)
	if err != nil {
		fmt.Printf("Error getting questions and results, %s", err)
		os.Exit(1)
	}

	go time.AfterFunc((time.Duration(*programDuration) * time.Second), func() {
		showScore(score, len(questionsAndResults))
		os.Exit(0)
	})

	score = service.PlayQuiz(questionsAndResults, &score)

	showScore(score, len(questionsAndResults))
}

func promptToStart() bool {
	var keyToStart string
	fmt.Println("Do you want to start? Press (Y) to start")
	fmt.Scan(&keyToStart)
	return strings.ToLower(keyToStart) == "y"
}

func showScore(score, amountQuestions int) {
	fmt.Printf("\nYou scored: %d out of %d", score, amountQuestions)
}
