/**
 *@Description     gin框架 练手项目 小清单
 *@author          liyajun
 *@create          2022-09-09 15:31
 */

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"list/DAO"
	"list/Models"
	"list/Routers"
	"list/Settings"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./list conf/config.ini")
		return
	}
	//连接数据库
	if err := Settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	if err := DAO.InitMysql(Settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	//退出关闭数据库
	defer DAO.Close()
	//模型绑定
	DAO.DB.AutoMigrate(&Models.Todo{})

	r := Routers.SetupRouter()

	err := r.Run(":9090")

	if err != nil {
		return
	}

}
