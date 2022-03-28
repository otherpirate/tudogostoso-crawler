package crawler

type RecipeUrl struct {
	Category string
	Link     string
}

type Recipe struct {
	Category    string       `json:"category"`
	Link        string       `json:"link"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name        string `json:"name"`
	Quantity    string `json:"quantity"`
	Unit        string `json:"unit"`
	IsOptional  bool   `json:"is_optional"`
	Description string `json:"description"`
}
