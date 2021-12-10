/*******
* @Author:qingmeng
* @Description:
* @File:post
* @Date2021/12/10
 */

package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddPost( post model.Post)error {
	return dao.InsertPost(post)
}

func GetPosts() ([]model.Post,error) {
	return dao.SelectPosts()
}

func GetPostById(postId int) (model.Post,error) {
	return dao.SelectPostById(postId)
}

func UpdatePost(post model.Post) error {
	return dao.UpdatePost(post)
}