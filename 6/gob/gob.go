package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer) // 構造体Bufferはバイトデータの可変バッファであり、メソッドReadとメソッドWriteを持っている
	encoder := gob.NewEncoder(buffer) // gobエンコーダの生成
	err := encoder.Encode(data) // dataをバッファ内にエンコードする
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600) // buffer.Bytesは、バッファの未読部分を保持する長さb.Len()のスライスを返します。
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw) // 生データからバッファを生成(生データをメソッドReadやメソッドWriteに渡している)
	// gobはバイナリデータにエンコード、デコードするライブラリ？
	dec := gob.NewDecoder(buffer) // NewDecoderは、io.Readerから読み込む新しいデコーダを返します。
	err = dec.Decode(data) // デコードしたものをdataに入れる？
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
