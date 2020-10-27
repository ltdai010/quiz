package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type User struct {
	UserName	string
}

const USER = "user"

func AddUser(user User) string {
	ref, _, err := client.Collection(USER).Add(ctx, user)
	if err != nil {
		return err.Error()
	}
	return ref.ID
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
	_, err := client.Collection(USER).Doc(id).Set(ctx, user, firestore.MergeAll)
	return err
}