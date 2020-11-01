package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"quiz/models"
)

// Operations about topic
type TopicController struct {
	beego.Controller
}

// @Title PostTopic
// @Description create users
// @Param	body		body 	models.Topic	true		"body for user content"
// @Success 200 {int} models.Topic.Name
// @Failure 403 body is empty
// @router /PostTopic [post]
func (u *TopicController) PostTopic() {
	var topic models.Topic
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &topic)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	name := models.AddTopic(&topic)
	u.Data["json"] = map[string]string{"Name": name}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all topic
// @Success 200 {object} models.Topic
// @router /GetAll [get]
func (u *TopicController) GetAllTopic() {
	topics := models.GetAllTopic()
	u.Data["json"] = topics
	u.ServeJSON()
}

// @Title GetTopic
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Topic
// @Failure 403 :id is empty
// @router /GetTopic/:id [get]
func (u *TopicController) GetTopic() {
	id := u.GetString(":id")
	if id != "" {
		topic, err := models.GetTopic(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = topic
		}
	}
	u.ServeJSON()
}

// @Title GetAllTopicOfQuiz
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /GetAllTopicOfQuiz/:id [get]
func (u *TopicController) GetAllTopicOfQuiz() {
	uid := u.Ctx.Input.Param(":id")
	if uid != "" {
		topics, err := models.GetALlTopicOfQuiz(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = topics
		}
	}
	u.ServeJSON()
}

// @Title PostQuizToTopic
// @Description create users
// @Param	body		body 	models.TopicQuiz	true		"body for user content"
// @Success 200 {string} models.Topic.Name
// @Failure 403 body is empty
// @router /PostQuizToTopic [post]
func (u *TopicController) PostQuizToTopic() {
	var topicQ models.TopicQuiz
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &topicQ)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	err = models.AddQuizToTopic(topicQ)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Ctx.WriteString("success")
}

