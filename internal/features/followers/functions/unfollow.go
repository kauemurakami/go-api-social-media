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

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Iniciar a transação
	tx, err := conn.Begin(context.Background())
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer tx.Rollback(context.Background()) // Rollback é chamado se a transação não for confirmada

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
	}
	var followers models.Followers
	if err = json.Unmarshal(body, &followers); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
	}
	// Executar a operação DELETE na tabela followers
	deleteQuery := "DELETE FROM followers WHERE user_id = $1 AND follower_id = $2"
	_, err = tx.Exec(context.Background(), deleteQuery, followers.UserID, followers.FollowerID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Confirmar a transação
	if err = tx.Commit(context.Background()); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Responder com uma resposta de sucesso
	responses.JSON(w, http.StatusOK, map[string]bool{"success": true})
}
