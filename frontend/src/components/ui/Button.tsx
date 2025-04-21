import React from "react";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "filled" | "outline" | "ghost" | "inverted";
  size?: "small" | "medium" | "large";
}

const getVariantClasses = (variant: string) => {
  switch (variant) {
    case "outline":
      return "border-2 border-indigo-600 text-indigo-600 hover:bg-indigo-100";
    case "ghost":
      return "text-indigo-600 hover:bg-indigo-100";
    case "inverted":
      return "bg-white hover:bg-indigo-50 text-indigo-700";
    case "filled":
    default:
      return "bg-indigo-600 hover:bg-indigo-400 text-white";
  }
};

const getSizeClasses = (size: string) => {
  switch (size) {
    case "small":
      return "px-2 py-1 text-sm";
    case "large":
      return "px-6 py-3 text-lg";
    case "medium":
    default:
      return "px-4 py-2 text-base";
  }
};

const Button: React.FC<ButtonProps> = ({
  variant = "filled",
  size = "medium",
  ...props
}) => {
  return (
    <button
      {...props}
      className={`transition font-semibold rounded-full flex flex-col items-center justify-center ${getVariantClasses(
        variant,
      )} ${getSizeClasses(size)} ${props.className} disabled:text-gray-300 disabled:bg-white disabled:border-gray-300
      `}
    >
      {props.children}
    </button>
  );
};

export default Button;
