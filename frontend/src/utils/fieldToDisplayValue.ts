import { Entry, Field } from "@/types/leaderboard";
import { formatDuration } from "./formatDuration";

export const fieldToDisplayValue = (row: Entry, field: Field): string => {
  if (!row.fields[field.id.toString()]) {
    return "";
  }
  const value = row.fields[field.id.toString()];
  switch (field.type) {
    case "TEXT":
      return shortenStr(value);
    case "NUMBER":
      return shortenNum(value);
    case "DURATION":
      return formatDuration(value);
    case "TIMESTAMP":
      return value;
    case "OPTION":
      return value;
    default:
      return "";
  }
};

const shortenStr = (str: string): string => {
  if (str.length > 15) return str.slice(0, 12) + "...";
  return str;
};

const shortenNum = (num: number): string => {
  if (num.toString().length > 15) return num.toExponential(9).toString();
  return num.toString();
};
