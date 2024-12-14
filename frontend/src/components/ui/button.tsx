import React from "react";

const Button: React.FC<React.ButtonHTMLAttributes<HTMLButtonElement>> = (
  props
) => {
  return (
    <button
      {...props}
      className={`text-white min-h-10 max-w-72 bg-indigo-600 hover:bg-indigo-400 transition font-semibold border-none rounded-full py-1 flex flex-col align-middle justify-center ${props.className}`}
    >
      {props.children}
    </button>
  );
};

export default Button;
