package models

import (
	"google.golang.org/api/iterator"
)

type DoneQuiz struct {
	QuizID	string
	UserID	string
	RightAns	int
	WrongAns	int
}

const DONE_QUIZ = "doneQuiz"

func AddDoneQuiz(d DoneQuiz) error {
	_, err := client.Collection(DONE_QUIZ).Doc(d.QuizID + "-" + d.UserID).Set(ctx, d)
	return err
}

func GetDoneQuizOfUser(userID string) (map[string]*DoneQuiz, error) {
	list := client.Collection(DONE_QUIZ).Where("UserID", "==", userID).Documents(ctx)
	listD := map[string]*DoneQuiz{}
	for {
		var d DoneQuiz
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		err = doc.DataTo(&d)
		if err != nil {
			return nil, err
		}
		listD[doc.Ref.ID] = &d
	}
	return listD, nil
}

func DeleteDoneQuiz(QuizID string) error {
	_, err := client.Collection(DONE_QUIZ).Doc(QuizID).Delete(ctx)
	return err
}
