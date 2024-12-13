import { Field } from "@/types/leaderboard";

export const isDigitField = (field: Field): boolean => {
  return ["REAL", "INTEGER", "DURATION"].includes(field.type);
};
