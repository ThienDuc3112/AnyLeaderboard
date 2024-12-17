import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import { Trash } from "lucide-react";
import React from "react";

const Field: React.FC = () => {
  return (
    <div className="px-4 py-4 sm:px-6">
      <div className="flex items-center justify-between">
        <Input placeholder="Display Name" className="flex-grow" value="" />
        <div className="ml-2 flex-shrink-0">
          <Button variant="ghost" size="small">
            <Trash className="h-5 w-5" />
          </Button>
        </div>
      </div>
      <div className="mt-3 sm:flex sm:justify-between">
        <div className="flex-1 mr-0 h-8 flex sm:mr-4">
          <select className="block w-full px-2 rounded-full border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50">
            <option value="TEXT">Text</option>
            <option value="SHORT_TEXT">Short Text</option>
            <option value="INTEGER">Integer</option>
            <option value="REAL">Real Number</option>
            <option value="DURATION">Duration</option>
            <option value="TIMESTAMP">Timestamp</option>
            <option value="OPTION">Option</option>
            <option value="USER">User</option>
          </select>
        </div>
        <div className="mt-3 flex items-center text-sm text-gray-500 sm:mt-0">
          <Switch label="Required" />
          <Switch label="Default Sort" className="ml-4" />
        </div>
      </div>
    </div>
  );
};

export default Field;
