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
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &quiz)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	name := models.AddQuiz(quiz)
	u.Data["json"] = map[string]string{"Name": name}
	u.ServeJSON()
}


// @Title PostQuestions
// @Description post questions
// @Param	id			path	string				true		"quizID"
// @Param	body		body	temp.mapQuestion	true		"body for user content"
// @Success 200 {string} models.Question.QuizName
// @Failure 403 body is empty
// @router /PostQuest/:id [post]
func (u *QuizController) PostQuestions() {
	id := u.Ctx.Input.Param(":id")
	var qt map[string]models.Question
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &qt)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	qname := models.AddQuestions(id, qt)
	u.Data["json"] = map[string]string{"quizName": qname}
	u.ServeJSON()
}

// @Title UpdateQuestion
// @Description post questions
// @Param	name		path	string				true		"name quiz"
// @Param	body		body	temp.mapQuestion	true		"body for user content"
// @Success 200 {string} models.Question.QuizName
// @Failure 403 body is empty
// @router /UpdateQuestion/:name [put]
func (u *QuizController) UpdateQuestion() {
	id := u.Ctx.Input.Param(":name")
	var qt map[string]models.Question
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &qt)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	qname := models.UpdateQuestion(id, qt)
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
// @Param	quizid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Question
// @Failure 403 :quizid is not exist
// @router /GetAllQuest/:quizid [get]
func (u *QuizController) GetAllQuest() {
	name := u.GetString(":quizid")
	if name != "" {
		users, err := models.GetAllQuestion(name)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = users
		}
	}
	u.ServeJSON()
}

// @Title PostImage
// @Description create users
// @Param	file		formData 	file	true		"image"
// @Param   name		query   string	true		"name"
// @Success 200 {string} done
// @Failure 403 body is empty
// @router /PostImage [post]
func (u *QuizController) PostImage() {
	file, _, err:=u.GetFile("file")
	name := u.GetString("name")
	if err!=nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	defer file.Close()
	err = models.UploadFile(file, name)
	if err!=nil {
		u.Ctx.WriteString("Upload failed")
	}else {
		u.Ctx.WriteString("Upload succeeded")
	}
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

// @Title GetAllQuizInTopic
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /GetAllQuizInTopic/:id [get]
func (u *QuizController) GetAllQuizInTopic() {
	uid := u.GetString(":id")
	if uid != "" {
		quiz, err := models.GetALlQuizInTopic(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = quiz
		}
	}
	u.ServeJSON()
}

// @Title Search
// @Description get user by uid
// @Param	key		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :key is empty
// @router /SearchQuiz/:key [get]
func (u *QuizController) Search() {
	key := u.GetString(":key")
	if key != "" {
		quizzes, err := models.SearchForQuiz(key)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = quizzes
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
