package web

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// func LogInfo(info string) {
// 	f, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
// 	infoLog.Println(info)
// 	fmt.Println(info)
// }

// func LogError(er error) {
// 	f, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	infoLog.Println(er)
// }
