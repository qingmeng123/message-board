/*******
* @Author:qingmeng
* @Description:
* @File:dao
* @Date2021/12/10
 */

package dao

import "database/sql"
import _"github.com/go-sql-driver/mysql"

var DB *sql.DB

var(
	name="root"
	pwd="@XUEHUI."
	host="localhost"
	port="3306"
	dbname="message_board"
)

func InitDB() {
	dsn:= name +":"+ pwd +"@tcp("+ host +":"+ port +")/"+ dbname+"?charset=utf8mb4&parseTime=True"
	db,err:=sql.Open("mysql",dsn)
	if err!=nil{
		panic(err)
	}
	DB=db
}
