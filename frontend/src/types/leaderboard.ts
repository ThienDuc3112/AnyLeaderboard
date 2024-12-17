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
  data: Record<string, { value: any }>[];
  allowAnonymous?: boolean;
  requiredVerification?: boolean;
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
}

export interface OptionField {
  name: string;
  fieldName: string;
  type: "OPTION";
  options: string[];
  required?: boolean;
  hidden?: boolean;
}

export interface UserField {
  name: string;
  fieldName: string;
  type: "USER";
  allowAnonymous: boolean;
  required?: boolean;
  hidden?: boolean;
}
