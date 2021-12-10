/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2021/12/10
 */

package model

import "time"

type Comment struct {
	PostId		int
	CommentId 	int
	Username  	string
	Txt			string
	CommentTime	time.Time
	ToCommentId	int		//回复某评论的id
	State		int		//状态：1表示存在，0表示不存在（即删除）
}
