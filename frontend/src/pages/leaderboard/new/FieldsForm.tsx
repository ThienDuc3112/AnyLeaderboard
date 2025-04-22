import Button from "@/components/ui/Button";
import { Plus } from "lucide-react";
import React from "react";
import Field from "./Field";
import { FieldArray, useFormikContext } from "formik";
import { SubmitType } from "./schema";

const FieldsForm: React.FC = () => {
  const p = useFormikContext<SubmitType>();
  return (
    <div className="border-t border-gray-400 pt-6 mb-6">
      <FieldArray name="fields">
        {({ push, remove }) => (
          <>
            <div className="flex items-center justify-between mb-4">
              <h3 className="text-lg font-medium text-gray-900">
                Leaderboard Fields
              </h3>
              <Button
                variant="filled"
                size="small"
                onClick={(e) => {
                  e.preventDefault();
                  e.stopPropagation();
                  push({
                    name: "",
                    type: "TEXT",
                    forRank: false,
                    hidden: false,
                    options: "",
                    required: false,
                  } as SubmitType["fields"][number]);
                }}
              >
                <span className="flex flex-row align-middle items-center gap-1">
                  <Plus className="h-5 w-5" />
                  Add Field
                </span>
              </Button>
            </div>

            <div className="bg-white shadow overflow-hidden rounded-2xl border border-indigo-400">
              <ul role="list" className="divide-y divide-indigo-400">
                {p.values.fields.map((_, index) => (
                  <li key={index}>
                    <Field index={index} remove={remove} />
                  </li>
                ))}
              </ul>
            </div>
          </>
        )}
      </FieldArray>
    </div>
  );
};

export default FieldsForm;
