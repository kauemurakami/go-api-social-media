package followers

import (
	functions "api-social-media/internal/features/followers/functions"
	"net/http"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	functions.FollowUser(w, r)
}
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	functions.UnfollowUser(w, r)
}
func FollowedByMe(w http.ResponseWriter, r *http.Request) {
	functions.FollowedByMe(w, r)
}
func MyFollowers(w http.ResponseWriter, r *http.Request) {
	functions.MyFollowers(w, r)
}
