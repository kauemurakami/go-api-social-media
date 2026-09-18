package followers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupFollowersRoutes(router *mux.Router) {

	router.HandleFunc("/follow", FollowUser).Methods(http.MethodPost)
	router.HandleFunc("/unfollow", UnfollowUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}/my-followers", MyFollowers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}/followed-by-me", FollowedByMe).Methods(http.MethodGet)
}
