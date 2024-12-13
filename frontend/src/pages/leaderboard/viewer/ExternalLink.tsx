import React from "react";
import { ExternalLink as ExternalLinkIcon } from "lucide-react";
import { ExternalLinkType } from "@/types/leaderboard";

interface PropType {
  link: ExternalLinkType;
}

const ExternalLink: React.FC<PropType> = ({ link }: PropType) => {
  return (
    <a
      key={link.url}
      href={link.url}
      target="_blank"
      rel="noopener noreferrer"
      className="inline-flex items-center h-7 px-2 text-xs font-semibold text-white bg-indigo-600 hover:bg-indigo-400 border border-gray-300 rounded-md transition focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    >
      <ExternalLinkIcon className="h-3 w-3 mr-1" />
      {link.displayValue}
    </a>
  );
};

export default ExternalLink;
