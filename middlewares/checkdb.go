package middlewares

import (
	"net/http"

	"github.com/nahuelsv/twitter-backend/bd"
)

/*CheckDB middleware to check database connection */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Lost connection with Database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
