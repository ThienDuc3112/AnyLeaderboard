export interface LeaderboardFull{
  name: string
  description?: string
  coverImageUrl?: string
  externalLinks?: ExternalLinkType[]
  fields: Field[]
  data: Record<string, any>[]
}

export interface ExternalLinkType {
  displayValue: string
  url: string
  icon?: string
}

export type Field = PrimitiveField | OptionField | UserField

interface PrimitiveField {
  name: string
  fieldName: string
  type: "TEXT" | "SHORT_TEXT" | "INTEGER" | "REAL" | "DURATION" | "TIMESTAMP"
  defaultSort?: boolean
}

interface OptionField {
  name: string
  fieldName: string
  type: "OPTION"
  options: string[]
}

interface UserField {
  name: string
  fieldName: string
  type: "USER"
  allowAnonymous: boolean
}
