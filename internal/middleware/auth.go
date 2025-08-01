package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AUTH_USER, ok := os.LookupEnv("AUTH_USER")
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

		}
		AUTH_PASS, ok := os.LookupEnv("AUTH_PASS")
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsername := sha256.Sum256([]byte(AUTH_USER))
			expectedPassword := sha256.Sum256([]byte(AUTH_PASS))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsername[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPassword[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
