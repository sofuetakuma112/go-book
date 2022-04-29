package gorm

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"` // Gormに対して非nullの列を作成するように指示する
	Comments  []Comment // 1対多
	CreatedAt time.Time // struct初期化時に値を指定しない場合は、Timeの初期値が当てられる
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int // 特定の形式のフィールド名の場合、Gormが外部キーであると想定してリレーションを作成する
	CreatedAt time.Time // CreatedAt?フィールドはデータベース中に新しいレコードが作成された際に必ず自動的に値が設定される
}

var Db *gorm.DB // gorm.DBはデータベースのハンドル

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "nice post!", Author: "Joe"}

	Db.Model(&post).Association("Comments").Append(comment) // Model: db操作を実行したいモデルを指定します。

	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
