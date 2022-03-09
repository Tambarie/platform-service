package domain

type State struct {
	Reference string `json:"reference,omitempty" bson:"reference"`
	Name      string `validate:"required" json:"name,omitempty" bson:"name"`
}
