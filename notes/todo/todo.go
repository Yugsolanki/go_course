package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Invalid Input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("\nTODO: %v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(json), 0644) // will be nil if everything works or error
}
