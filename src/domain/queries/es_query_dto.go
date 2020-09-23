package queries

//EsQuery type
type EsQuery struct {
	Equals []FieldValue
}

//FieldValue type
type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
