package auth

import (
	"crypto/tls"
	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/pkg/entities"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
)

// Session key type
type SessionKeyString string

// Types definitions
type SessionClaims struct {
	Iss   uuid.UUID `json:"iss"`
	Email string    `json:"email"`
	Exp   int64     `json:"exp"`
	jwt.RegisteredClaims
}

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	SigninWithEmailAndPassword(email string, password string) (string, error)
	SignupWithEmailAndPassword(entities.User) (string, error)
	RestoreAccount(email string) error
}

// Private repository struct reference to mongo collection
type repository struct {
	Env    *env.Env
	Client *pg.PgClient
}

// NewRepo is the single instance repo that is being created.
func NewRepo(pgClient *pg.PgClient, env *env.Env) Repository {
	return &repository{
		Env:    env,
		Client: pgClient,
	}
}

//******

// LoginWithEmailAndPassword is a mongo repository that helps to create time series locations
func (r *repository) SigninWithEmailAndPassword(email string, password string) (string, error) {
	// Create user container
	var user *entities.User
	// Find user exists on database
	r.Client.Db.Find(&user, "email = ?", email)
	log.Println(user)
	// Validate fetch not generated empty result
	if user.Email == "" {
		return "", &entities.NotFoundDataError{}
	}
	// Validate password
	if !checkPasswordHash(password, user.Password) {
		return "", &entities.EmailOrPasswordError{}
	}
	// Create the Claims
	claims := SessionClaims{
		Iss:   uuid.New(),
		Email: user.Email,
		Exp:   time.Now().Add(time.Minute * 10).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(r.Env.App.SigningKey))
	// Handle token generation error
	if err != nil {
		return "", err
	}
	// Return token
	return t, nil
}

// CreateLocation is a mongo repository that helps to create locations
func (r *repository) SignupWithEmailAndPassword(user entities.User) (string, error) {
	// Fetch user from database
	var fetchedUser *entities.User
	// Make query to database fetching user by email
	r.Client.Db.First(&fetchedUser, "email = ?", user.Email)
	log.Println(fetchedUser.Email)
	// Validate user not exist on database
	if fetchedUser.Email != "" && fetchedUser.Email == user.Email {
		return "", &entities.EmailAlreadyExistsError{}
	}
	// Store temporally user password
	password := user.Password
	// Hash user password
	hash, err := hashPassword(password)
	// Validate not has error hashing user password
	if err != nil {
		return "", nil
	}
	// Hash user password
	user.Password = hash
	// Create user on database
	r.Client.Db.Create(&user)
	// Message.
	message := []byte("Created account successfully...")
	// Dispatch mail notification
	go r.SendMailNotification(user, message)
	// Signin new user & return creadentials
	return r.SigninWithEmailAndPassword(user.Email, password)
}

// ReadLocation is a mongo repository that helps to fetch locations
func (r *repository) RestoreAccount(email string) error {
	return nil
}

// CheckPasswordHash compare password with hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword creates a hashed instance of user password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Send Mail notification
func (r *repository) SendMailNotification(user entities.User, message []byte) error {
	log.Println(r.Env)
	// msg        := message
	recipients := []string{user.Email}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "grupo@priver.app") //r.Env.SmtpBroker.Username)

	// Set E-Mail receivers
	m.SetHeader("To", recipients...)

	// Set E-Mail subject
	m.SetHeader("Subject", "DIGIMED Notification")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", string(message))

	// Settings for SMTP server
	d := gomail.NewDialer(
		r.Env.SmtpBroker.Host,
		int(r.Env.SmtpBroker.Port),
		r.Env.SmtpBroker.Username,
		r.Env.SmtpBroker.Password,
	)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Email Sent Successfully!")
	return nil
}

// ******
// Method to validate and return *jwt.Token pointer
func JwtValidate(stringToken string, env *env.Env) (*jwt.Token, error) {
	log.Printf("Received token: %s", stringToken)
	// Validate request has token parameter
	if stringToken == "" {
		return nil, &entities.InvalidOrExpiredTokenError{}
	}

	// Create the JWT key used to create the signature
	jwtKey := []byte(env.App.SigningKey)
	// Process token
	token, err := jwt.ParseWithClaims(stringToken, &SessionClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Handle errors
	if err != nil {
		log.Printf("1. %s", err.Error())
		return nil, &entities.InvalidOrExpiredTokenError{}
	}

	// Validate token state
	if !token.Valid {
		return nil, &entities.InvalidOrExpiredTokenError{}
	}

	// Extract claims
	claims, ok := token.Claims.(*SessionClaims)
	expTime := time.Unix(claims.Exp, 0)
	log.Printf("Session is active: %t", time.Now().Before(expTime))

	// Validate token not expired
	if !time.Now().Before(expTime) || !ok {
		return nil, &entities.InvalidOrExpiredTokenError{}
	}

	// Return claims
	return token, nil
}
