/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package model

type User struct {
	UserId         int    `json:"user_id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	SecurityAnswer string `json:"security_answer"`
}
