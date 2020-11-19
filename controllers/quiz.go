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

// @Title GetAllImageLink
// @Description get all questions
// @Param	quizid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Question
// @Failure 403 :quizid is not exist
// @router /GetAllImageLink/:quizid [get]
func (u *QuizController) GetAllImageLink() {
	name := u.GetString(":quizid")
	if name != "" {
		links, err := models.GetAllImageLinkInQuestion(name)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = links
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
// @Param	topicID		path 	string	true		"The key for staticblock"
// @Param	userID	query	string	true		"user ID"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /GetAllQuizInTopic/:topicID [get]
func (u *QuizController) GetAllQuizInTopic() {
	userID := u.GetString("userID")
	topicID := u.Ctx.Input.Param(":topicID")
	if topicID != "" {
		quiz, err := models.GetALlQuizInTopic(userID, topicID)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = quiz
		}
	}
	u.ServeJSON()
}

// @Title GetAllDoneQuizOfUser
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.DoneQuiz
// @Failure 403 :id is empty
// @router /GetAllDoneQuizOfUser/:id [get]
func (u *QuizController) GetAllDoneQuizOfUser() {
	uid := u.GetString(":id")
	if uid != "" {
		quiz, err := models.GetDoneQuizOfUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = quiz
		}
	}
	u.ServeJSON()
}

// @Title PostDoneQuiz
// @Description create users
// @Param	body		body 	models.DoneQuiz	true		"body for user content"
// @Success 200 {int} models.Quiz.Name
// @Failure 403 body is empty
// @router /PostDoneQuiz [post]
func (u *QuizController) PostDoneQuiz() {
	var quiz models.DoneQuiz
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &quiz)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	err = models.AddDoneQuiz(quiz)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title DeleteDoneQuiz
// @Description create users
// @Param	doneQuizID		path 	string	true		"done quiz id"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /DeleteDoneQuiz/:doneQuizID [post]
func (u *QuizController) DeleteDoneQuiz() {
	quizID := u.Ctx.Input.Param(":doneQuizID")
	err := models.DeleteDoneQuiz(quizID)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title GetRecentPlayedQuiz
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /GetRecentPlayedQuiz/:id [get]
func (u *QuizController) GetRecentPlayedQuiz() {
	uid := u.Ctx.Input.Param(":id")
	if uid != "" {
		quizs, err := models.GetRecentPlayedQuiz(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = quizs
		}
	}
	u.ServeJSON()
}

// @Title GetRecommendedQuiz
// @Description get user by uid
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 :id is empty
// @router /GetRecommendedQuiz/:id [get]
func (u *QuizController) GetRecommendedQuiz() {
	uid := u.Ctx.Input.Param(":id")
	if uid != "" {
		quizs, err := models.GetRecommendQuiz(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = quizs
		}
	}
	u.ServeJSON()
}

// @Title Search
// @Description get user by uid
// @Param	key		query 	string	true		"The key for staticblock"
// @Success 200 {object} models.Quiz
// @Failure 403 key is empty
// @router /SearchQuiz [get]
func (u *QuizController) Search() {
	key := u.GetString("key")
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

// @Title StartQuiz
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Success 200 {string} update done!
// @Failure 403 :uid is not int
// @router /StartQuiz/:uid [put]
func (u *QuizController) StartQuiz() {
	uid := u.GetString(":uid")
	if uid != "" {
		err := models.StartQuiz(uid)
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
	err := models.DeleteQuiz(uid)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
