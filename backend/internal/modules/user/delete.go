package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrIncorrectPassword = fmt.Errorf("incorrect password")

type DeleteParam struct {
	Password string
	UserID   int
}

func (s UserService) Delete(ctx context.Context, param DeleteParam) error {
	user, err := s.repo.GetUserByID(ctx, int32(param.UserID))
	if err == pgx.ErrNoRows {
		return ErrNoUsers
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return ErrIncorrectPassword
	}

	return s.repo.DeleteUser(ctx, int32(param.UserID))
}
