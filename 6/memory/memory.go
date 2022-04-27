package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	posts1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	posts2 := Post{Id: 2, Content: "Bonjor Monde!", Author: "Pierre"}
	posts3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	posts4 := Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"}

	store(posts1)
	store(posts2)
	store(posts3)
	store(posts4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
