package followers_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Quem eu sigo
func FollowedByMe(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Read the user ID from the URL
	userID := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("ID do usuário inválido"))
		return
	}

	// Query to get all users followed by the given user

	query := `
		SELECT u.id, u.name, u.nick, u.email
		FROM followers f
		JOIN users u ON f.follower_id = u.id
		WHERE f.follower_id = $1
	`

	rows, err := conn.Query(context.Background(), query, userUUID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var followedUsers []models.User
	for rows.Next() {
		var followed models.User
		if err := rows.Scan(&followed.ID, &followed.Name, &followed.Nick, &followed.Email); err != nil {
			responses.Err(w, http.StatusInternalServerError, err)
			return
		}
		followedUsers = append(followedUsers, followed)
	}

	if err := rows.Err(); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	var response interface{}
	if len(followedUsers) == 0 {
		response = []models.User{}
	} else {
		response = followedUsers
	}
	responses.JSON(w, http.StatusOK, response)
}
