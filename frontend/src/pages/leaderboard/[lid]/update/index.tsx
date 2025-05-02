import React from "react";
import Button from "@/components/ui/Button";
import { useAtomValue } from "jotai";
import { sessionAtom } from "@/contexts/user";
import { useNavigate, useParams } from "react-router";
import { useLeaderboard } from "@/hooks/useLeaderboard";

// Need to rework the api for this kind of updating
const UpdateLeaderboardPage: React.FC = () => {
  const session = useAtomValue(sessionAtom);
  const navigate = useNavigate();

  const { lid } = useParams();
  const { data, error, isLoading } = useLeaderboard(lid);


  if (!session) navigate("/signin", { replace: true });
  if (isLoading) return <p>Loading...</p>;
  if (error || !data) return <p>Internal error occured</p>;
  if (session?.user.id !== data.creatorId) return <p>Forbidden</p>;

  return (
    <div className="w-full">
      <div className="container max-w-3xl mx-auto my-12">
        <div className="shadow-md bg-indigo-50 rounded-lg overflow-hidden">
          <div className="bg-indigo-600 text-white px-6 py-4">
            <h2 className="text-2xl font-bold">Update Leaderboard</h2>
          </div>

          <form className="p-6" >
            {/* Meta data */}

            {/* Fields data */}

            {/* Submit */}
            <div className="border-t border-gray-400 pt-6 flex flex-col">
              <div className="flex justify-end">
                <Button variant="filled" size="medium" type="submit" disabled>
                  Update Leaderboard
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default UpdateLeaderboardPage;
