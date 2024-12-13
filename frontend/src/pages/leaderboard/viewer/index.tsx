import React from "react";
import { MockApiResponse } from "@/mocks/apiResponse";
import { useParams } from "react-router";
import LeaderboardHeader from "./LeaderboardHeader";
import LeaderBoardContent from "./LeaderBoardContent";

const LeaderboardViewer: React.FC = () => {
  const { lid } = useParams();
  console.log(lid);
  const data = MockApiResponse;
  return (
    <div className="w-full max-w-4xl mx-auto bg-indigo-50 rounded-lg shadow-md overflow-hidden">
      <LeaderboardHeader data={data} />
      <LeaderBoardContent data={data} />
    </div>
  );
};

export default LeaderboardViewer;
