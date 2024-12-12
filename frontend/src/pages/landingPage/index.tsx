import React from "react";
import Hero from "@/pages/landingPage/Hero";
import Features from "@/pages/landingPage/Features";
import CTA from "@/pages/landingPage/CTA";
import Footer from "@/pages/landingPage/Footer";

const LandingPage: React.FC = () => {
  return (
    <div className="bg-gray-100">
      <Hero />
      <Features />
      <CTA />
      <Footer />
    </div>
  );
};

export default LandingPage;
