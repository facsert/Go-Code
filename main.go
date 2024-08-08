package main

import (
	"fmt"
	"time"

	"learn/database"
)

type Article struct {
    // Id int `gorm:"primaryKey"`
    Title string `gorm:"column:title"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}


func main() {
	fmt.Println("main")

	database.Init()
    
    var article Article
	// res := (database.DB.Table("article").Select("id", "title", "created_at", "updated_at").Where("id = 6").Or("title = 'Prof.'").Find(&article))
	res := (database.DB.Table("article").Where("id = 6").Find(&article))
	fmt.Printf("articles: %#v\n", article)
	fmt.Printf("res: %v\n", res.RowsAffected)
}