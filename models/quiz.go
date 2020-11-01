package models

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
	"quiz/temp"
	"reflect"
	"strconv"
)

var (
	clientS *search.Client
	index	*search.Index
	bucket  *storage.BucketHandle
)

const QUIZ = "quiz"

func init() {
	clientS = search.NewClient("75E8OZCPI1", "f5c16b7cfd3f10ac4841cddcd762acb4")
	index	= clientS.InitIndex("quiz")
	config := &firebase.Config{
		StorageBucket: "quiz-010.appspot.com",
	}
	ctx := context.Background()
	sa := option.WithCredentialsFile("account/quiz-010-adafd5469f01.json")
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		log.Fatalln(err)
	}
	clientStorage, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	bucket, err = clientStorage.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
}

type Quiz struct {
	Creator 		 string
	Name 			 string
	NumberOfQuestion int
	Counter			 int
}

type Quiz_ struct {
	ObjectID string `json:"objectID"`
	ID 		 string
	Creator string
	Name string
	NumberOfQuestion int
}

type Counter struct {
	counter int
}

type PriorityQuiz struct {
	UserID	string
	PlayedQuiz map[string]string
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
	s, _, err := client.Collection(QUIZ).Add(ctx, map[string]interface{}{
		"Name":             q.Name,
		"NumberOfQuestion": q.NumberOfQuestion,
		"Creator":          q.Creator,
		"Question": map[string]interface{}{
		},
		"Counter": 0,
	})
	if err != nil {
		return err.Error()
	}
	q_ := Quiz_{
		ObjectID: s.ID,
		Name: q.Name,
		NumberOfQuestion: q.NumberOfQuestion,
		Creator: q.Creator,
	}
	_, err = index.SaveObject(q_)
	if err != nil {
		log.Printf("Failed adding alovelace: %v \n", err)
	}
	return s.ID
}

func GetRecentPlayedQuiz(userID string) (map[string]*Quiz, error) {
	doc, err := client.Collection(PRIORITY).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var p PriorityQuiz
	err = doc.DataTo(&p)
	if err != nil {
		return nil, err
	}
	mapQuiz := make(map[string]*Quiz)
	for _, v := range p.PlayedQuiz {
		doc, err = client.Collection(QUIZ).Doc(v).Get(ctx)
		if err != nil {
			return nil, err
		}
		var q Quiz
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		mapQuiz[v] = &q
	}
	return mapQuiz, nil
}

func GetRecommendQuiz(userID string) (map[string]*Quiz, error) {
	doc, err := client.Collection(PRIORITY).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var p PriorityQuiz
	err = doc.DataTo(&p)
	if err != nil {
		return nil, err
	}
	mapQuiz := make(map[string]*Quiz)
	mapTopic := make(map[string]string)
	for _, v := range p.PlayedQuiz {
		query := client.Collection(topicQuiz).Where("QuizID", "==", v).Documents(ctx)
		for {
			doc, err := query.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			tq, err := doc.DataAt("TopicID")
			if err != nil {
				return nil, err
			}
			mapTopic[tq.(string)] = tq.(string)
		}
	}
	for _, v := range mapTopic {
		query := client.Collection(topicQuiz).Where("TopicID", "==", v).Documents(ctx)
		for {
			doc, err := query.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			var tq Quiz
			quizID, err := doc.DataAt("QuizID")
			if err != nil {
				return nil, err
			}
			docQuiz, err := client.Collection(QUIZ).Doc(quizID.(string)).Get(ctx)
			if err != nil {
				return nil, err
			}
			err = docQuiz.DataTo(&tq)
			mapQuiz[quizID.(string)] = &tq
		}
	}
	return mapQuiz, nil
}


func StartQuiz(quizID string) error {
	doc, err := client.Collection(QUIZ).Doc(quizID).Get(ctx)
	if err != nil {
		return err
	}
	a, err := doc.DataAt("counter")
	if err != nil {
		_, err = client.Collection(QUIZ).Doc(quizID).Set(ctx, map[string]interface{} {
			"counter" : 1,
		}, firestore.MergeAll)
		return err
	}
	i := a.(int64)
	i++
	_, err = client.Collection(QUIZ).Doc(quizID).Set(ctx, map[string]interface{} {
		"counter" : i,
	}, firestore.MergeAll)
	return err
}

