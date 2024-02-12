package main

import (
	"context"
	"fmt"
	"httpserve/middleware"
	"httpserve/templates"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	noAuthRouter := router.PathPrefix("/").Subrouter()                         // for login, serving favicon, and static files
	noAuthRouter.PathPrefix("/static").Handler(http.FileServer(http.Dir("."))) // should i remove this?
	noAuthRouter.PathPrefix("/styles").Handler(http.FileServer(http.Dir("."))) // css shit
	noAuthRouter.PathPrefix("/js").Handler(http.FileServer(http.Dir(".")))     // js and dependencies
	noAuthRouter.PathPrefix("/assets").Handler(http.FileServer(http.Dir("."))) // hmmm, mostly images, banners and fonts shit

	noAuthRouter.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	}).Methods("GET")

	noAuthRouter.HandleFunc("/login", handleLogin).Methods("GET")
	noAuthRouter.HandleFunc("/login", handleLoginAuth).Methods("POST")
	noAuthRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	})

	noAuthRouter.HandleFunc("/thankyou", handleThankyou)

	authRouter := router.PathPrefix("/").Subrouter()      // for session, used for
	authRouter.Use(middleware.TokenAuthMiddleware)        // check all routes in this subrouter for cookie
	authRouter.Use(middleware.RefreshAuthTokenMiddleware) // refresh token on every request

	authRouter.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		template := templates.Form()
		if r.Header.Get("HX-Boosted") == "" {
			template = templates.WrapContent(template)
		}
		template.Render(context.Background(), w)
	}).Methods("GET", "POST")

	authRouter.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		template := templates.Home()
		if !isBoosted(r) {
			template = templates.WrapContent(template)
		}
		template.Render(context.Background(), w)
	})

	server := http.Server{
		Addr:              "127.0.0.1:8000",
		Handler:           router,
		ReadTimeout:       time.Second * 15,
		IdleTimeout:       time.Second * 15,
		WriteTimeout:      time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	template := templates.Login()
	if r.Header.Get("HX-Boosted") == "" {
		template = templates.WrapContent(template)
	}
	template.Render(context.Background(), w)
}

func handleLoginAuth(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")
	if username == "sivaram.asdf@gmail.com" && password == "123" {
		// write the cookie
		cookie := http.Cookie{
			Name:     "Authorization",
			Value:    "Bearer siva",
			Path:     "/",
			MaxAge:   int((time.Hour * 3).Seconds()),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
	}
}

func handleThankyou(w http.ResponseWriter, r *http.Request) {
	template := templates.ThankYou()
	if r.Header.Get("HX-Boosted") == "" {
		template = templates.WrapContent(template)
	}
	template.Render(context.Background(), w)
}

func isBoosted(r *http.Request) bool {
	return r.Header.Get("HX-Boosted") != ""
}
