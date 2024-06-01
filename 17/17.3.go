package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	logFileName := time.Now().String()
	logFileName = logFileName[:19] + " log.log"
	logFileName = strings.ReplaceAll(logFileName, ":", "-")
	fmt.Println("логирование в: ", logFileName)

	logFile, err := os.Create(logFileName)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer func() {
		logFile.Close()
	}()

	l := log.New(logFile, "httpServer: ", log.Ldate|log.Ltime)
	l.Println("start")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(l, logHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}
func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
	// l.Println(msg)
}
func authHandler(l *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")

		if xId != "my_secret" {
			l.Println("аларм. пользователь не авторизован")

			http.Error(w, "пользователь не авторизован",
				http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
