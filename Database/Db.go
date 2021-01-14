package Database

import (
	"Book/Pkg/Setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"
	"time"
)

var Db *gorm.DB

type Model struct {
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := Setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "数据库配置加载失败: %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("DB_HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if tablePrefix != "" {
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return tablePrefix + defaultTableName
		}
	}

	Db.SingularTable(true)

	if strings.ToLower(Setting.RunMode) == "debug" {
		Db.LogMode(true)
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Db.Close()
}
