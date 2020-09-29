package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/rand"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"google.golang.org/api/option"
	"log"
	"math/big"
	"quiz/temp"
	"strconv"
)

var (
	hosts map[int]*Host
	client *firestore.Client
	clientAl *search.Client
	ctx context.Context
)

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
	ref := client.Collection("host").Doc(strconv.Itoa(code))
	_, err := GetHostInfo(ctx, ref)
	if err != nil {
		return false
	}
	return true
}

func init() {
	hosts = make(map[int]*Host)
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
	//clientAl = search.NewClient("quiz", "quizapi")
	//index := clientAl.InitIndex("Name")
	//res, err := index.SaveObject([]Quiz{
	//	{Name: "math-01", NumberOfQuestion: 15},
	//})
}

func AddHost(host temp.HostUpdate) int {
	var h Host
	h.Code = generateCode()
	h.Name = host.Name
	h.NumberOfParticipant = host.NumberOfParticipant
	hosts[h.Code] = &h
	s := strconv.Itoa(h.Code)
	_, err := client.Collection("host").Doc(s).Set(ctx, map[string]interface{}{
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
	ref := client.Collection("host").Doc(strconv.Itoa(code))
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
	return hosts
}

func Update(Code int, host *temp.HostUpdate) (err error) {
	if v, ok := hosts[Code]; ok {
		v.Name = host.Name
		v.NumberOfParticipant = host.NumberOfParticipant
		return nil
	}
	return errors.New("code not exist")
}

func Delete(Code int) {
	delete(hosts, Code)
}

