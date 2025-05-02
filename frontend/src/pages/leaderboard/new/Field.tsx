import Button from "@/components/ui/Button";
import Dropdown from "@/components/ui/Dropdown";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import { useFormikContext } from "formik";
import { Trash } from "lucide-react";
import React from "react";
import { SubmitType } from "./schema";

interface PropType {
  index: number;
  remove: (index: number) => any;
}

const Field: React.FC<PropType> = ({ index, remove }) => {
  const p = useFormikContext<SubmitType>();
  return (
    <div className="px-4 py-4 sm:px-6">
      <div className="flex items-center justify-between">
        <Input
          name={`fields[${index}].name`}
          placeholder="Display Name"
          className="flex-grow"
          value={p.values.fields[index].name}
          onChange={p.handleChange}
          onBlur={p.handleBlur}
        />
        <div className="ml-2 flex-shrink-0">
          <Button
            variant="ghost"
            size="small"
            onClick={(e) => {
              e.preventDefault();
              e.stopPropagation();
              remove(index);
            }}
          >
            <Trash className="h-5 w-5" />
          </Button>
        </div>
      </div>
      <div className="mt-3 sm:flex sm:justify-between">
        <div className="flex-1 mr-0 h-8 flex sm:mr-4">
          <Dropdown
            options={["TEXT", "NUMBER", "DURATION", "TIMESTAMP", "OPTION"].map(
              (o) => ({
                text: `${o[0]}${o.slice(1).toLowerCase().replace("_", " ")}`,
                value: o,
              }),
            )}
            name={`fields[${index}].type`}
            value={p.values.fields[index].type}
            onChange={p.handleChange}
            onBlur={p.handleBlur}
          />
        </div>
        <div className="mt-3 flex items-center text-sm text-gray-500 sm:mt-0">
          <Switch
            name={`fields[${index}].required`}
            label="Required"
            checked={p.values.fields[index].required}
            onChange={p.handleChange}
          />
          <Switch
            name={`fields[${index}].forRank`}
            label="For rank"
            className="ml-4"
            checked={p.values.fields[index].forRank}
            onChange={p.handleChange}
          />
          <Switch
            name={`fields[${index}].hidden`}
            label="Hidden"
            className="ml-4"
            checked={p.values.fields[index].hidden}
            onChange={p.handleChange}
          />
        </div>
      </div>
      {/** Change this later */}
      {p.values.fields[index].type == "OPTION" && (
        <Input
          className="mt-3"
          placeholder="value 1, value 2,..."
          name={`fields[${index}].options`}
          value={p.values.fields[index].options}
          onChange={p.handleChange}
          onBlur={p.handleBlur}
        />
      )}
    </div>
  );
};

export default Field;
