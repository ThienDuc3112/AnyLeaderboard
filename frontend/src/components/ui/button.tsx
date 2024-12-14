import React from "react";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "filled" | "outline" | "ghost" | "inverted";
}

const Button: React.FC<ButtonProps> = ({ variant = "filled", ...props }) => {
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

  return (
    <button
      {...props}
      className={`min-h-10 max-w-72 transition font-semibold rounded-full flex flex-col items-center justify-center ${getVariantClasses(
        variant
      )} ${props.className}`}
    >
      {props.children}
    </button>
  );
};

export default Button;
