// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type AuthData struct {
	Token string `json:"Token"`
}

type Message struct {
	ID      string `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
}

// The `User` type defines the queryable fields for every user in our data source.
type User struct {
	ID                   int                `json:"id"`
	FullName             string             `json:"full_name"`
	Email                string             `json:"email"`
	Role                 Role               `json:"role"`
	Gender               Gender             `json:"gender"`
	IdentificationType   IdentificationType `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	CountryCode          string             `json:"country_code"`
	PhoneNumber          string             `json:"phone_number"`
	Ocupation            string             `json:"ocupation"`
	Weight               *float64           `json:"weight,omitempty"`
	Height               *float64           `json:"height,omitempty"`
	Birthday             *time.Time         `json:"birthday,omitempty"`
	ProfileImgURL        *string            `json:"profile_img_url,omitempty"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
	DeletedAt            *time.Time         `json:"deleted_at,omitempty"`
}

// The `UserInput` input type is used to create and update users.
type UserInput struct {
	FullName             string             `json:"full_name"`
	Email                string             `json:"email"`
	Password             string             `json:"password"`
	Gender               Gender             `json:"gender"`
	IdentificationType   IdentificationType `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	CountryCode          string             `json:"country_code"`
	PhoneNumber          string             `json:"phone_number"`
	Ocupation            string             `json:"ocupation"`
	Weight               *float64           `json:"weight,omitempty"`
	Height               *float64           `json:"height,omitempty"`
	Birthday             *time.Time         `json:"birthday,omitempty"`
	ProfileImg           *graphql.Upload    `json:"profile_img,omitempty"`
}

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFamele Gender = "Famele"
	GenderNone   Gender = "None"
)

var AllGender = []Gender{
	GenderMale,
	GenderFamele,
	GenderNone,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFamele, GenderNone:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IdentificationType string

const (
	IdentificationTypeV IdentificationType = "V"
	IdentificationTypeE IdentificationType = "E"
	IdentificationTypeP IdentificationType = "P"
	IdentificationTypeJ IdentificationType = "J"
)

var AllIdentificationType = []IdentificationType{
	IdentificationTypeV,
	IdentificationTypeE,
	IdentificationTypeP,
	IdentificationTypeJ,
}

func (e IdentificationType) IsValid() bool {
	switch e {
	case IdentificationTypeV, IdentificationTypeE, IdentificationTypeP, IdentificationTypeJ:
		return true
	}
	return false
}

func (e IdentificationType) String() string {
	return string(e)
}

func (e *IdentificationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IdentificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IdentificationType", str)
	}
	return nil
}

func (e IdentificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleOwner    Role = "Owner"
	RolePayer    Role = "Payer"
	RoleOperator Role = "Operator"
	RoleDriver   Role = "Driver"
)

var AllRole = []Role{
	RoleOwner,
	RolePayer,
	RoleOperator,
	RoleDriver,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleOwner, RolePayer, RoleOperator, RoleDriver:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}