import Button from "@/components/ui/button";
import TextInput from "@/components/ui/textInput";
import React from "react";

const Signup: React.FC = () => {
  return (
    <div className="w-full mt-32 flex justify-center items-center">
      <div className="border border-indigo-200 rounded-xl bg-indigo-50">
        <h1 className="font-bold text-2xl text-center my-5">Sign up</h1>
        <form className="flex flex-col py-3 px-5 my-1">
          <label htmlFor="username">Username</label>
          <TextInput id="username" />

          <label htmlFor="username">Email</label>
          <TextInput id="email" />

          <label htmlFor="password">Password</label>
          <TextInput type="password" id="password" />

          <label htmlFor="passwordVerify">Verify password</label>
          <TextInput type="password" id="passwordVerify" />

          <Button type="submit">Sign up</Button>
        </form>
      </div>
    </div>
  );
};

export default Signup;
