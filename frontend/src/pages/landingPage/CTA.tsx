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
        <a
          href="#"
          className="bg-white text-indigo-700 py-2 px-6 rounded-full text-lg font-semibold hover:bg-indigo-100 transition duration-300"
        >
          Start Free Trial
        </a>
      </div>
    </section>
  );
};

export default CTA;
