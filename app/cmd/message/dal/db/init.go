package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var PgDb *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s port='%s' user=%s password=%s dbname=%s TimeZone=Asia/Shanghai connect_timeout=10",
		os.Getenv("PGSQL_HOST"), os.Getenv("PGSQL_PORT"),
		os.Getenv("PGSQL_USER"), os.Getenv("PGSQL_PASSWORD"), os.Getenv("PGSQL_DBNAME"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		log.Fatal("TABLE_NAME can't be empty")
	}
	db = db.Table(tableName)

	err = db.Migrator().AutoMigrate(&Message{})
	if err != nil {
		log.Fatal(err)
	}

	PgDb = db
}
