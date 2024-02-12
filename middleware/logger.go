package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s strings.Builder
		now := time.Now()
		hour, minute, sec := now.Clock()
		s.WriteString(fmt.Sprintf("%d:%d:%d | %s %s |\n", hour, minute, sec, r.Method, r.URL.Path))
		// r.ParseForm()
		// if len(r.Form) != 0 {
		// 	for k := range r.Form {
		// 		s.WriteString(fmt.Sprintf("\t%v : %v\n", k, r.FormValue(k)))
		// 	}
		// }
		fmt.Print(s.String())
		next.ServeHTTP(w, r)
	})
}
