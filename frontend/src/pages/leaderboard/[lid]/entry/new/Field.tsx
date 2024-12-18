import Dropdown from "@/components/ui/Dropdown";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import { Field as FieldType } from "@/types/leaderboard";
import React from "react";

interface PropType {
  field: FieldType;
}

const FieldInput: React.FC<PropType> = ({ field }) => {
  switch (field.type) {
    case "TEXT":
    case "SHORT_TEXT":
      return <Input type="text" placeholder={`Enter ${field.name}`} />;

    case "INTEGER":
      return <Input type="number" placeholder={`Enter ${field.name}`} />;

    case "REAL":
      return (
        <Input type="number" step="0.01" placeholder={`Enter ${field.name}`} />
      );

    case "DURATION":
      return <Input type="text" placeholder="HH:MM:SS.mmm" />;

    case "OPTION":
      return (
        <div className="h-8 flex flex-1">
          <Dropdown
            options={field.options.map((o) => ({ text: o, value: o }))}
          />
        </div>
      );

    case "USER":
      return (
        <div className="space-y-4">
          <div className="flex items-center space-x-2 gap-2">
            <Input className="flex-grow" type="text" placeholder="Username" />
            <Switch label="Submit anonymously" />
          </div>
        </div>
      );

    default:
      return null;
  }
};

export default FieldInput;
