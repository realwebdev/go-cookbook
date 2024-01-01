// authentication service using Go kit. This service will have endpoint for user
// login and authentication.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type authServiceImpl struct{}

func (authServiceImpl) Login(ctx context.Context, username, password string) (string, error) {
	if username == "user" && password == "pass" {
		return "valid_token", nil
	}
	return "", fmt.Errorf("authentication failed")
}

func MakeLoginEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		token, err := svc.Login(ctx, req.Username, req.Password)
		return loginResponse{Token: token, Err: err}, nil
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
	Err   error  `json:"error,omitempty"`
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, nil
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := authServiceImpl{}

	loginEndpoint := MakeLoginEndpoint(svc)
	loginHandler := httptransport.NewServer(
		loginEndpoint,
		decodeLoginRequest,
		encodeResponse,
	)

	http.Handle("/login", loginHandler)

	// start the HTTP server

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
