package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1", data, 0644) // ファイルに書き込み(バイト列で渡しているが、書き込むファイルはテキストデータになる)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1") // ファイルから読み込む
	fmt.Print(string(read1))

	file1, _ := os.Create("data2")
	defer file1.Close() // deferはスタックで後入れ先出しで実行される

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2") // 構造体Fileを使ってファイルの読み書きをする
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
