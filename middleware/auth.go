package middleware

import (
	"net/http"
	"os"
	"rest-api-simple-auth/helpers"
	"strings"
)

/*
 * @visibility	private
 * @return		bool
 */
func headerAuthorization(r *http.Request) bool {

	secretKeys := strings.Split(os.Getenv("SECRET_KEY"), ",")

	if headerAuth := r.Header.Get("Authorization"); headerAuth != "" && secretKeys != nil {

		for _, value := range secretKeys {
			if value == headerAuth {
				return true
			}
		}
	}

	return false
}

//
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !headerAuthorization(r) {
			helpers.ErrorJsonResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
