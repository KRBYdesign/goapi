package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/KRBYdesign/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
  /$$$$$$   /$$$$$$         /$$$$$$  /$$$$$$$  /$$$$$$
 /$$__  $$ /$$__  $$       /$$__  $$| $$__  $$|_  $$_/
| $$  \__/| $$  \ $$      | $$  \ $$| $$  \ $$  | $$  
| $$ /$$$$| $$  | $$      | $$$$$$$$| $$$$$$$/  | $$  
| $$|_  $$| $$  | $$      | $$__  $$| $$____/   | $$  
| $$  \ $$| $$  | $$      | $$  | $$| $$        | $$  
|  $$$$$$/|  $$$$$$/      | $$  | $$| $$       /$$$$$$
 \______/  \______/       |__/  |__/|__/      |______/
 `)

	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Error(err)
	}
}