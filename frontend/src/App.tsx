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
import { useEffect, useState } from "react";
import { useSetAtom } from "jotai";
import { sessionAtom } from "./contexts/user";
import { api } from "./utils/api";
import UpdateLeaderboardPage from "./pages/leaderboard/[lid]/update";

function App() {
  const setSession = useSetAtom(sessionAtom);
  const [loading, setLoading] = useState(true);
  useEffect(() => {
    api
      .post("/auth/refresh", undefined, {
        withCredentials: true,
      })
      .then((res) => {
        setSession({
          activeToken: res.data.access_token,
          user: {
            ...res.data.user,
            createdAt: new Date(res.data.user.createdAt),
          },
        });
      })
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <p>Loading...</p>;

  return (
    <BrowserRouter>
      <div className="flex flex-col h-screen w-screen overflow-auto p-0 m-0">
        <Routes>
          {/** Page with navbar */}
          <Route path="/" element={<Layout />}>
            {/** ========== Completed routes ========== */}
            <Route index element={<LandingPage />} />
            <Route path="/signin" element={<SignInPage />} />
            <Route path="/signup" element={<SignupPage />} />
            <Route path="/leaderboard/:lid" element={<LeaderboardViewPage />} />
            <Route
              path="/leaderboard/:lid/entry/:eid"
              element={<EntryViewPage />}
            />

            {/** ========== Semi-completed routes ========== */}
            {/** Haven't styled the warning when error */}
            <Route
              path="/leaderboard/:lid/entry/new"
              element={<NewEntryPage />}
            />
            <Route path="/leaderboard/new" element={<NewLeaderboardPage />} />
            {/** Search option when login */}
            <Route path="/leaderboard" element={<BrowseLeaderboardPage />} />

            {/** ========== Incomplete routes ========== */}
            <Route path="/profile/me" />
            <Route path="/profile/:id" />
            <Route
              path="/leaderboard/:lid/update"
              element={<UpdateLeaderboardPage />}
            />
          </Route>
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
