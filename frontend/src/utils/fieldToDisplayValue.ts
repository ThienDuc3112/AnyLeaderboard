import { Entry, Field } from "@/types/leaderboard";
import { formatDuration } from "./formatDuration";

export const fieldToDisplayValue = (row: Entry, field: Field): string => {
  if (!row.fields[field.fieldName]) return "";
  const value = row.fields[field.fieldName].value;
  switch (field.type) {
    case "TEXT":
    case "SHORT_TEXT":
      return shortenStr(value);
    case "INTEGER":
    case "REAL":
      return shortenNum(value);
    case "DURATION":
      return formatDuration(value);
    case "TIMESTAMP":
      return value;
    case "OPTION":
      return value;
    case "USER":
      return value.username;
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
