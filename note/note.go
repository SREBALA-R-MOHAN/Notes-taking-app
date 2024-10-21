package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json: "title"` //struct tags(used as keys in json file generated)
	Content   string    `json: "content"`
	CreatedAt time.Time `json: "created_at"`
}

func (note Note) Display() {
	fmt.Printf("your note title: %v has the following content: \n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") //to modify filename
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note) //to convert data to json format
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) //0644: read and write permissions

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Input cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
