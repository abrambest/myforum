package pkg

import (
	"fmt"
	"log"
	"os"
)

func LogInfo(info string) {
	f, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Println(info)
	fmt.Println(info)
}

func LogError(er error) {
	f, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Println(er)
}
