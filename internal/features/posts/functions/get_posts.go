package posts_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Inicializar a conexão com o banco de dados
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Consulta SQL para obter todos os posts
	query := `
		SELECT *
		FROM posts
	`

	// Executar a consulta SQL
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	// Slice para armazenar todos os posts
	var posts []models.Post

	// Iterar sobre os resultados da consulta
	for rows.Next() {
		var post models.Post
		// Escanear os dados do post do banco de dados para a estrutura Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AuthorNick, &post.Likes, &post.CreatedAt); err != nil {
			responses.Err(w, http.StatusInternalServerError, err)
			return
		}
		// Adicionar o post ao slice de posts
		posts = append(posts, post)
	}

	// Verificar por erros no final do loop
	if err := rows.Err(); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	// Responder com a lista de posts
	responses.JSON(w, http.StatusOK, posts)
}
