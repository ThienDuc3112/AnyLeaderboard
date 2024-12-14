import Button from "@/components/ui/button";
import React from "react";

const CTA: React.FC = () => {
  return (
    <section className="bg-indigo-700 text-white py-20">
      <div className="container mx-auto px-4 text-center">
        <h2 className="text-3xl md:text-4xl font-bold mb-4">
          Ready to Boost Engagement?
        </h2>
        <p className="text-xl mb-8">
          Create your first leaderboard today and see the difference it makes!
        </p>
        <Button variant="inverted" className="mx-auto py-2 px-6 text-lg">
          Start Free Trial
        </Button>
      </div>
    </section>
  );
};

export default CTA;
