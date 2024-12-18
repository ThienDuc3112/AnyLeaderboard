import React from "react";
import { MockLeaderboardFull as leaderboard } from "@/mocks/leaderboardFull";
import { Field } from "@/types/leaderboard";
import LeaderboardHeader from "../../LeaderboardHeader";
import { MoveLeft } from "lucide-react";
import Button from "@/components/ui/Button";
import { useNavigate } from "react-router";

const EntryViewPage: React.FC = () => {
  const navigate = useNavigate();
  const entry = leaderboard.data[0];
  const renderFieldValue = (field: Field, value: any) => {
    switch (field.type) {
      case "TEXT":
      case "SHORT_TEXT":
      case "INTEGER":
      case "REAL":
        return <span className="text-gray-900">{value}</span>;
      case "DURATION": {
        const ms = parseInt(value);
        const hours = Math.floor(ms / 3600000);
        const minutes = Math.floor((ms % 3600000) / 60000);
        const seconds = Math.floor((ms % 60000) / 1000);
        const milliseconds = ms % 1000;
        return (
          <span className="text-gray-900">
            {`${hours.toString().padStart(2, "0")}:${minutes
              .toString()
              .padStart(2, "0")}:${seconds
              .toString()
              .padStart(2, "0")}.${milliseconds.toString().padStart(3, "0")}`}
          </span>
        );
      }
      case "OPTION":
        return <span className="text-gray-900">{value}</span>;
      case "USER":
        return (
          <span className="text-gray-900">
            {value.username}
            {value.userId ? ` (ID: ${value.userId})` : ""}
            {!value.userId && " (Anonymous)"}
          </span>
        );
      default:
        return <span className="text-gray-500">N/A</span>;
    }
  };

  return (
    <div className="w-full mt-12">
      <div className="max-w-5xl mx-auto bg-indigo-50 rounded-lg shadow-md overflow-hidden">
        <LeaderboardHeader data={leaderboard} />
        <div className="px-6 flex justify-between">
          <span className="font-semibold text-xl">Entry details</span>
          <Button
            size="small"
            variant="ghost"
            onClick={() => {
              navigate(`/leaderboard/${leaderboard.id}`);
            }}
          >
            <span className="flex flex-row align-middle items-center gap-2">
              <MoveLeft className="w-4 h-4" />
              Return
            </span>
          </Button>
        </div>
        <div className="p-6">
          <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
            {leaderboard.fields.map((field) => (
              <div key={field.fieldName} className="sm:col-span-1">
                <dt className="text-sm font-medium text-gray-500">
                  {field.name}
                </dt>
                <dd className="mt-1 text-sm">
                  {renderFieldValue(
                    field,
                    entry.fields[field.fieldName]?.value
                  )}
                </dd>
              </div>
            ))}
          </dl>
        </div>
      </div>
    </div>
  );
};

export default EntryViewPage;
