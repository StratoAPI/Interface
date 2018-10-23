package schema

type Processor interface {
	// Check if a resource schema exists
	ResourceExists(resource string) bool

	// Check if the resource structure is valid
	ResourceValid(resource string, data string, required bool) (bool, error)

	// Check if the resource structure is valid
	ResourceValidGo(resource string, data interface{}, required bool) (bool, error)

	// Get the schema of a resource
	GetSchema(resource string) *Schema
}

type Schema struct {
	Source map[string]interface{}
	Meta   ResourceMeta
}

type ResourceMeta struct {
	Resource string `json:"resource"`
	Store    string `json:"store"`
}
