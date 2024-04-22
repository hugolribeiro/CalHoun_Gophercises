package service

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"quizgame/entities"
	"time"
)

func GetQuestionsAndResult(csvFile *os.File, shuffleQuestions bool) ([]entities.QuestionAndResult, error) {
	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading records, %s", err)
		return []entities.QuestionAndResult{}, err
	}

	var actualQuestion entities.QuestionAndResult
	var allQuestionsAndResults []entities.QuestionAndResult
	for _, row := range records {
		actualQuestion = entities.QuestionAndResult{
			Question: row[0],
			Result:   row[1],
		}
		allQuestionsAndResults = append(allQuestionsAndResults, actualQuestion)
	}

	if shuffleQuestions {
		rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(allQuestionsAndResults), func(i, j int) {
			allQuestionsAndResults[i], allQuestionsAndResults[j] = allQuestionsAndResults[j], allQuestionsAndResults[i]
		})
	}

	return allQuestionsAndResults, nil
}

func PlayQuiz(questions []entities.QuestionAndResult, score *int) int {
	for i, q := range questions {
		answer := readAnswer(i+1, q.Question)
		if answer == q.Result {
			*score++
		}
	}
	return *score
}

func readAnswer(questionNumber int, question string) string {
	var answer string
	fmt.Printf("Problem #%d: %s = ", questionNumber, question)
	fmt.Scan(&answer)
	return answer
}
