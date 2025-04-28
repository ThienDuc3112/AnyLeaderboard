import React, { useState, useRef, useEffect, FC } from "react";
import { LucideProps } from "lucide-react";
import Button from "@/components/ui/Button";

type Action = {
  Icon: FC<LucideProps>;
  text: string;
  onClick: () => void;
};

interface ActionsDropdownProps {
  actions: Action[];
  Icon: FC<LucideProps>;
  text: string;
}

const ActionsDropdown: React.FC<ActionsDropdownProps> = ({
  actions,
  Icon,
  text,
}) => {
  const [open, setOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <div className="relative" ref={dropdownRef}>
      <Button variant="outline" onClick={() => setOpen(!open)}>
        <span className="flex flex-row align-middle items-center gap-2">
          <Icon className="w-4 h-4" />
          {text}
        </span>
      </Button>
      {open && (
        <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-2 z-20">
          {actions.map((action, i) => (
            <button
              onClick={() => {
                setOpen(false);
                action.onClick();
              }}
              key={i}
              className="w-full text-left px-4 py-2 hover:bg-indigo-100"
            >
              <span className="flex flex-row align-middle items-center gap-2">
                <action.Icon className="text-indigo-500" /> {action.text}
              </span>
            </button>
          ))}
        </div>
      )}
    </div>
  );
};

export default ActionsDropdown;
