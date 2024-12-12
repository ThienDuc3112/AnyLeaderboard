import React from "react";

const Login: React.FC = () => {
  return (
    <div className="w-full h-full flex justify-center items-center ">
      <div className="border border-indigo-200 rounded-xl bg-indigo-50">
        <form className="w-80 py-3 px-5 my-1">
          <label htmlFor="username">Username or email address</label>
          <input
            className="mt-1 h-10 w-72 px-2 rounded-full border-indigo-400 mb-6 focus:outline-none focus:ring-2 focus:ring-indigo-600"
            type="text"
            id="username"
          />

          <label htmlFor="password">Password</label>
          <input
            className="mt-1 h-10 w-72 px-2 rounded-full border-indigo-400 mb-6 focus:outline-none focus:ring-2 focus:ring-indigo-600"
            type="password"
            id="password"
          />

          <button
            className="text-white bg-indigo-600 font-semibold border-none rounded-lg w-full py-1 flex flex-col align-middle justify-center"
            type="submit"
          >
            Log in
          </button>
        </form>
      </div>
    </div>
  );
};

export default Login;
