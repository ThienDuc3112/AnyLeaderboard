import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import { MockLeaderboardPreview } from "@/mocks/leaderboardPreviews";
import { Clock, Search, TrendingUp, User } from "lucide-react";
import React, { useMemo } from "react";
import LeaderboardCard from "./LeaderboardCard";

interface FilterOption {
  icon: React.FC<{ className?: string }>;
  text: string;
  onClick: () => void;
}

const BrowseLeaderboardPage: React.FC = () => {
  const filterOptions = useMemo<FilterOption[]>(
    () => [
      {
        icon: Clock,
        text: "Recent",
        onClick: () => {},
      },
      {
        icon: TrendingUp,
        text: "Trending",
        onClick: () => {},
      },
      {
        icon: User,
        text: "Made by you",
        onClick: () => {},
      },
    ],
    []
  );
  return (
    <div className="w-full">
      <main className="container mx-auto px-4 py-8">
        <div className="mb-4 space-y-4">
          <h1 className="text-3xl font-bold">Browse Leaderboards</h1>
          <p className="text-muted-foreground">
            Discover and join competitive leaderboards from various games and
            communities
          </p>
        </div>

        {/* Search and Filters */}
        <div className="mb-8 bg-indigo-50 p-4 rounded-xl">
          <div className="mb-4 flex flex-wrap items-center gap-4 align-middle flex-row">
            <Input
              icon={<Search className="h-4 w-4" />}
              placeholder="Search leaderboards..."
              className="flex-1 my-0"
            />
            <Button className="w-fit">Search</Button>
          </div>
          <div className="flex flex-wrap items-center gap-4">
            {filterOptions.map((option, i) => (
              <Button key={i}>
                <span className="flex flex-row align-middle items-center gap-2">
                  <option.icon className="h-4 w-4" />
                  <span>{option.text}</span>
                </span>
              </Button>
            ))}
          </div>
        </div>

        {/* Leaderboard Grid */}
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {MockLeaderboardPreview.map((board) => (
            <LeaderboardCard key={board.id} board={board} />
          ))}
        </div>
      </main>
    </div>
  );
};

export default BrowseLeaderboardPage;
