package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"quiz/models"
	"quiz/temp"
)

// Operations about Quiz
type QuizController struct {
	beego.Controller
}

// @Title Post
// @Description create users
// @Param	body		body 	models.Quiz	true		"body for user content"
// @Success 200 {int} models.Quiz.Name
// @Failure 403 body is empty
// @router /PostQuiz [post]
func (u *QuizController) Post() {
	var quiz models.Quiz
	json.Unmarshal(u.Ctx.Input.RequestBody, &quiz)
	name := models.AddQuiz(quiz)
	u.Data["json"] = map[string]string{"Name": name}
	u.ServeJSON()
}


// @Title PostQuestions
// @Description post questions
// @Param	body		body	models.Question	true		"body for user content"
// @Success 200 {string} models.Question.QuizName
// @Failure 403 body is empty
// @router /PostQuest [post]
func (u *QuizController) PostQuestions() {
	var qt []models.Question
	json.Unmarshal(u.Ctx.Input.RequestBody, &qt)
	qname := models.AddQuestions(qt[0].QuizName, qt)
	u.Data["json"] = map[string]string{"quizName": qname}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Quiz
// @router /GetAll [get]
func (u *QuizController) GetAll() {
	users := models.GetAllQuiz()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title GetAllQuest
// @Description get all questions
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Question
// @Failure 403 :name is not exist
// @router /GetAllQuest/:name [get]
func (u *QuizController) GetAllQuest() {
	name := u.GetString(":name")
	if name != "" {
		users, err := models.GetAllQuestion(name)
		if err != nil {
			u.Data["json"] = users
		} else {
			u.Data["json"] = err.Error()
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /GetAQuiz/:id [get]
func (u *QuizController) Get() {
	uid := u.GetString(":id")
	if uid != "" {
		user, err := models.GetQuiz(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Search
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /SearchQuiz/:id [get]
func (u *QuizController) Search() {
	uid := u.GetString(":id")
	if uid != "" {
		user, err := models.GetQuiz(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Put
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	temp.QuizUpdate	true		"body for user content"
// @Success 200 {string} update done!
// @Failure 403 :uid is not int
// @router /UpdateQuiz/:uid [put]
func (u *QuizController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user temp.QuizUpdate
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		err := models.UpdateQuiz(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "update success"
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	qId		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 qId is empty
// @router /DeleteQuiz/:qId [delete]
func (u *QuizController) Delete() {
	uid := u.GetString(":qId")
	models.DeleteQuiz(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
