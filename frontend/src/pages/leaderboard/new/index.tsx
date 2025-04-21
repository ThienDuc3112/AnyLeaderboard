import React, { useCallback } from "react";
import { Plus, Trash } from "lucide-react";
import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Switch from "@/components/ui/Switch";
import Field from "./Field";
import * as y from "yup";
import {
  FieldArray,
  FormikHelpers,
  FormikProvider,
  setIn,
  useFormik,
} from "formik";
import { api } from "@/utils/api";
import { useAtomValue } from "jotai";
import { sessionAtom } from "@/globalState/user";
import { useNavigate } from "react-router";
import { AxiosError } from "axios";

const SubmitSchema = y.object().shape({
  name: y.string().required("Name is required"),
  description: y.string().max(256, "Cannot exceed 256 characters"),
  coverImageUrl: y.string().url("Must be an image url"),
  externalLinks: y
    .array(
      y.object().shape({
        displayValue: y
          .string()
          .required("Display value is required")
          .max(32, "Cannot exceed 32 characters"),
        url: y.string().url("Must be a valid url").required("url is required"),
      }),
    )
    .max(5, "Cannot exceed 5 external links")
    .required("external links must exist"),
  uniqueSubmission: y.boolean(),
  allowAnonymous: y.boolean().when("uniqueSubmission", {
    is: true,
    then: (s) =>
      s.isFalse("Cannot allow anonymous when unique submission is on"),
    otherwise: (s) => s,
  }),
  requiredVerification: y.boolean(),
  fields: y
    .array(
      y.object().shape({
        name: y
          .string()
          .required("field name is required")
          .max(32, "Cannot exceed 32 characters"),
        required: y.boolean(),
        hidden: y.boolean(),
        type: y
          .string()
          .required("Field type must be specified")
          .oneOf(["TEXT", "NUMBER", "DURATION", "TIMESTAMP", "OPTION"]),
        forRank: y
          .boolean()
          .required("for rank is required")
          .when("type", {
            is: "OPTION",
            then: (schema) => schema.isFalse(),
            otherwise: (schema) => schema,
          }),
        options: y.array(
          y
            .string()
            .min(1, "Option is required")
            .max(32, "Cannot exceed 32 characters"),
        ),
      }),
    )
    .required("Atleast 1 field must exist")
    .min(1, "Atleast 1 field must exist")
    .max(10, "A leaderboard cannot have more than 10 fields"),
});

type SubmitType = y.InferType<typeof SubmitSchema>;

const NewLeaderboardPage: React.FC = () => {
  const session = useAtomValue(sessionAtom);
  const navigate = useNavigate();

  const onSubmit = useCallback(
    async (
      values: SubmitType,
      { setStatus, setErrors }: FormikHelpers<SubmitType>,
    ) => {
      try {
        const res = await api.post(
          "/leaderboards",
          {
            ...values,
            fields: values.fields.map((field, i) => ({
              ...field,
              fieldOrder: i + 1,
            })),
          },
          {
            headers: { Authorization: `Bearer ${session?.activeToken}` },
          },
        );
        const data = res.data;
        if (!data.id)
          setStatus(
            "Internal server error, please check if the leaderboard have been created",
          );
        else navigate(`/leaderboard/${data.id}`);
      } catch (error) {
        if (error instanceof AxiosError && error.status == 400) {
          const data = error.response?.data;
          if (!data)
            setStatus(
              "Internal server error, please report this to the developer",
            );
          else if (!!data.error) setStatus(data.error);
          else {
            const formattedErrors = Object.entries(
              data as Record<string, string>,
            ).reduce((acc, [key, value]) => setIn(acc, key, value), {});
            setErrors(formattedErrors);
          }
        } else
          setStatus(
            "Internal server error, please report this to the developer",
          );
      }
    },
    [session],
  );

  const p = useFormik({
    initialValues: {
      name: "",
      description: "",
      coverImageUrl: "",
      allowAnonymous: false,
      uniqueSubmission: false,
      requiredVerification: false,

      fields: [],
      externalLinks: [],
    } as SubmitType,
    validationSchema: SubmitSchema,
    onSubmit: onSubmit,
  });

  if (!session) navigate("/signin");

  return (
    <div className="w-full">
      <div className="container max-w-3xl mx-auto my-12">
        <div className="shadow-md bg-indigo-50 rounded-lg overflow-hidden">
          <div className="bg-indigo-600 text-white px-6 py-4">
            <h2 className="text-2xl font-bold">Create New Leaderboard</h2>
          </div>

          <FormikProvider value={p}>
            <form className="p-6">
              {/* Meta data */}
              <div className="mb-6">
                <div className="space-y-4">
                  <div>
                    <label
                      htmlFor="name"
                      className="block text-sm font-medium text-gray-700"
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
                      className="block text-sm font-medium text-gray-700"
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
                      className="block text-sm font-medium text-gray-700"
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
                              p.values.externalLinks &&
                              p.values.externalLinks.length >= 5
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
                            <div
                              key={index}
                              className="grid grid-cols-12 gap-4"
                            >
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

              {/* Fields data */}
              <div className="border-t border-gray-400 pt-6 mb-6">
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

              {/* Submit */}
              <div className="border-t border-gray-400 pt-6 flex flex-col">
                <div>{JSON.stringify(p.errors)}</div>
                <div className="flex justify-end">
                  <Button variant="filled" size="medium">
                    Create Leaderboard
                  </Button>
                </div>
              </div>
            </form>
          </FormikProvider>
        </div>
      </div>
    </div>
  );
};

export default NewLeaderboardPage;
