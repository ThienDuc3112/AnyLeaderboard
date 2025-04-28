import { Entry, LeaderboardFull } from "@/types/leaderboard";
import { api } from "@/utils/api";
import { useQuery } from "@tanstack/react-query";

export const useEntry = (lid: string | undefined, eid: string | undefined) => {
  return useQuery<{ leaderboard: LeaderboardFull; entry: Entry }>({
    queryKey: ["leaderboard", lid, "entry", eid],
    queryFn: async () =>
      (await api.get(`/leaderboards/${lid}/entries/${eid}`)).data,
    enabled: !!lid && !!eid,
  });
};
