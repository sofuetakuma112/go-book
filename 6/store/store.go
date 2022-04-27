package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgresのドライバ(importした際にinit関数を呼び出して自分をpostgresのドライバとして登録する)
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error

	// sql.DBはデータベース接続のプールを内部に持っているだけで、クローズする必要はない
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable") // DB接続用の構造体を設定するだけで接続は行っていない

	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit) // インターフェースRows(イテレータ)を返す
	if err != nil {
		return
	}
	for rows.Next() { // RowsのメソッドNextを繰り返し呼び出せば、sql.Rowを返し、行がなくなるとio.EOFを返します。
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// （取得する投稿を表す）構造体Postが手元にあるわけではないので、
// その上にメソッドを定義せず、外部に関数を定義している
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement) // プリペアドステートメントを作成する
	if err != nil {
		return
	}
	defer stmt.Close()
	// 構造体sql.Rowへの参照をひとつだけ返すためにQueryRowを使っている
	// 構造体sql.Rowへの参照がひとつだけだから、Scan関数が呼び出せている
	// QueryRowはScanと合わせて使うケースが多いので、エラーは返さない
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	// 更新後のDBの値が必要ない場合は、高速に動作するDb.Execを用いる
	// 結果として返されるsql.Resultには、影響のあった行の数のほか、最後に挿入された行がある場合はそのidが設定される
	_, err = Db.Exec("UPDATE posts SET content = $2, author = $3 WHERE id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id) // 構造体Postを返す
	fmt.Println(readPost)

	readPost.Content = "Bonjor Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
}
