/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
)

//注册
func register(ctx *gin.Context) {
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")

	//检验用户名是否含有敏感词
	flag:=tool.CheckIfSensitive(username)
	if flag{
		tool.RespSensitiveError(ctx)
		return
	}

	user:=model.User{
		Username: username,
		Password: password,
	}
	flag,err:=service.IsExistUsername(username)
	if err!=nil{
		fmt.Println("judge exist username err:",err)
		tool.RespInternalError(ctx)
		return
	}

	if flag{
		tool.RespErrorWithData(ctx,"用户名已存在")
		return
	}
	flag=service.IsPasswordReasonable(password)
	if !flag{
		tool.RespErrorWithData(ctx,"密码不合理")
		return
	}

	err=service.Register(user)
	if err!=nil{
		fmt.Println("judge repeat username err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx,"注册成功")
}

//登陆
func login(ctx *gin.Context)  {
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")

	flag,err:=service.IsPasswordCorrect(username,password)
	if err!=nil{
		fmt.Println("judge password correct err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag{
		tool.RespErrorWithData(ctx,"密码错误")
		return
	}

	//设置cookie
	ctx.SetCookie("username",username,3600,"/","",false,false)

	tool.RespSuccessfulWithData(ctx,"欢迎您~"+username)
}

//登陆后密保更改
func changeSecurity(ctx *gin.Context) {
	answer:=ctx.PostForm("security_answer")
	iUsername,_:=ctx.Get("iUsername")
	username:=iUsername.(string)
	err:=service.UpdateSecurity(username,answer)
	if err!=nil{
		fmt.Println("update security err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

//登陆后修改密码
func changePassword(ctx *gin.Context) {
	oldPassword:=ctx.PostForm("old_password")
	newPassword:=ctx.PostForm("new_password")
	iUsername,_:=ctx.Get("iUsername")
	username:=iUsername.(string)

	//检验旧密码是否正确
	flag,err:=service.IsPasswordCorrect(username,oldPassword)
	if err!=nil{
		fmt.Println("judge password correct err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag{
		tool.RespErrorWithData(ctx,"旧密码错误")
		return
	}

	//验证新密码合理性
	flag=service.IsPasswordReasonable(newPassword)
	if !flag{
		tool.RespErrorWithData(ctx,"新密码不合理")
		return
	}

	//修改为新密码
	err=service.ChangePassword(username,newPassword)
	if err!=nil{
		fmt.Println("change password err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

//通过验证密保找回密码
func changePasswordBySecurity(ctx *gin.Context)  {
	username:=ctx.PostForm("username")
	answer:=ctx.PostForm("security_answer")
	newPassword:=ctx.PostForm("new_password")
	//验证用户名
	flag,err:=service.IsExistUsername(username)
	if err!=nil{
		fmt.Println("judge exist username err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag{
		tool.RespErrorWithData(ctx,"用户不存在")
		return
	}

	//验证密保
	flag, err = service.IsSecurityCorrect(username, answer)
	if err != nil {
		fmt.Println("judge security err:", err)
		return
	}
	if !flag{
		tool.RespErrorWithData(ctx,"密保答案错误")
		return
	}

	//重设密码
	flag=service.IsPasswordReasonable(newPassword)
	if !flag{
		tool.RespErrorWithData(ctx,"密码不合理")
		return
	}
	err=service.ChangePassword(username,newPassword)
	if err!=nil{
		fmt.Println("change password err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}

