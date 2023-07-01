package server

import (
	"log"
	"net/http"
	"strconv"
	"telbot/pkg/repository"

	"github.com/zhashkevych/go-pocket-sdk"
)

type AuthorizationServer struct {
	server          *http.Server
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectUrl     string
}

func NewAuthorizationServer(pocketClient *pocket.Client, tokenRepository repository.TokenRepository, redirecturl string) *AuthorizationServer {
	return &AuthorizationServer{
		pocketClient:    pocketClient,
		tokenRepository: tokenRepository,
		redirectUrl:     redirecturl,
	}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s,
	}
	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIdParam := r.URL.Query().Get("chat_id")

	if chatIdParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIdParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("chat_id", chatID)
	requestToken, err := s.tokenRepository.Get(chatID, repository.RequesTokens)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authResponse, err := s.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.tokenRepository.Save(chatID, authResponse.AccessToken, repository.AccessTokens); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("chat_id : %d\nrequest_token: %s\naccess_token: %s\n", chatID, requestToken, authResponse.AccessToken)
	w.Header().Add("Location", s.redirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)

}
