package main

import (
	fmt "fmt"

	"github.com/golang/protobuf/proto"
	"go_tools/data"
)

func main() {
	var text = []byte("hello")
	message := &data.Message{
		Text: text,
	}

	b, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	fmt.Println(b) // [10 5 104 101 108 108 111]

	newMessage := &data.Message{}
	err = proto.Unmarshal(b, newMessage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("newMessage = %+v\n", newMessage)
	fmt.Println(newMessage.GetText()) // [104 101 108 108 111]
}
