package middlewares

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/services"
	"net/http"
)

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := services.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
