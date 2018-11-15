package routers

import (
	"NewsWeb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//添加路由过滤器函数
	beego.InsertFilter("/article/*",beego.BeforeExec,funcFilter)

    beego.Router("/", &controllers.MainController{})
    //注册
    beego.Router("/reg",&controllers.UserController{},"get:ShowReg;post:HandlerReg")
	//登录
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandlerLogin")
	//文章列表
	beego.Router("/article/articlelist",&controllers.ArticleController{},"get:ShowArticleList")
    }

    //过滤器函数
    var funcFilter = func(ctx *context.Context) {
		sion:=ctx.Input.Session("userName")
		if  sion==nil {
			ctx.Redirect(302,"/login")
			return
		}
	}