package auth

import (
	"context"
	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/graph/model"
	"github.com/ayhamal/gogql-boilerplate/pkg/entities"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"
	"log"
)

func LoginWithEmailAndPasswordResolver(_ context.Context, email string, password string) (*model.AuthData, error) {
	// Load env variables
	env, err := env.New()
	// Check for errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Get pg client instance
	pgClient, err := pg.GetInstance()
	// Check for errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Create repository instance
	userRepo := NewRepo(pgClient, env)
	// Create service instance
	authService := NewService(userRepo)
	// Login user with email and password
	token, err := authService.LoginWithEmailAndPassword(email, password)
	// Check for errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Return results
	return &model.AuthData{
		Token: token,
	}, nil
}

func CreateUserResolver(_ context.Context, userInput model.UserInput) (*model.AuthData, error) {
	// Load env variables
	env, err := env.New()
	// Check for errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Get pg client instance
	pgClient, err := pg.GetInstance()
	// Check for errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Create repository instance
	userRepo := NewRepo(pgClient, env)
	// Create service instance
	authService := NewService(userRepo)
	// Create new user
	token, err := authService.RegisterWithEmailAndPassword(entities.User{
		FullName:             userInput.FullName,
		Email:                userInput.Email,
		Password:             userInput.Password,
		Gender:               string(userInput.Gender),
		Ocupation:            userInput.Ocupation,
		PhoneNumber:          userInput.PhoneNumber,
		CountryCode:          userInput.CountryCode,
		Weight:               userInput.Weight,
		Height:               userInput.Height,
		Birthday:             userInput.Birthday,
		IdentificationType:   string(userInput.IdentificationType),
		IdentificationNumber: userInput.IdentificationNumber,
	})
	// Check for errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Return results
	return &model.AuthData{
		Token: token,
	}, nil
}
