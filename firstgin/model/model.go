package model

import (

	//"fmt"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"firstgin/config"
)

var db *gorm.DB

//初始化数据库链接
func init(){
   
   var err error

   db, err = gorm.Open("mysql", "root:"+config.MYSQL_USER+"@tcp("+config.MYSQL_HOST+":"+config.MYSQL_PORT+")/"+config.MYSQL_DATABASE+"?charset=utf8&parseTime=True&loc=Local")
   if err != nil {
        panic("连接数据库失败")
	}
	
	//链接池
	db.DB().SetMaxIdleConns(config.MYSQL_MAX_IDES)
	db.DB().SetMaxOpenConns(config.MYSQL_MAX_CONNECTIONS)

	//表名不加s
	db.SingularTable(true)

}