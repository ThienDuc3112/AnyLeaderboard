import Button from "@/components/ui/Button";
import Dropdown from "@/components/ui/Dropdown";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import { Trash } from "lucide-react";
import React from "react";

interface PropType {
  index: number;
}

const Field: React.FC<PropType> = ({ index }) => {
  return (
    <div className="px-4 py-4 sm:px-6">
      <div className="flex items-center justify-between">
        <Input placeholder="Display Name" className="flex-grow" />
        <div className="ml-2 flex-shrink-0">
          <Button variant="ghost" size="small">
            <Trash className="h-5 w-5" />
          </Button>
        </div>
      </div>
      <div className="mt-3 sm:flex sm:justify-between">
        <div className="flex-1 mr-0 h-8 flex sm:mr-4">
          <Dropdown
            options={[
              "TEXT",
              "SHORT_TEXT",
              "INTEGER",
              "REAL",
              "DURATION",
              "TIMESTAMP",
              "OPTION",
            ].map((o) => ({
              text: `${o[0]}${o.slice(1).toLowerCase().replace("_", " ")}`,
              value: o,
            }))}
            value={index === 0 ? "OPTION" : undefined}
          />
        </div>
        <div className="mt-3 flex items-center text-sm text-gray-500 sm:mt-0">
          <Switch label="Required" />
          <Switch label="Default Sort" className="ml-4" />
          <Switch label="Hidden" className="ml-4" />
        </div>
      </div>
      {/** Change this later */}
      {index === 0 && (
        <Input className="mt-3" placeholder="value 1, value 2,..." />
      )}
    </div>
  );
};

export default Field;
