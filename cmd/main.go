package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"github.com/smailic05/diploma/internal/handler"
	repository "github.com/smailic05/diploma/internal/repository"
	"github.com/smailic05/diploma/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	logger.Info().Msgf("DEBUG: a sample jwt is %s", tokenString)

	dsn := "host=localhost user=postgres password=postgres dbname=backend port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := chi.NewRouter()
	authRepo := repository.NewAuth(db)
	authService := service.NewAuth(&logger, authRepo, []byte("secret"))

	repo := repository.New()
	serv := service.New(&logger, repo)
	h := handler.New(&logger, serv, authService)

	r.Group(func(r chi.Router) {
		r.Use(handler.JWT([]byte("secret")))

		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Get("/download", h.Download)
		r.Get("/auth", h.Download)
	})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Err(err)
	}
}
