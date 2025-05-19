import { LeaderboardPreview } from "@/types/leaderboard";
import { useInfiniteQuery } from "@tanstack/react-query";
import axios from "axios";
import { atom, useAtom } from "jotai";
import { useMemo } from "react";

const filteredAtom = atom<"recent" | "byUsername" | "favorite">("recent");
const searchAtom = atom<string>("");

export const useLeaderboards = () => {
  const [filter, setFilter] = useAtom(filteredAtom);
  const [search, setSearch] = useAtom(searchAtom);
  const baseUrl = import.meta.env.VITE_API_URL;

  const getInitialApiUrl = () => {
    const base =
      filter == "byUsername"
        ? `${baseUrl}/leaderboards/by-username`
        : filter == "favorite"
          ? `${baseUrl}/leaderboards/favorites`
          : search.trim()
            ? `${baseUrl}/leaderboards/search`
            : `${baseUrl}/leaderboards`;

    const url = new URL(base);
    if (search.trim()) url.searchParams.set("query", search.trim());
    return url.toString();
  };

  const { data, isLoading, error, fetchNextPage, hasNextPage } =
    useInfiniteQuery({
      queryKey: ["leaderboards", filter, search],
      queryFn: async ({ pageParam }) => {
        const url = pageParam || getInitialApiUrl();
        const response = await axios.get(url, { withCredentials: true });
        return response.data;
      },
      initialPageParam: "",
      getNextPageParam: (lastPage) => lastPage.next ?? null,
      maxPages: 20,
      staleTime: 30 * 1000,
    });

  const lbs = useMemo<LeaderboardPreview[]>(
    () => data?.pages.flatMap((page) => page.data) ?? [],
    [data],
  );

  const toggleFilter = (newFilter: "recent" | "byUsername" | "favorite") => {
    setFilter(newFilter);
  };

  const searchLeaderboards = (q: string) => {
    setSearch(q);
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
    search,
    searchLeaderboards,
  };
};
