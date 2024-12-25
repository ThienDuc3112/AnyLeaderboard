package leaderboard

// ============ Request body type ============
type createLeaderboardReqBody struct {
	Name                 string         `json:"name" validate:"required,isLBName"`
	Description          string         `json:"description" validate:"max=256"`
	CoverImageURL        string         `json:"cover_image_url" validate:"http_url"`
	ExternalLinks        []ExternalLink `json:"external_links" validate:"dive"`
	AllowAnonymous       bool           `json:"allow_anonymous"`
	RequiredVerification bool           `json:"required_verification"`
	Fields               []Field        `json:"fields" validate:"dive"`
}

type ExternalLink struct {
	DisplayValue string `json:"display_value" validate:"required,max=32"`
	URL          string `json:"url" validate:"required,http_url"`
}

type Field struct {
	Name       string   `json:"name" validate:"required"`
	Required   bool     `json:"required"`
	Hidden     bool     `json:"hidden"`
	FieldOrder int      `json:"field_order" validate:"required"`
	Type       string   `json:"type" validate:"required,oneof=TEXT SHORT_TEXT INTEGER REAL DURATION TIMESTAMP OPTION"`
	Options    []string `json:"options" validate:"required_if=Type OPTION"`
	ForRank    bool     `json:"for_rank"`
}
