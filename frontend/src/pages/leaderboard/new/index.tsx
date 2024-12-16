import React from "react";
import { Plus, Trash2 } from "lucide-react";
import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";

const NewLeaderboardPage: React.FC = () => {
  return (
    <div className="w-full">
      <div className="container max-w-3xl py-10 mx-auto">
        <div className="bg-white shadow-md rounded-lg overflow-hidden">
          <div className="p-6 border-b border-gray-200">
            <h1 className="text-2xl font-bold text-gray-900">
              Create New Leaderboard
            </h1>
          </div>
          <div className="p-6">
            <div className="mb-6">
              <div className="flex space-x-2 mb-4">
                <Button variant="filled" size="medium">
                  Basic Info
                </Button>
                <Button variant="outline" size="medium">
                  Fields
                </Button>
              </div>
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
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                    placeholder="Describe your leaderboard..."
                  ></textarea>
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
                <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
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
                <div className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
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
                  <div className="grid grid-cols-2 gap-4">
                    <Input placeholder="Display Text" />
                    <Input placeholder="URL" />
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
                  <span className="flex flex-row align-middle items-center gap-2">
                    <Plus className="h-5 w-5" />
                    Add Field
                  </span>
                </Button>
              </div>
              <div className="bg-white shadow overflow-hidden sm:rounded-md">
                <ul role="list" className="divide-y divide-gray-200">
                  <li>
                    <div className="px-4 py-4 sm:px-6">
                      <div className="flex items-center justify-between">
                        <div className="flex-1 grid grid-cols-2 gap-4">
                          <Input placeholder="Display Name" value="Score" />
                          <Input placeholder="Technical Name" value="score" />
                        </div>
                        <div className="ml-2 flex-shrink-0">
                          <Button variant="ghost" size="small">
                            <Trash2 className="h-5 w-5" />
                          </Button>
                        </div>
                      </div>
                      <div className="mt-2 sm:flex sm:justify-between">
                        <div className="flex-1 mt-2 flex">
                          <select className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
                            <option>Text</option>
                            <option>Short Text</option>
                            <option>Integer</option>
                            <option>Real Number</option>
                            <option>Duration</option>
                            <option>Timestamp</option>
                            <option>Option</option>
                            <option>User</option>
                          </select>
                        </div>
                        <div className="mt-2 flex items-center text-sm text-gray-500 sm:mt-0">
                          <Switch label="Required" />
                          <Switch label="Default Sort" className="ml-4" />
                        </div>
                      </div>
                    </div>
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
