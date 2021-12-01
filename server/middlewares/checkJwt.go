package middlewares

import (
	// "errors"
	// "log"
	"net/http"

	// jwtconfig "github.com/go-chat/server/jwtConfig"
	// "github.com/go-chat/server/utils"
	// "github.com/gorilla/mux"
)

func CheckJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// err:= jwtconfig.Verify(r)
		// if err != nil {
		// 	utils.ERROR(w, 401, errors.New("Unauthorized"))
		// 	return
		// }
		// cookie, err := r.Cookie("logged-in")
		// if err != nil {
		// 	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		// }
		// tokenJwt := cookie.Value
		//
		// err := jwtconfig.Verify(w, r)
		// if err != nil {
		// 	log.Println("has err")
		// 	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		// }
		next(w, r)
	}
}
