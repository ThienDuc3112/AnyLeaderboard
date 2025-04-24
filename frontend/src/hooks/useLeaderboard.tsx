import { LeaderboardFull } from "@/types/leaderboard";
import { api } from "@/utils/api";
import { useQuery } from "@tanstack/react-query";

export const useLeaderboard = (lid: string | undefined) => {
  return useQuery<LeaderboardFull>({
    queryKey: ["leaderboard", lid],
    queryFn: async () => (await api.get(`/leaderboards/${lid}`)).data,
  });
};
