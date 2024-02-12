package middleware

import (
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Authorization")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
			return
		}
		value := cookie.Value
		// fmt.Println("Got Auth header::", value)
		authorization := strings.Split(value, " ")
		if len(authorization) != 2 {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
			return
		}
		if !ValidateToken(authorization[1]) {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateToken(token string) bool {
	return (token == "siva")
}
