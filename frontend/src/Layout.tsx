import React from "react";
import NavBar from "@/components/NavBar";
import { Outlet } from "react-router";

const Layout: React.FC = () => {
  return (
    <div className="flex flex-col h-screen w-screen overflow-auto p-0 m-0">
      <NavBar />
      <Outlet />
    </div>
  );
};

export default Layout;
