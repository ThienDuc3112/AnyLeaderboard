import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import { sessionAtom } from "@/contexts/user";
import { api } from "@/utils/api";
import { AxiosError } from "axios";
import { useAtom } from "jotai";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router";

const ProfilePage = () => {
  const navigate = useNavigate();
  const [session, setSession] = useAtom(sessionAtom);
  const [password, setPassword] = useState("");

  const deleteAccount = useCallback(async () => {
    try {
      await api.delete(`/users`, {
        data: {
          password,
          id: session?.user.id,
        },
      });
      alert("Deleted");
      setSession(null);
      navigate("/");
    } catch (error) {
      console.error(error);
      if (error instanceof AxiosError) {
        if (error.status && error.status < 500) {
          const message: string = error.response?.data?.error;
          if (message) {
            alert(message);
          } else {
            alert("An error occured");
          }
        } else {
          alert("Internal server error");
        }
      } else {
        alert("Failed to delete user");
      }
    }
  }, [session, password]);

  return (
    <div>
      <p>Nothing here</p>
      <p>Delete your account?</p>
      <div className="flex flex-row align-middle">
        <div>
          <label htmlFor="password" className={`block text-sm font-medium`}>
            Password
          </label>
          <Input
            name="password"
            type="password"
            placeholder="Password to delete account"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <Button onClick={deleteAccount}>Delete</Button>
      </div>
    </div>
  );
};

export default ProfilePage;
