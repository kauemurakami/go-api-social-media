package users_functions

import (
	responses "api-social-media/internal/common/utils/response"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/domain/models"
	"context"
	"errors"
	"net/http"
	"strings"
)

func GetUsersByNickOrName(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Corrigir a query SQL
	query := "SELECT * FROM users WHERE lower(name) LIKE $1 OR lower(nick) LIKE $1"

	// Extrair o parâmetro de consulta 'user'
	user_param := r.URL.Query().Get("user")
	if user_param == "" {
		responses.Err(w, http.StatusBadRequest, errors.New("favor informar palavra chave para pesquisa"))
		return
	}

	nameOrNick := "%" + strings.ToLower(user_param) + "%"
	// abc
	rows, err := conn.Query(context.Background(), query, nameOrNick)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Pass, &user.CreatedAt); err != nil {
			responses.Err(w, http.StatusBadRequest, err)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	// Se não houver usuários encontrados, retornar um array vazio de usuários
	if len(users) == 0 {
		users = []models.User{}
	}

	responses.JSON(w, http.StatusOK, users)
}
