package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/rand"
	"errors"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"math/big"
	"strconv"
)

var (
	client *firestore.Client
	ctx    context.Context
)

const host = "host"

type Host struct {
	QuizID         string
	Started		   bool
	Owner          string
	MapParticipant map[string]string
	MapScore       map[string]int
}

func generateCode() int {
	var v int
	for {
		r, _ := rand.Int(rand.Reader, big.NewInt(899999))
		x := r.Int64()
		x += 100000
		v = int(x)
		if isExist(v) != false {
			continue
		}
		break
	}
	return v
}

func isExist(code int) bool {
	ref := client.Collection(host).Doc(strconv.Itoa(code))
	_, err := GetHostInfo(ctx, ref)
	if err != nil {
		return false
	}
	return true
}

func init() {
	ctx = context.Background()
	sa := option.WithCredentialsFile("account/quiz-010-adafd5469f01.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

}

func AddHost(ht Host) string {
	code := generateCode()
	s := fmt.Sprint(code)
	ht.Started = false
	_, err := client.Collection(host).Doc(s).Set(ctx, ht)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return s
}


func GetHostInfo(ctx context.Context, ref *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	doc, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetOne(code int) (object *Host, err error) {
	ref := client.Collection(host).Doc(strconv.Itoa(code))
	if doc, err := GetHostInfo(ctx, ref); err == nil {
		var h Host
		err := doc.DataTo(&h)
		if err != nil {
			return nil, errors.New("wrong type respond")
		}
		return &h, nil
	}
	return nil, errors.New("code not exist")
}

func GetAllHost() map[string]*Host {
	list := client.Collection(host).Documents(ctx)
	hosts := make(map[string]*Host)
	for {
		var h Host
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		err = doc.DataTo(&h)
		if err != nil {
			return nil
		}
		hosts[doc.Ref.ID] = &h
	}
	return hosts
}

func Update(code int, h *Host) (err error) {
	sCode := strconv.Itoa(code)
	ref := client.Collection(host).Doc(sCode)
	if _, err := GetHostInfo(ctx, ref); err == nil {
		_, err = client.Collection(host).Doc(sCode).Set(ctx, h, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("code not exist")
}

func PostScore(code string, score int, userID string) error {
	_, err := client.Collection(host).Doc(code).Set(ctx, map[string]interface{}{
		"MapScore": map[string]int{
			userID: score,
		},
	}, firestore.MergeAll)
	return err
}

func StartGame(code string) error {
	_, err := client.Collection(host).Doc(code).Set(ctx, map[string]interface{}{
		"Started" : true,
	})
	return err
}

func JoinHost(code string, userID string) error {
	doc, err := client.Collection(USER).Doc(userID).Get(ctx)
	if err != nil {
		return err
	}
	u := User{}
	err = doc.DataTo(&u)
	if err != nil {
		return err
	}
	_, err = client.Collection(host).Doc(code).Set(ctx, map[string]interface{}{
		"MapParticipant": map[string]string{
			userID: u.UserName,
		},
	}, firestore.MergeAll)
	return err
}

func Delete(Code int) {
	client.Collection(host).Doc(strconv.Itoa(Code)).Delete(ctx)
}
