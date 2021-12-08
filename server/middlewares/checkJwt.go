package middlewares

import (
	"context"
	"log"
	"net/http"

	jwtconfig "github.com/go-chat/server/jwtConfig"
)

func CheckJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		log.Println("Middlewares")
		user, err := jwtconfig.Verify(w, r)
		// log.Println(user)

		if err != nil {
			log.Println("Verify fail!!: ", err)
			http.Redirect(w, r, "user/login", http.StatusMovedPermanently)
		}
		context := context.WithValue(r.Context(), "User", user)
		r = r.WithContext(context)
		next.ServeHTTP(w, r)
	})
}
