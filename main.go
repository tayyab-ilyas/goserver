package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()
	corsMux := middlewareCors(m)
	m.HandleFunc("/", handlePage)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	addr := ":" + port

	srv := http.Server{
		Handler:      corsMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("server started on port ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Central-Allow-Methods", "GET")
	w.WriteHeader(200)
	const page = `<html><body><p>Hello From goserver</p></body></html>`

	
	w.Write([]byte(page))
}
func middlewareCors(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}