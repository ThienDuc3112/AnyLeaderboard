import React from "react";
import { BarChart, Users, Zap, Shield } from "lucide-react";

const features = [
  {
    icon: <BarChart className="h-8 w-8 text-indigo-600" />,
    title: "Real-time Updates",
    description: "See leaderboard changes as they happen with instant updates.",
  },
  {
    icon: <Users className="h-8 w-8 text-indigo-600" />,
    title: "Multiple Participants",
    description: "Support for unlimited participants in your leaderboards.",
  },
  {
    icon: <Zap className="h-8 w-8 text-indigo-600" />,
    title: "Customizable Design",
    description: "Personalize your leaderboard with custom colors and themes.",
  },
  {
    icon: <Shield className="h-8 w-8 text-indigo-600" />,
    title: "Secure & Private",
    description: "Keep your data safe with our robust security measures.",
  },
];

const Features: React.FC = () => {
  return (
    <section id="features" className="py-20 bg-white">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl font-bold text-center mb-12">
          Why Choose LeaderBoard Maker?
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature, index) => (
            <div key={index} className="text-center">
              <div className="mb-4 inline-block">{feature.icon}</div>
              <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
              <p className="text-gray-600">{feature.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Features;
