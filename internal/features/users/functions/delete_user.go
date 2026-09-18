package users_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extrair o ID do usuário da rota
	userID := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Excluir o usuário do banco de dados
	query := "DELETE FROM users WHERE id = $1"
	_, err = conn.Exec(context.Background(), query, userUUID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Responder com um JSON indicando o sucesso da exclusão
	response := map[string]bool{"success": true}
	json.NewEncoder(w).Encode(response)
}
