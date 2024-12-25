export interface LeaderboardPreview {
  id: string;
  name: string;
  description?: string;
  coverImageUrl?: string;
  externalLinks?: ExternalLinkType[];
  entryCount: number;
}

export interface LeaderboardFull extends LeaderboardPreview {
  fields: Field[];
  data: Entry[];
  allowAnonymous?: boolean;
  requiredVerification?: boolean;
}

export interface Entry {
  id: string;
  updatedAt: string | Date;
  createdAt: string | Date;
  fields: Record<string, { value: any }>;
}

export interface ExternalLinkType {
  displayValue: string;
  url: string;
}

export type Field = PrimitiveField | OptionField | UserField;
export interface PrimitiveField extends CommonFieldAttributes {
  type: "TEXT" | "SHORT_TEXT" | "NUMBER" | "DURATION" | "TIMESTAMP";
  for_rank?: boolean;
}

export interface OptionField extends CommonFieldAttributes {
  type: "OPTION";
  options: string[];
}

export interface UserField extends CommonFieldAttributes {
  type: "USER";
  allowAnonymous: boolean;
}

interface CommonFieldAttributes {
  name: string;
  fieldName: string;
  required?: boolean;
  hidden?: boolean;
  fieldOrder: number;
}
