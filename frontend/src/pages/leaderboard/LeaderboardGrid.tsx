import { LeaderboardPreview } from "@/types/leaderboard";
import React from "react";
import LeaderboardCard from "./LeaderboardCard";

type PropType = {
  lbs: LeaderboardPreview[];
};

const LeaderboardGrid: React.FC<PropType> = ({ lbs }) => {
  return (
    <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      {lbs.map((board) => (
        <LeaderboardCard key={board.id} board={board} />
      ))}
    </div>
  );
};

export default LeaderboardGrid;
