import axios from "axios";
import React, { useEffect } from "react";

const Test: React.FC = () => {
  useEffect(() => {
    (async () => {
      try {
        const res = await axios.post("https://localhost:9999/auth/signup", {
          username: "tester",
          displayName: "The best tester that ever lived",
          password: "password",
          email: "test@testing.com",
        });
        console.log(res.status, res.data);
      } catch (error) {
        if (error instanceof axios.AxiosError) {
          console.log(error.response?.status, error.response?.data);
        } else {
          console.error(error);
        }
      }
    })();
  }, []);

  return <div>Test</div>;
};

export default Test;
