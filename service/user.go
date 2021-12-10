/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

//注册
func Register(user model.User)error {
	err:=dao.InsertUser(user)
	return err
}

func IsPasswordCorrect(username,password string)(bool,error)  {
	user,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false, err
	}
	if user.Password!=password{
		return false, err
	}
	return true, err
}

//判断用户名是否存在
func IsExistUsername(username string)(bool,error)  {
	_,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false, err
	}
	return true,nil
}

func ChangePassword(username, newPassword string) error {
	err:=dao.UpdatePassword(username,newPassword)
	return err
}

//验证密码是否合理(可增加密码复杂性)
func IsPasswordReasonable(password string)bool  {
	if len(password)<6{
		return false
	}
	return true
}

//更新密保
func UpdateSecurity(username, answer string) error {
	err:=dao.UpdateSecurity(username,answer)
	return err
}

//验证密保
func IsSecurityCorrect(username,answer string) (bool ,error) {
	user,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false,err
	}
	if user.SecurityAnswer !=answer{
		return false,err
	}
	return true,nil
}

