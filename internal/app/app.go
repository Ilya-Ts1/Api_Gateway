package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"



	"github.com/go-chi/chi"
	pb "api_gateway/protos/gen/auth"
	controller "api_gateway/internal/controller"
)

func Run() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conent-Type", "text/plain")
		w.Write([]byte("Hello i am apigateway"))
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error open api_gateway http localhost:8080")
	}
}

type AuthMiddleware struct {
	authClient *controller.AuthClient
}

func NewAuthMiddleware(authClient *controller.AuthClient) *AuthMiddleware {
	return &AuthMiddleware{authClient: authClient}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем Bearer token
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Вызываем gRPC auth сервис
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		resp, err := m.authClient.Client.ValidateToken(ctx, &pb.ValidateTokenRequest{Token: token})
		if err != nil || resp.Valid == false {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Добавляем userID в контекст для downstream handlers
		ctx = context.WithValue(r.Context(), "userID", resp.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
