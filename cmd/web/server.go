package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

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

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	// mux.Handle("/static", http.NotFoundHandler()) ???
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/create", app.CreateSnippet)

	infoLog.Println("Запуск сервера на http://localhost:" + port)

	err := http.ListenAndServe(":"+port, mux)

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
