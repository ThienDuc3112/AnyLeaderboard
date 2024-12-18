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
  icon?: string;
}

export type Field = PrimitiveField | OptionField | UserField;

export interface PrimitiveField {
  name: string;
  fieldName: string;
  type: "TEXT" | "SHORT_TEXT" | "INTEGER" | "REAL" | "DURATION" | "TIMESTAMP";
  defaultSort?: boolean;
  required?: boolean;
  hidden?: boolean;
  fieldOrder: number;
}

export interface OptionField {
  name: string;
  fieldName: string;
  type: "OPTION";
  options: string[];
  required?: boolean;
  hidden?: boolean;
  fieldOrder: number;
}

export interface UserField {
  name: string;
  fieldName: string;
  type: "USER";
  allowAnonymous: boolean;
  required?: boolean;
  hidden?: boolean;
  fieldOrder: number;
}
