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
            Router: "/GetAllQuest/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "PostQuestions",
            Router: "/PostQuest",
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
            Router: "/SearchQuiz/:id",
            AllowHTTPMethods: []string{"get"},
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

}
