package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)

		r.ParseForm()
		if len(r.Form) != 0 {
			for k := range r.Form {
				fmt.Printf("\t%v : %v\n", k, r.FormValue(k))
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	router.Use(LoggingMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	router.PathPrefix("/static").Handler(http.FileServer(http.Dir(".")))
	router.PathPrefix("/styles").Handler(http.FileServer(http.Dir(".")))
	router.PathPrefix("/js").Handler(http.FileServer(http.Dir(".")))

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
