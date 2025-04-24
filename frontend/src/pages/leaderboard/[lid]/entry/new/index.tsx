import React, { useCallback, useMemo } from "react";
import Button from "@/components/ui/Button";
import FieldInput from "./Field";
import { useNavigate, useParams } from "react-router";
import { useLeaderboard } from "@/hooks/useLeaderboard";
import { FormikHelpers, FormikProvider, useFormik } from "formik";
import * as y from "yup";
import { api } from "@/utils/api";
import { AxiosError } from "axios";
import { useAtomValue } from "jotai";
import { sessionAtom } from "@/contexts/user";
import Input from "@/components/ui/Input";

const NewEntryPage: React.FC = () => {
  const { lid } = useParams();
  const { data: leaderboard, isLoading, error } = useLeaderboard(lid);
  const session = useAtomValue(sessionAtom);
  const navigate = useNavigate();

  const schema = useMemo(() => {
    if (!leaderboard) return y.object({} as Record<string, any>);
    const shape: Record<string, any> = {};
    leaderboard.fields.forEach((field) => {
      switch (field.type) {
        case "TEXT": {
          const fieldSchema = y.string().default("");
          if (field.required) fieldSchema.required("This field is required");
          shape[field.name] = fieldSchema;
          break;
        }
        case "TIMESTAMP":
        case "DURATION":
        case "NUMBER": {
          const fieldSchema = y.number().default(0);
          if (field.required) fieldSchema.required("This field is required");
          shape[field.name] = fieldSchema;
          break;
        }
        case "OPTION": {
          const fieldSchema = y
            .string()
            .oneOf(field.options)
            .default(field.options[0]);
          if (field.required) fieldSchema.required("This field is required");
          shape[field.name] = fieldSchema;
          break;
        }
      }
    });
    return y.object({
      ...shape,
      '"displayName': y.string().required("display name is required"),
    } as Record<string, any>);
  }, [leaderboard, lid]);

  const handleSubmit = useCallback(
    async (
      values: Record<string, any>,
      { setStatus }: FormikHelpers<Record<string, any>>,
    ) => {
      if (!leaderboard) return;
      const payload = { ...values };
      leaderboard.fields.forEach((field) => {
        if (field.type == "DURATION" && !!payload[field.name]) {
          payload[field.name] = new Date(payload[field.name]).valueOf();
        }
      });

      try {
        const res = await api.post(`/leaderboards/${lid}/entries`, payload);
        const { id } = res.data;
        if (!id)
          setStatus(
            "Internal server error, please check if the leaderboard have been created",
          );
        else navigate(`/leaderboard/${lid}`);
      } catch (error) {
        if (error instanceof AxiosError) {
          const errMsg = error.response?.data?.error;
          if (errMsg) setStatus(errMsg);
          else
            setStatus(
              "Internal server error, please report this to the developer",
            );
        } else {
          setStatus(
            "Internal server error, please report this to the developer",
          );
        }
      }
    },
    [leaderboard, lid],
  );

  const p = useFormik({
    onSubmit: handleSubmit,
    initialValues: schema.cast({
      '"displayName': session?.user.displayName ?? "",
    }),
    validationSchema: schema,
    enableReinitialize: true,
  });

  if (isLoading) return <p>Loading</p>;
  if (error || !leaderboard) return <p>Error</p>;
  if (!session && !leaderboard.allowAnonymous) {
    navigate("/login");
    return <p></p>;
  }

  return (
    <div className="w-full max-w-2xl mx-auto mt-12 bg-white shadow-md rounded-lg overflow-hidden">
      <div className="bg-indigo-600 text-white px-6 py-4">
        <h2 className="text-2xl font-bold">
          Add New Entry to {leaderboard.name}
        </h2>
      </div>
      <div className="p-6 bg-indigo-50">
        <FormikProvider value={p}>
          <form className="space-y-6" onSubmit={p.handleSubmit}>
            {leaderboard.fields
              .sort((a, b) => a.fieldOrder - b.fieldOrder)
              .map((field, i) => (
                <div key={i} className="space-y-2">
                  <label
                    htmlFor={field.name}
                    className="block text-sm font-medium text-gray-700"
                  >
                    {field.name}
                    {field.required && (
                      <span className="text-red-500 ml-1">*</span>
                    )}
                  </label>
                  <FieldInput field={field} />
                </div>
              ))}
            <div className="space-y-2">
              <label
                htmlFor={"username"}
                className="block text-sm font-medium text-gray-700"
              >
                Username
                <span className="text-red-500 ml-1">*</span>
              </label>
              <Input
                disabled={!!session}
                value={
                  session ? session.user.displayName : p.values['"displayName']
                }
                name={'"displayName'}
                onChange={p.handleChange}
                onBlur={p.handleBlur}
              />
            </div>

            <Button type="submit">Submit Entry</Button>
          </form>
        </FormikProvider>
        <div>{JSON.stringify(p.errors)}</div>
        <div>Status: {p.status}</div>
      </div>
    </div>
  );
};

export default NewEntryPage;
