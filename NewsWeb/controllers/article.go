package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

//显示文章列表
func (this *ArticleController)ShowArticleList()  {
/*	//获取 Session 返回值是 接口
	sion:=this.GetSession("userName")

	//如果接口为nil 返回 登录界面重新 登录
	if sion==nil {
		this.Redirect("/login",302)
		return
	}
	//接口赋值给字符串变量 需要接口断言
	username :=sion.(string)
	this.Data["username"] = username*/

	this.TplName = "index.html"
}