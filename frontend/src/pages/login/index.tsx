import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import { sessionAtom } from "@/globalState/user";
import { api } from "@/utils/api";
import { AxiosError } from "axios";
import { Formik } from "formik";
import { useSetAtom } from "jotai";
import React from "react";
import { useNavigate } from "react-router";
import * as y from "yup";

const SigninSchema = y.object().shape({
  username: y.string().required("Username must be filled in"),
  password: y.string().required("Password must be filled in"),
});

const SignInPage: React.FC = () => {
  const navigate = useNavigate();
  const setSession = useSetAtom(sessionAtom);

  return (
    <div className="w-full mt-32 flex justify-center items-center">
      <div className="border border-indigo-200 rounded-xl bg-indigo-50">
        <h1 className="font-bold text-2xl text-center my-5">Sign in</h1>
        <Formik
          initialValues={{ username: "", password: "" }}
          validationSchema={SigninSchema}
          onSubmit={async (values, { setStatus }) => {
            setStatus(undefined);
            try {
              const data = (
                await api.post("/auth/login", values, {
                  withCredentials: true,
                })
              ).data;
              console.log(data);
              setSession({
                user: {
                  ...data.user,
                  createdAt: new Date(data.user.createdAt),
                },
                activeToken: data.activeToken,
              });
              navigate("/leaderboard");
            } catch (error) {
              if (
                error instanceof AxiosError &&
                (error.status == 401 || error.status == 400)
              ) {
                setStatus("Invalid credentials");
              } else {
                console.error(error);
                setStatus("Some error occured, please contact the developer");
              }
            }
          }}
        >
          {(p) => (
            <form
              onSubmit={p.handleSubmit}
              className="flex flex-col py-3 px-5 my-1"
            >
              <label htmlFor="username">Username</label>
              <Input
                name="username"
                type="text"
                onChange={p.handleChange}
                value={p.values.username}
                onBlur={p.handleBlur}
                className="w-72"
              />
              {p.errors.username && p.touched.username && (
                <em className="text-red-500 text-sm">{p.errors.username}</em>
              )}

              <label htmlFor="password" className="mt-6">
                Password
              </label>
              <Input
                name="password"
                type="password"
                onChange={p.handleChange}
                value={p.values.password}
                onBlur={p.handleBlur}
                className="w-72"
              />
              {p.errors.password && p.touched.password && (
                <em className="text-red-500 text-sm">{p.errors.password}</em>
              )}

              {p.status && p.dirty && (
                <div className="text-red-500 mt-4 text-sm">{p.status}</div>
              )}

              <Button disabled={!p.isValid} type="submit" className="mt-6">
                Login
              </Button>
            </form>
          )}
        </Formik>
      </div>
    </div>
  );
};

export default SignInPage;
