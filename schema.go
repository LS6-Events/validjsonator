package validjsonator

type Schema struct {
	Ref                  string            `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Title                string            `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf           float64           `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum              float64           `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum     bool              `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum              float64           `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum     bool              `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength            int               `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength            int               `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern              string            `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems             int               `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems             int               `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems          bool              `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties        int               `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties        int               `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required             []string          `json:"required,omitempty" yaml:"required,omitempty"`
	Enum                 []interface{}     `json:"enum,omitempty" yaml:"enum,omitempty"`
	Type                 string            `json:"type,omitempty" yaml:"type,omitempty"`
	Format               string            `json:"format,omitempty" yaml:"format,omitempty"`
	AllOf                []Schema          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf                []Schema          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf                []Schema          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Not                  *Schema           `json:"not,omitempty" yaml:"not,omitempty"`
	Items                *Schema           `json:"items,omitempty" yaml:"items,omitempty"`
	Properties           map[string]Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	PatternProperties    map[string]Schema `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Description          string            `json:"description,omitempty" yaml:"description,omitempty"`
}
