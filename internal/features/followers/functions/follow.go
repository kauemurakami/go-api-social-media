package followers_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())
	// Init transaction
	tx, err := conn.Begin(context.Background())
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer tx.Rollback(context.Background())
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
	}
	var followers models.Followers
	if err = json.Unmarshal(body, &followers); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
	}
	query := "INSERT INTO followers (user_id, follower_id) VALUES ($1, $2) RETURNING *"

	err = tx.QueryRow(context.Background(),
		query,

		followers.UserID,
		followers.FollowerID,
	).Scan(
		&followers.ID,
		&followers.UserID,
		&followers.FollowerID,
	)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}
	// Commit the transaction
	if err = tx.Commit(context.Background()); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}
