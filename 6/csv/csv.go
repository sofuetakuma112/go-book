package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	posts1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	posts2 := Post{Id: 2, Content: "Bonjor Monde!", Author: "Pierre"}
	posts3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	posts4 := Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"}

	allPosts := []Post{
		posts1,
		posts2,
		posts3,
		posts4,
	}

	writer := csv.NewWriter(csvFile) // 引数のwriterに書き込む新しいwriterを作成する
	for _, post := range allPosts {
		// strconv.Itoaでint64型からstringへ変換している
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author} // CSVに書き込む一行をスライスで生成する
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() // バッファにあるデータを全て確実にファイルに書き込むために呼び出している

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // 負の数を渡すことで、レコード内に全てのフィールドが揃っていなくてもOKであるようにしている
	record, err := reader.ReadAll() // [][]string
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0) // 文字列の接頭辞(0b, 0xなど)から真の基数を解釈して、int型に変換する
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
