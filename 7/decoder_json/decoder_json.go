package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// io.Readerのストリームからデータが入ってくるときはDecoderを使う(Read関数等で継続的に取り出すケースの場合に有効)
	// 文字列データや、バイト列をメモリ内にすでに格納している場合はUnmarshalを使う
	decoder := json.NewDecoder(jsonFile) // JSONデータからデコーダを生成する
	for {
		var post Post
		err := decoder.Decode(&post) // JSONデータをデコードし、構造体に格納する
		if err == io.EOF { // EOFが検出されるまで繰り返す
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}
