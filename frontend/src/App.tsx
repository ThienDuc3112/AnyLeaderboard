import { BrowserRouter, Route, Routes } from "react-router";
import LandingPage from "@/pages/landingPage";
import SignInPage from "@/pages/login";
import Layout from "@/Layout";
import SignupPage from "@/pages/signup";
import LeaderboardViewerPage from "@/pages/leaderboard/[lid]";
import BrowseLeaderboardPage from "@/pages/leaderboard";
import NewLeaderboardPage from "./pages/leaderboard/new";
import NewEntryPage from "./pages/leaderboard/[lid]/entry/new";

function App() {
  return (
    <BrowserRouter>
      <div className="flex flex-col h-screen w-screen overflow-auto p-0 m-0">
        <Routes>
          {/** Page with navbar */}
          <Route path="/" element={<Layout />}>
            <Route index element={<LandingPage />} />

            {/** Auth routes */}
            <Route path="/signin" element={<SignInPage />} />
            <Route path="/signup" element={<SignupPage />} />

            {/** Profile page */}
            <Route path="/profile/me" />
            <Route path="/profile/:id" />

            {/** Leaderboard routes */}
            <Route path="/leaderboard" element={<BrowseLeaderboardPage />} />
            <Route
              path="/leaderboard/:lid"
              element={<LeaderboardViewerPage />}
            />
            <Route path="/leaderboard/new" element={<NewLeaderboardPage />} />
            <Route path="/leaderboard/:lid/update" />

            {/** Entry routes */}
            <Route path="/leaderboard/:id/entry/:eid" />
            <Route
              path="/leaderboard/:id/entry/new"
              element={<NewEntryPage />}
            />
          </Route>

          {/** Page without navbar */}
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
