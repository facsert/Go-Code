package database

import (
	"fmt"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBconnect struct {
	host string
	port int
    username string
	password string
	databaseName string
}

var (
	Base  = &DBconnect{
		databaseName: "base",
		host: "102.168.1.100",
		port: 5432,
		username: "username",
		password: "password",
		
	}

	Backup = &DBconnect{
		databaseName: "backup",
		host: "102.168.1.100",
		port: 5432,
		username: "username",
		password: "password",
		
	}

	DBList = []*DBconnect{Base, Backup}
	DBMap = make(map[string]*gorm.DB, 2)
	Default = Base
)

func Init() {
	for _, conn := range DBList {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			conn.host, conn.username, conn.password, conn.databaseName, conn.port,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database %s", conn.databaseName)
		}
		DBMap[conn.databaseName] = db
	}
}

func NewDB() *gorm.DB {
    return DBMap[Default.databaseName]
}

func SwitchDB(conn DBconnect) *gorm.DB {
    return DBMap[conn.databaseName]
}
