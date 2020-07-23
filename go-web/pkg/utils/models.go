package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deteled_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = viper.GetString("database.type")
	dbName = viper.GetString("database.name")
	user = viper.GetString("database.user")
	password = viper.GetString("database.pwd")
	host = viper.GetString("database.url")
	tablePrefix = viper.GetString("database.tablePrefix")
	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Panicln(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer DB.Close()
}
