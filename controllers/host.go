package controllers

import (
	"encoding/json"
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

// @Title Start
// @Description start host
// @Param	code		path 	string	true		"The object content"
// @Success 200 {string} models.Host.Name
// @Failure 403 body is empty
// @router /start/:code [put]
func (o *HostController) Start() {
	code := o.Ctx.Input.Param(":code")
	err := models.StartGame(code)
	if err != nil {
		o.Data["json"] = err
		o.ServeJSON()
		return
	}
	o.Data["json"] = "success"
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


// @Title JoinHost
// @Description create object
// @Param	code		path 	string	true		"The host code"
// @Param	userID		query	string	true		"The user code"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /:code/join [post]
func (o *HostController) JoinHost() {
	code := o.Ctx.Input.Param(":code")
	userID := o.GetString("userID")
	err := models.JoinHost(code, userID)
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
// @Param	score		query	int		true		"The code"
// @Param	userID	query	string	true		"The userID"
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
	userID := o.GetString("userID")
	err = models.PostScore(code, score, userID)
	if err != nil {
		o.Ctx.WriteString(err.Error())
		return
	}
	o.Data["json"] = "success"
	o.ServeJSON()
}