func UpdateQuestion(name string, questions map[string]Question) string {
	doc, err := client.Collection(QUIZ).Doc(name).Get(ctx)
	if err != nil {
		return "QUIZ not found"
	}
	list, err := doc.DataAt("Question")
	if err != nil {
		return err.Error()
	}
	creator, err := doc.DataAt("Creator")
	if err != nil {
		return err.Error()
	}
	number, err := doc.DataAt("NumberOfQuestion")
	if err != nil {
		return err.Error()
	}
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Map {
		return "wrong type"
	}
	_, err = client.Collection(QUIZ).Doc(name).Set(ctx, map[string]interface{}{
		"Creator":			creator,
		"Name"	:			name,
		"NumberOfQuestion":	number,
		"Question": map[string]interface{}{},
	})
	for i, v := range questions {
		_, err := client.Collection(QUIZ).Doc(name).Set(ctx, map[string]interface{}{
			"Creator":			creator,
			"Name"	:			name,
			"NumberOfQuestion":	number,
			"Question": map[string]interface{}{
				i: map[string]interface{}{
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

func AddQuestions(name string, questions map[string]Question) string {
	doc, err := client.Collection(QUIZ).Doc(name).Get(ctx)
	if err != nil {
		return "quiz not found"
	}
	list, err := doc.DataAt("Question")
	if err != nil {
		return err.Error()
	}
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Map {
		return "wrong type"
	}
	for i, v := range questions {
		_, err := client.Collection(QUIZ).Doc(name).Set(ctx, map[string]interface{}{
			"Question": map[string]interface{}{
				i: map[string]interface{}{
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

func UploadFile(file multipart.File, name string) error {
	wc := bucket.Object(name).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

func GetQuizInfo(ctx context.Context, ref *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	doc, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetQuiz(name string) (u *Quiz, err error) {
	ref := client.Collection(QUIZ).Doc(name)
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
	return nil, errors.New("QUIZ not exists")
}

func GetAllQuiz() map[string]*Quiz {
	list := client.Collection(QUIZ).Documents(ctx)
	quizzes := make(map[string]*Quiz)
	for{
		var q Quiz
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil
		}
		quizzes[doc.Ref.ID] = &q
	}
	return quizzes
}

func GetAllQuestion(quizName string) (map[string]*temp.QuestionUpdate, error) {
	ref := client.Collection(QUIZ).Doc(quizName)
	if doc, err := GetQuizInfo(ctx, ref); err == nil {
		questions := make(map[string]*temp.QuestionUpdate)
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
				questions[key.String()] = &q

			}
		}
		return questions, nil
	}
	return nil, errors.New("quiz not exists")
}

func UpdateQuiz(name string, q *temp.QuizUpdate) (err error) {
	ref := client.Collection(QUIZ).Doc(name)
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
		_, err = client.Collection(QUIZ).Doc(name).Set(ctx, map[string]interface{}{
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

func SearchForQuiz(key string) (map[string]*Quiz, error) {
	res, err := index.Search(key)
	if err != nil {
		return nil, err
	}
	var qs []*Quiz_
	quizzes := make(map[string]*Quiz)
	err = res.UnmarshalHits(&qs)
	if err != nil {
		return nil, err
	}
	for _, q := range qs {
		quiz := &Quiz{
			Creator:          q.Creator,
			Name:             q.Name,
			NumberOfQuestion: q.NumberOfQuestion,
		}
		quizzes[q.ObjectID] = quiz
	}
	return quizzes, nil
}

func DeleteQuiz(name string) error {
	_, err := client.Collection(QUIZ).Doc(name).Delete(ctx)
	if err != nil {
		return err
	}
	_, err = index.DeleteObject(name)
	return err
}

func GetALlQuizInTopic(topicID string) (map[string]*Quiz, error) {
	iter := client.Collection(topicQuiz).Where("TopicID", "==", topicID).Documents(ctx)
	mapq := make(map[string]*Quiz)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var tq TopicQuiz
		err = doc.DataTo(&tq)
		if err != nil {
			return nil, err
		}
		quizDoc, err := client.Collection(QUIZ).Doc(tq.QuizID).Get(ctx)
		if err != nil {
			return nil, err
		}
		var q Quiz
		err = quizDoc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		mapq[tq.QuizID] = &q
	}
	return mapq, nil
}
