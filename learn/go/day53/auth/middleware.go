package auth

import "net/http"

func RequireRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("X-Role")
		if got == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if got != role {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
