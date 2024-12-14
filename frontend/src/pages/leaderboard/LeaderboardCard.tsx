import { LeaderboardPreview } from "@/types/leaderboard";
import { Users } from "lucide-react";
import React from "react";
import { Link } from "react-router";

interface PropType {
  board: LeaderboardPreview;
}

const LeaderboardCard: React.FC<PropType> = ({ board }) => {
  return (
    <Link
      key={board.id}
      to={`/leaderboard/${board.id}`}
      className="group relative overflow-hidden rounded-lg border border-indigo-200 bg-card transition-colors hover:border-primary"
    >
      <div className="aspect-[2/1] overflow-hidden bg-muted">
        <img
          src={board.coverImageUrl}
          alt={"Image not availabe"}
          width={600}
          height={300}
          className="object-cover transition-transform group-hover:scale-105"
        />
      </div>
      <div className="p-4">
        <h3 className="font-semibold">{board.name}</h3>
        <p className="mt-1 text-sm text-muted-foreground">
          {board.description}
        </p>
        <div className="mt-4 flex items-center gap-4">
          <div className="flex items-center gap-1 text-sm text-muted-foreground">
            <Users className="h-4 w-4" />
            {board.entryCount} participants
          </div>
        </div>
      </div>
    </Link>
  );
};

export default LeaderboardCard;
