/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/12/10
 */

package main

import (
	"message-board/api"
	"message-board/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
