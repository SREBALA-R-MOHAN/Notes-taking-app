package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json: "text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo) //to convert data to json format
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) //0644: read and write permissions

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Input cannot be empty")
	}

	return Todo{
		Text: content,
	}, nil
}
