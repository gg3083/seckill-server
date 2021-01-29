package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"seckill-server/setting"
	"time"
)

var (
	err error
	mdb *sql.DB
)

var TableNameList = []string{"t_goods", "t_userinfo"}

func InitMysql() {
	uri := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "Qwe123.0", "tcp", "localhost", "3306", "seckill-server")
	mdb, err = sql.Open("mysql", uri)
	if err != nil {
		log.Fatalf("Open mysql failed,err:%v\n", err)
	}
	mdb.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	mdb.SetMaxOpenConns(100)                  //设置最大连接数
	mdb.SetMaxIdleConns(16)                   //设置闲置连接数
	existTable(TableNameList)
}

func existTable(tableNames []string) {
	for _, tableName := range tableNames {
		log.Printf("开始检测%s表是否存在\n", tableName)
		var name string
		row, err2 := mdb.Query("SELECT table_name FROM information_schema.TABLES WHERE table_name = ?", tableName)
		if err2 != nil {
			log.Fatalf("connect mysql failed,err:%v\n", err2.Error())
		}
		if row.Next() {
			if err3 := row.Scan(&name); err3 != nil{
				log.Println("err:", err3)
				//执行创建sql脚本初始化
				initTable(tableName)
			}else {
				log.Printf("表%s已存在\n", name)
			}
		}else {
			initTable(tableName)
		}

	}
}

func initTable(tableName string) {
	var script string
	switch tableName {
	case "t_goods":
		script = setting.SqlScript.Goods
	case "t_userinfo":
		script = setting.SqlScript.Userinfo
	default:
		script = ""
	}
	if script == "" {
		log.Fatalf("表%s不存在\n", tableName)
	}
	log.Printf("初始化表%s,脚本为%s\n", tableName, script)
	_, err2 := mdb.Exec(script)
	if err2 != nil {
		log.Fatalf("初始化表失败 %s", err2.Error())
	}
	log.Printf("表%s初始化成功！\n", tableName)
}
