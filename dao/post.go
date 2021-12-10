/*******
* @Author:qingmeng
* @Description:
* @File:post
* @Date2021/12/10
 */

package dao

import "message-board/model"

func InsertPost(post model.Post) error {
	_, err := DB.Exec("insert into message_board.post( username,txt,comment_num, post_time, update_time)"+
		" VALUES(?,?,?,?,?);", post.Username, post.Txt, post.CommentNum, post.PostTime, post.UpdateTime)
	return err
}

func SelectPostById(postId int)(model.Post,error){
	var post model.Post

	row:=DB.QueryRow("select post_id,username,txt,post_time,update_time,comment_num,state from message_board.post where post_id=?",postId)
	if row.Err()!=nil{
		return post,row.Err()
	}
	err:=row.Scan(&post.PostId,&post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum,&post.State)
	if err!=nil{
		return post,err
	}
	return post,err
}

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := DB.Query("SELECT post_id, username, txt, post_time, update_time, comment_num,state  FROM message_board.post where state=1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post
		err = rows.Scan(&post.PostId, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum,&post.State)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//留言更新
func UpdatePost(post model.Post)error  {
	_,err:=DB.Exec("update message_board.post set txt=? , comment_num=? , update_time=?,state=? where post_id=?;",post.Txt,post.CommentNum,post.UpdateTime,post.State,post.PostId)
	return err
}
