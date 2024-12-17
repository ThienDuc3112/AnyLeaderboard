import React from "react";

type PropType = React.InputHTMLAttributes<HTMLInputElement> & {
  icon?: React.ReactNode;
  variant?: "default" | "outline" | "filled";
};

const getVariantClasses = (variant: PropType["variant"] = "default") => {
  switch (variant) {
    case "outline":
      return "border-2 border-indigo-600";
    default:
      return "border border-indigo-400";
  }
};

const Input: React.FC<PropType> = (props) => {
  return (
    <div
      className={`px-3 bg-white rounded-full relative flex focus-within:ring-indigo-600 focus-within:outline-none focus-within:ring-1 items-center h-10 ${getVariantClasses(
        props.variant
      )} ${props.className}`}
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
