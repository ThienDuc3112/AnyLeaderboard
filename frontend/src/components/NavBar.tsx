import React from "react";
import { Trophy } from "lucide-react";
import { Link } from "react-router";

const NavBar: React.FC = () => {
  return (
    <header className="bg-white shadow-sm">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div className="flex items-center space-x-2">
          <Trophy className="h-8 w-8 text-indigo-600" />
          <Link to={"/"} className="text-xl font-bold text-gray-800">
            LeaderBoard Maker
          </Link>
        </div>
        <nav>
          <ul className="flex space-x-4">
            <li>
              <Link
                to={"/browse"}
                className="text-gray-600 hover:text-indigo-600"
              >
                Browse leaderboard
              </Link>
            </li>
            <li>
              <Link
                to={"/login"}
                className="text-gray-600 hover:text-indigo-600"
              >
                Log in
              </Link>
            </li>
            <li>
              <Link
                to={"/signup"}
                className="text-gray-600 hover:text-indigo-600"
              >
                Sign up
              </Link>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  );
};

export default NavBar;
