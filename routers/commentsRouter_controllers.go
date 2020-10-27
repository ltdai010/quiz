package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/DeleteAHost/:Id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/GetAHost/:code",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/GetALlHost",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/PostHost",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/UpdateAHost/:hostId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/DeleteQuiz/:qId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/GetAQuiz/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/GetAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetAllQuest",
            Router: "/GetAllQuest/:quizid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetAllQuizInTopic",
            Router: "/GetAllQuizInTopic/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "PostImage",
            Router: "/PostImage",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "PostQuestions",
            Router: "/PostQuest/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/PostQuiz",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/SearchQuiz/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "UpdateQuestion",
            Router: "/UpdateQuestion/:name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/UpdateQuiz/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:SaveGameController"] = append(beego.GlobalControllerRouter["quiz/controllers:SaveGameController"],
        beego.ControllerComments{
            Method: "DeleteSaveGame",
            Router: "/DeleteSaveGame/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:SaveGameController"] = append(beego.GlobalControllerRouter["quiz/controllers:SaveGameController"],
        beego.ControllerComments{
            Method: "GetAllOfUser",
            Router: "/GetAllSaveGameOfUser/:userID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:SaveGameController"] = append(beego.GlobalControllerRouter["quiz/controllers:SaveGameController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/GetSaveGame/:code",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:SaveGameController"] = append(beego.GlobalControllerRouter["quiz/controllers:SaveGameController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/PostSaveGame",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:TopicController"] = append(beego.GlobalControllerRouter["quiz/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetAllTopic",
            Router: "/GetAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:TopicController"] = append(beego.GlobalControllerRouter["quiz/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetTopic",
            Router: "/GetTopic/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:TopicController"] = append(beego.GlobalControllerRouter["quiz/controllers:TopicController"],
        beego.ControllerComments{
            Method: "PostQuizToTopic",
            Router: "/PostQuizToTopic",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:TopicController"] = append(beego.GlobalControllerRouter["quiz/controllers:TopicController"],
        beego.ControllerComments{
            Method: "PostTopic",
            Router: "/PostTopic",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:UserController"] = append(beego.GlobalControllerRouter["quiz/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/DeleteUser/:Id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:UserController"] = append(beego.GlobalControllerRouter["quiz/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAllUser",
            Router: "/GetAllUser",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:UserController"] = append(beego.GlobalControllerRouter["quiz/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: "/GetUser/:code",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:UserController"] = append(beego.GlobalControllerRouter["quiz/controllers:UserController"],
        beego.ControllerComments{
            Method: "PostUser",
            Router: "/PostUser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:UserController"] = append(beego.GlobalControllerRouter["quiz/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/UpdateUser/:userID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
