import React from "react";

interface PropType extends React.SelectHTMLAttributes<HTMLSelectElement> {
  options: { value: string; text: string }[];
  isError?: boolean;
  error?: string;
}

const Dropdown: React.FC<PropType> = ({ options, isError, ...props }) => {
  return (
    <select
      {...props}
      className={`block ${props.disabled ? "bg-gray-200" : "bg-white"} w-full px-3 rounded-full ${isError ? "border-red-500" : "border-gray-300"} shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 ${props.className}`}
    >
      {options.map((option, i) => (
        <option key={i} value={option.value}>
          {option.text}
        </option>
      ))}
    </select>
  );
};

export default Dropdown;
