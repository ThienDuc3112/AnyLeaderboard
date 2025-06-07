import React, { useCallback, useState } from "react";
import { useNavigate } from "react-router";
import { useAtomValue } from "jotai";
import { MoreVertical, Pencil, Plus, Trash } from "lucide-react";

import ExternalLink from "@/pages/leaderboard/[lid]/ExternalLink";
import Button from "@/components/ui/Button";
import ActionsDropdown from "@/components/ui/ActionDropdown";
import { sessionAtom } from "@/contexts/user";
import { LeaderboardFull } from "@/types/leaderboard";
import Popup from "@/components/ui/Popup";
import { api } from "@/utils/api";
import { AxiosError } from "axios";

interface PropType {
  data: LeaderboardFull;
}

const LeaderboardHeader: React.FC<PropType> = ({ data }) => {
  const navigate = useNavigate();
  const userSession = useAtomValue(sessionAtom);
  const [deletePopup, setDeletePopup] = useState(false);

  const deleteLeaderboard = useCallback(async () => {
    try {
      await api.delete(`/leaderboards/${data.id}`, {
        headers: {
          Authorization: `Bearer ${userSession?.activeToken}`,
        },
      });
      alert("Leaderboard deleted");
      navigate("/leaderboard");
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.status && error.status < 500) {
          const errorMsg: string = error.response?.data?.error;
          if (errorMsg) {
            alert(errorMsg);
            return;
          } else {
            console.error(error);
            alert("An error occurred");
          }
        } else {
          console.error(error);
          alert("Internal server error");
        }
      } else {
        console.error(error);
        alert(
          "Unknown error occur, check if your leaderboard have been deleted or not",
        );
        navigate("/leaderboard");
      }
    }
  }, [userSession, data]);

  return (
    <div className="p-6 space-y-4">
      <div className="flex gap-4">
        <div className="relative w-32 h-32 flex-shrink-0">
          <img
            src={data.coverImageUrl || "/placeholder.svg?height=128&width=128"}
            alt={`${data.name} cover`}
            className="rounded-lg object-cover w-full h-full"
          />
        </div>

        <div className="flex-grow">
          <div className="flex justify-between items-start">
            <h2 className="text-2xl font-bold">{data.name}</h2>
            {userSession && userSession.user.id === data.creatorId ? (
              <ActionsDropdown
                text="Actions"
                Icon={MoreVertical}
                actions={[
                  {
                    Icon: Plus,
                    text: "New entry",
                    onClick: () => {
                      navigate(`/leaderboard/${data.id}/entry/new`);
                    },
                  },
                  {
                    Icon: Pencil,
                    text: "Edit leaderboard",
                    onClick: () => {
                      navigate(`/leaderboard/${data.id}/update`);
                    },
                  },
                  {
                    Icon: Trash,
                    text: "Delete",
                    onClick() {
                      setDeletePopup(true);
                    },
                  },
                ]}
              />
            ) : (
              <Button
                variant="outline"
                onClick={() => {
                  navigate(`/leaderboard/${data.id}/entry/new`);
                }}
              >
                <span className="flex flex-row align-middle items-center gap-2">
                  <Plus className="h-4 w-4" />
                  New Entry
                </span>
              </Button>
            )}
          </div>
          {data.description && (
            <p className="text-sm text-gray-600 mt-1">{data.description}</p>
          )}
          <div className="mt-2 text-xs text-gray-600">
            <span className="font-mono">{data.entriesCount}</span> entries
          </div>
          <div className="flex flex-wrap gap-2 mt-2">
            {data.externalLinks?.map((link, i) => (
              <ExternalLink key={i} link={link} />
            ))}
          </div>
          <Popup isOpen={deletePopup} onClose={() => setDeletePopup(false)}>
            <p>Are you sure?</p>
            <div className="m-4 flex justify-between">
              <Button variant="filled" onClick={deleteLeaderboard}>
                Pretty sure
              </Button>
              <Button variant="outline" onClick={() => setDeletePopup(false)}>
                Nah
              </Button>
            </div>
          </Popup>
        </div>
      </div>
    </div>
  );
};

export default LeaderboardHeader;
