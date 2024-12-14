import React from "react";
import { Trophy } from "lucide-react";
import { Link } from "react-router";
import { useAtom } from "jotai";
import { userAtom } from "@/context/user";

interface NavbarOption {
  to: string;
  text: string;
}
const SignInOptions: NavbarOption[] = [
  {
    text: "Browse leaderboards",
    to: "/leaderboard",
  },
  {
    text: "Create new leaderboard",
    to: "/leaderboard/new",
  },
  {
    text: "Profile",
    to: "/profile/me",
  },
];

const SignOutOptions: NavbarOption[] = [
  {
    text: "Browse leaderboards",
    to: "/leaderboard",
  },
  {
    text: "Sign in",
    to: "/signin",
  },
  {
    text: "Sign up",
    to: "/signup",
  },
];

const NavBar: React.FC = () => {
  const [user] = useAtom(userAtom);
  return (
    <nav className="bg-white shadow-sm">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div className="flex items-center space-x-2">
          <Trophy className="h-8 w-8 text-indigo-600" />
          <Link to={"/"} className="text-xl font-bold text-gray-800">
            LeaderBoard Maker
          </Link>
        </div>
        <nav>
          <ul className="flex space-x-4">
            {(user ? SignInOptions : SignOutOptions).map((option, i) => {
              return (
                <li key={i}>
                  <Link
                    to={option.to}
                    className="text-gray-600 hover:text-indigo-600"
                  >
                    {option.text}
                  </Link>
                </li>
              );
            })}
          </ul>
        </nav>
      </div>
    </nav>
  );
};

export default NavBar;
