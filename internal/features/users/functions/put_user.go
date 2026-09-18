package users_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do usuário da rota
	userID := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Ler o corpo da solicitação
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Decodificar o corpo JSON para obter os dados atualizados do usuário
	var updatedUser models.User
	if err := json.Unmarshal(body, &updatedUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	// Consultar o banco de dados para obter os dados do usuário atual
	var currentUser models.User
	query := "SELECT * FROM users WHERE id = $1"
	err = conn.QueryRow(context.Background(), query, userUUID).Scan(
		&currentUser.ID, &currentUser.Name, &currentUser.Nick, &currentUser.Email, &currentUser.Pass, &currentUser.CreatedAt,
	)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Atualizar os campos do usuário com base nos dados fornecidos na solicitação
	if updatedUser.Name != "" {
		currentUser.Name = updatedUser.Name
	}
	if updatedUser.Nick != "" {
		currentUser.Nick = updatedUser.Nick
	}
	if updatedUser.Email != "" {
		currentUser.Email = updatedUser.Email
	}
	if updatedUser.Pass != "" {
		currentUser.Pass = updatedUser.Pass
	}

	// Atualizar o usuário no banco de dados
	query = "UPDATE users SET name = $1, nick = $2, email = $3, pass = $4 WHERE id = $5 RETURNING *"
	err = conn.QueryRow(context.Background(), query, currentUser.Name, currentUser.Nick, currentUser.Email, currentUser.Pass, currentUser.ID).Scan(
		&currentUser.ID, &currentUser.Name, &currentUser.Nick, &currentUser.Email, &currentUser.Pass, &currentUser.CreatedAt,
	)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, currentUser)
}
