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

// Quem me segue
func MyFollowers(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Read the user ID from the request body
	userID := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("ID do usuário inválido"))
		return
	}

	query := `
		SELECT u.id, u.name, u.nick, u.email
		FROM followers f
		JOIN users u ON f.follower_id = u.id
		WHERE f.user_id = $1
	`

	rows, err := conn.Query(context.Background(), query, userUUID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User
		if err := rows.Scan(&follower.ID, &follower.Name, &follower.Nick, &follower.Email); err != nil {
			responses.Err(w, http.StatusInternalServerError, err)
			return
		}
		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}
