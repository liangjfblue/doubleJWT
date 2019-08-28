package models

import (
    "doubleJWT/config"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/sirupsen/logrus"
    "time"
)

// BaseModel 基本model
type BaseModel struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

var MysqlPool *gorm.DB

func InitMysql(mysqlConf *config.MysqlConfig) {
    var err error
    str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Password, mysqlConf.Addr, mysqlConf.Db)
    MysqlPool, err = gorm.Open("mysql", str)
    if err != nil {
        logrus.Error("open error : ", err)
        panic(err)
    }

    MysqlPool.LogMode(true)
    MysqlPool.DB().SetMaxIdleConns(50)
    MysqlPool.DB().SetMaxOpenConns(500)
}

func Close() {
    _ = MysqlPool.Close()
}