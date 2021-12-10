/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2021/12/10
 */

package dao

import (
	"message-board/model"
)

func InsertComment(comment model.Comment) error {
	_,err:=DB.Exec("insert into message_board.comment( post_id,username, txt,comment_time,to_comment_id) VALUES(?,?,?,?,?);",comment.PostId,comment.Username,comment.Txt,comment.CommentTime,comment.ToCommentId)
	return err
}

func SelectCommentByPostId(postId int)([]model.Comment,error)  {
	var comments []model.Comment

	rows,err:=DB.Query("select * from message_board.comment where post_id=?",postId)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var comment model.Comment
		err=rows.Scan(&comment.CommentId,&comment.PostId,&comment.Username,&comment.Txt,&comment.CommentTime,&comment.ToCommentId,&comment.State)
		if err!=nil{
			return nil, err
		}
		comments=append(comments,comment)
	}
	return comments,err
}

func SelectCommentById(commentId int)(model.Comment,error)  {
	var comment model.Comment
	row:=DB.QueryRow("select * from message_board.comment where comment_id=?",commentId)
	if row.Err()!=nil{
		return comment,row.Err()
	}
	err:=row.Scan(&comment.CommentId,&comment.PostId,&comment.Username,&comment.Txt,&comment.CommentTime,&comment.ToCommentId,&comment.State)
	if err!=nil{
		return model.Comment{}, err
	}
	return comment, err
}

func UpdateComment(comment model.Comment) error {
	_,err:=DB.Exec("update message_board.comment set txt=?,comment_time=?,state=? where comment_id=?;",comment.Txt,comment.CommentTime,comment.State,comment.CommentId)
	return err
}