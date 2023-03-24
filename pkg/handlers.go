package pkg

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello!!!"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("show snippet!!!"))
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("GET - method prohibited!"))
	}
	w.Write([]byte("create snippet!!!"))
}
