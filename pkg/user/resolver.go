package user

import (
	"context"
	"log"

	"github.com/ayhamal/gogql-boilerplate/graph/model"
)

func GetUserFromCtxResolver(ctx context.Context) (*model.User, error) {
	// Get user service instance
	usersService, err := GetInstance()
	// Handle possible errors
	if err != nil {
		log.Println("Gql - Error getting user service instance: ", err)
		return nil, err
	}
	// Get user from context
	user, err := usersService.GetUserFromCtx(ctx)
	// Validate fetch not has errors
	if err != nil {
		return nil, err
	}
	// Result user
	ur := user.ToGql()
	// Resolve user
	return ur, nil
}
