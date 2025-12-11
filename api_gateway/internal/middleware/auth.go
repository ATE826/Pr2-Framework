package middleware

import (
	"net/http"
	"strings"

	"pr2/pkg/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем свободный доступ к регистрации и логину
		if strings.HasPrefix(r.URL.Path, "/v1/users/register") || strings.HasPrefix(r.URL.Path, "/v1/users/login") {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Передаем user_id в заголовке
		r.Header.Set("X-User-ID", claims["user_id"].(string))
		next.ServeHTTP(w, r)
	})
}
