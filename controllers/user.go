package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"quiz/models"
)

// Operations about host
type UserController struct {
	beego.Controller
}

// @Title PostUser
// @Description create object
// @Param	body		body 	models.User	true		"The object content"
// @Success 200 {string} models.User.UserID
// @Failure 403 body is empty
// @router /PostUser [post]
func (o *UserController) PostUser() {
	var ob models.User
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	id := models.AddUser(ob)
	o.Data["json"] = map[string]string{"Id": id}
	o.ServeJSON()
}

// @Title AddPlayedQuiz
// @Description create object
// @Param	userID		path 	string	true		"The object content"
// @Param	quizID		query	string	true		"The quizID"
// @Success 200 {string} models.User.UserID
// @Failure 403 body is empty
// @router /AddPlayedQuiz/:userID [post]
func (o *UserController) AddPlayedQuiz() {
	id := o.Ctx.Input.Param(":userID")
	qid := o.GetString("quizID")
	err := models.AddPlayedQuiz(qid, id)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = "success"
	o.ServeJSON()
}

// @Title GetUser
// @Description find object by code
// @Param	code		path 	string	true		"the code you want to get"
// @Success 200 {code} models.User
// @Failure 403 :code is empty
// @router /GetUser/:code [get]
func (o *UserController) GetUser() {
	objectId := o.Ctx.Input.Param(":code")
	ob, err := models.GetUser(objectId)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = ob
	o.ServeJSON()
}

// @Title GetAllUser
// @Description get all objects
// @Success 200 {object} models.Host
// @Failure 403 empty
// @router /GetAllUser [get]
func (o *UserController) GetAllUser() {
	obs := models.GetAllUser()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title UpdateUser
// @Description update the object
// @Param	userID		path 	string	true		"The host you want to update"
// @Param	body		body 	models.User	true		"The body"
// @Success 200 {object} models.Host
// @Failure 403 :hostId is empty
// @router /UpdateUser/:userID [put]
func (o *UserController) UpdateUser() {
	objectId := o.Ctx.Input.Param(":userID")
	var ob models.User
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := models.UpdateUser(objectId, ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title DeleteUser
// @Description delete the object
// @Param	Id		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /DeleteUser/:Id [delete]
func (o *UserController) DeleteUser() {
	objectId := o.Ctx.Input.Param(":Id")
	err := models.DeleteUser(objectId)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

