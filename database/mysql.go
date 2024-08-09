package database

// import (
//     "fmt"
//     "log"

//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
// )

// const (
//     host     = "localhost"
//     port     = 3306
//     username = "root"
//     password = "admin"
//     dbname   = "db"
// )

// var DB *gorm.DB

// func main() {
//     dsn := fmt.Sprintf(
//         "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//         username, password, host, port, dbname,
//     )
//     db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatalf("failed to connect database: %v", err)
//     }
//     DB = db
// }