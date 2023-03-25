package web

import (
	"forum_alem_01/pkg"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Server() {
	port := ""
	if len(os.Args) > 2 {
		log.Println("Usage: ./cmd $port")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	// mux.Handle("/static", http.NotFoundHandler()) ???
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)

	pkg.LogInfo("Запуск сервера на http://localhost:" + port)
	// infoLog.Println("Запуск сервера на http://localhost:" + port)

	err := http.ListenAndServe(":"+port, mux)
	pkg.LogError(err)
	errorLog.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
