import axios from "axios";
import React, { useEffect } from "react";

const Test: React.FC = () => {
  useEffect(() => {
    (async () => {
      try {
        const res = await axios.post("http://localhost:9999/auth/signup", {
          username: "not_huyen",
          displayName: "huyen the chuni guy",
          password: "12345678",
          email: "b@b.com",
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
