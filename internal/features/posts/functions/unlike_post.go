package posts_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	// Obter o ID do post da URL usando o Mux
	postID := mux.Vars(r)["id"]

	// Converter o postID para o tipo UUID
	id, err := uuid.Parse(postID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("ID do post inválido"))
		return
	}

	// Inicializar a conexão com o banco de dados
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Consulta SQL para decrementar o contador de likes do post, garantindo que nunca seja menor que zero
	query := `
		UPDATE posts
		SET likes = GREATEST(likes - 1, 0)
		WHERE id = $1
	`

	// Executar a consulta SQL para decrementar o contador de likes
	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Responder com uma mensagem de sucesso
	responses.JSON(w, http.StatusOK, map[string]string{"message": "Like removido com sucesso"})
}
