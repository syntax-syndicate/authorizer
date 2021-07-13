// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BasicAuthLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BasicAuthLoginResponse struct {
	Success      bool     `json:"success"`
	Message      string   `json:"message"`
	Errors       []*Error `json:"errors"`
	StatusCode   int      `json:"statusCode"`
	RefreshToken *string  `json:"refreshToken"`
	User         *User    `json:"user"`
}

type BasicAuthSignupInput struct {
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	CofirmPassword string  `json:"cofirmPassword"`
	Image          *string `json:"image"`
}

type BasicAuthSignupResponse struct {
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Errors     []*Error `json:"errors"`
	StatusCode int      `json:"statusCode"`
	User       *User    `json:"user"`
}

type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type Response struct {
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Errors     []*Error `json:"errors"`
	StatusCode int      `json:"statusCode"`
}

type User struct {
	ID              string  `json:"id"`
	Email           string  `json:"email"`
	SignUpMethod    string  `json:"SignUpMethod"`
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	EmailVerifiedAt *int64  `json:"emailVerifiedAt"`
	Password        *string `json:"password"`
	Image           *string `json:"image"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
}

type VerificationRequest struct {
	ID         string  `json:"id"`
	Identifier *string `json:"identifier"`
	Token      *string `json:"token"`
	Email      *string `json:"email"`
	Expires    *int64  `json:"expires"`
	CreatedAt  *int64  `json:"createdAt"`
	UpdatedAt  *int64  `json:"updatedAt"`
}

type VerifySignupTokenInput struct {
	Token string `json:"token"`
}
