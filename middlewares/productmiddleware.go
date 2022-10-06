package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vernandodev/go-restapi-jwt-mux/config"
	"github.com/vernandodev/go-restapi-jwt-mux/helper"
)

func ProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")

		// mengecek apakah token tersedia
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Akses tidak diizinkan"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		// jika token tersedia maka akan mengambil value dari token tersebut
		tokenString := c.Value

		claims := &config.JWTclaim{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				response := map[string]string{"message": "Akses tidak diizinkan"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				// token expired
				response := map[string]string{"message": "Akses tidak diizinkan, Token kadaluarsa"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Akses tidak diizinkan"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		if !token.Valid {
			response := map[string]string{"message": "Akses tidak diizinkan"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
