package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

// If you have interface that requires only one method then interface name will be method name + 'er' at the end
// e.g. :- If Method = Jump, then Interface = Jumper
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
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

	err = outputData(userTodo)
	if err != nil {
		return
	}
	fmt.Println("Saving the todo succeded")
}

// Any type: basically accepts any type of value
func printSomething(value interface{}) {

	// `value.(type)` only works inside `switch`
	// the other way is below, it give back two things, `ok` is a boolean which is important
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("Integer: ", stringVal)
		return
	}

	// if you don't wanna use `if` also can go for `switch`
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case todo.Todo:
		fmt.Println("Todo: ", value)
	case string:
		fmt.Println(value)
	default:
		return
	}

	fmt.Println(value)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
		return err
	}
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
