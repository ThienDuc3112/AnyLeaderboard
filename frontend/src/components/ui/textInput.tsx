import React from "react";

type PropType = React.InputHTMLAttributes<HTMLInputElement> & {
  icon?: React.ReactNode;
};

const Input: React.FC<PropType> = (props) => {
  return (
    <div
      className={`min-w-72 px-3 bg-white rounded-full relative flex focus-within:ring-indigo-600 focus-within:outline-none focus-within:ring-1 items-center border border-indigo-400 h-10 ${props.className}`}
    >
      {props.icon && <span className="mr-3">{props.icon}</span>}
      <input
        {...props}
        className={`h-full w-full bg-inherit focus:outline-none rounded-lg`}
      />
    </div>
  );
};

export default Input;
