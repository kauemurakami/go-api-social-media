package routes

import (
	"api-social-media/internal/features/auth"
	"api-social-media/internal/features/followers"
	"api-social-media/internal/features/posts"
	"api-social-media/internal/features/users"

	"github.com/gorilla/mux"
)

func SetupAppRoutes() *mux.Router {
	router := mux.NewRouter()
	auth.SetupAuthRoutes(router)
	users.SetupUserRoutes(router)
	followers.SetupFollowersRoutes(router)
	posts.SetupPostsRoutes(router)
	return router
}
