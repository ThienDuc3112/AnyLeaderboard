package leaderboard

import "context"

type SearchParam struct {
	Term   string
	UserId int32
}

func (s LeaderboardService) Search(ctx context.Context, param SearchParam) {

}
