package models

import (
	"errors"
	"google.golang.org/api/iterator"
)

type Topic struct {
	Name string
}

type TopicQuiz struct {
	QuizID  string
	TopicID string
}

const topic = "topic"
const topicQuiz = "topicQuiz"

func AddTopic(t *Topic) string {
	docs, _, err := client.Collection(topic).Add(ctx, map[string]interface{}{
		"Name": t.Name,
	})
	if err != nil {
		return err.Error()
	}
	return docs.ID
}

func GetAllTopic() map[string]*Topic {
	list := client.Collection(topic).Documents(ctx)
	topics := make(map[string]*Topic)
	for {
		var t Topic
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		err = doc.DataTo(&t)
		if err != nil {
			return nil
		}
		topics[doc.Ref.ID] = &t
	}
	return topics
}

func AddQuizToTopic(topicQ TopicQuiz) error {
	_, _, err := client.Collection(topicQuiz).Add(ctx, map[string]interface{}{
		"QuizID":  topicQ.QuizID,
		"TopicID": topicQ.TopicID,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetTopic(topicID string) (*Topic, error) {
	ref := client.Collection(topic).Doc(topicID)
	if doc, err := GetHostInfo(ctx, ref); err == nil {
		var t Topic
		err := doc.DataTo(&t)
		if err != nil {
			return nil, errors.New("wrong type respond")
		}
		return &t, nil
	}
	return nil, errors.New("code not exist")
}
