import React, { useCallback } from "react";
import Button from "@/components/ui/Button";
import { FormikHelpers, FormikProvider, setIn, useFormik } from "formik";
import { api } from "@/utils/api";
import { useAtomValue } from "jotai";
import { sessionAtom } from "@/contexts/user";
import { useNavigate, useParams } from "react-router";
import { AxiosError } from "axios";
import { SubmitSchema, SubmitType } from "../schema";
import Metadatas from "../Metadatas";
import FieldsForm from "../FieldsForm";
import { useLeaderboard } from "@/hooks/useLeaderboard";

// Need to rework the api for this kind of updating
const UpdateLeaderboardPage: React.FC = () => {
  const session = useAtomValue(sessionAtom);
  const navigate = useNavigate();

  const { lid } = useParams();
  const { data, error, isLoading } = useLeaderboard(lid);

  const handleSubmit = useCallback(
    async (
      values: SubmitType,
      { setStatus, setErrors, resetForm }: FormikHelpers<SubmitType>,
    ) => {
      try {
        const res = await api.post(
          `/leaderboards/${lid}/update`,
          {
            ...values,
            fields: values.fields.map((field, i) => ({
              ...field,
              fieldOrder: i + 1,
              options: field.options
                ? field.options
                    .split(",")
                    .map((str) => str.trim())
                    .filter((str) => str != "")
                : undefined,
            })),
          },
          {
            headers: { Authorization: `Bearer ${session?.activeToken}` },
          },
        );
        const data = res.data;
        resetForm();
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
      name: data?.name ?? "",
      description: data?.description ?? "",
      coverImageUrl: data?.coverImageUrl ?? "",
      allowAnonymous: data?.allowAnonymous ?? false,
      uniqueSubmission: data?.uniqueSubmission ?? false,
      requiredVerification: data?.requiredVerification ?? false,

      externalLinks: data?.externalLinks ?? [],
      fields:
        data?.fields.map((field) => ({
          type: field.type ?? "NUMBER",
          name: field.name ?? "",
          forRank: (field as any).forRank ?? false,
          hidden: field.hidden ?? false,
          required: field.required ?? false,
          options: (field as any).options?.join(",") ?? "",
        })) ?? [],
    } as SubmitType,
    validationSchema: SubmitSchema,
    onSubmit: handleSubmit,
    enableReinitialize: true,
  });

  if (!session) navigate("/signin", { replace: true });
  if (isLoading) return <p>Loading...</p>;
  if (error || !data) return <p>Internal error occured</p>;
  if (session?.user.id !== data.creatorId) return <p>Forbidden</p>;

  return (
    <div className="w-full">
      <div className="container max-w-3xl mx-auto my-12">
        <div className="shadow-md bg-indigo-50 rounded-lg overflow-hidden">
          <div className="bg-indigo-600 text-white px-6 py-4">
            <h2 className="text-2xl font-bold">Update Leaderboard</h2>
          </div>

          <FormikProvider value={p}>
            <form className="p-6" onSubmit={p.handleSubmit}>
              {/* Meta data */}
              <Metadatas />

              {/* Fields data */}
              <FieldsForm />

              {/* Submit */}
              <div className="border-t border-gray-400 pt-6 flex flex-col">
                <div>{JSON.stringify(p.errors)}</div>
                <div>Status: {p.status}</div>
                <div className="flex justify-end">
                  <Button variant="filled" size="medium" type="submit" disabled>
                    Update Leaderboard
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

export default UpdateLeaderboardPage;
