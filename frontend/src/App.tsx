import { BrowserRouter, Route, Routes } from "react-router";
import LandingPage from "@/pages/landingPage";
import Login from "@/pages/login";
import Layout from "@/Layout";
import Signup from "@/pages/signup";

function App() {
  return (
    <BrowserRouter>
      <div className="flex flex-col h-screen w-screen overflow-auto p-0 m-0">
        <Routes>
          <Route path="/" element={<Layout />}>
            {/** Page with navbar */}
            <Route index element={<LandingPage />} />
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/browse" />
            <Route path="/profile/:id" />
          </Route>

          {/** Page without navbar */}
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
