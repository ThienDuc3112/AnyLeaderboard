package auth

import (
	"anylbapi/internal/database"
	"context"
	"database/sql"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (s authService) signup(context context.Context, body signUpParam) error {
	// Clean data
	body.Username = strings.ToLower(body.Username)
	body.Email = strings.ToLower(body.Email)

	// Check duplicate Username
	_, err := s.repo.GetUserByUsername(context, body.Username)
	if err == nil {
		return errUsernameTaken
	}
	if err != sql.ErrNoRows {
		return err
	}

	// Check duplicate Email
	_, err = s.repo.GetUserByEmail(context, body.Email)
	if err == nil {
		return errEmailUsed
	}
	if err != sql.ErrNoRows {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.repo.CreateUser(context, database.CreateUserParams{
		Username:    body.Username,
		DisplayName: body.DisplayName,
		Email:       body.Email,
		Password:    string(hashedPassword),
	})
	if err != nil {
		return err
	}

	return nil
}
