import { BrowserRouter, Route, Routes } from "react-router";
import LandingPage from "@/pages/landingPage";
import SignInPage from "@/pages/login";
import Layout from "@/Layout";
import SignupPage from "@/pages/signup";
import LeaderboardViewPage from "@/pages/leaderboard/[lid]";
import BrowseLeaderboardPage from "@/pages/leaderboard";
import NewLeaderboardPage from "@/pages/leaderboard/new";
import NewEntryPage from "@/pages/leaderboard/[lid]/entry/new";
import EntryViewPage from "@/pages/leaderboard/[lid]/entry/[eid]";
import { useEffect } from "react";
import { useSetAtom } from "jotai";
import { sessionAtom } from "./globalState/user";
import { api } from "./utils/api";

function App() {
  const setSession = useSetAtom(sessionAtom);
  useEffect(() => {
    api
      .post("/auth/refresh", undefined, {
        withCredentials: true,
      })
      .then((res) => {
        setSession({
          activeToken: res.data.activeToken,
          user: {
            ...res.data.user,
            createdAt: new Date(res.data.user.createdAt),
          },
        });
      })
      .catch(() => {});
  }, []);
  return (
    <BrowserRouter>
      <div className="flex flex-col h-screen w-screen overflow-auto p-0 m-0">
        <Routes>
          {/** Page with navbar */}
          <Route path="/" element={<Layout />}>
            <Route index element={<LandingPage />} /> {/* Completed */}
            {/** Auth routes */}
            <Route path="/signin" element={<SignInPage />} /> {/* Completed */}
            <Route path="/signup" element={<SignupPage />} /> {/* Completed */}
            {/** Profile page */}
            <Route path="/profile/me" />
            <Route path="/profile/:id" />
            {/** Leaderboard routes */}
            <Route path="/leaderboard" element={<BrowseLeaderboardPage />} />
            <Route
              path="/leaderboard/:lid"
              element={<LeaderboardViewPage />}
            />{" "}
            {/* Completed, surprisingly */}
            <Route path="/leaderboard/new" element={<NewLeaderboardPage />} />
            <Route path="/leaderboard/:lid/update" />
            {/** Entry routes */}
            <Route
              path="/leaderboard/:id/entry/:eid"
              element={<EntryViewPage />}
            />
            <Route
              path="/leaderboard/:id/entry/new"
              element={<NewEntryPage />}
            />
          </Route>
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
