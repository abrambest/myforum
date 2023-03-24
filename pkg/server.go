package pkg

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)

	log.Println("Запуск сервера на http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
