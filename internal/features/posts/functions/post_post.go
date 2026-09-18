package posts_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Decodificar o JSON da solicitação para um struct Post
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Verificar se todos os campos necessários estão presentes
	if post.AuthorID == uuid.Nil || post.Title == "" || post.Content == "" || post.AuthorNick == "" {
		http.Error(w, "Preencha todos os campos obrigatórios", http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Preparar a consulta SQL para inserir o post
	query := `
		INSERT INTO posts (title, content, author_id, author_nick)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	// Executar a consulta SQL e escanear o resultado para obter o ID e a data de criação do post
	var insertedPost models.Post
	err = conn.QueryRow(context.Background(), query,
		post.Title, post.Content, post.AuthorID, post.AuthorNick,
	).Scan(
		&insertedPost.ID, &insertedPost.CreatedAt,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Atualizar os campos do post inserido com os campos do post enviado
	insertedPost.Title = post.Title
	insertedPost.Content = post.Content
	insertedPost.AuthorID = post.AuthorID
	insertedPost.AuthorNick = post.AuthorNick
	insertedPost.Likes = post.Likes

	// Responder com o post inserido e atualizado
	responses.JSON(w, http.StatusOK, insertedPost)

}
