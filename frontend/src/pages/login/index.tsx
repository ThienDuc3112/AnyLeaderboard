import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import { userAtom } from "@/context/user";
import { useAtom } from "jotai";
import React from "react";
import { useNavigate } from "react-router";

const SignInPage: React.FC = () => {
  const navigate = useNavigate();
  const [, setUser] = useAtom(userAtom);
  return (
    <div className="w-full mt-32 flex justify-center items-center">
      <div className="border border-indigo-200 rounded-xl bg-indigo-50">
        <h1 className="font-bold text-2xl text-center my-5">Sign in</h1>
        <form
          className="flex flex-col py-3 px-5 my-1"
          onSubmit={(e) => {
            e.preventDefault();
            setUser({
              displayName: "huyen",
              email: "test@mail.com",
              id: "adlsjkf",
              username: "huyen",
            });
            navigate("/leaderboard");
          }}
        >
          <label htmlFor="username">Username or email address</label>
          <Input className="w-72" id="username" />

          <label htmlFor="password" className="mt-6">
            Password
          </label>
          <Input className="w-72" type="password" id="password" />

          <Button type="submit" className="mt-6">
            Sign in
          </Button>
        </form>
      </div>
    </div>
  );
};

export default SignInPage;
