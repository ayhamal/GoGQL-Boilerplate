package user

import (
	"context"
	"log"

	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/pkg/entities"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	// GetUserByEmail(email string) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
	// EnrollUserToCompany(ctx context.Context, enrollment entities.Enrollment) (*entities.User, error)
	GetUserFromCtx(ctx context.Context) (*entities.User, error)
}

// Private service struct reference to mongo collection
type service struct {
	repository Repository
}

// GetInstance is used to create a Router & single instance of the service
func GetInstance() (Service, error) {
	// Get mongodb instance
	pgClient, err := pg.GetInstance()
	// Handle get instance error
	if err != nil {
		log.Println("UserService - Cannot get mongo users instance...")
		return nil, err
	}
	// Create repository instance
	usersRepo := NewRepo(pgClient)
	// Create service instance
	usersService := NewService(usersRepo)
	// Return results
	return usersService, nil
}

// New is used to create a Router & single instance of the service
func New(pgClient *pg.PgClient, env *env.Env) Service {
	// Create repository instance
	userRepo := NewRepo(pgClient)
	// Create service instance
	userService := NewService(userRepo)
	// Create user api
	// apiAuth := fiberApp.Group("/users")
	// Return results
	return userService
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// GetUserFromCtx used to fetch user data based on request context
func (s *service) GetUserFromCtx(ctx context.Context) (*entities.User, error) {
	return s.repository.GetUserFromCtx(ctx)
}
