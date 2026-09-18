package posts //use same package name inside your package
import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupPostsRoutes(router *mux.Router) {
	router.HandleFunc("/posts", CreatePost).Methods(http.MethodPost)
	router.HandleFunc("/posts", GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", GetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", DeletePost).Methods(http.MethodDelete)
	router.HandleFunc("/posts", UpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/posts/{id}/like", LikePost).Methods(http.MethodPut)
	router.HandleFunc("/posts/{id}/unlike", UnlikePost).Methods(http.MethodPut)
}
