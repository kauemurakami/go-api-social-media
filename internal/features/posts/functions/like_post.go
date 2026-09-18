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

func LikePost(w http.ResponseWriter, r *http.Request) {
	// Obter o ID do post da URL usando o Mux
	postID := mux.Vars(r)["id"]

	// Converter o postID para o tipo UUID
	id, err := uuid.Parse(postID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("id do post inválido"))
		return
	}

	// Inicializar a conexão com o banco de dados
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Consulta SQL para incrementar o contador de likes do post
	query := `
		UPDATE posts
		SET likes = likes + 1
		WHERE id = $1
	`

	// Executar a consulta SQL para incrementar o contador de likes
	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Responder com uma mensagem de sucesso
	responses.JSON(w, http.StatusOK, map[string]string{"message": "Like adicionado com sucesso"})
}
