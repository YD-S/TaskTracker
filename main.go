package main

import (
	"fmt"
)

type Data struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func main() {
	var dbFile string

	fmt.Println("Enter database file name (e.g., data.json):")
	_, err := fmt.Scanln(&dbFile)
	if err != nil {
		return
	}
}
