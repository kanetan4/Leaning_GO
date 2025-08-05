package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"ex4_noteservice/kitex_gen/notes"
	"ex4_noteservice/kitex_gen/notes/noteservice"
	"github.com/cloudwego/kitex/client"
)

func main() {
	cli, err := noteservice.NewClient("note-service", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Note Client!")
	fmt.Println("Available commands: create | get | list | delete | exit")

	for {
		fmt.Print("\nEnter command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "create":
			fmt.Print("User: ")
			user, _ := reader.ReadString('\n')
			fmt.Print("Title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Content: ")
			content, _ := reader.ReadString('\n')

			req := &notes.CreateNoteRequest{
				User:    strings.TrimSpace(user),
				Title:   strings.TrimSpace(title),
				Content: strings.TrimSpace(content),
			}

			resp, err := cli.CreateNote(context.Background(), req)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("✅ Note created:")
				printNote(resp)
			}

		case "get":
			fmt.Print("Note ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)

			fmt.Print("User: ")
			user, _ := reader.ReadString('\n')

			req := &notes.GetNoteRequest{
				User: strings.TrimSpace(user),
				Id:   id,
			}

			resp, err := cli.GetNote(context.Background(), req)
			if err != nil {
				fmt.Println("Error:", err)
			} else if resp == nil {
				fmt.Println("❌ Note not found.")
			} else {
				fmt.Println("✅ Note found:")
				printNote(resp)
			}

		case "list":
			fmt.Print("User: ")
			user, _ := reader.ReadString('\n')
			user = strings.TrimSpace(user)

			resp, err := cli.ListNotes(context.Background(), user)
			if err != nil {
				fmt.Println("Error:", err)
			} else if len(resp) == 0 {
				fmt.Println("📭 No notes found for user.")
			} else {
				fmt.Printf("📋 Listing %d notes:\n", len(resp))
				for _, note := range resp {
					printNote(note)
				}
			}

		case "delete":
			fmt.Print("Note ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)

			fmt.Print("User: ")
			user, _ := reader.ReadString('\n')

			req := &notes.DeleteNoteRequest{
				User: strings.TrimSpace(user),
				Id:   id,
			}

			resp, err := cli.DeleteNote(context.Background(), req)
			if err != nil {
				fmt.Println("Error:", err)
			} else if resp {
				fmt.Println("🗑️ Note deleted successfully.")
			} else {
				fmt.Println("❌ Note not found or failed to delete.")
			}

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command.")
		}
	}
}

func printNote(n *notes.Note) {
	fmt.Println("------------------------------------------------")
	fmt.Println("ID:        ", n.Id)
	fmt.Println("User:      ", n.User)
	fmt.Println("Title:     ", n.Title)
	fmt.Println("Content:   ", n.Content)
	fmt.Println("Created At:", time.Unix(n.CreateTime, 0))
}
