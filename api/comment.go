/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2021/12/10
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

//回复评论
func reply(ctx *gin.Context) {

	addComment(ctx)
}

func addComment(ctx *gin.Context) {
	iUsername,_:=ctx.Get("iUsername")
	username:=iUsername.(string)
	txt:=ctx.PostForm("txt")

	//检验信息是否含有敏感词
	flag:=tool.CheckIfSensitive(txt)
	if flag{
		tool.RespSensitiveError(ctx)
		return
	}
	postIdString:=ctx.PostForm("post_id")
	postId,err:=strconv.Atoi(postIdString)
	if err!=nil{
		fmt.Println("post id string to int err:",err)
		tool.RespErrorWithData(ctx,"文章id有误")
		return
	}

	//是否在回复其他评论,无则默认0
	toCommentId,err:=strconv.Atoi(ctx.PostForm("to_comment_id"))
	if err!=nil{
		fmt.Println("to_comment_id to int err:",err)
		tool.RespErrorWithData(ctx,"其他评论id有误")
		return
	}
	comment:=model.Comment{
		PostId: postId,
		Txt: txt,
		Username: username,
		CommentTime: time.Now(),
		ToCommentId: toCommentId,
	}
	err=service.AddComment(comment)
	if err!=nil{
		fmt.Println("add comment err:",err)
		tool.RespInternalError(ctx)
		return
	}

	//给文章增加评论数，更新时间
	var post model.Post
	post,err=service.GetPostById(postId)
	if err!=nil{
		fmt.Println("get post err:",err)
		tool.RespInternalError(ctx)
		return
	}
	post.CommentNum+=1
	post.UpdateTime=time.Now()
	err=service.UpdatePost(post)
	if err!=nil{
		fmt.Println("update post err:",err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

//更新评论
func UpdateComment(ctx *gin.Context) {
	commentIdString:=ctx.Param("comment_id")
	commentId,err:=strconv.Atoi(commentIdString)
	if err!=nil{
		fmt.Println("commentId string to int err:",err)
		tool.RespErrorWithData(ctx,"comment_id格式有误")
		return
	}

	//通过comment_id找出该条评论
	comment,err:=dao.SelectCommentById(commentId)
	if err!=nil{
		fmt.Println("select comment by id err:",err)
		tool.RespInternalError(ctx)
		return
	}

	//更新评论
	txt:=ctx.PostForm("txt")
	//检验信息是否含有敏感词
	flag:=tool.CheckIfSensitive(txt)
	if flag{
		tool.RespSensitiveError(ctx)
		return
	}
	comment.Txt=txt
	comment.State,err=strconv.Atoi(ctx.PostForm("state"))
	if err!=nil{
		fmt.Println("state to int err:",err)
		tool.RespInternalError(ctx)
		return
	}
	comment.CommentTime=time.Now()
	err=service.UpdateComment(comment)
	if err!=nil{
		fmt.Println("update comment err:",err)
		tool.RespInternalError(ctx)
	}

	tool.RespSuccessfulWithData(ctx,"更新评论成功")
}