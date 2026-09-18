package posts_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro da URL usando o Mux
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

	// Consulta SQL para obter o post pelo ID
	query := `
		SELECT id, title, content, author_id, author_nick, likes, created_at
		FROM posts
		WHERE id = $1
	`

	// Executar a consulta SQL
	row := conn.QueryRow(context.Background(), query, id)

	// Estrutura para armazenar o post recuperado
	var post models.Post

	// Escanear os dados do post do banco de dados para a estrutura Post
	err = row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AuthorNick, &post.Likes, &post.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			responses.Err(w, http.StatusNotFound, errors.New("post não encontrado"))
		} else {
			responses.Err(w, http.StatusInternalServerError, err)
		}
		return
	}

	// Responder com o post recuperado
	responses.JSON(w, http.StatusOK, post)
}
