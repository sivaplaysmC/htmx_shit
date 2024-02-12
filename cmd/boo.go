package main

import (
	"fmt"
	"httpserve/templates"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:     "auth_token",
			Value:    "Bear grylls",
			Path:     "/",
			Expires:  time.Now().Add(time.Second * 60),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		fmt.Fprintln(w, "<a href='/page2'> goto page2 </a>")
	}).Methods(http.MethodGet)

	router.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<a href='/page1'> goto page1 </a>")
	})

	router.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The End!!")
	})

	http.ListenAndServe(":8080", router)
}

func WrapComponent(component templ.Component, r *http.Request) templ.Component {
	if r.Header.Get("HX-Boosted") == "" {
		return templates.WrapContent(component)
	}
	return component
}
