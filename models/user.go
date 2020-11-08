package models

import (
	"google.golang.org/api/iterator"
	"strconv"
)

type User struct {
	UserName	string
}

const USER = "user"
const PRIORITY = "priority"

func AddUser(user User) string {
	ref, _, err := client.Collection(USER).Add(ctx, user)
	if err != nil {
		return err.Error()
	}
	return ref.ID
}

func AddPlayedQuiz(quizID string, userID string) error {
	doc, err := client.Collection(PRIORITY).Doc(userID).Get(ctx)
	if err != nil {
		_, err = client.Collection(PRIORITY).Doc(userID).Set(ctx, PriorityQuiz{
			UserID:        userID,
			PlayedQuiz: map[string]string{
				"0" : quizID,
			},
		})
		return err
	}
	d, err := doc.DataAt("PlayedQuiz")
	if err != nil {
		return err
	}
	mapQuizes := d.(map[string]interface{})
	temp := quizID
	for i := 0; i < 10; i++ {
		if v, ok := mapQuizes[strconv.Itoa(i)]; ok {
			temp = v.(string)
			mapQuizes[strconv.Itoa(i)] = quizID
		} else {
			mapQuizes[strconv.Itoa(i)] = temp
			break
		}
	}
	return nil
}


func GetUser(id string) (*User, error) {
	ref := client.Collection(USER).Doc(id)
	doc, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	u := &User{}
	err = doc.DataTo(u)
	return u, err
}

func GetAllUser() map[string]*User {
	list := client.Collection(USER).Documents(ctx)
	users := make(map[string]*User)
	for{
		var q User
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil
		}
		users[doc.Ref.ID] = &q
	}
	return users
}

func DeleteUser(id string) error {
	_, err := client.Collection(USER).Doc(id).Delete(ctx)
	return err
}

func UpdateUser(id string, user User) error {
	_, err := client.Collection(USER).Doc(id).Set(ctx, User{UserName: user.UserName})
	return err
}