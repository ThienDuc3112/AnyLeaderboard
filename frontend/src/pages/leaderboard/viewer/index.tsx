import React from "react";
import { MockLeaderboardFull } from "@/mocks/leaderboardFull";
import { useParams } from "react-router";
import LeaderboardHeader from "./LeaderboardHeader";
import LeaderboardContent from "./LeaderboardContent";

const LeaderboardViewer: React.FC = () => {
  const { lid } = useParams();
  console.log(lid);
  const data = MockLeaderboardFull;
  return (
    <div className="w-full mt-12">
      <div className="max-w-5xl mx-auto bg-indigo-50 rounded-lg shadow-md overflow-hidden">
        <LeaderboardHeader data={data} />
        <LeaderboardContent data={data} />
      </div>
    </div>
  );
};

export default LeaderboardViewer;
