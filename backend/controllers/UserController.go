package controllers

import (
	"cms/backend/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//用户绑定结构
type NewUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (this *UserController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	b.Handle("GET", "/login", "Login")
}

func (this *UserController) Login(ctx iris.Context) {
	// var ctx context.Context
	videoServerice := services.NewVideoServerice()
	id := "095b6a010f1d75957f0e7a551b2526c5"
	video, _ := videoServerice.GetById(id)
	ctx.JSON(video)
	// return "profile"
}

func (this *UserController) Json(ctx iris.Context) {
	peter := NewUser{
		Firstname: "John",
		Lastname:  "Doe",
		City:      "Neither FBI knows!!!",
		Age:       25,
	}
	ctx.JSON(peter)
}

func Getprofile() string {
	return "profile"
}
