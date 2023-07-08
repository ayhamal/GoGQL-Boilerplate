package auth

import (
	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/pkg/entities"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"

)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	LoginWithEmailAndPassword(email string, password string) (string,error)
	RegisterWithEmailAndPassword(user entities.User) (string,error)
	RestoreAccount(email string) error
}

// Private service struct reference to mongo collection
type service struct {
	repository Repository
}

// New is used to create a Router & single instance of the service
func New(pgClient *pg.PgClient, env *env.Env) (Service) {
	// Create repository instance
	authRepo := NewRepo(pgClient, env)
	// Create service instance
	authService := NewService(authRepo)
	// Create user api
	// apiAuth := fiberApp.Group("/auth")
	// Return results
	return authService
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// LoginWithEmailAndPassword is s service layer to provide session using email and password
func (s *service) LoginWithEmailAndPassword(email string, password string) (string,error) {
	return s.repository.SigninWithEmailAndPassword(email, password)
}

// RegisterWithEmailAndPassword is a service layer that helps register new user using email and password
func (s *service) RegisterWithEmailAndPassword(user entities.User) (string,error) {
	return s.repository.SignupWithEmailAndPassword(user)
}

// RestoreAccount is a service layer that helps user restore accound
func (s *service) RestoreAccount(email string) error {
	return s.repository.RestoreAccount(email)
}

