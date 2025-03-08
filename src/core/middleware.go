// middleware.go
package core

import (
	"net/http"
)

func MuxCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Permitir todos los métodos HTTP
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")

		// Permitir encabezados personalizados
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")

		// Permitir credenciales
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Manejar solicitudes preflight (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pasar la solicitud al siguiente handler
		next.ServeHTTP(w, r)
	})
}
