package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const PORT = "PORT"

type Author struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
}

func LoadEnv(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}

func createServer(port string) {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		var books = []Book{
			{
				ID:   1,
				Name: "Senhor dos Aneis e a sociedade do Anel",
				Author: Author{
					ID:     1,
					Name:   "J. R. R. Tolkien",
					Age:    81,
					Active: false,
				},
			},
			{
				ID:   2,
				Name: "Harry Potter e o Enigma do Princípe",
				Author: Author{
					ID:     1,
					Name:   "J. K. Rowling",
					Age:    55,
					Active: true,
				},
			},
		}

		content, _ := json.Marshal(books)
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(content))
	})

	http.ListenAndServe(":"+port, nil)
}

func main() {
	port, _ := LoadEnv(PORT)
	createServer(port)
}
