package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"quiz/temp"
	"reflect"
	"strconv"
)

var (
	QuizList map[string]*Quiz
	QuestList map[string][]*Question
)

const quiz = "quiz"

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

func AddQuiz(q Quiz) string {
	_, err := client.Collection(quiz).Doc(q.Name).Set(ctx, map[string]interface{}{
		"Name":             q.Name,
		"NumberOfQuestion": q.NumberOfQuestion,
		"Creator":          q.Creator,
	})
	if err != nil {
		log.Printf("Failed adding alovelace: %v \n", err)
	}
	return q.Name
}

func AddQuestions(name string, questions []Question) string {
	doc, err := client.Collection(quiz).Doc(name).Get(ctx)
	if err != nil {
		return "quiz not found"
	}
	list, err := doc.DataAt("Question")
	if err != nil {
		return err.Error()
	}
	length := 0
	v := reflect.ValueOf(list)
	if v.Kind() == reflect.Map {
		length = len(v.MapKeys())
	}
	for i, v := range questions {
		num := strconv.Itoa(length + i)
		_, err := client.Collection(quiz).Doc(name).Set(ctx, map[string]interface{}{
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

func GetQuizInfo(ctx context.Context, ref *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	doc, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetQuiz(name string) (u *Quiz, err error) {
	ref := client.Collection(quiz).Doc(name)
	if doc, err := GetQuizInfo(ctx, ref); err == nil {
		var q Quiz
		name, err := doc.DataAt("Name")
		if err != nil {
			return nil, err
		}
		creator, err := doc.DataAt("Creator")
		if err != nil {
			return nil, err
		}
		numberOfQuestion, err := doc.DataAt("NumberOfQuestion")
		if err != nil {
			return nil, err
		}
		q.Name = fmt.Sprint(name)
		q.Creator = fmt.Sprint(creator)
		q.NumberOfQuestion, err = strconv.Atoi(fmt.Sprint(numberOfQuestion))
		if err != nil {
			return nil,err
		}
		return &q, nil
	}
	return nil, errors.New("quiz not exists")
}

func GetAllQuiz() map[string]*Quiz {
	list := client.Collection(quiz).Documents(ctx)
	quizzes := make(map[string]*Quiz)
	for{
		var q Quiz
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		name, err := doc.DataAt("Name")
		if err != nil {
			return nil
		}
		creator, err := doc.DataAt("Creator")
		if err != nil {
			return nil
		}
		numberOfQuestion, err := doc.DataAt("NumberOfQuestion")
		if err != nil {
			return nil
		}
		q.Name = fmt.Sprint(name)
		q.Creator = fmt.Sprint(creator)
		q.NumberOfQuestion, err = strconv.Atoi(fmt.Sprint(numberOfQuestion))
		if err != nil {
			return nil
		}
		quizzes[q.Name] = &q
	}
	return quizzes
}

func GetAllQuestion(quizName string) ([]*temp.QuestionUpdate, error) {
	ref := client.Collection(quiz).Doc(quizName)
	if doc, err := GetQuizInfo(ctx, ref); err == nil {
		var questions []*temp.QuestionUpdate
		list, err := doc.DataAt("Question")
		v := reflect.ValueOf(list)
		if err != nil {
			return nil, err
		}
		if v.Kind() == reflect.Map {
			for _, key := range v.MapKeys() {
				var q temp.QuestionUpdate
				value := v.MapIndex(key).Interface().(map[string]interface{})
				q.Answer = int(value["Answer"].(int64))
				q.Question = value["Question"].(string)
				q.Choice1 = value["Choice1"].(string)
				q.Choice2 = value["Choice2"].(string)
				q.Choice3 = value["Choice3"].(string)
				q.Choice4 = value["Choice4"].(string)
				questions = append(questions, &q)
			}
		}
		return questions, nil
	}
	return nil, errors.New("quiz not exists")
}

func UpdateQuiz(name string, q *temp.QuizUpdate) (err error) {
	ref := client.Collection(quiz).Doc(name)
	doc, err := ref.Get(ctx)
	if err != nil {
		return err
	}
	v, err := doc.DataAt("Creator")
	if err != nil {
		return err
	}
	s := fmt.Sprint(v)
	if _, err := GetQuizInfo(ctx, ref); err == nil {
		_, err = client.Collection(quiz).Doc(name).Set(ctx, map[string]interface{}{
			"Name":    name,
			"NumberOfQuestion":   q.NumberOfQuestion,
			"Creator": s,
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("quiz not exist")
}

func DeleteQuiz(name string) {
	client.Collection(quiz).Doc(name).Delete(ctx)
}
