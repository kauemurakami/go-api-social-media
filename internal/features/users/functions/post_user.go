package users_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	conn := db.SetupDB()
	defer conn.Close(context.Background())
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
	}
	if err = user.Prepare("cadastro"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	query := "INSERT INTO users (name, email, pass, nick) VALUES ($1, $2, $3, $4) RETURNING *"
	var insertedUser models.User
	err = conn.QueryRow(context.Background(),
		query,
		user.Name,
		user.Email,
		user.Pass,
		user.Nick,
	).Scan(
		&insertedUser.ID,
		&insertedUser.Name,
		&insertedUser.Nick,
		&insertedUser.Email,
		&insertedUser.Pass,
		&insertedUser.CreatedAt,
	)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)

	}
	responses.JSON(w, http.StatusOK, insertedUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
}
