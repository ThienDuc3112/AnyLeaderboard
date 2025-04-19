import Button from "@/components/ui/Button";
import TextInput from "@/components/ui/Input";
import { api } from "@/utils/api";
import { AxiosError } from "axios";
import { Formik } from "formik";
import React from "react";
import { useNavigate } from "react-router";
import * as y from "yup"

const SignupSchema = y.object().shape({
  username: y.string()
    .required("Username is required")
    .min(3, "Username must be 3 characters or more")
    .max(64, "Username cannot go pass 64 characters"),
  email: y.string().required("Email is required").email("Must be valid email"),
  password: y.string()
    .required("Password is required")
    .min(8, "Password must be between 8 to 64 characters")
    .max(64, "Password must be between 8 to 64 characters"),
  passwordVerify: y.string()
    .required("Please confirm your password")
    .oneOf([y.ref("password")], "Password must match")
})

const SignupPage: React.FC = () => {
  const navigate = useNavigate();
  return (
    <div className="w-full mt-32 flex justify-center items-center">
      <div className="border border-indigo-200 rounded-xl bg-indigo-50">
        <h1 className="font-bold text-2xl text-center my-5">Sign up</h1>

        <Formik
          initialValues={{
            username: "",
            email: "",
            password: "",
            passwordVerify: ""
          }}
          validationSchema={SignupSchema}
          onSubmit={async (values, { setStatus, setErrors }) => {
            setStatus(undefined)
            try {
              await api.post("/auth/signup", {
                username: values.username,
                password: values.password,
                displayName: values.username,
                email: values.email,
              })
              navigate("/signin");
            } catch (error) {
              if (error instanceof AxiosError && (error.status == 400)) {
                const data = error.response?.data
                if (!data) {
                  console.error(error)
                  setStatus("Some error occured, please contact the developer")
                } else {
                  if (data.username) setErrors({ username: data.username })
                  if (data.email) setErrors({ email: data.email })
                  if (data.error) setStatus(data.error)
                }
              } else {
                console.error(error)
                setStatus("Some error occured, please contact the developer")
              }
            }

          }}
        >
          {
            p => (
              <form onSubmit={p.handleSubmit} className="flex flex-col py-3 px-5 my-1">
                <label htmlFor="username">Username</label>
                <TextInput
                  className="w-72"
                  name="username"
                  value={p.values.username}
                  onChange={p.handleChange}
                  onBlur={p.handleBlur}
                />
                {p.errors.username && p.touched.username && <em
                  className="text-red-500 text-sm"
                >{p.errors.username}</em>}

                <label htmlFor="email" className="mt-6">Email</label>
                <TextInput
                  className="w-72"
                  name="email"
                  value={p.values.email}
                  onChange={p.handleChange}
                  onBlur={p.handleBlur}
                />
                {p.errors.email && p.touched.email && <em
                  className="text-red-500 text-sm"
                >{p.errors.email}</em>}

                <label htmlFor="password" className="mt-6">Password</label>
                <TextInput
                  className="w-72"
                  type="password"
                  name="password"
                  value={p.values.password}
                  onChange={p.handleChange}
                  onBlur={p.handleBlur}
                />
                {p.errors.password && p.touched.password && <em
                  className="text-red-500 text-sm"
                >{p.errors.password}</em>}

                <label htmlFor="passwordVerify" className="mt-6">Verify password</label>
                <TextInput
                  className="w-72"
                  type="password"
                  id="passwordVerify"
                  value={p.values.passwordVerify}
                  onChange={p.handleChange}
                  onBlur={p.handleBlur}
                />
                {p.errors.passwordVerify && p.touched.passwordVerify && <em
                  className="text-red-500 text-sm"
                >{p.errors.passwordVerify}</em>}

                {p.status && p.dirty && <div
                  className="text-red-500 mt-4 text-sm"
                >{p.status}</div>}

                <Button type="submit" className="mt-6">
                  Sign up
                </Button>
              </form>
            )
          }
        </Formik>
      </div>
    </div>
  );
};

export default SignupPage;
