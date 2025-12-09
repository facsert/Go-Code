package db

import (
	"fmt"
	"log"
	"time"
	"sync"

	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConn struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type DBName string

const (
	Base   DBName = "base"
	Backup DBName = "backup"
)
var ConnMap = map[DBName]*DBConn{
	Base: {
		Host:     "localhost",
		Port:     5432,
		Username: "root",
		Password: "admin",
		DBName:   "base",
	},
	// Backup: &DBConn{
	// 	Host: "localhost",
	// 	Port: 5432,
	// 	Username: "root",
	// 	Password: "admin",
	// 	DBName: "backup",
	// },
}
var (
	DBMap = make(map[DBName]*sql.DB, len(ConnMap))
	defaultDB *sql.DB
	mutex     sync.RWMutex
)

// init all database connection
func Init() {
	for name, conn := range ConnMap {
		cursor, err := sql.Connect(
			"postgres",
			fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
				conn.Username, conn.Password, conn.Host, conn.Port, conn.DBName,
			))
		if err != nil {
			log.Fatalf("connect database %s err: %v", conn.DBName, err)
		}
		cursor.SetMaxOpenConns(10)
		cursor.SetMaxIdleConns(5)
		cursor.SetConnMaxLifetime(5 * time.Minute)
		DBMap[name] = cursor
	}

	defaultDB = DBMap[Base]
}

// get default database
func New() *sql.DB {
	mutex.RLock()
	defer mutex.RUnlock()
	return defaultDB
}

// get non default database
func Temp(name DBName) *sql.DB {
	mutex.RLock()
	defer mutex.RUnlock()
	return DBMap[name]
}
// set default database
func SetDefault(name DBName) {
	mutex.Lock()
	defer mutex.Unlock()
	defaultDB = DBMap[name]
}
// Close all database connection
func Close() {
	mutex.Lock()
	defer mutex.Unlock()
	for _, db := range DBMap {
		if db == nil {
			continue
		}
        db.Close()
	}
}
