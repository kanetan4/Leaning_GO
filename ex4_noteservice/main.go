package main

import (
	notes "ex4_noteservice/kitex_gen/notes/noteservice"
	"log"
)

func main() {
	svr := notes.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
