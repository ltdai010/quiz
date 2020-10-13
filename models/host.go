package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/rand"
	"errors"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"math/big"
	"quiz/temp"
	"strconv"
)

var (
	client *firestore.Client
	ctx context.Context
	sa  option.ClientOption
)

const host = "host"

type Host struct {
	Name 	   			string
	Code       			int
	NumberOfParticipant int
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
	_, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
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

func AddHost(ht temp.HostUpdate) int {
	var h Host
	h.Code = generateCode()
	h.Name = ht.Name
	h.NumberOfParticipant = ht.NumberOfParticipant
	s := strconv.Itoa(h.Code)
	_, err := client.Collection(host).Doc(s).Set(ctx, map[string]interface{}{
		"Name":    h.Name,
		"NumberOfParticipant":   h.NumberOfParticipant,
		"Code": h.Code,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return h.Code
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

func GetAllHost() map[int]*Host {
	list := client.Collection(host).Documents(ctx)
	hosts := make(map[int]*Host)
	for{
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
		hosts[h.Code] = &h
	}
	return hosts
}

func Update(code int, h *temp.HostUpdate) (err error) {
	sCode := strconv.Itoa(code)
	ref := client.Collection(host).Doc(sCode)
	if _, err := GetHostInfo(ctx, ref); err == nil {
		_, err = client.Collection(host).Doc(sCode).Set(ctx, map[string]interface{}{
			"Name":    h.Name,
			"NumberOfParticipant":   h.NumberOfParticipant,
			"Code": code,
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("code not exist")
}

func Delete(Code int) {
	client.Collection(host).Doc(strconv.Itoa(Code)).Delete(ctx)
}

