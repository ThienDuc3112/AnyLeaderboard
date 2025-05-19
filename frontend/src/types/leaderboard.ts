export interface LeaderboardPreview {
  id: number;
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
  uniqueSubmission?: boolean;
  creatorId: number;
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
  id: number;
  displayValue: string;
  url: string;
}

export type Field = PrimitiveField | OptionField | TextField;
export interface PrimitiveField extends CommonFieldAttributes {
  type: "NUMBER" | "DURATION" | "TIMESTAMP";
  for_rank?: boolean;
}
export interface TextField extends CommonFieldAttributes {
  type: "TEXT";
}

export interface OptionField extends CommonFieldAttributes {
  type: "OPTION";
  options: string[];
}

interface CommonFieldAttributes {
  id: number;
  name: string;
  required?: boolean;
  hidden?: boolean;
  fieldOrder: number;
}
