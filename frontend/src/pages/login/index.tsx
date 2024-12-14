import Button from "@/components/ui/button";
import Input from "@/components/ui/textInput";
import React from "react";

const SignIn: React.FC = () => {
  return (
    <div className="w-full mt-32 flex justify-center items-center">
      <div className="border border-indigo-200 rounded-xl bg-indigo-50">
        <h1 className="font-bold text-2xl text-center my-5">Sign in</h1>
        <form className="flex flex-col py-3 px-5 my-1">
          <label htmlFor="username">Username or email address</label>
          <Input id="username" />

          <label htmlFor="password" className="mt-6">
            Password
          </label>
          <Input type="password" id="password" />

          <Button type="submit" className="mt-6">
            Sign in
          </Button>
        </form>
      </div>
    </div>
  );
};

export default SignIn;
