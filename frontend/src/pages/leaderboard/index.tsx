import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import { Clock, Search, TrendingUp, User } from "lucide-react";
import React, {
  FormEventHandler,
  useCallback,
  useEffect,
  useMemo,
  useState,
} from "react";
import LeaderboardGrid from "./LeaderboardGrid";
import LoadMore from "./LoadMore";
import { useLeaderboards } from "./state";
import { useSearchParams } from "react-router";

interface FilterOption {
  icon: React.FC<{ className?: string }>;
  text: string;
  onClick: () => void;
  disabled: boolean;
  active: boolean;
}

const BrowseLeaderboardPage: React.FC = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [searchTerm, setSearchTerm] = useState("");
  const {
    lbs,
    isLoading,
    error,
    fetchNextPage,
    hasNextPage,
    toggleFilter,
    filter,
    searchLeaderboards,
  } = useLeaderboards();

  useEffect(() => {
    const q = searchParams.get("q");
    if (q !== null) {
      searchLeaderboards(q);
      setSearchTerm(q);
    }

    return () => {
      searchLeaderboards("");
    };
  }, [searchParams]);

  const filterOptions = useMemo<FilterOption[]>(
    () => [
      {
        icon: Clock,
        text: "Recent",
        onClick: () => {
          toggleFilter("recent");
        },
        disabled: false,
        active: filter == "recent",
      },
      {
        icon: TrendingUp,
        text: "Favorite",
        onClick: () => {},
        disabled: true,
        active: filter == "favorite",
      },
      {
        icon: User,
        text: "Made by you",
        onClick: () => {},
        disabled: true,
        active: filter == "byUsername",
      },
    ],
    [],
  );

  const search = useCallback<FormEventHandler<HTMLFormElement>>(
    (e) => {
      e.preventDefault();
      setSearchParams({ q: searchTerm });
    },
    [searchTerm],
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
          <form
            onSubmit={search}
            className="mb-4 flex flex-wrap items-center gap-4 align-middle flex-row"
          >
            <Input
              icon={<Search className="h-4 w-4" />}
              placeholder="Search leaderboards..."
              className="flex-1 my-0"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
            <Button type="submit" className="w-fit">
              Search
            </Button>
          </form>
          <div className="flex flex-wrap items-center gap-4">
            {filterOptions.map((option, i) => (
              <Button
                variant={option.active ? "filled" : "outline"}
                key={i}
                disabled={option.disabled}
              >
                <span className="flex flex-row align-middle items-center gap-2">
                  <option.icon className="h-4 w-4" />
                  <span>{option.text}</span>
                </span>
              </Button>
            ))}
          </div>
        </div>

        {/* Leaderboard Grid */}
        {isLoading ? (
          <p>Loading...</p>
        ) : error ? (
          <p>An error occured</p>
        ) : (
          <LeaderboardGrid lbs={lbs} />
        )}

        {!isLoading && <LoadMore hasMore={hasNextPage} fn={fetchNextPage} />}
      </main>
    </div>
  );
};

export default BrowseLeaderboardPage;
