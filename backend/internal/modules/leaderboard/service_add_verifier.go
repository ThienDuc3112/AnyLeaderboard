package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s leaderboardService) addVerifier(ctx context.Context, param addVerifierParam) error {
	user, err := s.repo.GetUserByUsername(ctx, param.username)
	if err == pgx.ErrNoRows {
		return ErrNoUser
	} else if err != nil {
		return err
	}

	err = s.repo.AddVerifier(ctx, database.AddVerifierParams{LeaderboardID: param.lid, UserID: user.ID})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Printf("%+v\n", *pgErr)
			if pgErr.Code == "23505" {
				return ErrAlreadyVerifier
			}
		} else {
			return err
		}
	}
	return nil
}
