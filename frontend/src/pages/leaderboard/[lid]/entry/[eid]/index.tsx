import React, { useCallback } from "react";
import { Field } from "@/types/leaderboard";
import LeaderboardHeader from "../../LeaderboardHeader";
import { MoveLeft } from "lucide-react";
import Button from "@/components/ui/Button";
import { useNavigate, useParams } from "react-router";
import { useEntry } from "@/hooks/useEntry";

const EntryViewPage: React.FC = () => {
  const navigate = useNavigate();
  const { lid, eid } = useParams();
  const { data, isLoading, error } = useEntry(lid, eid);
  const renderFieldValue = useCallback((field: Field, value: any) => {
    switch (field.type) {
      case "TEXT":
      case "NUMBER":
      case "OPTION":
        return String(value);
      case "DURATION": {
        const ms = parseInt(value);
        const hours = Math.floor(ms / 3600000);
        const minutes = Math.floor((ms % 3600000) / 60000);
        const seconds = Math.floor((ms % 60000) / 1000);
        const milliseconds = ms % 1000;
        return `${hours.toString().padStart(2, "0")}:${minutes
          .toString()
          .padStart(2, "0")}:${seconds
          .toString()
          .padStart(2, "0")}.${milliseconds.toString().padStart(3, "0")}`;
      }
      case "TIMESTAMP":
        return new Date(value).toUTCString();
      default:
        return "N/A";
    }
  }, []);

  if (isLoading) return <p>Loading...</p>;
  if (!data || error) return <p>An error occured</p>;

  return (
    <div className="w-full mt-12">
      <div className="max-w-5xl mx-auto bg-indigo-50 rounded-lg shadow-md overflow-hidden">
        <LeaderboardHeader data={data.leaderboard} />

        <div className="px-6 flex justify-between">
          <span className="font-semibold text-xl">Entry details</span>
          <Button
            size="small"
            variant="ghost"
            onClick={() => {
              navigate(-1);
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
            {data.leaderboard.fields.map((field) => (
              <div key={field.name} className="sm:col-span-1">
                <dt className="text-sm font-medium text-gray-500">
                  {field.name}
                </dt>
                <dd className="mt-1 text-sm">
                  <span className="text-gray-900">
                    {renderFieldValue(field, data.entry.fields[field.name])}
                  </span>
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
