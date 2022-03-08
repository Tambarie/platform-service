package domain

type Category struct {
	Reference   string `json:"reference,omitempty" bson:"reference"`
	Name        string `validate:"required" json:"name,omitempty" bson:"name"`
	Description string `json:"description,omitempty" bson:"description"`
}

type SubCategory struct {
	Reference   string `json:"reference,omitempty" bson:"reference"`
	Name        string `validate:"required" json:"name,omitempty" bson:"name"`
	Description string `json:"description,omitempty" bson:"description"`
}
