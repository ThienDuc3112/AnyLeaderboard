import React from "react";
import { useParams } from "react-router";
import LeaderboardHeader from "./LeaderboardHeader";
import LeaderboardContent from "./LeaderboardContent";
import { useQuery } from "@tanstack/react-query";
import { api } from "@/utils/api";
import { LeaderboardFull } from "@/types/leaderboard";

const LeaderboardViewPage: React.FC = () => {
  const { lid } = useParams();
  const { data, isLoading, error } = useQuery<LeaderboardFull>({
    queryKey: ["leaderboard", lid],
    queryFn: async () => (await api.get(`/leaderboards/${lid}`)).data
  })

  if (isLoading) return <p>Loading...</p>;
  if (error || !data) return <p>An error occured</p>;
  return (
    <div className="w-full mt-12">
      <div className="max-w-5xl mx-auto bg-indigo-50 rounded-lg shadow-md overflow-hidden">
        <LeaderboardHeader data={data} />
        {
          <LeaderboardContent data={data} />
        }
      </div>
    </div>
  );
};

export default LeaderboardViewPage;
