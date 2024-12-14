import { BrowserRouter, Route, Routes } from "react-router";
import LandingPage from "@/pages/landingPage";
import SignIn from "@/pages/login";
import Layout from "@/Layout";
import Signup from "@/pages/signup";
import LeaderboardViewer from "@/pages/leaderboard/viewer";
import BrowseLeaderboard from "@/pages/leaderboard";

function App() {
  return (
    <BrowserRouter>
      <div className="flex flex-col h-screen w-screen overflow-auto p-0 m-0">
        <Routes>
          {/** Page with navbar */}
          <Route path="/" element={<Layout />}>
            <Route index element={<LandingPage />} />

            {/** Auth routes */}
            <Route path="/signin" element={<SignIn />} />
            <Route path="/signup" element={<Signup />} />

            {/** Profile page */}
            <Route path="/profile/me" />
            <Route path="/profile/:id" />

            {/** Leaderboard routes */}
            <Route path="/leaderboard" element={<BrowseLeaderboard />} />
            <Route path="/leaderboard/:lid" element={<LeaderboardViewer />} />
            <Route path="/leaderboard/new" />
            <Route path="/leaderboard/:lid/update" />

            {/** Entry routes */}
            <Route path="/leaderboard/:id/entry/:eid" />
            <Route path="/leaderboard/:id/entry/new" />
          </Route>

          {/** Page without navbar */}
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
