package service

import (
	"os"
	"quizgame/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetQuestionsAndResult(t *testing.T) {
	var fileInput *os.File

	fileInput, _ = os.Open("../file/problems_test.csv")

	defer fileInput.Close()

	questionsAndResult, err := GetQuestionsAndResult(fileInput, false)
	expectedQuestionsAndResult := []entities.QuestionAndResult{
		{
			Question: "5+5",
			Result:   "10",
		},
		{
			Question: "2+7",
			Result:   "9",
		},
	}

	assert.Equal(t, expectedQuestionsAndResult, questionsAndResult)
	assert.NoError(t, err)
}

func Test_GetQuestionsAndResult_Error_Reader(t *testing.T) {
	var fileInput *os.File

	questionsAndResult, err := GetQuestionsAndResult(fileInput, false)
	expectedQuestionsAndResult := []entities.QuestionAndResult{}

	assert.Equal(t, expectedQuestionsAndResult, questionsAndResult)
	assert.Error(t, err, assert.AnError)
}
