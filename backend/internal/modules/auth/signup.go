package auth

import (
	"anylbapi/internal/database"
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s AuthService) Signup(context context.Context, param SignUpParam) error {
	// Clean data
	param.Username = strings.ToLower(param.Username)
	param.Email = strings.ToLower(param.Email)

	// Check duplicate Username
	_, err := s.repo.GetUserByUsername(context, param.Username)
	if err == nil {
		return ErrUsernameTaken
	}
	if err != pgx.ErrNoRows {
		return err
	}

	// Check duplicate Email
	_, err = s.repo.GetUserByEmail(context, param.Email)
	if err == nil {
		return ErrEmailUsed
	}
	if err != pgx.ErrNoRows {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(context, database.CreateUserParams{
		Username:    param.Username,
		DisplayName: param.DisplayName,
		Email:       param.Email,
		Password:    string(hashedPassword),
	})

}
