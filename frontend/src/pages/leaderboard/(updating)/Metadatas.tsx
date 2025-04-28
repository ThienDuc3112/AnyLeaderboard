import Input from "@/components/ui/Input";
import { FieldArray, useFormikContext } from "formik";
import React from "react";
import { SubmitType } from "./schema";
import Switch from "@/components/ui/Switch";
import Button from "@/components/ui/Button";
import { Plus, Trash } from "lucide-react";

const Metadatas: React.FC = () => {
  const p = useFormikContext<SubmitType>();
  return (
    <div className="mb-6">
      <div className="space-y-4">
        <div>
          <label
            htmlFor="name"
            className={`block text-sm font-medium ${p.touched.name && p.errors.name ? "text-red-500" : "text-gray-700"}`}
          >
            Name
          </label>
          <Input
            name="name"
            placeholder="My Awesome Leaderboard"
            value={p.values.name}
            onChange={(e) => p.handleChange(e)}
            onBlur={(e) => p.handleBlur(e)}
          />
        </div>

        <div>
          <label
            htmlFor="description"
            className={`block text-sm font-medium ${p.touched.description && p.errors.description ? "text-red-500" : "text-gray-700"}`}
          >
            Description
          </label>
          <textarea
            name="description"
            rows={3}
            className="mt-1 block w-full rounded-2xl px-3 py-2 border border-indigo-400 shadow-sm focus-within:ring-indigo-600 focus-within:outline-none focus-within:ring-1"
            placeholder="Describe your leaderboard..."
            value={p.values.description}
            onChange={(e) => p.handleChange(e)}
            onBlur={(e) => p.handleBlur(e)}
          />
        </div>

        <div>
          <label
            htmlFor="coverImageUrl"
            className={`block text-sm font-medium ${p.touched.coverImageUrl && p.errors.coverImageUrl ? "text-red-500" : "text-gray-700"}`}
          >
            Cover Image URL
          </label>
          <Input
            name="coverImageUrl"
            placeholder="https://..."
            value={p.values.coverImageUrl}
            onChange={(e) => p.handleChange(e)}
            onBlur={(e) => p.handleBlur(e)}
          />
        </div>

        <div className="flex items-center justify-between p-4 bg-white rounded-2xl border border-indigo-400">
          <div>
            <h3 className="text-sm font-medium text-gray-900">
              Allow Anonymous Entries
            </h3>
            <p className="text-sm text-gray-500">
              Let users submit entries without an account
            </p>
          </div>
          <Switch
            name="allowAnonymous"
            checked={p.values.allowAnonymous}
            onChange={p.handleChange}
          />
        </div>

        <div className="flex items-center justify-between p-4 bg-white rounded-2xl border border-indigo-400">
          <div>
            <h3 className="text-sm font-medium text-gray-900">
              Require Verification
            </h3>
            <p className="text-sm text-gray-500">
              Entries must be verified by a moderator
            </p>
          </div>
          <Switch
            name="requiredVerification"
            checked={p.values.requiredVerification}
            onChange={p.handleChange}
          />
        </div>

        <div className="flex items-center justify-between p-4 bg-white rounded-2xl border border-indigo-400">
          <div>
            <h3 className="text-sm font-medium text-gray-900">
              Unique submission
            </h3>
            <p className="text-sm text-gray-500">
              Only 1 submission per user will show up
            </p>
          </div>
          <Switch
            name="uniqueSubmission"
            checked={p.values.uniqueSubmission}
            onChange={p.handleChange}
          />
        </div>

        <FieldArray name="externalLinks">
          {({ push, remove }) => (
            <>
              <div className="flex items-center justify-between mb-2">
                <h3 className="text-sm font-medium text-gray-700">
                  External Links
                </h3>
                <Button
                  variant="ghost"
                  size="small"
                  className="inline-flex items-center"
                  disabled={
                    p.values.externalLinks && p.values.externalLinks.length >= 5
                  }
                  onClick={(e) => {
                    e.preventDefault();
                    e.stopPropagation();
                    push({ displayValue: "", url: "" });
                  }}
                >
                  <span className="flex flex-row align-middle items-center gap-2">
                    <Plus className="h-4 w-4" />
                    Add Link
                  </span>
                </Button>
              </div>
              {/* External links list */}
              <div className="flex flex-col gap-2">
                {p.values.externalLinks.length == 0 ? (
                  <span className="text-sm font-medium text-gray-700">
                    No external links
                  </span>
                ) : null}
                {p.values.externalLinks.map((link, index) => (
                  <div key={index} className="grid grid-cols-12 gap-4">
                    <Input
                      placeholder="Display Text"
                      className="col-span-3"
                      name={`externalLinks[${index}].displayValue`}
                      value={link.displayValue}
                      onChange={p.handleChange}
                      onBlur={p.handleBlur}
                    />
                    <Input
                      placeholder="URL"
                      className="col-span-8"
                      name={`externalLinks[${index}].url`}
                      value={link.url}
                      onChange={p.handleChange}
                      onBlur={p.handleBlur}
                    />
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
                ))}
              </div>
            </>
          )}
        </FieldArray>
      </div>
    </div>
  );
};

export default Metadatas;
