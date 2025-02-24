package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("input valid input")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func Save(note *Note) error {
	json, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile("project.json", json, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
