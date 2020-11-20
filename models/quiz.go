package models

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
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
	Name 			 string
	NumberOfQuestion int
	Counter			 int
}

type ResQuiz struct {
	Name 			 string
	NumberOfQuestion int
	Counter			 int
	Playing			 bool
}

type Quiz_ struct {
	ObjectID string `json:"objectID"`
	ID 		 string
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
		"NumberOfQuestion": 0,
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
		NumberOfQuestion: 0,
	}
	_, err = index.SaveObject(q_)
	if err != nil {
		log.Printf("Failed adding alovelace: %v \n", err)
	}
	return s.ID
}

func GetRecentPlayedQuiz(userID string) (map[string]*ResQuiz, error) {
	doc, err := client.Collection(PRIORITY).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var p PriorityQuiz
	err = doc.DataTo(&p)
	if err != nil {
		return nil, err
	}
	mapQuiz := make(map[string]*ResQuiz)
	for _, v := range p.PlayedQuiz {
		doc, err = client.Collection(QUIZ).Doc(v).Get(ctx)
		if err != nil {
			return nil, err
		}
		var q ResQuiz
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.Playing = false
		if Playing(userID, v) {
			q.Playing = true
		}
		mapQuiz[v] = &q
	}
	return mapQuiz, nil
}

func GetRecommendQuiz(userID string) (map[string]*ResQuiz, error) {
	doc, err := client.Collection(PRIORITY).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var p PriorityQuiz
	err = doc.DataTo(&p)
	if err != nil {
		return nil, err
	}
	mapQuiz := make(map[string]*ResQuiz)
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
			var tq ResQuiz
			quizID, err := doc.DataAt("QuizID")
			if err != nil {
				return nil, err
			}
			docQuiz, err := client.Collection(QUIZ).Doc(quizID.(string)).Get(ctx)
			if err != nil {
				return nil, err
			}
			tq.Playing = false
			if Playing(userID, quizID.(string)) {
				tq.Playing = true
			}
			err = docQuiz.DataTo(&tq)
			if err != nil {
				return nil, err
			}
			mapQuiz[quizID.(string)] = &tq
		}
	}
	return mapQuiz, nil
}

func Playing(userID string, quizID string) bool {
	mapSave, err := GetAllSaveGameByUser(userID)
	if err != nil {
		return false
	}
	for _, i := range mapSave {
		if i.QuizID == quizID {
			return true
		}
	}
	return false
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
	_, err = client.Collection(QUIZ).Doc(name).Set(ctx, map[string]interface{}{
		"NumberOfQuestion" : len(questions),
	}, firestore.MergeAll)
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
		numberOfQuestion, err := doc.DataAt("NumberOfQuestion")
		if err != nil {
			return nil, err
		}
		q.Name = fmt.Sprint(name)
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
		if err != nil {
			return nil
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

func GetAllImageLinkInQuestion(quizName string) (map[string]string, error) {
	list, err := client.Collection(QUIZ).Doc(quizName).Get(ctx)
	if err != nil {
		return nil, err
	}
	a, err := list.DataAt("Question")
	quests := a.(map[string]interface{})
	if err != nil {
		return nil, err
	}
	mapLink := map[string]string{}
	for k, _ := range quests {
		mapLink[k] = "storage.googleapis.com/quiz-010.appspot.com/" + quizName + "-" + k
	}
	return mapLink, nil
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

func SearchForQuiz(userID, key string) (map[string]*ResQuiz, error) {
	res, err := index.Search(key)
	if err != nil {
		return nil, err
	}
	var qs []*Quiz_
	quizzes := make(map[string]*ResQuiz)
	err = res.UnmarshalHits(&qs)
	if err != nil {
		return nil, err
	}
	for _, q := range qs {
		quiz := &ResQuiz{
			Name:             q.Name,
			NumberOfQuestion: q.NumberOfQuestion,
			Playing: false,
		}
		if Playing(userID, q.ObjectID) {
			quiz.Playing = true
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


func GetALlQuizInTopic(userID, topicID string) (map[string]*ResQuiz, error) {
	iter := client.Collection(topicQuiz).Where("TopicID", "==", topicID).Documents(ctx)
	mapq := make(map[string]*ResQuiz)
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
			continue
		}
		var q ResQuiz
		err = quizDoc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.Playing = false
		if Playing(userID, tq.QuizID) {
			q.Playing = true
		}
		mapq[tq.QuizID] = &q
	}
	return mapq, nil
}
