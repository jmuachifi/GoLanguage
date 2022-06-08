package routers

import (
	"my-first-beego-project/controllers"

	"github.com/astaxie/beego"

	"my-first-beego-project/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{},
		"get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{},
		"get:Dashbaord")
	beego.Router("/login", &controllers.SessionController{},
		"get:Login")
	beego.Router("/logout", &controllers.SessionController{},
		"get:Logout")
	beego.Router("/home", &controllers.SessionController{},
		"get:Home")

	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)

	beego.Router("/employee", &controllers.FirstController{},
		"get:GetEmployee")

	beego.Router("/getFromCache", &controllers.CacheController{},
		"get:GetFromCache")
}
