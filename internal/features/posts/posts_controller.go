package posts

import (
	functions "api-social-media/internal/features/posts/functions"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	functions.CreatePost(w, r)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	functions.GetPosts(w, r)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	functions.GetPost(w, r)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	functions.DeletePost(w, r)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	functions.UpdatePost(w, r)
}
func LikePost(w http.ResponseWriter, r *http.Request) {
	functions.LikePost(w, r)
}
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	functions.UnlikePost(w, r)
}
