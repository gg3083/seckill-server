package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"seckill-server/setting"
	"strings"
	"time"
)

var (
	err error
	mdb *sql.DB
)

var TableNameList = []string{"t_goods", "t_user_info", "t_user_address", "t_user_fund", "t_user_fund_record", "t_order"}

func InitMysql() {
	uri := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "Qwe123.0", "tcp", "localhost", "3306", "seckill-server")
	mdb, err = sql.Open("mysql", uri)
	if err != nil {
		log.Fatalf("Open mysql failed,err:%v\n", err)
	}
	mdb.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	mdb.SetMaxOpenConns(50)                   //设置最大连接数
	mdb.SetMaxIdleConns(16)                   //设置闲置连接数
	existTable(TableNameList)
}

func existTable(tableNames []string) {
	for _, tableName := range tableNames {
		log.Printf("开始检测%s表是否存在\n", tableName)
		var count int
		err2 := mdb.QueryRow("SELECT count(1) FROM information_schema.TABLES WHERE table_schema = ? and table_name = ?", "seckill-server", tableName).Scan(&count)
		if err2 != nil {
			log.Fatalf("connect mysql failed,err:%v\n", err2.Error())
		}
		if count == 0 {
			//执行创建sql脚本初始化
			initTable(tableName)
		} else {
			log.Printf("表%s已存在\n", tableName)
			if setting.SqlScript.IsFlush {
				clearTable(tableName)
				initTable(tableName)
			}
		}
	}
}

func initTable(tableName string) {
	var script string
	switch tableName {
	case "t_goods":
		script = setting.SqlScript.Goods
	case "t_user_info":
		script = setting.SqlScript.UserInfo
	case "t_user_address":
		script = setting.SqlScript.UserAddress
	case "t_user_fund":
		script = setting.SqlScript.UserFund
	case "t_user_fund_record":
		script = setting.SqlScript.UserFundRecord
	case "t_order":
		script = setting.SqlScript.Order
	default:
		script = ""
	}
	if script == "" {
		log.Fatalf("表%s不存在，且不存在初始化脚本\n", tableName)
	}
	log.Printf("初始化表%s,脚本为%s\n", tableName, script)
	_, err2 := mdb.Exec(script)
	if err2 != nil {
		log.Fatalf("初始化表失败 %s", err2.Error())
	}
	log.Printf("表%s初始化成功！\n", tableName)
}

func clearTable(tableName string) {
	var script strings.Builder
	script.WriteString("drop table ")
	switch tableName {
	case "t_goods":
		script.WriteString(tableName)
	case "t_user_info":
		script.WriteString(tableName)
	case "t_user_address":
		script.WriteString(tableName)
	case "t_user_fund":
		script.WriteString(tableName)
	case "t_user_fund_record":
		script.WriteString(tableName)
	case "t_order":
		script.WriteString(tableName)
	default:
		script.Reset()
	}
	if script.String() == "" {
		log.Fatalf("表%s不存在，且不存在初始化脚本\n", tableName)
	}
	log.Printf("初始化表%s,脚本为%s\n", tableName, script.String())
	_, err2 := mdb.Exec(script.String())
	if err2 != nil {
		log.Fatalf("初始化表失败 %s", err2.Error())
	}
	log.Printf("表%s初始化成功！\n", tableName)
}
