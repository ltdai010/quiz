package controllers

import (
	"encoding/json"
	"log"
	"quiz/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about host
type HostController struct {
	beego.Controller
}

// @Title Post
// @Description create object
// @Param	body		body 	models.Host	true		"The object content"
// @Success 200 {string} models.Host.Name
// @Failure 403 body is empty
// @router /PostHost [post]
func (o *HostController) Post() {
	var ob models.Host
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	id := models.AddHost(ob)
	o.Data["json"] = map[string]string{"Id": id}
	o.ServeJSON()
}

// @Title Get
// @Description find object by code
// @Param	code		path 	string	true		"the code you want to get"
// @Success 200 {code} models.Host
// @Failure 403 :code is empty
// @router /GetAHost/:code [get]
func (o *HostController) Get() {
	objectId := o.Ctx.Input.Param(":code")
	code, err := strconv.Atoi(objectId)
	log.Println(code)
	if err == nil {
		ob, err := models.GetOne(code)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Host
// @Failure 403 :objectId is empty
// @router /GetALlHost [get]
func (o *HostController) GetAll() {
	obs := models.GetAllHost()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	hostId		path 	string	true		"The host you want to update"
// @Param	body		body 	models.Host	true		"The body"
// @Success 200 {object} models.Host
// @Failure 403 :hostId is empty
// @router /UpdateAHost/:hostId [put]
func (o *HostController) Put() {
	objectId := o.Ctx.Input.Param(":hostId")
	code, err := strconv.Atoi(objectId)
	if err != nil {
		o.Data["json"] = err.Error()
	}
	var ob models.Host
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err = models.Update(code, &ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	Id		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /DeleteAHost/:Id [delete]
func (o *HostController) Delete() {
	objectId := o.Ctx.Input.Param(":Id")
	id, err := strconv.Atoi(objectId)
	if err != nil {
		o.Data["json"] = "wrong type"
	}
	models.Delete(id)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

// @Title AddUser
// @Description create object
// @Param	code		path 	string	true		"The host code"
// @Param	userID		query	string	true		"The user code"
// @Param	username	query	string	true		"The username"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /:code/AddUser [post]
func (o *HostController) AddUser() {
	code := o.Ctx.Input.Param(":code")
	userID := o.GetString("userID")
	username := o.GetString("username")
	err := models.AddUserToHost(code, userID, username)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = "success"
	o.ServeJSON()
}

// @Title PostScore
// @Description create object
// @Param	code		path 	string	true		"The host code"
// @Param	score		query	int		true		"The user code"
// @Param	username	query	string	true		"The username"
// @Success 200 {string} models.Host.Name
// @Failure 403 body is empty
// @router /:code/PostScore [post]
func (o *HostController) PostScore() {
	code := o.Ctx.Input.Param(":code")
	score, err := o.GetInt("score")
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	username := o.GetString("username")
	err = models.PostScore(code, score, username)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = "success"
	o.ServeJSON()
}