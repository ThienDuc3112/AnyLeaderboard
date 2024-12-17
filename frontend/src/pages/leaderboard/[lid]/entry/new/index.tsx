import React from "react";
import { Field } from "@/types/leaderboard";
import { MockLeaderboardFull as leaderboard } from "@/mocks/leaderboardFull";
import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import Dropdown from "@/components/ui/Dropdown";

const NewEntryPage: React.FC = () => {
  const renderField = (field: Field) => {
    switch (field.type) {
      case "TEXT":
      case "SHORT_TEXT":
        return <Input type="text" placeholder={`Enter ${field.name}`} />;

      case "INTEGER":
        return <Input type="number" placeholder={`Enter ${field.name}`} />;

      case "REAL":
        return (
          <Input
            type="number"
            step="0.01"
            placeholder={`Enter ${field.name}`}
          />
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

  return (
    <div className="w-full max-w-2xl mx-auto bg-white shadow-md rounded-lg overflow-hidden">
      <div className="bg-indigo-600 text-white px-6 py-4">
        <h2 className="text-2xl font-bold">
          Add New Entry to {leaderboard.name}
        </h2>
      </div>
      <div className="p-6">
        <form className="space-y-6">
          {leaderboard.fields.map((field) => (
            <div key={field.fieldName} className="space-y-2">
              <label
                htmlFor={field.fieldName}
                className="block text-sm font-medium text-gray-700"
              >
                {field.name}
                {field.required && <span className="text-red-500 ml-1">*</span>}
              </label>
              {renderField(field)}
            </div>
          ))}
          <Button>Submit Entry</Button>
        </form>
      </div>
    </div>
  );
};

export default NewEntryPage;
