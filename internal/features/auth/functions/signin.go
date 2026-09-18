package auth_functions

import (
	responses "api-social-media/internal/common/utils/response"
	sec "api-social-media/internal/core/security"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/jackc/pgx"
)

func Signin(w http.ResponseWriter, r *http.Request) {

	conn := db.SetupDB()
	defer conn.Close(context.Background())
	tx, err := conn.Begin(context.Background())
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer tx.Rollback(context.Background())
	// Ler o corpo da requisição
	var requestBody struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("erro ao ler o corpo da requisição"))
		return
	}

	email := strings.ToLower(requestBody.Email)
	if email == "" {
		responses.Err(w, http.StatusBadRequest, errors.New("favor informar o email para pesquisa"))
		return
	}

	// Consultar o banco de dados para obter o usuário pelo email
	query := "SELECT id, email, pass FROM users WHERE lower(email) = $1"
	var user models.User
	err = tx.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Email, &user.Pass)
	if err != nil {
		if err == pgx.ErrNoRows {
			responses.Err(w, http.StatusNotFound, errors.New("usuário não encontrado"))
		} else {
			responses.Err(w, http.StatusInternalServerError, err)
		}
		return
	}

	/// dividir funções
	if err = sec.ComparePassHash(requestBody.Pass, user.Pass); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
	}

	query = "SELECT id, name, email, pass, nick, created_at FROM users WHERE id = $1"
	err = tx.QueryRow(context.Background(), query, user.ID).Scan(
		&user.ID, &user.Name, &user.Email, &user.Pass, &user.Nick, &user.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			responses.Err(w, http.StatusNotFound, errors.New("usuário não encontrado"))
		} else {
			responses.Err(w, http.StatusInternalServerError, err)
		}
		return
	}

	query = "SELECT token FROM user_token WHERE user_id = $1"
	var token string
	err = tx.QueryRow(context.Background(), query, user.ID).Scan(&token)
	if err != nil {
		if err == pgx.ErrNoRows {
			responses.Err(w, http.StatusNotFound, errors.New("token não encontrado para o usuário"))
		} else {
			responses.Err(w, http.StatusInternalServerError, err)
		}
		return
	}

	//Responder com um JSON contendo o ID, email e senha do usuário
	response := map[string]interface{}{
		"user": map[string]string{
			"id":    user.ID.String(),
			"name":  user.Name,
			"email": user.Email,
			"nick":  user.Nick,
		},
		"token": token,
	}

	responses.JSON(w, http.StatusOK, response)

}
