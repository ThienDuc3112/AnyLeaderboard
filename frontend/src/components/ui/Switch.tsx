import React from "react";

interface SwitchProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  checked?: boolean;
}

const Switch: React.FC<SwitchProps> = ({
  label,
  checked = false,
  onChange,
  ...props
}) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (onChange) {
      onChange(e);
    }
  };

  return (
    <label className={`flex items-center cursor-pointer ${props.className}`}>
      <div className="relative">
        <input
          {...props}
          type="checkbox"
          checked={checked}
          onChange={handleChange}
          className="sr-only"
        />
        <div
          className={`block w-14 h-8 rounded-full transition-colors duration-300 ease-in-out ${
            checked ? "bg-indigo-600" : "bg-gray-300"
          }`}
        />
        <div
          className={`absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition-transform duration-300 ease-in-out ${
            checked ? "transform translate-x-full" : ""
          }`}
        />
      </div>
      {label && (
        <span className="ml-3 text-sm font-medium text-gray-900">{label}</span>
      )}
    </label>
  );
};

export default Switch;
