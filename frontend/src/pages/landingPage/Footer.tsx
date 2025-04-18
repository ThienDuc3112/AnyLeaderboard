import React from "react";

const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-800 text-white py-8">
      <div className="container mx-auto px-4 text-center">
        <p>&copy; 2024 LeaderBoard Maker. All rights reserved.</p>
        <div className="mt-4">
          <a className="text-gray-400 hover:text-white mx-2">Privacy Policy</a>
          <a className="text-gray-400 hover:text-white mx-2">
            Terms of Service
          </a>
          <a className="text-gray-400 hover:text-white mx-2">Contact Us</a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
