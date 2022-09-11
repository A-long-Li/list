/**
 *@Description
 *@author          liyajun
 *@create          2022-09-09 17:04
 */

package DAO

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"list/Settings"
)

var (
	DB *gorm.DB
)

func InitMysql(cfg *Settings.MySQLConfig) (err error) {
	msg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase)
	DB, err = gorm.Open("mysql", msg)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Bind(str interface{}) {
	DB.AutoMigrate(str)
}

func Close() {
	err := DB.Close()
	if err != nil {
		return
	}
}
