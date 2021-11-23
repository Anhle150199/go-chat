package middlewares

import (
	// "errors"
	"net/http"

	// jwtconfig "github.com/go-chat/server/jwtConfig"
	// "github.com/go-chat/server/utils"
	// "github.com/gorilla/mux"
)

func CheckJwt(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		// err:= jwtconfig.Verify(r)
		// if err != nil {
		// 	utils.ERROR(w, 401, errors.New("Unauthorized"))
		// 	return
		// }
		next(w, r)
	}
}	