package main

import (
	_ "github.com/go-sql-driver/mysql" //手动引入mysql驱动,变量常量,运行init函数
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	// 驱动类型
	DbEngine, err := xorm.NewEngine("mysql", "root:Acjq.100@(127.0.0.1:3306/chat?charset=utf8")

	//显示sql语句,方便调试
	DbEngine.ShowSQL(true)

	// 设置数据库连接数,直接关乎数据库操作的性能
	DbEngine.SetMaxOpenConns(2)

	//自动创建表
	//User结构体
	DbEngine.Sync2(new(User))

}
func main() {

}
