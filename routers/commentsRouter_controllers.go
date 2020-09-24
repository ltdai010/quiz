package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:code",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:HostController"] = append(beego.GlobalControllerRouter["quiz/controllers:HostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "GetAllQuest",
            Router: "/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quiz/controllers:QuizController"] = append(beego.GlobalControllerRouter["quiz/controllers:QuizController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
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

}
