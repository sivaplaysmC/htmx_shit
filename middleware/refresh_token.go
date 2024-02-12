package middleware

import (
	"net/http"
	"time"
)

func RefreshAuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("Authorization")

		newCookie := *cookie

		newCookie.MaxAge = int((time.Hour * 3).Seconds())
		newCookie.SameSite = http.SameSiteStrictMode
		http.SetCookie(w, &newCookie)
		next.ServeHTTP(w, r)
	})
}
