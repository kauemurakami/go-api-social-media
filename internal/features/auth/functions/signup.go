package auth_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/data/services"
	"api-social-media/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {

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
	err = tx.QueryRow(context.Background(),
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

	token, err := services.Createtoken(user.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}
	fmt.Println(token, user.ID)
	insertTokenQuery := "INSERT INTO user_token (token, user_id) VALUES ($1, $2)"
	_, err = tx.Exec(context.Background(), insertTokenQuery,
		token,
		insertedUser.ID,
	)

	if err != nil {
		fmt.Println(err)
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Commit the transaction
	if err = tx.Commit(context.Background()); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]interface{}{
		"user": map[string]string{
			"id":    insertedUser.ID.String(),
			"name":  insertedUser.Name,
			"email": insertedUser.Email,
			"nick":  insertedUser.Nick,
		},
		"token": token,
	}

	responses.JSON(w, http.StatusOK, response)
}
