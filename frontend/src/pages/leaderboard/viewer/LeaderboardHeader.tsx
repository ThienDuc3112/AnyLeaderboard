import React from "react";
import ExternalLink from "@/pages/leaderboard/viewer/ExternalLink";
import { LeaderboardFull } from "@/types/leaderboard";

interface PropType {
  data: LeaderboardFull;
}

const LeaderboardHeader: React.FC<PropType> = ({ data }) => {
  return (
    <div className="p-6 space-y-4">
      <div className="flex gap-4">
        <div className="relative w-32 h-32 flex-shrink-0">
          <img
            src={data.coverImageUrl || "/placeholder.svg?height=128&width=128"}
            alt={`${data.name} cover`}
            className="rounded-lg object-cover w-full h-full"
          />
        </div>

        <div className="flex-grow">
          <h2 className="text-2xl font-bold">{data.name}</h2>
          {data.description && (
            <p className="text-sm text-gray-600 mt-1">{data.description}</p>
          )}
          <div className="mt-2 text-xs text-gray-600">
            <span className="font-mono">{data.entryCount}</span> entries
          </div>
          <div className="flex flex-wrap gap-2 mt-2">
            {data.externalLinks?.map((link, i) => (
              <ExternalLink key={i} link={link} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default LeaderboardHeader;