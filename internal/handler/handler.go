package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/smailic05/diploma/internal/model"
)

const (
	//TODO delete const
	params = "Persian.glb"

	downloadError = "Error when trying to download"
)

type Handler struct {
	logger      *zerolog.Logger
	service     Service
	authService AuthService
}

type AuthService interface {
	Authenticate(string, string) (string, string, error)
}

type Service interface {
	Download(context.Context, string) ([]byte, error)
}

func New(logger *zerolog.Logger, srv Service, authSrv AuthService) *Handler {
	return &Handler{
		logger:      logger,
		service:     srv,
		authService: authSrv,
	}
}

func (h *Handler) Download(w http.ResponseWriter, r *http.Request) {
	//TODO params here
	data, err := h.service.Download(r.Context(), params)
	if err != nil {
		h.logger.Error().Err(err).Msg(downloadError)
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Bad request"})
	}
	//TODO decompress here
	writeResponseFile(w, http.StatusOK, params, data)

}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	req := &model.AuthRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		h.logger.Error().Err(err).Msg("Invalid incoming data")
		writeResponse(w, http.StatusBadRequest, model.Error{Error: "Bad request"})
		return
	}

	accessToken, refreshToken, err := h.authService.Authenticate(req.Username, req.Password)
	if err != nil {
		h.logger.Error().Err(err).Msg("Authentication error")
		writeResponse(w, http.StatusForbidden, model.Error{Error: "Forbidden"})
		return
	}

	res := &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	writeResponse(w, http.StatusOK, res)
}
