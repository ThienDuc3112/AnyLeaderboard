package leaderboard

import (
	"anylbapi/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type GetEntryParam struct {
	Eid int32
	Lid int32
}

type GetEntryReturn struct {
	Entry       models.Entry
	Leaderboard models.LeaderboardFull
}

func (s LeaderboardService) GetEntry(ctx context.Context, param GetEntryParam) (*GetEntryReturn, error) {
	// TODO: TEST THE HECK OUT OF CUSTOM GET ENTRIES MY GOD IT SO BUGGY
	lb, err := s.GetLeaderboard(ctx, param.Lid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNoLeaderboard
	}
	if err != nil {
		return nil, err
	}

	entry, err := s.repo.GetLeaderboardEntryById(ctx, param.Eid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNoEntry
	}
	if err != nil {
		return nil, err
	}
	if entry.LeaderboardID != int32(lb.ID) {
		return nil, ErrNoEntry
	}

	e := models.Entry{
		Id:         int(entry.ID),
		CreatedAt:  entry.CreatedAt.Time,
		UpdatedAt:  entry.UpdatedAt.Time,
		Fields:     entry.CustomFields,
		Verified:   entry.Verified,
		Username:   entry.Username,
		VerifiedAt: nil,
		VerifiedBy: "",
	}
	if entry.VerifiedAt.Valid {
		e.VerifiedAt = &entry.VerifiedAt.Time
	}
	if entry.VerifiedBy.Valid {
		user, err := s.repo.GetUserByID(ctx, entry.VerifiedBy.Int32)
		if err != nil {
			return nil, err
		}
		e.VerifiedBy = user.Username
	}

	return &GetEntryReturn{
		Entry:       e,
		Leaderboard: lb,
	}, nil
}
