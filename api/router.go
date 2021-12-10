/*******
* @Author:qingmeng
* @Description:
* @File:router
* @Date2021/12/10
 */

package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine:=gin.Default()

	engine.POST("/user",register)		//注册
	engine.GET("/user",login)			//登陆

	userGroup:=engine.Group("/user")
	{
		userGroup.POST("/security",auth,changeSecurity)	//更新密保
		userGroup.POST("/password")

		passwordGroup:=userGroup.Group("/password")
		{
			passwordGroup.POST("/",auth,changePassword)	//登陆后的直接修改密码
			passwordGroup.POST("/security",auth,changePasswordBySecurity)//密保找回密码
		}
	}

	postGroup := engine.Group("/post")
	{
		postGroup.Use(auth)
		postGroup.POST("/", addPost)            //发布新留言
		postGroup.POST("/:post_id", updatePost) //修改留言

		postGroup.GET("/", briefPosts)         //查看全部留言概略
		postGroup.GET("/:post_id", postDetail) //查看一条留言详细信息
		postGroup.DELETE("/:post_id",updatePost)	//删除留言
	}

	commentGroup:=engine.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/",addComment)		//发送评论，包括回复别人的评论
		commentGroup.POST("/:comment_id",UpdateComment)		//修改评论
		commentGroup.DELETE("/:comment_id",UpdateComment)	//删除评论
	}

	engine.Run()
}
