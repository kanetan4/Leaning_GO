package main

import (
    "context"
    "fmt"
	"ex4_noteservice/kitex_gen/notes/noteservice"
    "ex4_noteservice/kitex_gen/notes"
    "github.com/cloudwego/kitex/client"
)

func main() {
    client, err := noteservice.NewClient("note-service", client.WithHostPorts("127.0.0.1:8888"))
    if err != nil {
        panic(err)
    }

    req := &notes.CreateNoteRequest{
        User: "kane",
        Title: "From user-service",
        Content: "Learning Kitex",
    }

    res, err := client.CreateNote(context.Background(), req)
    if err != nil {
        panic(err)
    }

    fmt.Println("Note created:", res)
}
