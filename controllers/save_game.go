package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"quiz/models"
)

// Operations about save game
type SaveGameController struct {
	beego.Controller
}

// @Title Post
// @Description create object
// @Param	body		body 	models.SaveGame	true		"The object content"
// @Success 200 {string} models.SaveGame
// @Failure 403 body is empty
// @router /PostSaveGame [post]
func (o *SaveGameController) Post() {
	var ob models.SaveGame
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	log.Println(err)
	id := models.AddSaveGame(ob)
	o.Data["json"] = map[string]string{"Id": id}
	o.ServeJSON()
}

// @Title Get
// @Description find object by code
// @Param	code		path 	string	true		"the code you want to get"
// @Success 200 {code} models.SaveGame
// @Failure 403 :code is empty
// @router /GetSaveGame/:code [get]
func (o *SaveGameController) Get() {
	objectId := o.Ctx.Input.Param(":code")
	ob, err := models.GetSaveGame(objectId)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = ob
	}
	o.ServeJSON()
}

// @Title GetSaveGameByUserQuiz
// @Description find savegame
// @Param	userID		query 	string	true		"the user you want to get"
// @Param	quizID		query 	string	true		"the quiz you want to get"
// @Success 200 {object} models.SaveGame
// @Failure 403 :bad request
// @router /GetSaveGameByUserQuiz/ [get]
func (o *SaveGameController) GetSaveGameByUserQuiz() {
	userID := o.GetString("userID")
	quizID := o.GetString("quizID")
	ob, err := models.GetSaveGameByUserQuiz(userID, quizID)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = ob
	}
	o.ServeJSON()
}

// @Title GetAllOfUser
// @Description get all objects
// @Param	userID	path	string	true	"The userID"
// @Success 200 {object} models.SaveGame
// @Failure 403 :userID is empty
// @router /GetAllSaveGameOfUser/:userID [get]
func (o *SaveGameController) GetAllOfUser() {
	id := o.Ctx.Input.Param(":userID")
	obs, err := models.GetAllSaveGameByUser(id)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title DeleteSaveGame
// @Description delete the object
// @Param	id		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /DeleteSaveGame/:id [delete]
func (o *SaveGameController) DeleteSaveGame() {
	objectId := o.Ctx.Input.Param(":id")
	err := models.DeleteSaveGame(objectId)
	if err != nil {
		o.Ctx.Input.Param("id")
		return
	}
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

// @Title UpdateSaveGame
// @Description update the object
// @Param	savegameID		path 	string			true		"The host you want to update"
// @Param	body			body 	[]int	true		"The body"
// @Success 200 {string} success
// @Failure 403 :savegameID is empty
// @router /UpdateSaveGame/:savegameID [put]
func (o *SaveGameController) UpdateSaveGame() {
	objectId := o.Ctx.Input.Param(":savegameID")
	var list []int
	json.Unmarshal(o.Ctx.Input.RequestBody, &list)
	err := models.UpdateSaveGame(objectId, list)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}