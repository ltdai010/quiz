package models

import (
	"errors"
	"quiz/temp"
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
	Name string
	NumberOfQuestion int
	QuestionList []*int
}

type Question struct {
	Id int
	Question string
	Choice1 string
	Choice2 string
	Choice3 string
	Choice4 string
	answer int
}

func AddQuiz(quiz Quiz) string {
	QuizList[quiz.Name] = &quiz
	return quiz.Name
}

func AddQuestions(name string, questions []*Question) string {
	QuestList[name] = questions
	return name
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

func UpdateQuiz(name string, quiz *temp.QuizUpdate) (err error) {
	if u, ok := QuizList[name]; ok {
		u.NumberOfQuestion = quiz.NumberOfQuestion
		u.QuestionList = quiz.QuestionList
	} else {
		return errors.New("quiz not found")
	}
	return nil
}

func DeleteQuiz(name string) {
	delete(QuizList, name)
}
