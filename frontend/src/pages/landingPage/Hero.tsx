import Button from "@/components/ui/Button";
import React from "react";

const Hero: React.FC = () => {
  return (
    <section className="bg-indigo-700 text-white py-20">
      <div className="container mx-auto px-4 text-center">
        <h1 className="text-4xl md:text-5xl font-bold mb-4">
          Create Engaging Leaderboards in Minutes
        </h1>
        <p className="text-xl mb-8">
          Motivate your team, track progress, and boost competition with our
          easy-to-use leaderboard maker.
        </p>
        <Button variant="inverted" className="py-2 px-6 text-lg mx-auto">
          Get Started Free
        </Button>
      </div>
    </section>
  );
};

export default Hero;
