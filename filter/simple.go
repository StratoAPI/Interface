package filter

type Operation string

const (
	// Equals
	OpEQ Operation = "eq"

	// Not Equals
	OpNEQ Operation = "neq"

	// Less Than
	OpLT Operation = "lt"

	// Less Than or Equals
	OpLTE Operation = "lte"

	// Greater Than
	OpGT Operation = "gt"

	// Greater Than or Equals
	OpGTE Operation = "gte"
)

type Simple struct {
	// The key to filter by
	Key string `json:"key"`

	// The filter operation
	Operation Operation `json:"op"`

	// The filter value
	Value interface{} `json:"val"`
}
