package models

import (
	"google.golang.org/api/iterator"
)

const SAVE_GAME = "saveGame"

type SaveGame struct {
	UserID    string
	QuizID    string
	DoneQuest int
	QuizDone  bool
	MapQuest  map[string]DoneQuest
}

type DoneQuest struct {
	ChoosedAnswer int
	Answer        int
}

func AddSaveGame(saveGame SaveGame) string {
	ref, _, err := client.Collection(SAVE_GAME).Add(ctx, saveGame)
	if err != nil {
		return err.Error()
	}
	return ref.ID
}

func GetSaveGame(id string) (*SaveGame, error) {
	ref := client.Collection(SAVE_GAME).Doc(id)
	doc, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}
	sg := &SaveGame{}
	err = doc.DataTo(sg)
	return sg, err
}

func DeleteSaveGame(id string) error {
	_, err := client.Collection(SAVE_GAME).Doc(id).Delete(ctx)
	return err
}

func GetAllSaveGameByUser(userID string) (map[string]*SaveGame, error) {
	iter := client.Collection(SAVE_GAME).Where("UserID", "==", userID).Documents(ctx)
	mapS := make(map[string]*SaveGame)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		s := &SaveGame{}
		err = doc.DataTo(s)
		if err != nil {
			return nil, err
		}
		mapS[doc.Ref.ID] = s
	}
	return mapS, nil
}