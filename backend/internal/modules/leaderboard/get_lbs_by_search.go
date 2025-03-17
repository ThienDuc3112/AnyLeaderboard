package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/utils"
	"context"
)

type SearchParam struct {
	Term         string
	UserId       int32
	PageSize     int32
	SearchCursor float32
}

type SearchParamReturn struct {
	GetLBsReturn
	Rank []float32
}

func (s LeaderboardService) Search(ctx context.Context, param SearchParam) (SearchParamReturn, error) {
	if param.UserId != 0 {
		rows, err := s.repo.SearchFavoriteLeaderboards(ctx, database.SearchFavoriteLeaderboardsParams{
			UserID:     param.UserId,
			Limit:      param.PageSize,
			Language:   utils.DetectLanguage(param.Term),
			Query:      param.Term,
			RankCursor: param.SearchCursor,
		})
		if err != nil {
			return SearchParamReturn{}, err
		}

		res := make([]models.Leaderboard, len(rows))
		counts := make([]int, len(rows))
		ranks := make([]float32, len(rows))

		for i, row := range rows {
			ranks[i] = row.Rank
			counts[i] = int(row.EntriesCount)
			res[i] = models.Leaderboard{
				ID:                   int(row.ID),
				Name:                 row.Name,
				Description:          row.Description,
				CoverImageUrl:        row.CoverImageUrl.String,
				CreatedAt:            row.CreatedAt.Time,
				Creator:              int(row.Creator),
				UpdatedAt:            row.UpdatedAt.Time,
				AllowAnonymous:       row.AllowAnonymous,
				RequiredVerification: row.RequireVerification,
				UniqueSubmission:     row.UniqueSubmission,
			}
		}

		return SearchParamReturn{
			GetLBsReturn: GetLBsReturn{
				Leaderboards: res,
				EntryCounts:  counts,
			},
			Rank: ranks,
		}, nil

	} else {
		rows, err := s.repo.SearchLeaderboards(ctx, database.SearchLeaderboardsParams{
			Limit:      param.PageSize,
			Language:   utils.DetectLanguage(param.Term),
			Query:      param.Term,
			RankCursor: param.SearchCursor,
		})
		if err != nil {
			return SearchParamReturn{}, err
		}

		res := make([]models.Leaderboard, len(rows))
		counts := make([]int, len(rows))
		ranks := make([]float32, len(rows))

		for i, row := range rows {
			ranks[i] = row.Rank
			counts[i] = int(row.EntriesCount)
			res[i] = models.Leaderboard{
				ID:                   int(row.ID),
				Name:                 row.Name,
				Description:          row.Description,
				CoverImageUrl:        row.CoverImageUrl.String,
				CreatedAt:            row.CreatedAt.Time,
				Creator:              int(row.Creator),
				UpdatedAt:            row.UpdatedAt.Time,
				AllowAnonymous:       row.AllowAnonymous,
				RequiredVerification: row.RequireVerification,
				UniqueSubmission:     row.UniqueSubmission,
			}
		}

		return SearchParamReturn{
			GetLBsReturn: GetLBsReturn{
				Leaderboards: res,
				EntryCounts:  counts,
			},
			Rank: ranks,
		}, nil
	}
}
