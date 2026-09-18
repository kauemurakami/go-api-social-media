package auth

import (
	functions "api-social-media/internal/features/auth/functions"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	functions.Signin(w, r)
}
func Signup(w http.ResponseWriter, r *http.Request) {
	functions.Signup(w, r)
}
