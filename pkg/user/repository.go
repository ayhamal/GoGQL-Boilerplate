package user

import (
	"context"
	"log"

	"github.com/ayhamal/gogql-boilerplate/pkg/auth"
	"github.com/ayhamal/gogql-boilerplate/pkg/entities"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	// GetUserByEmail(email string) (*entities.User, error)
	// CreateUser(user *entities.User) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
	GetUserFromCtx(ctx context.Context) (*entities.User, error)
}

// Private repository struct reference to mongo collection
type repository struct {
	Client *pg.PgClient
}

// NewRepo is the single instance repo that is being created.
func NewRepo(pgClient *pg.PgClient) Repository {
	return &repository{
		Client: pgClient,
	}
}

// Static - GetUserFromCtx func helps to fetch user details from request context
func (r *repository) GetUserFromCtx(ctx context.Context) (*entities.User, error) {
	// Get user claims from context
	claims, _ := ctx.Value(auth.SessionKeyString("current_user")).(*auth.SessionClaims)
	// Validate user claims
	if claims == nil {
		log.Println("Gql - User not authored")
		return nil, &entities.InvalidOrExpiredTokenError{}
	}
	// Create user container
	var user *entities.User
	// Find user exists on database
	r.Client.Db.Find(&user, "email = ?", claims.Email)
	// Log user details
	log.Println(user)
	// Validate fetch not generated empty result
	if user.Email == "" {
		return nil, &entities.NotFoundDataError{}
	}
	// Return user details
	return user, nil
}
