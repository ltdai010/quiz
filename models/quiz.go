package models

import (
	"cloud.google.com/go/firestore"
	"errors"
	"log"
	"quiz/temp"
	"strconv"
)

var (
	QuizList map[string]*Quiz
	QuestList map[string][]*Question
)

func init() {
	QuizList = make(map[string]*Quiz)
	QuestList = make(map[string][]*Question)
}

type Quiz struct {
	Creator string
	Name string
	NumberOfQuestion int
}

type Question struct {
	QuizName string
	Question string
	Choice1  string
	Choice2  string
	Choice3  string
	Choice4  string
	Answer   int
}

func AddQuiz(quiz Quiz) string {
	_, err := client.Collection("quiz").Doc(quiz.Name).Set(ctx, map[string]interface{}{
		"Name":    quiz.Name,
		"NumberOfQuestion":   quiz.NumberOfQuestion,
		"Creator": quiz.Creator,
	})
	if err != nil {
		log.Printf("Failed adding alovelace: %v \n", err)
	}
	return quiz.Name
}

func AddQuestions(name string, questions []Question) string {
	for i, v := range questions {
		num := strconv.Itoa(i)
		_, err := client.Collection("quiz").Doc(name).Set(ctx, map[string]interface{}{
			"Question": map[string]interface{}{
				num: map[string]interface{}{
					"Question": v.Question,
					"Choice1" : v.Choice1,
					"Choice2" : v.Choice2,
					"Choice3" : v.Choice3,
					"Choice4" : v.Choice4,
					"Answer"  : v.Answer,
				},
			},
		}, firestore.MergeAll)
		if err != nil {
			log.Printf("Failed adding alovelace: %v \n", err)
		}
	}
	return "done"
}

func GetQuiz(name string) (u *Quiz, err error) {
	if u, ok := QuizList[name]; ok {
		return u, nil
	}
	return nil, errors.New("quiz not exists")
}

func GetAllQuiz() map[string]*Quiz {
	return QuizList
}

func GetAllQuestion(quizName string) ([]*Question, error) {
	if list, ok := QuestList[quizName]; ok {
		return list, nil
	}
	return nil, errors.New("quiz not exists")
}

func UpdateQuiz(name string, quiz *temp.QuizUpdate) (err error) {
	if q, ok := QuizList[name]; ok {
		q.NumberOfQuestion = quiz.NumberOfQuestion
	} else {
		return errors.New("quiz not found")
	}
	return nil
}

func DeleteQuiz(name string) {
	_, err := client.Collection("quiz").Doc(name).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}
