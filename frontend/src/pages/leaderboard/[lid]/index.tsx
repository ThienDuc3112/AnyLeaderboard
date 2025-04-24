import React from "react";
import { useParams } from "react-router";
import LeaderboardHeader from "./LeaderboardHeader";
import LeaderboardContent from "./LeaderboardContent";
import { useLeaderboard } from "@/hooks/useLeaderboard";

const LeaderboardViewPage: React.FC = () => {
  const { lid } = useParams();
  const { data, isLoading, error } = useLeaderboard(lid);

  if (isLoading) return <p>Loading...</p>;
  if (error || !data) return <p>An error occured</p>;
  return (
    <div className="w-full mt-12">
      <div className="max-w-5xl mx-auto bg-indigo-50 rounded-lg shadow-md overflow-hidden">
        <LeaderboardHeader data={data} />
        {<LeaderboardContent data={data} />}
      </div>
    </div>
  );
};

export default LeaderboardViewPage;
