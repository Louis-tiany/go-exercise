/*
* File     : model/model.go
* Author   : *
* Mail     : *
* Creation : Sun 05 Mar 2023 04:12:49 PM CST
 */
package model

import (
	"fmt"
	"stock/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Symbol      string  `gorm:"column:symbol;comment:股票代码;type:varchar(100);" json:"symbol"`
	Name        string  `gorm:"column:name;comment:name;type:varchar(100);" json:"name"`
	Trade       float64 `gorm:"column:trade;comment:trade;type:float;" json:"trade"`
	RatioAmount float64 `gorm:"column:ratioamount;comment:ratioamount;type:float;" json:"ratioamount"`
	InAmont     float64 `gorm:"column:inamount;comment:inamout;type:float;" json:"inamount"`
	OutAmont    float64 `gorm:"column:outamount;comment:inamout;type:float;" json:"outamount"`
	NetAmount   float64 `gorm:"column:netamount;comment:inamout;type:float;" json:"netamount"`
}

type DbConn struct {
	Db *gorm.DB
}

var Db *DbConn

func NewStockDb() *DbConn {
	InitDb()
	Db.Db.AutoMigrate(&Stock{})
	return Db
}

func InitDb() *DbConn {
	Db = &DbConn{Db: connectdb()}
	return Db
}

func connectdb() *gorm.DB {
	var err error
	dsn := config.Config.DbConfig.Host + ":" + config.Config.DbConfig.Username + "@tcp" + "(" + config.Config.DbConfig.Host + ":" + config.Config.DbConfig.Port + ")/" + config.Config.DbConfig.DbName + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}
