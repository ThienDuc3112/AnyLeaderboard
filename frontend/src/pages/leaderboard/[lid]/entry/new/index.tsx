import React from "react";
import { MockLeaderboardFull as leaderboard } from "@/mocks/leaderboardFull";
import Button from "@/components/ui/Button";
import FieldInput from "./Field";

const NewEntryPage: React.FC = () => {
  return (
    <div className="w-full max-w-2xl mx-auto mt-12 bg-white shadow-md rounded-lg overflow-hidden">
      <div className="bg-indigo-600 text-white px-6 py-4">
        <h2 className="text-2xl font-bold">
          Add New Entry to {leaderboard.name}
        </h2>
      </div>
      <div className="p-6 bg-indigo-50">
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
              <FieldInput field={field} />
            </div>
          ))}
          <Button>Submit Entry</Button>
        </form>
      </div>
    </div>
  );
};

export default NewEntryPage;
