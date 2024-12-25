import { Field } from "@/types/leaderboard";

export const isDigitField = (field: Field): boolean => {
  return ["NUMBER", "DURATION"].includes(field.type);
};
