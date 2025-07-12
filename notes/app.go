package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeded")

	fmt.Println("--------------------------------------")

	todoText := getUserInput("Todo content: ")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	userTodo.Display()
	err = userTodo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}
	fmt.Println("Saving the todo succeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
