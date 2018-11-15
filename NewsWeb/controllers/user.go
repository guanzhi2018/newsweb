package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"NewsWeb/models"
	"encoding/base64"
)

type  UserController struct {
	beego.Controller
}

//显示注册页面
func (this *UserController )ShowReg ()  {
	this.TplName = "register.html"
}

//处理注册页面 
func (this *UserController) HandlerReg()  {
	//接收数据
	userName:=this.GetString("userName")
	password:=this.GetString("password")

	//校验数据
	if userName=="" || password==""{
		this.Data["errMsg"]="用户名或者密码不能为空"
		this.TplName = "register.html"
		return
	}

	//处理数据
	//插入操作
	//获取orm对象
	o:= orm.NewOrm()
	//获取插入对象
	var user models.User
	//给插入对象赋值
	user.UserName = userName
	user.Pwd = password
	//插入
	_,err2:=o.Insert(&user)
	if err2!=nil {
		this.Data["errMsg"]="注册用户失败"
		this.TplName="register.html"
		return
	}


	//返回数据
	//this.Ctx.WriteString("插入数据库成功！")
	//跳转 状态码 302
	this.Redirect("/login",302)
	
}

//显示登录页面
func(this * UserController)ShowLogin(){
	userName:=this.Ctx.GetCookie("userName")
	//如果获取Cookie不为空 则显示用户名 并选中记住密码选择框
	if userName!="" {
		//base64解密 参数:字符串 ,返回值: []byte
		dec,_:=base64.StdEncoding.DecodeString(userName)
		this.Data["userName"] = string(dec)
		this.Data["checked"]="checked"
	}else {
		//否则必须对应把值置空
		this.Data["userName"] = ""
		this.Data["checked"]=""
	}
	this.TplName = "login.html"
}

//处理登录页面
func (this * UserController)HandlerLogin()  {
	//请求数据
	userName:=this.GetString("userName")
	password:= this.GetString("password")
	remember:= this.GetString("remember")

	//校验数据
	if  userName==""||password==""{
		this.Data["msgErr"]="用户名密码不能为空"
		this.TplName="login.html"
		return
	}
	//处理数据
	//查询操作
	//获取orm对象
	o:=orm.NewOrm()
	//获取查询对象
	var user models.User
	//给查询对象赋值
	user.UserName = userName
	//查询
	err:=o.Read(&user,"UserName")
	if err!=nil {
		this.Data["msgErr"]="用户名不存在"
		this.TplName="login.html"
		return
	}
	//查询的对象获取的密码 和 前端输入的密码进行比较
	if user.Pwd != password {
		this.Data["msgErr"]="密码不正确"
		this.TplName = "login.html"
		return
	}



	//返回数据

	//beego.Info(remember)
	//登录成功的时候 如果用户选中了记住用户名
	//把数据存到Cookie中
	if remember == "on"  {
		//base64加密 参数 : []byte ,返回值 :字符串
		enc:=base64.StdEncoding.EncodeToString([]byte(userName))
		//设置 Cookie 参数一 键 参数二 值 参数三 保存时间 秒
		this.Ctx.SetCookie("userName",enc,3600)
	}else {
		//如果没有选中 删除 Cookie 参数一 键 参数二 可以是任意值 参数三 时间设置为-1
		this.Ctx.SetCookie("userName",userName,-1)
	}

	//登录成功设置 Session
	//设置 Session 参数一 键 参数二 值
	this.SetSession("userName",userName)
	//this.Ctx.WriteString("登录成功")
	this.Redirect("/article/articlelist",302)

}
