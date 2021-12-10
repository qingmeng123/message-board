/*******
* @Author:qingmeng
* @Description:
* @File:posg
* @Date2021/12/10
 */

package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

//查看某一留言的详情
func postDetail(ctx *gin.Context) {
	postIdString:=ctx.Param("post_id")
	postId,err:=strconv.Atoi(postIdString)
	if err!=nil{
		fmt.Println("post id string to int err:",err)
		tool.RespErrorWithData(ctx,"post_id格式有误")
		return
	}

	//根据postId拿到post
	post,err:=service.GetPostById(postId)
	if err!=nil{
		fmt.Println("get post by id err:",err)
		return
	}

	//找到它的评论
	comments,err:=service.GetPostComments(postId)
	if err!=nil{
		if err!=sql.ErrNoRows{
			fmt.Println("get post comments err:",err)
			tool.RespInternalError(ctx)
			return
		}
	}
	var postDetail model.PostDetail
	postDetail.Post=post
	postDetail.Comments=comments

	tool.RespSuccessfulWithData(ctx,postDetail)
}

//所有留言简况
func briefPosts(ctx *gin.Context){
	posts,err:=service.GetPosts()
	if err!=nil{
		fmt.Println("get posts err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,posts)
}

//增加留言
func addPost(ctx *gin.Context) {
	iUsername,_:=ctx.Get("iUsername")
	username:=iUsername.(string)
	txt:=ctx.PostForm("txt")

	//检验信息是否含有敏感词
	flag:=tool.CheckIfSensitive(txt)
	if flag{
		tool.RespSensitiveError(ctx)
		return
	}
	post:=model.Post{
		CommentNum: 0,
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err:=service.AddPost(post)
	if err!=nil{
		fmt.Println("add post err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

//留言更新(可通过改变状态来删除)
func updatePost(ctx *gin.Context)  {
	postIdString:=ctx.Param("post_id")
	postId,err:=strconv.Atoi(postIdString)
	if err!=nil{
		fmt.Println("post id string to int err:",err)
		tool.RespErrorWithData(ctx,"post_id格式有误")
		return
	}
	post,err:=service.GetPostById(postId)
	if err!=nil{
		fmt.Println("get post by id err:",err)
		tool.RespInternalError(ctx)
		return
	}

	//更新信息
	txt:=ctx.PostForm("txt")

	//检验信息是否含有敏感词
	flag:=tool.CheckIfSensitive(txt)
	if flag{
		tool.RespSensitiveError(ctx)
		return
	}
	post.Txt=txt
	post.State,err=strconv.Atoi(ctx.PostForm("state"))
	if err!=nil{
		fmt.Println("state to int err:",err)
		tool.RespInternalError(ctx)
		return
	}
	post.UpdateTime=time.Now()
	err=service.UpdatePost(post)
	if err!=nil{
		fmt.Println("update post err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,"更新留言成功")
}
