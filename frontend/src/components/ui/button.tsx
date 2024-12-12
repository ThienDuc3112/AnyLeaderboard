import React from "react";

const Button: React.FC<React.ButtonHTMLAttributes<HTMLButtonElement>> = (
  props
) => {
  return (
    <button
      {...props}
      className="text-white min-h-10 min-w-72 bg-indigo-600 font-semibold border-none rounded-full w-full py-1 flex flex-col align-middle justify-center"
    >
      {props.children}
    </button>
  );
};

export default Button;
