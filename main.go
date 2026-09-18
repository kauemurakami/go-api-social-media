package main

import (
	app_config "api-social-media/internal/core/config"
	"api-social-media/internal/core/middlewares"
	"api-social-media/internal/data/providers/database"
	"api-social-media/internal/routes"
	"fmt"
	"log"
	"net/http"
)

// gerando secret p token
// func init() {
// 	key := make([]byte, 64)
// 	// slice de 64 posições será copulado com valores aleatorios
// 	// para gerar e verificar nossos tokens
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal("Erro na gerações")
// 	}
// 	fmt.Println(key) // array de eslices

// 	//convertendo numa string base64
// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64) // array de eslices

// }

func main() {
	app_config.SetupEnvironments()
	db.SetupDB()
	fmt.Printf("Run API :%s", app_config.API_port)
	router := routes.SetupAppRoutes()
	router.Use(middlewares.SetupHeadersMiddleware)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", app_config.API_port), router))

}
