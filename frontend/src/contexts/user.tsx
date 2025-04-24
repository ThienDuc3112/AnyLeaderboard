import { User, UserSession } from "@/types/user";
import { atom } from "jotai";

export const userAtom = atom<User | undefined | null>(undefined);

export const sessionAtom = atom<UserSession | undefined | null>(undefined);
