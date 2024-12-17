import React from "react";
import { Plus } from "lucide-react";
import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import Field from "./Field";

const NewLeaderboardPage: React.FC = () => {
  return (
    <div className="w-full">
      <div className="container max-w-3xl sm:py-10 mx-auto">
        <div className="shadow-md bg-indigo-50 rounded-lg overflow-hidden">
          <div className="p-6 border-b border-indigo-400">
            <h1 className="text-2xl font-bold text-gray-900">
              Create New Leaderboard
            </h1>
          </div>
          <div className="p-6">
            <div className="mb-6">
              <div className="space-y-4">
                <div>
                  <label
                    htmlFor="name"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Name
                  </label>
                  <Input id="name" placeholder="My Awesome Leaderboard" />
                </div>
                <div>
                  <label
                    htmlFor="description"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Description
                  </label>
                  <textarea
                    id="description"
                    rows={3}
                    className="mt-1 block w-full rounded-2xl px-3 py-2 border border-indigo-400 shadow-sm focus-within:ring-indigo-600 focus-within:outline-none focus-within:ring-1"
                    placeholder="Describe your leaderboard..."
                  />
                </div>
                <div>
                  <label
                    htmlFor="coverImageUrl"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Cover Image URL
                  </label>
                  <Input id="coverImageUrl" placeholder="https://..." />
                </div>
                <div className="flex items-center justify-between p-4 bg-gray-50 rounded-2xl border border-indigo-400">
                  <div>
                    <h3 className="text-sm font-medium text-gray-900">
                      Allow Anonymous Entries
                    </h3>
                    <p className="text-sm text-gray-500">
                      Let users submit entries without an account
                    </p>
                  </div>
                  <Switch />
                </div>
                <div className="flex items-center justify-between p-4 bg-gray-50 rounded-2xl border border-indigo-400">
                  <div>
                    <h3 className="text-sm font-medium text-gray-900">
                      Require Verification
                    </h3>
                    <p className="text-sm text-gray-500">
                      Entries must be verified by a moderator
                    </p>
                  </div>
                  <Switch />
                </div>
                <div>
                  <div className="flex items-center justify-between mb-2">
                    <h3 className="text-sm font-medium text-gray-700">
                      External Links
                    </h3>
                    <Button
                      variant="ghost"
                      size="small"
                      className="inline-flex items-center"
                    >
                      <span className="flex flex-row align-middle items-center gap-2">
                        <Plus className="h-4 w-4" />
                        Add Link
                      </span>
                    </Button>
                  </div>
                  <div className="grid grid-cols-3 gap-4">
                    <Input placeholder="Display Text" />
                    <Input placeholder="URL" className="col-span-2" />
                  </div>
                </div>
              </div>
            </div>
            <div className="border-t border-gray-200 pt-6">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-medium text-gray-900">
                  Leaderboard Fields
                </h3>
                <Button variant="filled" size="small">
                  <span className="flex flex-row align-middle items-center gap-1">
                    <Plus className="h-5 w-5" />
                    Add Field
                  </span>
                </Button>
              </div>
              <div className="bg-white shadow overflow-hidden rounded-2xl border border-indigo-400">
                <ul role="list" className="divide-y divide-indigo-400">
                  <li>
                    <Field index={0} />
                  </li>
                  <li>
                    <Field index={1} />
                  </li>
                </ul>
              </div>
            </div>
            <div className="mt-6 flex justify-end">
              <Button variant="filled" size="medium">
                Create Leaderboard
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NewLeaderboardPage;
