import React from "react";

const Input: React.FC<React.InputHTMLAttributes<HTMLInputElement>> = (
  props
) => {
  return (
    <input
      {...props}
      className="mt-1 h-10 w-72 px-2 rounded-full border-indigo-400 mb-6 focus:outline-none focus:ring-2 focus:ring-indigo-600"
    />
  );
};

export default Input;
