import { LeaderboardPreview } from "@/types/leaderboard";
import { useInfiniteQuery } from "@tanstack/react-query";
import axios from "axios";
import { atom, useAtom } from "jotai";
import { useMemo } from "react";

const filteredAtom = atom<"recent" | "byUsername" | "favorite">("recent");

export const useLeaderboards = () => {
  const [filter, setFilter] = useAtom(filteredAtom);
  const baseUrl = import.meta.env.VITE_API_URL;

  const getInitialApiUrl = () => {
    switch (filter) {
      case "byUsername":
        return `${baseUrl}/leaderboards/by-username`;
      case "favorite":
        return `${baseUrl}/leaderboards/favorites`;
      default:
        return `${baseUrl}/leaderboards`;
    }
  };

  const { data, isLoading, error, fetchNextPage, hasNextPage } =
    useInfiniteQuery({
      queryKey: ["leaderboards", filter],
      queryFn: async ({ pageParam }) => {
        const url = pageParam || getInitialApiUrl();
        const response = await axios.get(url, { withCredentials: true });
        return response.data;
      },
      initialPageParam: "",
      getNextPageParam: (lastPage) => lastPage.next ?? null,
      maxPages: 100,
    });

  const lbs = useMemo<LeaderboardPreview[]>(
    () => data?.pages.flatMap((page) => page.data) ?? [],
    [data],
  );

  const toggleFilter = (newFilter: "recent" | "byUsername" | "favorite") => {
    setFilter(newFilter);
  };

  return {
    lbs,
    data,
    isLoading,
    error,
    fetchNextPage,
    hasNextPage,
    toggleFilter,
    filter,
  };
};
