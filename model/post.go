/*******
* @Author:qingmeng
* @Description:
* @File:post
* @Date2021/12/10
 */

package model

import "time"

type Post struct {
	PostId 		int 		`json:"post_id"`
	Username	string		`json:"username"`
	Txt			string 		`json:"txt"`
	CommentNum	int 		`json:"comment_num"`
	PostTime	time.Time 	`json:"post_time"`
	UpdateTime	time.Time	`json:"update_time"`
	State       int 		`json:"state"`			//状态：1表示存在，0表示不存在（即删除）
}

type PostDetail struct {
	Post
	Comments []Comment
}
