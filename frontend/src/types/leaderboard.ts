export interface LeaderboardPreview {
  id: string;
  name: string;
  description?: string;
  coverImageUrl?: string;
  entriesCount: number;
}

export interface LeaderboardFull extends LeaderboardPreview {
  fields: Field[];
  data: Entry[];
  allowAnonymous?: boolean;
  externalLinks?: ExternalLinkType[];
  requiredVerification?: boolean;
}

export interface Entry {
  id: string;
  updatedAt: string | Date;
  createdAt: string | Date;
  fields: Record<string, any>;
  username: string;
  verified: boolean;
}

export interface ExternalLinkType {
  displayValue: string;
  url: string;
}

export type Field = PrimitiveField | OptionField;
export interface PrimitiveField extends CommonFieldAttributes {
  type: "TEXT" | "SHORT_TEXT" | "NUMBER" | "DURATION" | "TIMESTAMP";
  for_rank?: boolean;
}

export interface OptionField extends CommonFieldAttributes {
  type: "OPTION";
  options: string[];
}

interface CommonFieldAttributes {
  name: string;
  required?: boolean;
  hidden?: boolean;
  fieldOrder: number;
}
